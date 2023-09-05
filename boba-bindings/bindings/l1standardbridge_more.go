// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const L1StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1007\"},{\"astId\":1006,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)46_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1007\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1StandardBridgeStorageLayout = new(solc.StorageLayout)

var L1StandardBridgeDeployedBin = "0x6080604052600436106101635760003560e01c806387087623116100c0578063a9f9e67511610074578063c4d66de811610059578063c4d66de814610491578063c89701a2146103ed578063e11013dd146104b157600080fd5b8063a9f9e6751461045e578063b1a1a8821461047e57600080fd5b806391c49bf8116100a557806391c49bf8146103ed578063927ede2d146104205780639a2ac6d51461044b57600080fd5b806387087623146103875780638f601f66146103a757600080fd5b8063540abf731161011757806358a997f6116100fc57806358a997f6146103135780637f46ddb214610333578063838b25201461036757600080fd5b8063540abf73146102d157806354fd4d50146102f157600080fd5b80631532ec34116101485780631532ec34146102545780631635f5fd146102675780633cb747bf1461027a57600080fd5b80630166a07a1461022157806309fc88431461024157600080fd5b3661021c57333b156101fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b61021a333362030d40604051806020016040528060008152506104c4565b005b600080fd5b34801561022d57600080fd5b5061021a61023c366004612558565b6104d7565b61021a61024f366004612609565b61089e565b61021a61026236600461265c565b610975565b61021a61027536600461265c565b610989565b34801561028657600080fd5b506003546102a79073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102dd57600080fd5b5061021a6102ec3660046126cf565b610dff565b3480156102fd57600080fd5b50610306610e44565b6040516102c891906127bc565b34801561031f57600080fd5b5061021a61032e3660046127cf565b610ee7565b34801561033f57600080fd5b506102a77f000000000000000000000000000000000000000000000000000000000000000081565b34801561037357600080fd5b5061021a6103823660046126cf565b610fbb565b34801561039357600080fd5b5061021a6103a23660046127cf565b611000565b3480156103b357600080fd5b506103df6103c2366004612852565b600260209081526000928352604080842090915290825290205481565b6040519081526020016102c8565b3480156103f957600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102a7565b34801561042c57600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff166102a7565b61021a61045936600461288b565b6110d4565b34801561046a57600080fd5b5061021a610479366004612558565b611116565b61021a61048c366004612609565b611125565b34801561049d57600080fd5b5061021a6104ac3660046128ee565b6111f6565b61021a6104bf36600461288b565b611269565b6104d184843485856112ac565b50505050565b60035473ffffffffffffffffffffffffffffffffffffffff16331480156105c65750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa15801561058a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ae919061290b565b73ffffffffffffffffffffffffffffffffffffffff16145b610678576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101f3565b61068187611492565b156107cf5761069087876114f4565b610742576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101f3565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b1580156107b257600080fd5b505af11580156107c6573d6000803e3d6000fd5b50505050610851565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461080d908490612957565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c1683529390529190912091909155610851908585611614565b610895878787878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116e892505050565b50505050505050565b333b1561092d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b6109703333348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112ac92505050565b505050565b6109828585858585610989565b5050505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610a785750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610a3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a60919061290b565b73ffffffffffffffffffffffffffffffffffffffff16145b610b2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101f3565b823414610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101f3565b3073ffffffffffffffffffffffffffffffffffffffff851603610c5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101f3565b60035473ffffffffffffffffffffffffffffffffffffffff90811690851603610d09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101f3565b610d4b85858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061177692505050565b6000610d68855a86604051806020016040528060008152506117e9565b905080610df7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101f3565b505050505050565b61089587873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061180392505050565b6060610e6f7f0000000000000000000000000000000000000000000000000000000000000000611b4b565b610e987f0000000000000000000000000000000000000000000000000000000000000000611b4b565b610ec17f0000000000000000000000000000000000000000000000000000000000000000611b4b565b604051602001610ed39392919061296e565b604051602081830303815290604052905090565b333b15610f76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b610df786863333888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c8892505050565b61089587873388888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c8892505050565b333b1561108f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b610df786863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061180392505050565b6104d133858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104c492505050565b610895878787878787876104d7565b333b156111b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b61097033338585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104c492505050565b610102600055600261120782611c97565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6104d13385348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112ac92505050565b82341461133b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c7565000060648201526084016101f3565b61134785858584611d75565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b9085907f0000000000000000000000000000000000000000000000000000000000000000907f1635f5fd00000000000000000000000000000000000000000000000000000000906113c6908b908b9086908a906024016129e4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261145992918890600401612a2d565b6000604051808303818588803b15801561147257600080fd5b505af1158015611486573d6000803e3d6000fd5b50505050505050505050565b60006114be827f1d1d8b6300000000000000000000000000000000000000000000000000000000611de8565b806114ee57506114ee827fec4fc8e300000000000000000000000000000000000000000000000000000000611de8565b92915050565b6000611520837f1d1d8b6300000000000000000000000000000000000000000000000000000000611de8565b156115c9578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611570573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611594919061290b565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161490506114ee565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611570573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109709084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611e0b565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b386868660405161176093929190612a72565b60405180910390a4610df7868686868686611f17565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e63184846040516117d5929190612ab0565b60405180910390a36104d184848484611f9f565b600080600080845160208601878a8af19695505050505050565b61180c87611492565b1561195a5761181b87876114f4565b6118cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101f3565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b15801561193d57600080fd5b505af1158015611951573d6000803e3d6000fd5b505050506119ee565b61197c73ffffffffffffffffffffffffffffffffffffffff881686308661200c565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546119ba908490612ac9565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b6119fc87878787878661206a565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b907f0000000000000000000000000000000000000000000000000000000000000000907f0166a07a0000000000000000000000000000000000000000000000000000000090611a7d908b908d908c908c908c908b90602401612ae1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611b1092918790600401612a2d565b600060405180830381600087803b158015611b2a57600080fd5b505af1158015611b3e573d6000803e3d6000fd5b5050505050505050505050565b606081600003611b8e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611bb85780611ba281612b3c565b9150611bb19050600a83612ba3565b9150611b92565b60008167ffffffffffffffff811115611bd357611bd3612bb7565b6040519080825280601f01601f191660200182016040528015611bfd576020820181803683370190505b5090505b8415611c8057611c12600183612957565b9150611c1f600a86612be6565b611c2a906030612ac9565b60f81b818381518110611c3f57611c3f612bfa565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611c79600a86612ba3565b9450611c01565b949350505050565b61089587878787878787611803565b600054610100900460ff16611d2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016101f3565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f238484604051611dd4929190612ab0565b60405180910390a36104d1848484846120f8565b6000611df383612157565b8015611e045750611e0483836121bb565b9392505050565b6000611e6d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661228a9092919063ffffffff16565b8051909150156109705780806020019051810190611e8b9190612c29565b610970576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f3565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd868686604051611f8f93929190612a72565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051611ffe929190612ab0565b60405180910390a350505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526104d19085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611666565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d03968686866040516120e293929190612a72565b60405180910390a4610df7868686868686612299565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051611ffe929190612ab0565b6000612183827f01ffc9a7000000000000000000000000000000000000000000000000000000006121bb565b80156114ee57506121b4827fffffffff000000000000000000000000000000000000000000000000000000006121bb565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015612273575060208210155b801561227f5750600081115b979650505050505050565b6060611c808484600085612311565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf868686604051611f8f93929190612a72565b6060824710156123a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f3565b73ffffffffffffffffffffffffffffffffffffffff85163b612421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f3565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161244a9190612c4b565b60006040518083038185875af1925050503d8060008114612487576040519150601f19603f3d011682016040523d82523d6000602084013e61248c565b606091505b509150915061227f828286606083156124a6575081611e04565b8251156124b65782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f391906127bc565b73ffffffffffffffffffffffffffffffffffffffff8116811461250c57600080fd5b50565b60008083601f84011261252157600080fd5b50813567ffffffffffffffff81111561253957600080fd5b60208301915083602082850101111561255157600080fd5b9250929050565b600080600080600080600060c0888a03121561257357600080fd5b873561257e816124ea565b9650602088013561258e816124ea565b9550604088013561259e816124ea565b945060608801356125ae816124ea565b93506080880135925060a088013567ffffffffffffffff8111156125d157600080fd5b6125dd8a828b0161250f565b989b979a50959850939692959293505050565b803563ffffffff8116811461260457600080fd5b919050565b60008060006040848603121561261e57600080fd5b612627846125f0565b9250602084013567ffffffffffffffff81111561264357600080fd5b61264f8682870161250f565b9497909650939450505050565b60008060008060006080868803121561267457600080fd5b853561267f816124ea565b9450602086013561268f816124ea565b935060408601359250606086013567ffffffffffffffff8111156126b257600080fd5b6126be8882890161250f565b969995985093965092949392505050565b600080600080600080600060c0888a0312156126ea57600080fd5b87356126f5816124ea565b96506020880135612705816124ea565b95506040880135612715816124ea565b94506060880135935061272a608089016125f0565b925060a088013567ffffffffffffffff8111156125d157600080fd5b60005b83811015612761578181015183820152602001612749565b838111156104d15750506000910152565b6000815180845261278a816020860160208601612746565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e046020830184612772565b60008060008060008060a087890312156127e857600080fd5b86356127f3816124ea565b95506020870135612803816124ea565b945060408701359350612818606088016125f0565b9250608087013567ffffffffffffffff81111561283457600080fd5b61284089828a0161250f565b979a9699509497509295939492505050565b6000806040838503121561286557600080fd5b8235612870816124ea565b91506020830135612880816124ea565b809150509250929050565b600080600080606085870312156128a157600080fd5b84356128ac816124ea565b93506128ba602086016125f0565b9250604085013567ffffffffffffffff8111156128d657600080fd5b6128e28782880161250f565b95989497509550505050565b60006020828403121561290057600080fd5b8135611e04816124ea565b60006020828403121561291d57600080fd5b8151611e04816124ea565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561296957612969612928565b500390565b60008451612980818460208901612746565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516129bc816001850160208a01612746565b600192019182015283516129d7816002840160208801612746565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525083604083015260806060830152612a236080830184612772565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201526000612a5c6060830185612772565b905063ffffffff83166040830152949350505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000612aa76060830184612772565b95945050505050565b828152604060208201526000611c806040830184612772565b60008219821115612adc57612adc612928565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152612b3060c0830184612772565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612b6d57612b6d612928565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612bb257612bb2612b74565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082612bf557612bf5612b74565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215612c3b57600080fd5b81518015158114611e0457600080fd5b60008251612c5d818460208701612746565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1StandardBridgeStorageLayoutJSON), L1StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1StandardBridge"] = L1StandardBridgeStorageLayout
	deployedBytecodes["L1StandardBridge"] = L1StandardBridgeDeployedBin
}
