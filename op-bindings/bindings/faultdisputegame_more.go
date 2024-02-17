// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const FaultDisputeGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"createdAt\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1001,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"resolvedAt\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1002,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"status\",\"offset\":16,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1011\"},{\"astId\":1003,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_struct(ClaimData)1012_storage)dyn_storage\"},{\"astId\":1004,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"credit\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1005,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1014,t_bool)\"},{\"astId\":1006,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"subgames\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\"},{\"astId\":1007,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"bonds\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1008,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"subgameAtRootResolved\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_bool\"},{\"astId\":1009,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"initialized\",\"offset\":1,\"slot\":\"6\",\"type\":\"t_bool\"},{\"astId\":1010,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"safetyMode\",\"offset\":2,\"slot\":\"6\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(ClaimData)1012_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IFaultDisputeGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1012_storage\"},\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_enum(GameStatus)1011\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint256)dyn_storage\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1014,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1014\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.ClaimData\",\"numberOfBytes\":\"160\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1013\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1014\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1015\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Position)1016\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1017\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var FaultDisputeGameStorageLayout = new(solc.StorageLayout)

var FaultDisputeGameDeployedBin = "0x6080604052600436106101e35760003560e01c80638d450a9511610102578063d8cc1a3c11610095578063f8f43ff611610064578063f8f43ff61461071d578063fa24f7431461073d578063fa315aa914610761578063fdffbb281461079457600080fd5b8063d8cc1a3c1461066c578063e1f0c3761461068c578063ec5e6308146106bf578063f3f7214e146106f257600080fd5b8063c55cd0c7116100d1578063c55cd0c714610581578063c6f0308c14610594578063cf09e0d01461061e578063d5d44d801461063f57600080fd5b80638d450a95146104a9578063bbdc02db146104dc578063bcef3b551461051d578063c395e1ca1461055a57600080fd5b806360e274641161017a5780638129fc1c116101495780638129fc1c146104375780638980e0cc1461043f5780638987c876146104545780638b85902b1461046957600080fd5b806360e2746414610391578063632247ea146103b15780636361506d146103c457806368800abf1461040457600080fd5b806335fef567116101b657806335fef567146102c05780633a768463146102d557806354fd4d5014610326578063609d33341461037c57600080fd5b80630356fe3a146101e857806319effeb41461022a578063200d2ed2146102705780632810e1d6146102ab575b600080fd5b3480156101f457600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b6040519081526020015b60405180910390f35b34801561023657600080fd5b506000546102579068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610221565b34801561027c57600080fd5b5060005461029e90700100000000000000000000000000000000900460ff1681565b6040516102219190613636565b3480156102b757600080fd5b5061029e6107a7565b6102d36102ce366004613677565b6109a4565b005b3480156102e157600080fd5b5060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610221565b34801561033257600080fd5b5061036f6040518060400160405280600581526020017f302e362e3000000000000000000000000000000000000000000000000000000081525081565b6040516102219190613704565b34801561038857600080fd5b5061036f6109b4565b34801561039d57600080fd5b506102d36103ac366004613739565b6109c7565b6102d36103bf366004613772565b610b36565b3480156103d057600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360200135610217565b34801561041057600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610217565b6102d3611358565b34801561044b57600080fd5b50600154610217565b34801561046057600080fd5b506102d3611709565b34801561047557600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360400135610217565b3480156104b557600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610217565b3480156104e857600080fd5b5060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610221565b34801561052957600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335610217565b34801561056657600080fd5b506102176105753660046137a7565b50662386f26fc1000090565b6102d361058f366004613677565b61182b565b3480156105a057600080fd5b506105b46105af3660046137d9565b611837565b6040805163ffffffff909816885273ffffffffffffffffffffffffffffffffffffffff968716602089015295909416948601949094526fffffffffffffffffffffffffffffffff9182166060860152608085015291821660a08401521660c082015260e001610221565b34801561062a57600080fd5b506000546102579067ffffffffffffffff1681565b34801561064b57600080fd5b5061021761065a366004613739565b60026020526000908152604090205481565b34801561067857600080fd5b506102d361068736600461383b565b6118ce565b34801561069857600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610257565b3480156106cb57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610217565b3480156106fe57600080fd5b506040516fffffffffffffffffffffffffffffffff8152602001610221565b34801561072957600080fd5b506102d36107383660046138c5565b611eac565b34801561074957600080fd5b50610752612345565b604051610221939291906138f1565b34801561076d57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610217565b6102d36107a23660046137d9565b6123a2565b600080600054700100000000000000000000000000000000900460ff1660028111156107d5576107d5613607565b1461080c576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065460ff16610848576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660016000815481106108745761087461391f565b6000918252602090912060059091020154640100000000900473ffffffffffffffffffffffffffffffffffffffff16146108af5760016108b2565b60025b6000805467ffffffffffffffff421668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff82168117835592935083927fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffff000000000000000000ffffffffffffffff9091161770010000000000000000000000000000000083600281111561096357610963613607565b02179055600281111561097857610978613607565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a290565b6109b082826000610b36565b5050565b60606109c260406020612803565b905090565b6000546109ee90620151809068010000000000000000900467ffffffffffffffff1661397d565b421015610a27576040517f8af80e8300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065460009062010000900460ff1615610a6b575073ffffffffffffffffffffffffffffffffffffffff811660009081526005602052604081208054919055610a97565b5073ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040812080549190555b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610af1576040519150601f19603f3d011682016040523d82523d6000602084013e610af6565b606091505b5050905080610b31576040517f83e6cc6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b60008054700100000000000000000000000000000000900460ff166002811115610b6257610b62613607565b14610b99576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060018481548110610bae57610bae61391f565b600091825260208083206040805160e0810182526005909402909101805463ffffffff808216865273ffffffffffffffffffffffffffffffffffffffff6401000000009092048216948601949094526001820154169184019190915260028101546fffffffffffffffffffffffffffffffff90811660608501526003820154608085015260049091015480821660a0850181905270010000000000000000000000000000000090910490911660c0840152919350909190610c73908390869061289a16565b90506000610d13826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169050861580610d555750610d527f0000000000000000000000000000000000000000000000000000000000000000600261397d565b81145b8015610d5f575084155b15610d96576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000811115610df0576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e1b7f0000000000000000000000000000000000000000000000000000000000000000600161397d565b8103610e2d57610e2d868885886128a2565b34662386f26fc100001115610e6e576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835160009063ffffffff90811614610ece576001856000015163ffffffff1681548110610e9d57610e9d61391f565b906000526020600020906005020160040160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b60c0850151600090610ef29067ffffffffffffffff165b67ffffffffffffffff1690565b67ffffffffffffffff1642610f1c610ee5856fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610f30919061397d565b610f3a9190613995565b90507f000000000000000000000000000000000000000000000000000000000000000060011c677fffffffffffffff1667ffffffffffffffff82161115610fad576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b421760008a8152608087901b6fffffffffffffffffffffffffffffffff8d1617602052604081209192509060008181526003602052604090205490915060ff161561102b576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016003600083815260200190815260200160002060006101000a81548160ff02191690831515021790555060016040518060e001604052808d63ffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001346fffffffffffffffffffffffffffffffff1681526020018c8152602001886fffffffffffffffffffffffffffffffff168152602001846fffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055506080820151816003015560a08201518160040160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060c08201518160040160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550505034600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112eb919061397d565b909155505060008b81526004602052604090206001805461130c9190613995565b8154600181018355600092835260208320015560405133918c918e917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a45050505050505050505050565b600654610100900460ff161561139a576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036040013511611451576040517ff40239db000000000000000000000000000000000000000000000000000000008152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560048201526024015b60405180910390fd5b60663611156114685763c407e0256000526004601cfd5b6040805160e08101825263ffffffff8082526000602080840182815232858701818152346fffffffffffffffffffffffffffffffff81811660608a01908152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560808b01908152600160a08c0181815242851660c08e0190815282548084018455928c529c517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6600593840290810180549b5192909e167fffffffffffffffff000000000000000000000000000000000000000000000000909b169a909a1764010000000073ffffffffffffffffffffffffffffffffffffffff9283160217909c5595517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7890180547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909c1617909a5590517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf8870180547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016918416919091179055517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf9860155915197519782167001000000000000000000000000000000009890921697909702177fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfa909201919091558252919091529182208054919290916116a090849061397d565b90915550506000805467ffffffffffffffff42167fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909116179055600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663452a93206040518163ffffffff1660e01b8152600401602060405180830381865afa158015611774573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061179891906139ac565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146117fc576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff1662010000179055565b6109b082826001610b36565b6001818154811061184757600080fd5b60009182526020909120600590910201805460018201546002830154600384015460049094015463ffffffff8416955064010000000090930473ffffffffffffffffffffffffffffffffffffffff908116949216926fffffffffffffffffffffffffffffffff91821692918082169170010000000000000000000000000000000090041687565b60008054700100000000000000000000000000000000900460ff1660028111156118fa576118fa613607565b14611931576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600187815481106119465761194661391f565b6000918252602082206005919091020160048101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506119a57f0000000000000000000000000000000000000000000000000000000000000000600161397d565b611a41826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1614611a82576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808915611b7157611ad57f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000613995565b6001901b611af4846fffffffffffffffffffffffffffffffff16612a63565b67ffffffffffffffff16611b0891906139f8565b15611b4557611b3c611b2d60016fffffffffffffffffffffffffffffffff8716613a0c565b865463ffffffff166000612b09565b60030154611b67565b7f00000000000000000000000000000000000000000000000000000000000000005b9150849050611b9b565b60038501549150611b98611b2d6fffffffffffffffffffffffffffffffff86166001613a3d565b90505b600882901b60088a8a604051611bb2929190613a71565b6040518091039020901b14611bf3576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611bfe8c612bed565b90506000611c0d836003015490565b6040517fe14ced320000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e14ced3290611c87908f908f908f908f908a90600401613aca565b6020604051808303816000875af1158015611ca6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cca9190613b04565b600485015491149150600090600290611d75906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611e11896fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611e1b9190613b1d565b611e259190613b3e565b67ffffffffffffffff161590508115158103611e6d576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505085547fffffffffffffffff0000000000000000000000000000000000000000ffffffff163364010000000002179095555050505050505050505050565b60008054700100000000000000000000000000000000900460ff166002811115611ed857611ed8613607565b14611f0f576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600080611f1e86612c1c565b93509350935093506000611f3485858585613049565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637dc0d1d06040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc791906139ac565b9050600189036120bf5773ffffffffffffffffffffffffffffffffffffffff81166352f0f3ad8a84612023367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036020013590565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815260048101939093526024830191909152604482015260206064820152608481018a905260a4015b6020604051808303816000875af1158015612095573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120b99190613b04565b5061233a565b600289036120eb5773ffffffffffffffffffffffffffffffffffffffff81166352f0f3ad8a8489612023565b600389036121175773ffffffffffffffffffffffffffffffffffffffff81166352f0f3ad8a8487612023565b6004890361228f5760006fffffffffffffffffffffffffffffffff8616156121af576121756fffffffffffffffffffffffffffffffff87167f0000000000000000000000000000000000000000000000000000000000000000613108565b61219f907f000000000000000000000000000000000000000000000000000000000000000061397d565b6121aa90600161397d565b6121d1565b7f00000000000000000000000000000000000000000000000000000000000000005b905073ffffffffffffffffffffffffffffffffffffffff82166352f0f3ad8b8560405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526004810192909252602482015260c084901b604482015260086064820152608481018b905260a4016020604051808303816000875af1158015612264573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122889190613b04565b505061233a565b60058903612308576040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018a9052602481018390524660c01b6044820152600860648201526084810188905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a401612076565b6040517fff137e6500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050505050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335606061239b6109b4565b9050909192565b60008054700100000000000000000000000000000000900460ff1660028111156123ce576123ce613607565b14612405576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006001828154811061241a5761241a61391f565b60009182526020822060059190910201600481015490925061245c90700100000000000000000000000000000000900460401c67ffffffffffffffff16610ee5565b600483015490915060009061248e90700100000000000000000000000000000000900467ffffffffffffffff16610ee5565b6124989042613b1d565b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c166124d28284613b65565b67ffffffffffffffff1611612513576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848152600460205260409020805485158015612533575060065460ff165b1561256a576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8015801561257757508515155b156125dc578454640100000000900473ffffffffffffffffffffffffffffffffffffffff16600081156125aa57816125c6565b600187015473ffffffffffffffffffffffffffffffffffffffff165b90506125d281886131bd565b5050505050505050565b60006fffffffffffffffffffffffffffffffff815b8381101561272257600085828154811061260d5761260d61391f565b6000918252602080832090910154808352600490915260409091205490915015612663576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600182815481106126785761267861391f565b600091825260209091206005909102018054909150640100000000900473ffffffffffffffffffffffffffffffffffffffff161580156126d1575060048101546fffffffffffffffffffffffffffffffff908116908516115b1561270f576001810154600482015473ffffffffffffffffffffffffffffffffffffffff90911695506fffffffffffffffffffffffffffffffff1693505b50508061271b90613b88565b90506125f1565b5061276a73ffffffffffffffffffffffffffffffffffffffff8316156127485782612764565b600188015473ffffffffffffffffffffffffffffffffffffffff165b886131bd565b86547fffffffffffffffff0000000000000000000000000000000000000000ffffffff1664010000000073ffffffffffffffffffffffffffffffffffffffff84160217875560008881526004602052604081206127c6916135cd565b876000036125d257600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555050505050505050565b6060600061283a84367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900361397d565b90508267ffffffffffffffff1667ffffffffffffffff81111561285f5761285f613bc0565b6040519080825280601f01601f191660200182016040528015612889576020820181803683370190505b509150828160208401375092915050565b151760011b90565b60006128c16fffffffffffffffffffffffffffffffff84166001613a3d565b905060006128d182866001612b09565b9050600086901a83806129c4575061290a60027f00000000000000000000000000000000000000000000000000000000000000006139f8565b60048301546002906129ae906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6129b89190613b3e565b67ffffffffffffffff16145b15612a1c5760ff8116600114806129de575060ff81166002145b612a17576040517ff40239db00000000000000000000000000000000000000000000000000000000815260048101889052602401611448565b612a5a565b60ff811615612a5a576040517ff40239db00000000000000000000000000000000000000000000000000000000815260048101889052602401611448565b50505050505050565b600080612af0837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b60008082612b5257612b4d6fffffffffffffffffffffffffffffffff86167f00000000000000000000000000000000000000000000000000000000000000006132aa565b612b6d565b612b6d856fffffffffffffffffffffffffffffffff16613471565b905060018481548110612b8257612b8261391f565b906000526020600020906005020191505b60048201546fffffffffffffffffffffffffffffffff828116911614612be557815460018054909163ffffffff16908110612bd057612bd061391f565b90600052602060002090600502019150612b93565b509392505050565b6000806000806000612bfe86612c1c565b9350935093509350612c1284848484613049565b9695505050505050565b6000806000806000859050600060018281548110612c3c57612c3c61391f565b600091825260209091206004600590920201908101549091507f000000000000000000000000000000000000000000000000000000000000000090612d13906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1611612d54576040517fb34b5c2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000815b60048301547f000000000000000000000000000000000000000000000000000000000000000090612e1b906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169250821115612e9757825463ffffffff16612e617f0000000000000000000000000000000000000000000000000000000000000000600161397d565b8303612e6b578391505b60018181548110612e7e57612e7e61391f565b9060005260206000209060050201935080945050612d58565b600481810154908401546fffffffffffffffffffffffffffffffff91821691166000816fffffffffffffffffffffffffffffffff16612f00612eeb856fffffffffffffffffffffffffffffffff1660011c90565b6fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff161490508015612fe5576000612f38836fffffffffffffffffffffffffffffffff16612a63565b67ffffffffffffffff161115612f9b576000612f72612f6a60016fffffffffffffffffffffffffffffffff8616613a0c565b896001612b09565b6003810154600490910154909c506fffffffffffffffffffffffffffffffff169a50612fbf9050565b7f00000000000000000000000000000000000000000000000000000000000000009a505b600386015460048701549099506fffffffffffffffffffffffffffffffff16975061303b565b6000613007612f6a6fffffffffffffffffffffffffffffffff85166001613a3d565b6003808901546004808b015492840154930154909e506fffffffffffffffffffffffffffffffff9182169d50919b50169850505b505050505050509193509193565b60006fffffffffffffffffffffffffffffffff841681036130af5782826040516020016130929291909182526fffffffffffffffffffffffffffffffff16602082015260400190565b604051602081830303815290604052805190602001209050613100565b60408051602081018790526fffffffffffffffffffffffffffffffff8087169282019290925260608101859052908316608082015260a0016040516020818303038152906040528051906020012090505b949350505050565b600080613195847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1690508083036001841b600180831b0386831b17039250505092915050565b60028101546fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000001810161322d576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280830180547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff17905573ffffffffffffffffffffffffffffffffffffffff841660009081526020919091526040812080548392906132a090849061397d565b9091555050505050565b600081613349846fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161161338a576040517fb34b5c2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61339383613471565b905081613432826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161161346b5761346861344f83600161397d565b6fffffffffffffffffffffffffffffffff83169061351d565b90505b92915050565b60008119600183011681613505827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b6000806135aa847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169050808303600180821b0385821b179250505092915050565b50805460008255906000526020600020908101906135eb91906135ee565b50565b5b8082111561360357600081556001016135ef565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310613671577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561368a57600080fd5b50508035926020909101359150565b6000815180845260005b818110156136bf576020818501810151868301820152016136a3565b818111156136d1576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006134686020830184613699565b73ffffffffffffffffffffffffffffffffffffffff811681146135eb57600080fd5b60006020828403121561374b57600080fd5b813561375681613717565b9392505050565b8035801515811461376d57600080fd5b919050565b60008060006060848603121561378757600080fd5b833592506020840135915061379e6040850161375d565b90509250925092565b6000602082840312156137b957600080fd5b81356fffffffffffffffffffffffffffffffff8116811461375657600080fd5b6000602082840312156137eb57600080fd5b5035919050565b60008083601f84011261380457600080fd5b50813567ffffffffffffffff81111561381c57600080fd5b60208301915083602082850101111561383457600080fd5b9250929050565b6000806000806000806080878903121561385457600080fd5b863595506138646020880161375d565b9450604087013567ffffffffffffffff8082111561388157600080fd5b61388d8a838b016137f2565b909650945060608901359150808211156138a657600080fd5b506138b389828a016137f2565b979a9699509497509295939492505050565b6000806000606084860312156138da57600080fd5b505081359360208301359350604090920135919050565b63ffffffff841681528260208201526060604082015260006139166060830184613699565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156139905761399061394e565b500190565b6000828210156139a7576139a761394e565b500390565b6000602082840312156139be57600080fd5b815161375681613717565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082613a0757613a076139c9565b500690565b60006fffffffffffffffffffffffffffffffff83811690831681811015613a3557613a3561394e565b039392505050565b60006fffffffffffffffffffffffffffffffff808316818516808303821115613a6857613a6861394e565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b606081526000613ade606083018789613a81565b8281036020840152613af1818688613a81565b9150508260408301529695505050505050565b600060208284031215613b1657600080fd5b5051919050565b600067ffffffffffffffff83811690831681811015613a3557613a3561394e565b600067ffffffffffffffff80841680613b5957613b596139c9565b92169190910692915050565b600067ffffffffffffffff808316818516808303821115613a6857613a6861394e565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613bb957613bb961394e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(FaultDisputeGameStorageLayoutJSON), FaultDisputeGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["FaultDisputeGame"] = FaultDisputeGameStorageLayout
	deployedBytecodes["FaultDisputeGame"] = FaultDisputeGameDeployedBin
	immutableReferences["FaultDisputeGame"] = true
}
