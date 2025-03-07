# boba-chain-ops

## boba-migrate
This package performs state regenesis. It takes the following input:

1. An `alloaction.json` file that contains a list of pre-allocated accounts.
2. A `genesis.json` file that contains the genesis block configuration.
3. A list of addresses that transacted on the network prior to this past regenesis.
4. A list of addresses that performed approvals on prior versions of the OVM ETH contract.
5. A list of msg information that stores the cross domain message from L2 to L1

It creates an initialized Bedrock erigon database as output. It does this by performing the following steps:

1. Create genesis (types.Genesis) from `genesis.json` and `allocation.json` in memory.
2. Set up the proxy contract for the predeploy contracts and update the bytecode of the predeploy contracts.
3. Migrate Boba legacy proxy implementation contract the new proxy contract and delete the slots
4. Migrate the ETH balance from the storage of OVM ETH contract to the balance field in genesis and delete the ETH balance storage slot.
5. Migrate the msg information from the storage of OVM_CrossDomainMessenger contract to the new slot.

It performs the following integrity checks:

1. OVM ETH storage slots must be completely accounted for.
2. The total supply of OVM ETH migrated must match the total supply of the OVM ETH contract.

### Compilation

Run `make boba-migrate`

## boba-rollover

This package performs state regenesis for creating a legacy chain in erigon. The new chain is only readable and is not compatible with v3. It takes the following input:

1. An `alloaction.json` file that contains a list of pre-allocated accounts.
2. A `genesis.json` file that contains the genesis block configuration.

3. A list of addresses that transacted on the network prior to this past regenesis.

It creates an initialized erigon database as output. It does this by performing the following steps:

1. Migrate the ETH balance from the storage of OVM ETH contract to the balance field in genesis and delete the ETH balance storage slot.

It performs the following integrity checks:

1. OVM ETH storage slots must be completely accounted for.
2. The total supply of OVM ETH migrated must match the total supply of the OVM ETH contract.

### Compilation

Run `make boba-rollover`

## boba-regenerate

This pacakge performs the chain regeneration via calling engine api. It does this by performing the following steps:

1. Call the legacy block chain to get the block and transaction information.
2. Build `PayloadAttributes` and call `engine_forkchoiceUpdatedV1` to get `PayloadID`
3. Call `engine_getPayloadV1` with the `PayloadID` from the last step to get `executionPayload`
4. Call `engine_newPayloadV1` with the `executionPayload` to execute the transaction
5. Call `engine_forkchoiceUpdatedV1` to build the block

It performs the following integrity checks:

1. The block hash and transaction hash must match the legacy block chain

### Compilation

Run `make boba-regenerate`

## boba-crawler

This package performs the process of getting addresses that send or receive ETH from the legacy block chain. It does this by performing the following steps:

1. Call `debug_traceTransaction` to find out addresses that send and receive ETH from internal transactions.
2. Call `eth_getLogs` to find out addresses that receive ETH from the ` ETH Mint` event

It performs the following integrity checks:

1. The address list can be used to compute the all storage keys of `OVM_ETH` contract from the allocation file.

### Compilation

Run `make boba-crawler`

## boba-devnet

This package generates a clean genesis file for devent. It only includes the predeployed contracts for L2. It takes the following input to generate the genesis file:

1. The deployment configuration for the l2
2. The hardhat deployment path
3. The l1 PRC endpoint for quering the block information

### Compilation

Run `make boba-devnet`

## boba-connect

This package generates a transition block between the legacy and new systems. It does this by performing the following steps:

1. A configuration file to get the timestamp of the transition block.
2. Use the engine api to create an empty block with the right block timestamp

### Compilation

Run `make boba-connect`

# Migration

## Preparation

### Files

Before the migration, we need these following files

* `witness.txt` - this includes the set of withdrawals through the `L2CrossDomainMessenger` from after the evm equivalence upgrade. To generate this file, we add the following environment variable in the l2geth client. 

  ```yaml
  environment:
    L2GETH_STATE_DUMP_PATH: "/dump/witness.txt"
  volumes:
    - VOLUME_LOCATION:/dump
  ```

* `eth-addresses.json` - this includes all addresses that have transferred or received ETH. When we migrate the slots from the `OVM_ETH` contract to the state, we need this file to compute the storage key of each address. Thes file is generated by `eth-crawler` in `boba-chain-ops` via running

  ```bash
  go run ./cmd/boba-crawler --rpc-url=https://mainnet.boba.network --output-path=./eth-addresses.json
  ```

  Another way to generate this file is by adding the following environment variable to the l2geth client, which is faster.

  ```yaml
  image: bobanetwork/l2geth:eth-dump
  environment:
    L2GETH_ETH_DUMP_PATH: "/dump/eth-address.txt"
  volumes:
    - VOLUME_LOCATION:/dump
  ```

* `eth-allowance.json` - this includes all allowances of the `OVM_ETH` contract.

  For Sepolia L2, this file is **EMPTY**. 

  For Mainnet L2, this file is needed and **SHOULDN'T** be empty. You can download this file from [s3](https://boba-db.s3.us-east-2.amazonaws.com/mainnet/eth-allowances.json).

* `ovm-message.json` - this includes the set of withdrawals through the `L2CrossDomainMessenger` from before the evm equivalence upgrade. 

  This file is **EMPTY**.

### Verification

Before the actual migration, we can ensure that `eth-addresses.json` and `eth-allowance.json` are able to generate all slots in the `OVM_ETH` contract, and that `witness.txt` can regenerate the cross-chain messages in the `L2CrossChainMessenger` contract. To verify this, you need to run [geth-dump](https://github.com/bobanetwork/boba/blob/develop/op-chain-ops/cmd/geth-dump/main.go) in `op-chain-ops` to generate the allocation file from the `l2geth` database. After obtaining the allocation file, you can execute `boba-crawler` in `boba-chain-ops` by running:

```bash
go run ./cmd/boba-crawler --eth-addresses-output-path=./eth-addresses.json --eth-allowances-output-path=./eth-allowance.json --witness-file-path=./witness.txt --alloc-path=alloc.json --post-check-only=true
```

### Contracts

#### Copy Deployment Files

To reduce upgrade time, we can deploy the L1 contracts first, without updating the implementation contracts or setting the L2 start time. We will copy and paste the deployment files from the `boba_legacy repo` to the `boba` repo. For example, we will copy and paste the [deployment files](https://github.com/bobanetwork/boba_legacy/tree/develop/packages/contracts/deployments/mainnet) of three contracts `Lib_AddressManager`, `Proxy__L1CrossDomainMessenger` and `Proxy__L1StandardBridge` to [contracts-bedrock](https://github.com/bobanetwork/boba/tree/develop/packages/contracts-bedrock/deployments) in the `boba` repo.

The names of the deployment files for `Proxy__L1StandardBridge` and `Proxy__L1CrossDomainMessenger` must be changed to `Proxy__OVM_L1CrossDomainMessenger.json` and `Proxy__OVM_L1StandardBridge.json`. Therefore, the deployment process can skip these two contracts.

**!! DON'T COPY OTHER DEPLOYMENT FILES TO THE FOLDER !!**

#### Add Network Configuration

The new network settings should be added to [hardhat.config.ts](https://github.com/bobanetwork/boba/blob/develop/packages/contracts-bedrock/hardhat.config.ts) and the configuration file should be added to the [deploy-config](https://github.com/bobanetwork/boba/tree/develop/packages/contracts-bedrock/deploy-config) folder.

To avoid the upgrade, delete the [024-SystemDictatorInit.ts](https://github.com/bobanetwork/boba/blob/develop/packages/contracts-bedrock/deploy/024-SystemDictatorInit.ts), [025-SystemDictatorSteps-1.ts](https://github.com/bobanetwork/boba/blob/develop/packages/contracts-bedrock/deploy/025-SystemDictatorSteps-1.ts) and [026-SystemDictatorSteps-2.ts](https://github.com/bobanetwork/boba/blob/develop/packages/contracts-bedrock/deploy/026-SystemDictatorSteps-2.ts) run

```bash
yarn deploy:hardhat --network boba-mainnet
```

In this process, we won't update the implementation contracts and initialize it. We will run the initialization after getting the actual `l1StartingBlockTag`, `l2OutputOracleStartingBlockNumber` and `l2OutputOracleStartingTimestamp` for the production.

### Erigon

A genesis file is needed when we generate the database. An example is

```json
{
  "config": {
    "chainId": 288,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "muirGlacierBlock": 0,
    "berlinBlock": 1000000000, // Future blocks
    "londonBlock": 1000000000, // Future blocks
    "arrowGlacierBlock": 1000000000, // Future blocks
    "grayGlacierBlock": 1000000000, // Future blocks
    "mergeNetsplitBlock": 1000000000, // Future blocks
    "bedrockBlock": 1000000000, // Future blocks
    "terminalTotalDifficulty": 0,
    "terminalTotalDifficultyPassed": true,
    "optimism": {
      "eip1559Elasticity": 6,
      "eip1559Denominator": 50,
      "eip1559DenominatorCanyon": 250
    },
    "regolithTime": 1811059381, // Future timestamp
    "shanghaiTime": 1811059381, // Future timestamp
    "canyonTime": 1811059381 // Future timestamp
  },
  "nonce": "0x0",
  "timestamp": "0x0",
  "extraData": "0x",
  "gasLimit": "0xA7D8C0", // 11,000,000
  "difficulty": "0x0",
  "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "coinbase": "0x4200000000000000000000000000000000000011",
  "alloc":{}, // MUST BE EMPTY
  "number": "0x0",
  "gasUsed": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "baseFeePerGas": "0x1"
}


```

**!! THE `alloc` MUST BE EMPTY !!**

Once we add the genesis block and the genesis file, we can create the db via

``` bash
./build/bin/erigon init --datadir db genesis.json
```

Now we can start the specific erigon built from the [boba-migration](https://github.com/bobanetwork/op-erigon/tree/bedrock-migration) branch.

```bash
./build/bin/erigon --datadir db --private.api.addr=localhost:9090 --http.addr=0.0.0.0 --http.port=9545 --http.corsdomain="*" --http.vhosts="*" --authrpc.addr=0.0.0.0 --authrpc.port=8551 --authrpc.vhosts="*" --authrpc.jwtsecret=jwt --rollup.disabletxpoolgossip=true --chain=dev --networkid=911 --http.api=eth,debug,net,engine,erigon,web3 --torrent.port=42068 --rollup.historicalrpc=http://localhost:8547 --log.console.verbosity dbug
```

>The `--rollup.historicalrpc` is required to re-generate the legacy chain
>
>*The `--chain` and `--networkid=` should be fixed*

Once we verify that the genesis block is the same as the legacy chain, we can start the [boba-regenerate](https://github.com/bobanetwork/boba/blob/develop/boba-chain-ops/cmd/boba-regenerate/main.go) in `boba-chain-ops` using `engine_api` to re-build the legacy block chain in erigon.

```bash
go run ./cmd/boba-regenerate --l2-legacy-endpoint=http://localhost:8547 --jwt-secret-path=jwt
```

## Process

Once we decide the hardfork block for the bedrock, we stop the legacy blockchain and dump the final state from it using the [geth-dump](https://github.com/bobanetwork/boba/blob/develop/op-chain-ops/cmd/geth-dump/main.go) in `op-chain-ops`. This creates a final state file for our migration.

```bash
go run ./cmd/geth-dump --db-path=db --output-path=genesis.json
```

> The **hardfork block number** is the **last block number of the legacy chain + 1**.

We now can stop the [boba-regenerate](https://github.com/bobanetwork/boba/blob/develop/boba-chain-ops/cmd/boba-regenerate/main.go) and the specific erigon built from the [bedrock-migration](https://github.com/bobanetwork/op-erigon/pull/58) branch once it reaches the hardfork block.

---

**!! Important !!**

**An L1 block and the hardfork block should be selected at this moment. The block hash and timestamp should be updated in the configuration file for the migration and L1 contract deployment.** For example, the values for  `l1StartingBlockTag` and `l2OutputOracleStartingTimestamp` should be set to the chosen block hash and timestamp respectively in [hardhat-local.json](https://github.com/bobanetwork/boba/blob/develop/packages/contracts-bedrock/deploy-config/hardhat-local.json). The `l2OutputOracleStartingBlockNumber`  should be set to the **hardfork block number** in [hardhat-local.json](https://github.com/bobanetwork/boba/blob/develop/packages/contracts-bedrock/deploy-config/hardhat-local.json). The `l1BobaTokenAddress`  should be updated to reflect the address where we deployed on L1.

---

Once we update the configuration file, we can continue the L1 contract deployment by running

```
yarn deploy:hardhat --network boba-mainnet
```

The command will finish the rest of configuration settings.

Then we create a genesis file for the migration. For example,

```json
{
  "config": {
    "chainId": 288,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "muirGlacierBlock": 0,
    "berlinBlock": 1135204, // Transition block
    "londonBlock": 1135204, // Transition block
    "arrowGlacierBlock": 1135204, // Transition block
    "grayGlacierBlock": 1135204, // Transition block
    "mergeNetsplitBlock": 1135204, // Transition block
    "bedrockBlock": 1135204, // Transition block
    "terminalTotalDifficulty": 0,
    "terminalTotalDifficultyPassed": true,
    "optimism": {
      "eip1559Elasticity": 6,
      "eip1559Denominator": 50,
      "eip1559DenominatorCanyon": 250
    },
    "regolithTime": 1711654404, // l2OutputOracleStartingTimestamp
    "shanghaiTime": 1711654404, // l2OutputOracleStartingTimestamp
    "canyonTime": 1711654404, // l2OutputOracleStartingTimestamp
    "cancunTime": 1711654405,  // l2OutputOracleStartingTimestamp + 1
    "ecotoneTime": 1711654405 // l2OutputOracleStartingTimestamp + 1 
  },
  "nonce": "0x0",
  "timestamp": "0x0",
  "extraData": "0x",
  "gasLimit": "0x1C9C380", // 30,000,000
  "difficulty": "0x22A4C7", // The difficulty of the last block
  "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "coinbase": "0x4200000000000000000000000000000000000011",
  "alloc":{},
  "number": "0x0",
  "gasUsed": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "baseFeePerGas": "0x3b9aca00"
}
```

>The `BEDROCK_BLOCK` must be the same as  `l2OutputOracleStartingBlockNumber`  which is **hardfork block number**

Once we prepare everything, we can run the [boba-mirgate](https://github.com/bobanetwork/boba/blob/develop/boba-chain-ops/cmd/boba-migrate/main.go) to insert the transition block.

```bash
go run ./cmd/boba-migrate
--l1-rpc-url=l1-node \
--db-path=erigon-db \
--alloc-path=geth-allocation.json \
--genesis-config-path=genesis.json \
--ovm-addresses=eth-addresses.json \
--ovm-allowances=eth-allowance.json \
--ovm-messages=ovm-message.json \
--witness-file=witness.txt \
--deploy-config=config.json \
--network=boba-mainnet \
--hardhat-deployments=deployment-folder \
--outfile-rollup=/files/rollup.json
--outfile-genesis=/files/alloc-transition.json
```

If all checks pass during the above process, we are ready to lanuch the bedrock!

## DevOps

Since we will run the op-node and erigon in **an EC instance**, is it possible to do the all following operations in this **EC2 instance**?

### Replica

We will run a replica node and add this new configuration to generate the `witness.txt`.

```yaml
environment:
  L2GETH_STATE_DUMP_PATH: "/dump/witness.txt"
  L2GETH_ETH_DUMP_PATH: "/dump/eth-address.txt"
volumes:
  - VOLUME_LOCATION:/dump
```

### Erigon

We will run an erigon node and re-generate the legacy block chain. Once we reach the hardfork block, we will do the migration and insert the transition block to this db. This db will be used to start the bedrock and distributed to all partners.

### boba-crawler

We will run a `boba-crawler` program to get all ETH addresses pointing at the **replica node**.

### boba-regenerate

We will run a `boba-regenerate` program to re-generate the legacy block chain pointing at the **erigon node**.

### Note

We will distribute the erigon db to our partners. What's the best way to handle this case?

## Abort migration

We can restore the legacy system by resetting the addresses in `Lib_AddressManager` and restoring the implementation contracts of the proxy contracts.

* Register following contracts

  ```json
  {
    "OVM_CanonicalTransactionChain",
    "OVM_L2CrossDomainMessenger",
    "OVM_DecompressionPrecompileAddress",
    "OVM_Sequencer",
    "OVM_Proposer",
    "OVM_ChainStorageContainer-CTC-batches",
    "OVM_ChainStorageContainer-CTC-queue",
    "OVM_CanonicalTransactionChain",
    "OVM_StateCommitmentChain",
    "OVM_BondManager",
    "OVM_ExecutionManager",
    "OVM_FraudVerifier",
    "OVM_StateManagerFactory",
    "OVM_StateTransitionerFactory",
    "OVM_SafetyChecker",
    "OVM_L1MultiMessageRelayer",
    "BondManager"
  }
  ```

* Update the following implementation contracts

  ```json
  {
  	"Proxy__L1CrossDomainMessenger",
  	"Proxy__L1StandardBridge"
  }
  ```
