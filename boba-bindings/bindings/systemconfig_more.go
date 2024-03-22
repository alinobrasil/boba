// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/boba/boba-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1010_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106102415760003560e01c8063935f029e11610145578063dac6e63a116100bd578063f45e65d81161008c578063f8c68de011610071578063f8c68de014610597578063fd32aa0f1461059f578063ffa1ad74146105a757600080fd5b8063f45e65d81461057a578063f68016b71461058357600080fd5b8063dac6e63a1461054e578063e0e2016d14610556578063e81b2c6d1461055e578063f2fde38b1461056757600080fd5b8063bc49ce5f11610114578063c71973f6116100f9578063c71973f6146103f4578063c9b26f6114610407578063cc731b021461041a57600080fd5b8063bc49ce5f146103e4578063c4e8ddfa146103ec57600080fd5b8063935f029e146103ae5780639b7d7f0a146103c1578063a7119869146103c9578063b40a817c146103d157600080fd5b806348cd4cb1116101d857806354fd4d50116101a757806361d157681161018c57806361d1576814610380578063715018a6146103885780638da5cb5b1461039057600080fd5b806354fd4d501461032f5780635d73369c1461037857600080fd5b806348cd4cb1146102d75780634add321d146102df5780634d9f1559146103005780634f16540b1461030857600080fd5b80630c18c162116102145780630c18c162146102ab57806318d13918146102b457806319f5cea8146102c75780631fd19ee1146102cf57600080fd5b806306c9265714610246578063078f29cf146102615780630a49cb031461028e5780630bbb796814610296575b600080fd5b61024e6105af565b6040519081526020015b60405180910390f35b6102696105dd565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610258565b610269610616565b6102a96102a4366004611928565b610646565b005b61024e60655481565b6102a96102c2366004611a55565b610a44565b61024e610a58565b610269610a83565b61024e610aad565b6102e7610add565b60405167ffffffffffffffff9091168152602001610258565b610269610b03565b61024e7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b61036b6040518060400160405280600681526020017f312e31322e30000000000000000000000000000000000000000000000000000081525081565b6040516102589190611ae2565b61024e610b33565b61024e610b5e565b6102a9610b89565b60335473ffffffffffffffffffffffffffffffffffffffff16610269565b6102a96103bc366004611af5565b610b9d565b610269610bb3565b610269610be3565b6102a96103df366004611b17565b610c13565b61024e610c24565b610269610c4f565b6102a9610402366004611b32565b610c7f565b6102a9610415366004611b4e565b610c90565b6104de6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516102589190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b610269610ca1565b61024e610cd1565b61024e60675481565b6102a9610575366004611a55565b610cfc565b61024e60665481565b6068546102e79067ffffffffffffffff1681565b61024e610db0565b61024e610ddb565b61024e600081565b6105da60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611b96565b81565b600061061161060d60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611b96565b5490565b905090565b600061061161060d60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611b96565b600054610100900460ff16158080156106665750600054600160ff909116105b806106805750303b158015610680575060005460ff166001145b610711576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561076f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610777610e2a565b6107808a610cfc565b61078987610ec9565b6107938989610f2a565b61079c86610fbb565b6107c57f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08869055565b6107f86107f360017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611b96565b849055565b61082c61082660017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611b96565b83519055565b61086361085a60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611b96565b60208401519055565b61089a61089160017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611b96565b60408401519055565b6108d16108c860017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611b96565b60608401519055565b6109086108ff60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611b96565b60808401519055565b61093f61093660017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611b96565b60a08401519055565b610947611099565b61095084611101565b610958610add565b67ffffffffffffffff168667ffffffffffffffff1610156109d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610708565b8015610a3857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b610a4c611575565b610a55816115f6565b50565b6105da60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611b96565b60006106117f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b600061061161060d60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a0611b96565b6069546000906106119063ffffffff6a0100000000000000000000820481169116611bad565b600061061161060d60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611b96565b6105da60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611b96565b6105da60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611b96565b610b91611575565b610b9b600061167a565b565b610ba5611575565b610baf8282610f2a565b5050565b600061061161060d60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611b96565b600061061161060d60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611b96565b610c1b611575565b610a5581610fbb565b6105da60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611b96565b600061061161060d60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611b96565b610c87611575565b610a5581611101565b610c98611575565b610a5581610ec9565b600061061161060d60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611b96565b6105da60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a0611b96565b610d04611575565b73ffffffffffffffffffffffffffffffffffffffff8116610da7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610708565b610a558161167a565b6105da60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611b96565b6105da60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611b96565b9055565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b5490565b600054610100900460ff16610ec1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610708565b610b9b6116f1565b60678190556040805160208082018490528251808303909101815290820190915260005b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610f1e9190611ae2565b60405180910390a35050565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610fae9190611ae2565b60405180910390a3505050565b610fc3610add565b67ffffffffffffffff168167ffffffffffffffff161015611040576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610708565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610eed565b6110c761060d60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a0611b96565b600003610b9b57610b9b6110fc60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a0611b96565b439055565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff1611156111b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d6178206261736500000000000000000000006064820152608401610708565b6001816040015160ff1611611248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e203100000000000000000000000000000000006064820152608401610708565b6068546080820151825167ffffffffffffffff909216916112699190611bd9565b63ffffffff1611156112d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610708565b6000816020015160ff161161136e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f74206265203000000000000000000000000000000000006064820152608401610708565b8051602082015163ffffffff82169160ff9091169061138e908290611bf8565b6113989190611c42565b63ffffffff161461142b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d69740000000000000000006064820152608401610708565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b60335473ffffffffffffffffffffffffffffffffffffffff163314610b9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610708565b61161f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08829055565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290506003610eed565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611788576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610708565b610b9b3361167a565b803573ffffffffffffffffffffffffffffffffffffffff811681146117b557600080fd5b919050565b803567ffffffffffffffff811681146117b557600080fd5b60405160c0810167ffffffffffffffff8111828210171561181c577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803563ffffffff811681146117b557600080fd5b803560ff811681146117b557600080fd5b600060c0828403121561185957600080fd5b60405160c0810181811067ffffffffffffffff821117156118a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040529050806118b283611822565b81526118c060208401611836565b60208201526118d160408401611836565b60408201526118e260608401611822565b60608201526118f360808401611822565b608082015260a08301356fffffffffffffffffffffffffffffffff8116811461191b57600080fd5b60a0919091015292915050565b6000806000806000806000806000898b0361026081121561194857600080fd5b6119518b611791565b995060208b0135985060408b0135975060608b0135965061197460808c016117ba565b955061198260a08c01611791565b94506119918c60c08d01611847565b93506119a06101808c01611791565b925060c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60820112156119d257600080fd5b506119db6117d2565b6119e86101a08c01611791565b81526119f76101c08c01611791565b6020820152611a096101e08c01611791565b6040820152611a1b6102008c01611791565b6060820152611a2d6102208c01611791565b6080820152611a3f6102408c01611791565b60a0820152809150509295985092959850929598565b600060208284031215611a6757600080fd5b611a7082611791565b9392505050565b6000815180845260005b81811015611a9d57602081850181015186830182015201611a81565b81811115611aaf576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611a706020830184611a77565b60008060408385031215611b0857600080fd5b50508035926020909101359150565b600060208284031215611b2957600080fd5b611a70826117ba565b600060c08284031215611b4457600080fd5b611a708383611847565b600060208284031215611b6057600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611ba857611ba8611b67565b500390565b600067ffffffffffffffff808316818516808303821115611bd057611bd0611b67565b01949350505050565b600063ffffffff808316818516808303821115611bd057611bd0611b67565b600063ffffffff80841680611c36577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff80831681851681830481118215151615611c6557611c65611b67565b0294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
