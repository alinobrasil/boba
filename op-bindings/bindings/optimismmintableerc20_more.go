// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const OptimismMintableERC20StorageLayoutJSON = "{\"storage\":[{\"astId\":28386,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":28392,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":28394,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":28396,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":28398,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":25564,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"remoteToken\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_address\"},{\"astId\":25567,\"contract\":\"contracts/universal/OptimismMintableERC20.sol:OptimismMintableERC20\",\"label\":\"bridge\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var OptimismMintableERC20StorageLayout = new(solc.StorageLayout)

var OptimismMintableERC20DeployedBin = "0x608060405234801561001057600080fd5b50600436106101365760003560e01c806395d89b41116100b2578063ae1f6aaf11610081578063d6c0b2c411610066578063d6c0b2c4146102bb578063dd62ed3e146102db578063e78cea921461032157600080fd5b8063ae1f6aaf1461025e578063c01e1bd61461029d57600080fd5b806395d89b411461021d5780639dc29fac14610225578063a457c2d714610238578063a9059cbb1461024b57600080fd5b806323b872dd1161010957806339509351116100ee57806339509351146101bf57806340c10f19146101d257806370a08231146101e757600080fd5b806323b872dd1461019d578063313ce567146101b057600080fd5b806301ffc9a71461013b57806306fdde0314610163578063095ea7b31461017857806318160ddd1461018b575b600080fd5b61014e610149366004611080565b610341565b60405190151581526020015b60405180910390f35b61016b610432565b60405161015a91906110c9565b61014e610186366004611165565b6104c4565b6002545b60405190815260200161015a565b61014e6101ab36600461118f565b6104dc565b6040516012815260200161015a565b61014e6101cd366004611165565b610500565b6101e56101e0366004611165565b61054c565b005b61018f6101f53660046111cb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b61016b610656565b6101e5610233366004611165565b610665565b61014e610246366004611165565b61075e565b61014e610259366004611165565b61082f565b60065473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b60055473ffffffffffffffffffffffffffffffffffffffff16610278565b6005546102789073ffffffffffffffffffffffffffffffffffffffff1681565b61018f6102e93660046111e6565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6006546102789073ffffffffffffffffffffffffffffffffffffffff1681565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007f0bc32271000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000085168314806103fa57507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b8061042957507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b60606003805461044190611219565b80601f016020809104026020016040519081016040528092919081815260200182805461046d90611219565b80156104ba5780601f1061048f576101008083540402835291602001916104ba565b820191906000526020600020905b81548152906001019060200180831161049d57829003601f168201915b5050505050905090565b6000336104d281858561083d565b5060019392505050565b6000336104ea8582856109f1565b6104f5858585610ac8565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104d2908290869061054790879061129b565b61083d565b60065473ffffffffffffffffffffffffffffffffffffffff1633146105f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b6106028282610d7b565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161064a91815260200190565b60405180910390a25050565b60606004805461044190611219565b60065473ffffffffffffffffffffffffffffffffffffffff16331461070c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084016105ef565b6107168282610e9b565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161064a91815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016105ef565b6104f5828686840361083d565b6000336104d2818585610ac8565b73ffffffffffffffffffffffffffffffffffffffff83166108df576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff8216610982576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ac25781811015610ab5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016105ef565b610ac2848484840361083d565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610b6b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff8216610c0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610d0890849061129b565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d6e91815260200190565b60405180910390a3610ac2565b73ffffffffffffffffffffffffffffffffffffffff8216610df8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105ef565b8060026000828254610e0a919061129b565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610e4490849061129b565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8216610f3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015610ff4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016105ef565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906110309084906112b3565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016109e4565b60006020828403121561109257600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146110c257600080fd5b9392505050565b600060208083528351808285015260005b818110156110f6578581018301518582016040015282016110da565b81811115611108576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461116057600080fd5b919050565b6000806040838503121561117857600080fd5b6111818361113c565b946020939093013593505050565b6000806000606084860312156111a457600080fd5b6111ad8461113c565b92506111bb6020850161113c565b9150604084013590509250925092565b6000602082840312156111dd57600080fd5b6110c28261113c565b600080604083850312156111f957600080fd5b6112028361113c565b91506112106020840161113c565b90509250929050565b600181811c9082168061122d57607f821691505b602082108103611266577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156112ae576112ae61126c565b500190565b6000828210156112c5576112c561126c565b50039056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC20StorageLayoutJSON), OptimismMintableERC20StorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC20"] = OptimismMintableERC20StorageLayout
	deployedBytecodes["OptimismMintableERC20"] = OptimismMintableERC20DeployedBin
}
