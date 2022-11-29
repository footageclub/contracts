// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClubERC20MetaData contains all meta data concerning the ClubERC20 contract.
var ClubERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200291838038062002918833981810160405281019062000037919062000346565b828281600390816200004a919062000621565b5080600490816200005c919062000621565b5050506000600560006101000a81548160ff0219169083151502179055506200009a6200008e620000aa60201b60201c565b620000b260201b60201c565b8060068190555050505062000708565b600033905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001e18262000196565b810181811067ffffffffffffffff82111715620002035762000202620001a7565b5b80604052505050565b60006200021862000178565b9050620002268282620001d6565b919050565b600067ffffffffffffffff821115620002495762000248620001a7565b5b620002548262000196565b9050602081019050919050565b60005b838110156200028157808201518184015260208101905062000264565b60008484015250505050565b6000620002a46200029e846200022b565b6200020c565b905082815260208101848484011115620002c357620002c262000191565b5b620002d084828562000261565b509392505050565b600082601f830112620002f057620002ef6200018c565b5b8151620003028482602086016200028d565b91505092915050565b6000819050919050565b62000320816200030b565b81146200032c57600080fd5b50565b600081519050620003408162000315565b92915050565b60008060006060848603121562000362576200036162000182565b5b600084015167ffffffffffffffff81111562000383576200038262000187565b5b6200039186828701620002d8565b935050602084015167ffffffffffffffff811115620003b557620003b462000187565b5b620003c386828701620002d8565b9250506040620003d6868287016200032f565b9150509250925092565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200043357607f821691505b602082108103620004495762000448620003eb565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004b37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000474565b620004bf868362000474565b95508019841693508086168417925050509392505050565b6000819050919050565b600062000502620004fc620004f6846200030b565b620004d7565b6200030b565b9050919050565b6000819050919050565b6200051e83620004e1565b620005366200052d8262000509565b84845462000481565b825550505050565b600090565b6200054d6200053e565b6200055a81848462000513565b505050565b5b8181101562000582576200057660008262000543565b60018101905062000560565b5050565b601f821115620005d1576200059b816200044f565b620005a68462000464565b81016020851015620005b6578190505b620005ce620005c58562000464565b8301826200055f565b50505b505050565b600082821c905092915050565b6000620005f660001984600802620005d6565b1980831691505092915050565b6000620006118383620005e3565b9150826002028217905092915050565b6200062c82620003e0565b67ffffffffffffffff811115620006485762000647620001a7565b5b6200065482546200041a565b6200066182828562000586565b600060209050601f83116001811462000699576000841562000684578287015190505b62000690858262000603565b86555062000700565b601f198416620006a9866200044f565b60005b82811015620006d357848901518255600182019150602085019450602081019050620006ac565b86831015620006f35784890151620006ef601f891682620005e3565b8355505b6001600288020188555050505b505050505050565b61220080620007186000396000f3fe6080604052600436106101355760003560e01c806370a08231116100ab57806395d89b411161006f57806395d89b41146103cf5780639dfde201146103fa578063a457c2d714610425578063a9059cbb14610462578063dd62ed3e1461049f578063f2fde38b146104dc57610135565b806370a0823114610310578063715018a61461034d57806379cc6790146103645780638456cb591461038d5780638da5cb5b146103a457610135565b806339509351116100fd57806339509351146102355780633ccfd60b146102725780633f4ba83a1461028957806340c10f19146102a057806342966c68146102bc5780635c975abb146102e557610135565b806306fdde031461013a578063095ea7b31461016557806318160ddd146101a257806323b872dd146101cd578063313ce5671461020a575b600080fd5b34801561014657600080fd5b5061014f610505565b60405161015c91906115ce565b60405180910390f35b34801561017157600080fd5b5061018c60048036038101906101879190611689565b610597565b60405161019991906116e4565b60405180910390f35b3480156101ae57600080fd5b506101b76105ba565b6040516101c4919061170e565b60405180910390f35b3480156101d957600080fd5b506101f460048036038101906101ef9190611729565b6105c4565b60405161020191906116e4565b60405180910390f35b34801561021657600080fd5b5061021f6105f3565b60405161022c9190611798565b60405180910390f35b34801561024157600080fd5b5061025c60048036038101906102579190611689565b6105fc565b60405161026991906116e4565b60405180910390f35b34801561027e57600080fd5b50610287610633565b005b34801561029557600080fd5b5061029e6106ea565b005b6102ba60048036038101906102b59190611689565b6106fc565b005b3480156102c857600080fd5b506102e360048036038101906102de91906117b3565b61074f565b005b3480156102f157600080fd5b506102fa610763565b60405161030791906116e4565b60405180910390f35b34801561031c57600080fd5b50610337600480360381019061033291906117e0565b61077a565b604051610344919061170e565b60405180910390f35b34801561035957600080fd5b506103626107c2565b005b34801561037057600080fd5b5061038b60048036038101906103869190611689565b6107d6565b005b34801561039957600080fd5b506103a26107f6565b005b3480156103b057600080fd5b506103b9610808565b6040516103c6919061181c565b60405180910390f35b3480156103db57600080fd5b506103e4610832565b6040516103f191906115ce565b60405180910390f35b34801561040657600080fd5b5061040f6108c4565b60405161041c919061170e565b60405180910390f35b34801561043157600080fd5b5061044c60048036038101906104479190611689565b6108ca565b60405161045991906116e4565b60405180910390f35b34801561046e57600080fd5b5061048960048036038101906104849190611689565b610941565b60405161049691906116e4565b60405180910390f35b3480156104ab57600080fd5b506104c660048036038101906104c19190611837565b610964565b6040516104d3919061170e565b60405180910390f35b3480156104e857600080fd5b5061050360048036038101906104fe91906117e0565b6109eb565b005b606060038054610514906118a6565b80601f0160208091040260200160405190810160405280929190818152602001828054610540906118a6565b801561058d5780601f106105625761010080835404028352916020019161058d565b820191906000526020600020905b81548152906001019060200180831161057057829003601f168201915b5050505050905090565b6000806105a2610a6e565b90506105af818585610a76565b600191505092915050565b6000600254905090565b6000806105cf610a6e565b90506105dc858285610c3f565b6105e7858585610ccb565b60019150509392505050565b60006012905090565b600080610607610a6e565b90506106288185856106198589610964565b6106239190611906565b610a76565b600191505092915050565b61063b610f4a565b60003373ffffffffffffffffffffffffffffffffffffffff16476040516106619061196b565b60006040518083038185875af1925050503d806000811461069e576040519150601f19603f3d011682016040523d82523d6000602084013e6106a3565b606091505b50509050806106e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106de906119cc565b60405180910390fd5b50565b6106f2610f4a565b6106fa610fc8565b565b600654341015610741576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073890611a38565b60405180910390fd5b61074b828261102b565b5050565b61076061075a610a6e565b8261118a565b50565b6000600560009054906101000a900460ff16905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6107ca610f4a565b6107d46000611360565b565b6107e8826107e2610a6e565b83610c3f565b6107f2828261118a565b5050565b6107fe610f4a565b610806611426565b565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610841906118a6565b80601f016020809104026020016040519081016040528092919081815260200182805461086d906118a6565b80156108ba5780601f1061088f576101008083540402835291602001916108ba565b820191906000526020600020905b81548152906001019060200180831161089d57829003601f168201915b5050505050905090565b60065481565b6000806108d5610a6e565b905060006108e38286610964565b905083811015610928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091f90611aca565b60405180910390fd5b6109358286868403610a76565b60019250505092915050565b60008061094c610a6e565b9050610959818585610ccb565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6109f3610f4a565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5990611b5c565b60405180910390fd5b610a6b81611360565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ae5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adc90611bee565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4b90611c80565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610c32919061170e565b60405180910390a3505050565b6000610c4b8484610964565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610cc55781811015610cb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cae90611cec565b60405180910390fd5b610cc48484848403610a76565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3190611d7e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610da9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da090611e10565b60405180910390fd5b610db4838383611489565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610e3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3190611ea2565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ecd9190611906565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f31919061170e565b60405180910390a3610f448484846114a1565b50505050565b610f52610a6e565b73ffffffffffffffffffffffffffffffffffffffff16610f70610808565b73ffffffffffffffffffffffffffffffffffffffff1614610fc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbd90611f0e565b60405180910390fd5b565b610fd06114a6565b6000600560006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611014610a6e565b604051611021919061181c565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361109a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109190611f7a565b60405180910390fd5b6110a660008383611489565b80600260008282546110b89190611906565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461110d9190611906565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611172919061170e565b60405180910390a3611186600083836114a1565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036111f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f09061200c565b60405180910390fd5b61120582600083611489565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561128b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112829061209e565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546112e291906120be565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611347919061170e565b60405180910390a361135b836000846114a1565b505050565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61142e6114ef565b6001600560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611472610a6e565b60405161147f919061181c565b60405180910390a1565b6114916114ef565b61149c838383611539565b505050565b505050565b6114ae610763565b6114ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114e49061213e565b60405180910390fd5b565b6114f7610763565b15611537576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161152e906121aa565b60405180910390fd5b565b505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561157857808201518184015260208101905061155d565b60008484015250505050565b6000601f19601f8301169050919050565b60006115a08261153e565b6115aa8185611549565b93506115ba81856020860161155a565b6115c381611584565b840191505092915050565b600060208201905081810360008301526115e88184611595565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611620826115f5565b9050919050565b61163081611615565b811461163b57600080fd5b50565b60008135905061164d81611627565b92915050565b6000819050919050565b61166681611653565b811461167157600080fd5b50565b6000813590506116838161165d565b92915050565b600080604083850312156116a05761169f6115f0565b5b60006116ae8582860161163e565b92505060206116bf85828601611674565b9150509250929050565b60008115159050919050565b6116de816116c9565b82525050565b60006020820190506116f960008301846116d5565b92915050565b61170881611653565b82525050565b600060208201905061172360008301846116ff565b92915050565b600080600060608486031215611742576117416115f0565b5b60006117508682870161163e565b93505060206117618682870161163e565b925050604061177286828701611674565b9150509250925092565b600060ff82169050919050565b6117928161177c565b82525050565b60006020820190506117ad6000830184611789565b92915050565b6000602082840312156117c9576117c86115f0565b5b60006117d784828501611674565b91505092915050565b6000602082840312156117f6576117f56115f0565b5b60006118048482850161163e565b91505092915050565b61181681611615565b82525050565b6000602082019050611831600083018461180d565b92915050565b6000806040838503121561184e5761184d6115f0565b5b600061185c8582860161163e565b925050602061186d8582860161163e565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806118be57607f821691505b6020821081036118d1576118d0611877565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061191182611653565b915061191c83611653565b9250828201905080821115611934576119336118d7565b5b92915050565b600081905092915050565b50565b600061195560008361193a565b915061196082611945565b600082019050919050565b600061197682611948565b9150819050919050565b7f4661696c656420746f2073656e6420456e746572000000000000000000000000600082015250565b60006119b6601483611549565b91506119c182611980565b602082019050919050565b600060208201905081810360008301526119e5816119a9565b9050919050565b7f45746865722076616c75652073656e74206973206e6f7420636f727265637400600082015250565b6000611a22601f83611549565b9150611a2d826119ec565b602082019050919050565b60006020820190508181036000830152611a5181611a15565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611ab4602583611549565b9150611abf82611a58565b604082019050919050565b60006020820190508181036000830152611ae381611aa7565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611b46602683611549565b9150611b5182611aea565b604082019050919050565b60006020820190508181036000830152611b7581611b39565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611bd8602483611549565b9150611be382611b7c565b604082019050919050565b60006020820190508181036000830152611c0781611bcb565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c6a602283611549565b9150611c7582611c0e565b604082019050919050565b60006020820190508181036000830152611c9981611c5d565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611cd6601d83611549565b9150611ce182611ca0565b602082019050919050565b60006020820190508181036000830152611d0581611cc9565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611d68602583611549565b9150611d7382611d0c565b604082019050919050565b60006020820190508181036000830152611d9781611d5b565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611dfa602383611549565b9150611e0582611d9e565b604082019050919050565b60006020820190508181036000830152611e2981611ded565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611e8c602683611549565b9150611e9782611e30565b604082019050919050565b60006020820190508181036000830152611ebb81611e7f565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611ef8602083611549565b9150611f0382611ec2565b602082019050919050565b60006020820190508181036000830152611f2781611eeb565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611f64601f83611549565b9150611f6f82611f2e565b602082019050919050565b60006020820190508181036000830152611f9381611f57565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611ff6602183611549565b915061200182611f9a565b604082019050919050565b6000602082019050818103600083015261202581611fe9565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000612088602283611549565b91506120938261202c565b604082019050919050565b600060208201905081810360008301526120b78161207b565b9050919050565b60006120c982611653565b91506120d483611653565b92508282039050818111156120ec576120eb6118d7565b5b92915050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000612128601483611549565b9150612133826120f2565b602082019050919050565b600060208201905081810360008301526121578161211b565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000612194601083611549565b915061219f8261215e565b602082019050919050565b600060208201905081810360008301526121c381612187565b905091905056fea2646970667358221220fe2da43ec039cd4574d4e8151904f25f7cbae5be1b87d8161c0289df24bf769f64736f6c63430008100033",
}

// ClubERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ClubERC20MetaData.ABI instead.
var ClubERC20ABI = ClubERC20MetaData.ABI

// ClubERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClubERC20MetaData.Bin instead.
var ClubERC20Bin = ClubERC20MetaData.Bin

// DeployClubERC20 deploys a new Ethereum contract, binding an instance of ClubERC20 to it.
func DeployClubERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _price *big.Int) (common.Address, *types.Transaction, *ClubERC20, error) {
	parsed, err := ClubERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClubERC20Bin), backend, _name, _symbol, _price)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClubERC20{ClubERC20Caller: ClubERC20Caller{contract: contract}, ClubERC20Transactor: ClubERC20Transactor{contract: contract}, ClubERC20Filterer: ClubERC20Filterer{contract: contract}}, nil
}

// ClubERC20 is an auto generated Go binding around an Ethereum contract.
type ClubERC20 struct {
	ClubERC20Caller     // Read-only binding to the contract
	ClubERC20Transactor // Write-only binding to the contract
	ClubERC20Filterer   // Log filterer for contract events
}

// ClubERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ClubERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClubERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ClubERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClubERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClubERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClubERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClubERC20Session struct {
	Contract     *ClubERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClubERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClubERC20CallerSession struct {
	Contract *ClubERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ClubERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClubERC20TransactorSession struct {
	Contract     *ClubERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ClubERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ClubERC20Raw struct {
	Contract *ClubERC20 // Generic contract binding to access the raw methods on
}

// ClubERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClubERC20CallerRaw struct {
	Contract *ClubERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ClubERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClubERC20TransactorRaw struct {
	Contract *ClubERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewClubERC20 creates a new instance of ClubERC20, bound to a specific deployed contract.
func NewClubERC20(address common.Address, backend bind.ContractBackend) (*ClubERC20, error) {
	contract, err := bindClubERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClubERC20{ClubERC20Caller: ClubERC20Caller{contract: contract}, ClubERC20Transactor: ClubERC20Transactor{contract: contract}, ClubERC20Filterer: ClubERC20Filterer{contract: contract}}, nil
}

// NewClubERC20Caller creates a new read-only instance of ClubERC20, bound to a specific deployed contract.
func NewClubERC20Caller(address common.Address, caller bind.ContractCaller) (*ClubERC20Caller, error) {
	contract, err := bindClubERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClubERC20Caller{contract: contract}, nil
}

// NewClubERC20Transactor creates a new write-only instance of ClubERC20, bound to a specific deployed contract.
func NewClubERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ClubERC20Transactor, error) {
	contract, err := bindClubERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClubERC20Transactor{contract: contract}, nil
}

// NewClubERC20Filterer creates a new log filterer instance of ClubERC20, bound to a specific deployed contract.
func NewClubERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ClubERC20Filterer, error) {
	contract, err := bindClubERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClubERC20Filterer{contract: contract}, nil
}

// bindClubERC20 binds a generic wrapper to an already deployed contract.
func bindClubERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClubERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClubERC20 *ClubERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClubERC20.Contract.ClubERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClubERC20 *ClubERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClubERC20.Contract.ClubERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClubERC20 *ClubERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClubERC20.Contract.ClubERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClubERC20 *ClubERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClubERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClubERC20 *ClubERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClubERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClubERC20 *ClubERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClubERC20.Contract.contract.Transact(opts, method, params...)
}

// Price is a free data retrieval call binding the contract method 0x9dfde201.
//
// Solidity: function Price() view returns(uint256)
func (_ClubERC20 *ClubERC20Caller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "Price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x9dfde201.
//
// Solidity: function Price() view returns(uint256)
func (_ClubERC20 *ClubERC20Session) Price() (*big.Int, error) {
	return _ClubERC20.Contract.Price(&_ClubERC20.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0x9dfde201.
//
// Solidity: function Price() view returns(uint256)
func (_ClubERC20 *ClubERC20CallerSession) Price() (*big.Int, error) {
	return _ClubERC20.Contract.Price(&_ClubERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ClubERC20 *ClubERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ClubERC20 *ClubERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ClubERC20.Contract.Allowance(&_ClubERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ClubERC20 *ClubERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ClubERC20.Contract.Allowance(&_ClubERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ClubERC20 *ClubERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ClubERC20 *ClubERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ClubERC20.Contract.BalanceOf(&_ClubERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ClubERC20 *ClubERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ClubERC20.Contract.BalanceOf(&_ClubERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ClubERC20 *ClubERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ClubERC20 *ClubERC20Session) Decimals() (uint8, error) {
	return _ClubERC20.Contract.Decimals(&_ClubERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ClubERC20 *ClubERC20CallerSession) Decimals() (uint8, error) {
	return _ClubERC20.Contract.Decimals(&_ClubERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ClubERC20 *ClubERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ClubERC20 *ClubERC20Session) Name() (string, error) {
	return _ClubERC20.Contract.Name(&_ClubERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ClubERC20 *ClubERC20CallerSession) Name() (string, error) {
	return _ClubERC20.Contract.Name(&_ClubERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClubERC20 *ClubERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClubERC20 *ClubERC20Session) Owner() (common.Address, error) {
	return _ClubERC20.Contract.Owner(&_ClubERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClubERC20 *ClubERC20CallerSession) Owner() (common.Address, error) {
	return _ClubERC20.Contract.Owner(&_ClubERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClubERC20 *ClubERC20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClubERC20 *ClubERC20Session) Paused() (bool, error) {
	return _ClubERC20.Contract.Paused(&_ClubERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClubERC20 *ClubERC20CallerSession) Paused() (bool, error) {
	return _ClubERC20.Contract.Paused(&_ClubERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ClubERC20 *ClubERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ClubERC20 *ClubERC20Session) Symbol() (string, error) {
	return _ClubERC20.Contract.Symbol(&_ClubERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ClubERC20 *ClubERC20CallerSession) Symbol() (string, error) {
	return _ClubERC20.Contract.Symbol(&_ClubERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ClubERC20 *ClubERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClubERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ClubERC20 *ClubERC20Session) TotalSupply() (*big.Int, error) {
	return _ClubERC20.Contract.TotalSupply(&_ClubERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ClubERC20 *ClubERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ClubERC20.Contract.TotalSupply(&_ClubERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Approve(&_ClubERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Approve(&_ClubERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ClubERC20 *ClubERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ClubERC20 *ClubERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Burn(&_ClubERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ClubERC20 *ClubERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Burn(&_ClubERC20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ClubERC20 *ClubERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ClubERC20 *ClubERC20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.BurnFrom(&_ClubERC20.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ClubERC20 *ClubERC20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.BurnFrom(&_ClubERC20.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ClubERC20 *ClubERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ClubERC20 *ClubERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.DecreaseAllowance(&_ClubERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ClubERC20 *ClubERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.DecreaseAllowance(&_ClubERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ClubERC20 *ClubERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ClubERC20 *ClubERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.IncreaseAllowance(&_ClubERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ClubERC20 *ClubERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.IncreaseAllowance(&_ClubERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) payable returns()
func (_ClubERC20 *ClubERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) payable returns()
func (_ClubERC20 *ClubERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Mint(&_ClubERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) payable returns()
func (_ClubERC20 *ClubERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Mint(&_ClubERC20.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClubERC20 *ClubERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClubERC20 *ClubERC20Session) Pause() (*types.Transaction, error) {
	return _ClubERC20.Contract.Pause(&_ClubERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClubERC20 *ClubERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _ClubERC20.Contract.Pause(&_ClubERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClubERC20 *ClubERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClubERC20 *ClubERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _ClubERC20.Contract.RenounceOwnership(&_ClubERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClubERC20 *ClubERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClubERC20.Contract.RenounceOwnership(&_ClubERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Transfer(&_ClubERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.Transfer(&_ClubERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.TransferFrom(&_ClubERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ClubERC20 *ClubERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClubERC20.Contract.TransferFrom(&_ClubERC20.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClubERC20 *ClubERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClubERC20 *ClubERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClubERC20.Contract.TransferOwnership(&_ClubERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClubERC20 *ClubERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClubERC20.Contract.TransferOwnership(&_ClubERC20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClubERC20 *ClubERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClubERC20 *ClubERC20Session) Unpause() (*types.Transaction, error) {
	return _ClubERC20.Contract.Unpause(&_ClubERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClubERC20 *ClubERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _ClubERC20.Contract.Unpause(&_ClubERC20.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ClubERC20 *ClubERC20Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClubERC20.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ClubERC20 *ClubERC20Session) Withdraw() (*types.Transaction, error) {
	return _ClubERC20.Contract.Withdraw(&_ClubERC20.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ClubERC20 *ClubERC20TransactorSession) Withdraw() (*types.Transaction, error) {
	return _ClubERC20.Contract.Withdraw(&_ClubERC20.TransactOpts)
}

// ClubERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ClubERC20 contract.
type ClubERC20ApprovalIterator struct {
	Event *ClubERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClubERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClubERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClubERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubERC20Approval represents a Approval event raised by the ClubERC20 contract.
type ClubERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ClubERC20 *ClubERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ClubERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ClubERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ClubERC20ApprovalIterator{contract: _ClubERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ClubERC20 *ClubERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ClubERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ClubERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubERC20Approval)
				if err := _ClubERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ClubERC20 *ClubERC20Filterer) ParseApproval(log types.Log) (*ClubERC20Approval, error) {
	event := new(ClubERC20Approval)
	if err := _ClubERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ClubERC20 contract.
type ClubERC20OwnershipTransferredIterator struct {
	Event *ClubERC20OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClubERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubERC20OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClubERC20OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClubERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubERC20OwnershipTransferred represents a OwnershipTransferred event raised by the ClubERC20 contract.
type ClubERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClubERC20 *ClubERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClubERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClubERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClubERC20OwnershipTransferredIterator{contract: _ClubERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClubERC20 *ClubERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClubERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClubERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubERC20OwnershipTransferred)
				if err := _ClubERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClubERC20 *ClubERC20Filterer) ParseOwnershipTransferred(log types.Log) (*ClubERC20OwnershipTransferred, error) {
	event := new(ClubERC20OwnershipTransferred)
	if err := _ClubERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubERC20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ClubERC20 contract.
type ClubERC20PausedIterator struct {
	Event *ClubERC20Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClubERC20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubERC20Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClubERC20Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClubERC20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubERC20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubERC20Paused represents a Paused event raised by the ClubERC20 contract.
type ClubERC20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ClubERC20 *ClubERC20Filterer) FilterPaused(opts *bind.FilterOpts) (*ClubERC20PausedIterator, error) {

	logs, sub, err := _ClubERC20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ClubERC20PausedIterator{contract: _ClubERC20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ClubERC20 *ClubERC20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ClubERC20Paused) (event.Subscription, error) {

	logs, sub, err := _ClubERC20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubERC20Paused)
				if err := _ClubERC20.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ClubERC20 *ClubERC20Filterer) ParsePaused(log types.Log) (*ClubERC20Paused, error) {
	event := new(ClubERC20Paused)
	if err := _ClubERC20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ClubERC20 contract.
type ClubERC20TransferIterator struct {
	Event *ClubERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClubERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClubERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClubERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubERC20Transfer represents a Transfer event raised by the ClubERC20 contract.
type ClubERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ClubERC20 *ClubERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ClubERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClubERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClubERC20TransferIterator{contract: _ClubERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ClubERC20 *ClubERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ClubERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClubERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubERC20Transfer)
				if err := _ClubERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ClubERC20 *ClubERC20Filterer) ParseTransfer(log types.Log) (*ClubERC20Transfer, error) {
	event := new(ClubERC20Transfer)
	if err := _ClubERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubERC20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ClubERC20 contract.
type ClubERC20UnpausedIterator struct {
	Event *ClubERC20Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClubERC20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubERC20Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClubERC20Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClubERC20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubERC20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubERC20Unpaused represents a Unpaused event raised by the ClubERC20 contract.
type ClubERC20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ClubERC20 *ClubERC20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*ClubERC20UnpausedIterator, error) {

	logs, sub, err := _ClubERC20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ClubERC20UnpausedIterator{contract: _ClubERC20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ClubERC20 *ClubERC20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ClubERC20Unpaused) (event.Subscription, error) {

	logs, sub, err := _ClubERC20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubERC20Unpaused)
				if err := _ClubERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ClubERC20 *ClubERC20Filterer) ParseUnpaused(log types.Log) (*ClubERC20Unpaused, error) {
	event := new(ClubERC20Unpaused)
	if err := _ClubERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
