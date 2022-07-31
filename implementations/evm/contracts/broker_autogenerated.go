// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BrokerBooking is an auto generated low-level Go binding around an user-defined struct.
type BrokerBooking struct {
	Index          *big.Int
	VmTypeId       *big.Int
	Miner          common.Address
	User           common.Address
	PricePerSecond *big.Int
	BookedAt       *big.Int
	BookedTill     *big.Int
}

// BrokerVMOffer is an auto generated low-level Go binding around an user-defined struct.
type BrokerVMOffer struct {
	Index             *big.Int
	Miner             common.Address
	PricePerSecond    *big.Int
	MachinesAvailable *big.Int
	VmTypeId          *big.Int
}

// BrokerMetaData contains all meta data concerning the Broker contract.
var BrokerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"communityAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingStopped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_signature\",\"type\":\"bytes20\"}],\"name\":\"SetMtlsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"abortType\",\"type\":\"uint8\"}],\"name\":\"abortBooking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"bookVM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"claimExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"extendBooking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"findBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vmTypeId\",\"type\":\"uint256\"}],\"name\":\"getAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoinDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMtlsHash\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUsersBookings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCommunityAddress\",\"type\":\"address\"}],\"name\":\"setCommunityContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCommunityFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setStablecoinAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"updateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060015560006003553480156200001b57600080fd5b5060405162004d9538038062004d958339818101604052810190620000419190620000f3565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000125565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000bb826200008e565b9050919050565b620000cd81620000ae565b8114620000d957600080fd5b50565b600081519050620000ed81620000c2565b92915050565b6000602082840312156200010c576200010b62000089565b5b60006200011c84828501620000dc565b91505092915050565b614c6080620001356000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806362c296171161010f578063bb0090c9116100a2578063cbb6bfbd11610071578063cbb6bfbd1461061d578063cf696d141461064d578063ea5f39841461066b578063f44f11cd1461069b576101f0565b8063bb0090c914610581578063bf1527651461059f578063c700ff0a146105bd578063ca3b3b4b146105ed576101f0565b80639ecd0500116100de5780639ecd0500146104d55780639f60c50c14610505578063a35310d414610521578063ab49ec4514610551576101f0565b806362c296171461044d5780637bef4f54146104695780637ee080c61461048757806390ad688b146104b7576101f0565b8063301bcaae1161018757806354fd4d501161015657806354fd4d50146103b1578063569e2e0f146103cf5780635a610b57146103ed578063626bbb271461041d576101f0565b8063301bcaae146103035780633b3e328b1461033357806348146113146103635780634eb264e814610381576101f0565b806320255c7e116101c357806320255c7e1461027b57806320e56d93146102995780632d0d6c33146102b55780632dc8a2b9146102d3576101f0565b80630d12bbdb146101f55780631032c61014610225578063117298e41461024157806318c5e5021461025d575b600080fd5b61020f600480360381019061020a9190613693565b6106b7565b60405161021c91906136db565b60405180910390f35b61023f600480360381019061023a91906136f6565b610770565b005b61025b600480360381019061025691906137a1565b61084c565b005b6102656108ba565b60405161027291906137dd565b60405180910390f35b6102836108c2565b60405161029091906137dd565b60405180910390f35b6102b360048036038101906102ae9190613871565b610967565b005b6102bd61107c565b6040516102ca91906137dd565b60405180910390f35b6102ed60048036038101906102e891906138b1565b61111f565b6040516102fa91906139bc565b60405180910390f35b61031d60048036038101906103189190613a03565b611234565b60405161032a9190613b6d565b60405180910390f35b61034d600480360381019061034891906136f6565b6114fe565b60405161035a91906137dd565b60405180910390f35b61036b6115f4565b60405161037891906137dd565b60405180910390f35b61039b60048036038101906103969190613693565b61163b565b6040516103a891906136db565b60405180910390f35b6103b96117cc565b6040516103c691906137dd565b60405180910390f35b6103d76117d1565b6040516103e49190613b9e565b60405180910390f35b61040760048036038101906104029190613a03565b6117fb565b6040516104149190613b6d565b60405180910390f35b61043760048036038101906104329190613a03565b611ac5565b6040516104449190613b6d565b60405180910390f35b610467600480360381019061046291906138b1565b611d8f565b005b61047161237d565b60405161047e9190613c18565b60405180910390f35b6104a1600480360381019061049c9190613693565b6123a7565b6040516104ae9190613d4a565b60405180910390f35b6104bf6125e5565b6040516104cc91906137dd565b60405180910390f35b6104ef60048036038101906104ea9190613daa565b6125ef565b6040516104fc91906136db565b60405180910390f35b61051f600480360381019061051a9190613dd7565b6126cb565b005b61053b60048036038101906105369190613a03565b612ac6565b60405161054891906136db565b60405180910390f35b61056b60048036038101906105669190613a03565b612ba2565b6040516105789190613e30565b60405180910390f35b610589612beb565b60405161059691906137dd565b60405180910390f35b6105a7612c32565b6040516105b491906137dd565b60405180910390f35b6105d760048036038101906105d29190613a03565b612cc3565b6040516105e49190613d4a565b60405180910390f35b61060760048036038101906106029190613693565b612f2b565b60405161061491906136db565b60405180910390f35b61063760048036038101906106329190613e4b565b613074565b60405161064491906137dd565b60405180910390f35b610655613475565b6040516106629190613e9a565b60405180910390f35b61068560048036038101906106809190613a03565b61350d565b6040516106929190613ec4565b60405180910390f35b6106b560048036038101906106b09190613f0b565b613563565b005b600080821180156106c85750606582105b80156107215750600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610760576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075790613fbb565b60405180910390fd5b81600a8190555060009050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610813576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080a9061404d565b60405180910390fd5b80600080858152602001908152602001600020600301819055508160008085815260200190815260200160002060040181905550505050565b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c021790555050565b600042905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161092192919061406d565b602060405180830381865afa15801561093e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096291906140ab565b905090565b6000600260008467ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815260200160068201548152505090506000610a7b6108ba565b9050816060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610aef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae690614124565b60405180910390fd5b60018360ff161480610b04575060028360ff16145b610b43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3a906141b6565b60405180910390fd5b8160c001518110610b89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8090614248565b60405180910390fd5b60008060018560ff1603610ba4576032915060329050610bbb565b600a546064610bb39190614297565b9150600a5490505b60008460a0015184610bcd9190614297565b8560800151610bdc91906142cb565b9050600060648483610bee91906142cb565b610bf89190614354565b905060008183610c089190614297565b905082600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c599190614297565b925050819055508660a001518760c00151610c749190614297565b8760800151610c8391906142cb565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cd19190614297565b92505081905550600260008860000151815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600482016000905560058201600090556006820160009055505060018860ff1603610e0257866060015173ffffffffffffffffffffffffffffffffffffffff16876040015173ffffffffffffffffffffffffffffffffffffffff167f26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41848a600001518b60a001518b610de09190614297565b8c60200151604051610df59493929190614385565b60405180910390a3610e8d565b866060015173ffffffffffffffffffffffffffffffffffffffff16876040015173ffffffffffffffffffffffffffffffffffffffff167f232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1848a600001518b60a001518b610e6f9190614297565b8c60200151604051610e849493929190614385565b60405180910390a35b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8860400151846040518363ffffffff1660e01b8152600401610eee9291906143ca565b6020604051808303816000875af1158015610f0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f31919061441f565b610f70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6790614498565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610fef9291906143ca565b6020604051808303816000875af115801561100e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611032919061441f565b611071576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106890614504565b60405180910390fd5b505050505050505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016110d99190613b9e565b602060405180830381865afa1580156110f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111a91906140ab565b905090565b6111276135aa565b600260008367ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815250509050919050565b6060600060035467ffffffffffffffff81111561125457611253614524565b5b60405190808252806020026020018201604052801561128d57816020015b61127a6135aa565b8152602001906001900390816112725790505b509050600080600090505b600354811015611448578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361143557600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815260200160068201548152505083838151811061141a57611419614553565b5b60200260200101819052506001826114329190614582565b91505b8080611440906145d8565b915050611298565b508067ffffffffffffffff81111561146357611462614524565b5b60405190808252806020026020018201604052801561149c57816020015b6114896135aa565b8152602001906001900390816114815790505b50925060005b818110156114f6578281815181106114bd576114bc614553565b5b60200260200101518482815181106114d8576114d7614553565b5b602002602001018190525080806114ee906145d8565b9150506114a2565b505050919050565b60006040518060a0016040528060015481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018381526020018481525060008060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015560808201518160040155905050600160008154809291906115d8906145d8565b9190505550600180546115eb9190614297565b90509392505050565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b600081611646612c32565b1015611687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167e9061466c565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016116e49291906143ca565b6020604051808303816000875af1158015611703573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611727919061441f565b61173457600090506117c7565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461177f9190614297565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600181565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060035467ffffffffffffffff81111561181b5761181a614524565b5b60405190808252806020026020018201604052801561185457816020015b6118416135aa565b8152602001906001900390816118395790505b509050600080600090505b600354811015611a0f578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036119fc57600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815250508383815181106119e1576119e0614553565b5b60200260200101819052506001826119f99190614582565b91505b8080611a07906145d8565b91505061185f565b508067ffffffffffffffff811115611a2a57611a29614524565b5b604051908082528060200260200182016040528015611a6357816020015b611a506135aa565b815260200190600190039081611a485790505b50925060005b81811015611abd57828181518110611a8457611a83614553565b5b6020026020010151848281518110611a9f57611a9e614553565b5b60200260200101819052508080611ab5906145d8565b915050611a69565b505050919050565b6060600060035467ffffffffffffffff811115611ae557611ae4614524565b5b604051908082528060200260200182016040528015611b1e57816020015b611b0b6135aa565b815260200190600190039081611b035790505b509050600080600090505b600354811015611cd9578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611cc657600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152602001600682015481525050838381518110611cab57611caa614553565b5b6020026020010181905250600182611cc39190614582565b91505b8080611cd1906145d8565b915050611b29565b508067ffffffffffffffff811115611cf457611cf3614524565b5b604051908082528060200260200182016040528015611d2d57816020015b611d1a6135aa565b815260200190600190039081611d125790505b50925060005b81811015611d8757828181518110611d4e57611d4d614553565b5b6020026020010151848281518110611d6957611d68614553565b5b60200260200101819052508080611d7f906145d8565b915050611d33565b505050919050565b6000600260008367ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815260200160068201548152505090506000611ea36108ba565b9050816040015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f0e906146d8565b60405180910390fd5b8160c00151811015611f5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f5590614744565b60405180910390fd5b60008260a001518360c00151611f749190614297565b8360800151611f8391906142cb565b905060006064600a546064611f989190614297565b83611fa391906142cb565b611fad9190614354565b905060008183611fbd9190614297565b90508260066000876060015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546120129190614297565b925050819055508260076000876060015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461206c9190614297565b92505081905550600260008660000151815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160009055600582016000905560068201600090555050846060015173ffffffffffffffffffffffffffffffffffffffff16856040015173ffffffffffffffffffffffffffffffffffffffff167f75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c8488600001518960a001518a60c001516121749190614297565b8a602001516040516121899493929190614385565b60405180910390a3600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8660400151846040518363ffffffff1660e01b81526004016121f29291906143ca565b6020604051808303816000875af1158015612211573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612235919061441f565b612274576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161226b90614498565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016122f39291906143ca565b6020604051808303816000875af1158015612312573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612336919061441f565b612375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161236c90614504565b60405180910390fd5b505050505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060015467ffffffffffffffff8111156123c7576123c6614524565b5b60405190808252806020026020018201604052801561240057816020015b6123ed613613565b8152602001906001900390816123e55790505b509050600080600090505b60015481101561252f57846000808381526020019081526020016000206004015414801561244e5750600080600083815260200190815260200160002060030154115b1561251c576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152505083838151811061250157612500614553565b5b60200260200101819052506001826125199190614582565b91505b8080612527906145d8565b91505061240b565b508067ffffffffffffffff81111561254a57612549614524565b5b60405190808252806020026020018201604052801561258357816020015b612570613613565b8152602001906001900390816125685790505b50925060005b818110156125dd578281815181106125a4576125a3614553565b5b60200260200101518482815181106125bf576125be614553565b5b602002602001018190525080806125d5906145d8565b915050612589565b505050919050565b6000600a54905090565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612681576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612678906147d6565b60405180910390fd5b81600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b6000600260008467ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152602001600682015481525050905060006127df6108ba565b9050816060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612853576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161284a90614842565b60405180910390fd5b8160c001518110612899576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612890906148d4565b60405180910390fd5b8282608001516128a991906142cb565b6128b1612c32565b10156128f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128e990614966565b60405180910390fd5b82826080015161290291906142cb565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546129509190614582565b92505081905550828260c001516129679190614582565b5081600260008667ffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a0820151816005015560c08201518160060155905050816060015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff167f075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd9029814746884600001518560200151604051612ab8929190614986565b60405180910390a350505050565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612b58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b4f90614a21565b60405180910390fd5b81600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612cbe9190614297565b905090565b6060600060015467ffffffffffffffff811115612ce357612ce2614524565b5b604051908082528060200260200182016040528015612d1c57816020015b612d09613613565b815260200190600190039081612d015790505b509050600080600090505b600154811015612e75578473ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612e62576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201548152602001600482015481525050838381518110612e4757612e46614553565b5b6020026020010181905250600182612e5f9190614582565b91505b8080612e6d906145d8565b915050612d27565b508067ffffffffffffffff811115612e9057612e8f614524565b5b604051908082528060200260200182016040528015612ec957816020015b612eb6613613565b815260200190600190039081612eae5790505b50925060005b81811015612f2357828181518110612eea57612ee9614553565b5b6020026020010151848281518110612f0557612f04614553565b5b60200260200101819052508080612f1b906145d8565b915050612ecf565b505050919050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401612f8c93929190614a41565b6020604051808303816000875af1158015612fab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fcf919061441f565b612fdc576000905061306f565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546130279190614582565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b6000806000808581526020019081526020016000206003015411801561309a5750600082115b80156130cd575081600080858152602001908152602001600020600201546130c291906142cb565b6130ca612c32565b10155b61310c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161310390614b36565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036131b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131a790614bc8565b60405180910390fd5b81600080858152602001908152602001600020600201546131d191906142cb565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461321f9190614582565b9250508190555060006040518060e00160405280600354815260200160008087815260200190815260200160002060040154815260200160008087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200160008087815260200190815260200160002060020154815260200142815260200184426132f09190614582565b815250905080600260006003548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a0820151816005015560c08201518160060155905050600360008154809291906133e1906145d8565b9190505550806060015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff167f246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b683600001518460200151604051613455929190614986565b60405180910390a3600160035461346c9190614297565b91505092915050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135089190614bfd565b905090565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6040518060e001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b6136708161365d565b811461367b57600080fd5b50565b60008135905061368d81613667565b92915050565b6000602082840312156136a9576136a8613658565b5b60006136b78482850161367e565b91505092915050565b60008115159050919050565b6136d5816136c0565b82525050565b60006020820190506136f060008301846136cc565b92915050565b60008060006060848603121561370f5761370e613658565b5b600061371d8682870161367e565b935050602061372e8682870161367e565b925050604061373f8682870161367e565b9150509250925092565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b61377e81613749565b811461378957600080fd5b50565b60008135905061379b81613775565b92915050565b6000602082840312156137b7576137b6613658565b5b60006137c58482850161378c565b91505092915050565b6137d78161365d565b82525050565b60006020820190506137f260008301846137ce565b92915050565b600067ffffffffffffffff82169050919050565b613815816137f8565b811461382057600080fd5b50565b6000813590506138328161380c565b92915050565b600060ff82169050919050565b61384e81613838565b811461385957600080fd5b50565b60008135905061386b81613845565b92915050565b6000806040838503121561388857613887613658565b5b600061389685828601613823565b92505060206138a78582860161385c565b9150509250929050565b6000602082840312156138c7576138c6613658565b5b60006138d584828501613823565b91505092915050565b6138e78161365d565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613918826138ed565b9050919050565b6139288161390d565b82525050565b60e08201600082015161394460008501826138de565b50602082015161395760208501826138de565b50604082015161396a604085018261391f565b50606082015161397d606085018261391f565b50608082015161399060808501826138de565b5060a08201516139a360a08501826138de565b5060c08201516139b660c08501826138de565b50505050565b600060e0820190506139d1600083018461392e565b92915050565b6139e08161390d565b81146139eb57600080fd5b50565b6000813590506139fd816139d7565b92915050565b600060208284031215613a1957613a18613658565b5b6000613a27848285016139ee565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60e082016000820151613a7260008501826138de565b506020820151613a8560208501826138de565b506040820151613a98604085018261391f565b506060820151613aab606085018261391f565b506080820151613abe60808501826138de565b5060a0820151613ad160a08501826138de565b5060c0820151613ae460c08501826138de565b50505050565b6000613af68383613a5c565b60e08301905092915050565b6000602082019050919050565b6000613b1a82613a30565b613b248185613a3b565b9350613b2f83613a4c565b8060005b83811015613b60578151613b478882613aea565b9750613b5283613b02565b925050600181019050613b33565b5085935050505092915050565b60006020820190508181036000830152613b878184613b0f565b905092915050565b613b988161390d565b82525050565b6000602082019050613bb36000830184613b8f565b92915050565b6000819050919050565b6000613bde613bd9613bd4846138ed565b613bb9565b6138ed565b9050919050565b6000613bf082613bc3565b9050919050565b6000613c0282613be5565b9050919050565b613c1281613bf7565b82525050565b6000602082019050613c2d6000830184613c09565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a082016000820151613c7560008501826138de565b506020820151613c88602085018261391f565b506040820151613c9b60408501826138de565b506060820151613cae60608501826138de565b506080820151613cc160808501826138de565b50505050565b6000613cd38383613c5f565b60a08301905092915050565b6000602082019050919050565b6000613cf782613c33565b613d018185613c3e565b9350613d0c83613c4f565b8060005b83811015613d3d578151613d248882613cc7565b9750613d2f83613cdf565b925050600181019050613d10565b5085935050505092915050565b60006020820190508181036000830152613d648184613cec565b905092915050565b6000613d778261390d565b9050919050565b613d8781613d6c565b8114613d9257600080fd5b50565b600081359050613da481613d7e565b92915050565b600060208284031215613dc057613dbf613658565b5b6000613dce84828501613d95565b91505092915050565b60008060408385031215613dee57613ded613658565b5b6000613dfc85828601613823565b9250506020613e0d8582860161367e565b9150509250929050565b6000819050919050565b613e2a81613e17565b82525050565b6000602082019050613e456000830184613e21565b92915050565b60008060408385031215613e6257613e61613658565b5b6000613e708582860161367e565b9250506020613e818582860161367e565b9150509250929050565b613e9481613838565b82525050565b6000602082019050613eaf6000830184613e8b565b92915050565b613ebe81613749565b82525050565b6000602082019050613ed96000830184613eb5565b92915050565b613ee881613e17565b8114613ef357600080fd5b50565b600081359050613f0581613edf565b92915050565b600060208284031215613f2157613f20613658565b5b6000613f2f84828501613ef6565b91505092915050565b600082825260208201905092915050565b7f636f6d6d756e697479206665652073686f756c6420626520696e2072616e676560008201527f206f66203120746f203130300000000000000000000000000000000000000000602082015250565b6000613fa5602c83613f38565b9150613fb082613f49565b604082019050919050565b60006020820190508181036000830152613fd481613f98565b9050919050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000614037602283613f38565b915061404282613fdb565b604082019050919050565b600060208201905081810360008301526140668161402a565b9050919050565b60006040820190506140826000830185613b8f565b61408f6020830184613b8f565b9392505050565b6000815190506140a581613667565b92915050565b6000602082840312156140c1576140c0613658565b5b60006140cf84828501614096565b91505092915050565b7f6f6e6c79206f776e6572206f6620626f6f6b696e672063616e207265706f7274600082015250565b600061410e602083613f38565b9150614119826140d8565b602082019050919050565b6000602082019050818103600083015261413d81614101565b9050919050565b7f7265706f7274206f722073746f7020626f6f6b696e672069732072657175697260008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b60006141a0602283613f38565b91506141ab82614144565b604082019050919050565b600060208201905081810360008301526141cf81614193565b9050919050565b7f626f6f6b696e6720697320657870697265642c206d696e65722073686f756c6460008201527f20636c61696d206578706972656420626f6f6b696e6700000000000000000000602082015250565b6000614232603683613f38565b915061423d826141d6565b604082019050919050565b6000602082019050818103600083015261426181614225565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006142a28261365d565b91506142ad8361365d565b9250828210156142c0576142bf614268565b5b828203905092915050565b60006142d68261365d565b91506142e18361365d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561431a57614319614268565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061435f8261365d565b915061436a8361365d565b92508261437a57614379614325565b5b828204905092915050565b600060808201905061439a60008301876137ce565b6143a760208301866137ce565b6143b460408301856137ce565b6143c160608301846137ce565b95945050505050565b60006040820190506143df6000830185613b8f565b6143ec60208301846137ce565b9392505050565b6143fc816136c0565b811461440757600080fd5b50565b600081519050614419816143f3565b92915050565b60006020828403121561443557614434613658565b5b60006144438482850161440a565b91505092915050565b7f636f756c64206e6f74207061796f757420666f72206d696e6572000000000000600082015250565b6000614482601a83613f38565b915061448d8261444c565b602082019050919050565b600060208201905081810360008301526144b181614475565b9050919050565b7f636f756c64206e6f74207061796f757420666f7220636f6d6d756e6974790000600082015250565b60006144ee601e83613f38565b91506144f9826144b8565b602082019050919050565b6000602082019050818103600083015261451d816144e1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061458d8261365d565b91506145988361365d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156145cd576145cc614268565b5b828201905092915050565b60006145e38261365d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361461557614614614268565b5b600182019050919050565b7f696e73756666696369656e742066756e647320746f2077697468647261770000600082015250565b6000614656601e83613f38565b915061466182614620565b602082019050919050565b6000602082019050818103600083015261468581614649565b9050919050565b7f6f6e6c79206d696e6572206f6620626f6f6b696e672063616e20636c61696d00600082015250565b60006146c2601f83613f38565b91506146cd8261468c565b602082019050919050565b600060208201905081810360008301526146f1816146b5565b9050919050565b7f626f6f6b696e67206973207374696c6c20696e20757365000000000000000000600082015250565b600061472e601783613f38565b9150614739826146f8565b602082019050919050565b6000602082019050818103600083015261475d81614721565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f737461626c65636f696e00000000000000000000000000000000000000000000602082015250565b60006147c0602a83613f38565b91506147cb82614764565b604082019050919050565b600060208201905081810360008301526147ef816147b3565b9050919050565b7f6f6e6c79206f776e6572206f6620626f6f6b696e672063616e20657874656e64600082015250565b600061482c602083613f38565b9150614837826147f6565b602082019050919050565b6000602082019050818103600083015261485b8161481f565b9050919050565b7f626f6f6b696e6720697320657870697265642c2069742063616e206265206f6e60008201527f6c7920636c61696d6564206279206d696e657200000000000000000000000000602082015250565b60006148be603383613f38565b91506148c982614862565b604082019050919050565b600060208201905081810360008301526148ed816148b1565b9050919050565b7f696e73756666696369656e742066756e647320746f20657874656e6420626f6f60008201527f6b696e6700000000000000000000000000000000000000000000000000000000602082015250565b6000614950602483613f38565b915061495b826148f4565b604082019050919050565b6000602082019050818103600083015261497f81614943565b9050919050565b600060408201905061499b60008301856137ce565b6149a860208301846137ce565b9392505050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f6e657720636f6d6d756e69747920636f6e747261637400000000000000000000602082015250565b6000614a0b603683613f38565b9150614a16826149af565b604082019050919050565b60006020820190508181036000830152614a3a816149fe565b9050919050565b6000606082019050614a566000830186613b8f565b614a636020830185613b8f565b614a7060408301846137ce565b949350505050565b7f4e6f206d616368696e657320617661696c61626c65204f52207365637320736860008201527f6f756c6420626520706f736974697665204f52207573657220646f6573206e6f60208201527f74206861766520726571756972656420616d6f756e7420746f20626f6f6b206d60408201527f616368696e650000000000000000000000000000000000000000000000000000606082015250565b6000614b20606683613f38565b9150614b2b82614a78565b608082019050919050565b60006020820190508181036000830152614b4f81614b13565b9050919050565b7f6f666665722077697468207468617420696e64657820646f6573206e6f74206560008201527f7869737400000000000000000000000000000000000000000000000000000000602082015250565b6000614bb2602483613f38565b9150614bbd82614b56565b604082019050919050565b60006020820190508181036000830152614be181614ba5565b9050919050565b600081519050614bf781613845565b92915050565b600060208284031215614c1357614c12613658565b5b6000614c2184828501614be8565b9150509291505056fea26469706673582212202dedfc11c24b283c24c3fd39c45afcceeb08fb3d59d41b1931684057dbe51ed264736f6c634300080f0033",
}

// BrokerABI is the input ABI used to generate the binding from.
// Deprecated: Use BrokerMetaData.ABI instead.
var BrokerABI = BrokerMetaData.ABI

// BrokerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrokerMetaData.Bin instead.
var BrokerBin = BrokerMetaData.Bin

// DeployBroker deploys a new Ethereum contract, binding an instance of Broker to it.
func DeployBroker(auth *bind.TransactOpts, backend bind.ContractBackend, communityAddress common.Address) (common.Address, *types.Transaction, *Broker, error) {
	parsed, err := BrokerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrokerBin), backend, communityAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// Broker is an auto generated Go binding around an Ethereum contract.
type Broker struct {
	BrokerCaller     // Read-only binding to the contract
	BrokerTransactor // Write-only binding to the contract
	BrokerFilterer   // Log filterer for contract events
}

// BrokerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokerSession struct {
	Contract     *Broker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokerCallerSession struct {
	Contract *BrokerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrokerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokerTransactorSession struct {
	Contract     *BrokerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokerRaw struct {
	Contract *Broker // Generic contract binding to access the raw methods on
}

// BrokerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokerCallerRaw struct {
	Contract *BrokerCaller // Generic read-only contract binding to access the raw methods on
}

// BrokerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokerTransactorRaw struct {
	Contract *BrokerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBroker creates a new instance of Broker, bound to a specific deployed contract.
func NewBroker(address common.Address, backend bind.ContractBackend) (*Broker, error) {
	contract, err := bindBroker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// NewBrokerCaller creates a new read-only instance of Broker, bound to a specific deployed contract.
func NewBrokerCaller(address common.Address, caller bind.ContractCaller) (*BrokerCaller, error) {
	contract, err := bindBroker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerCaller{contract: contract}, nil
}

// NewBrokerTransactor creates a new write-only instance of Broker, bound to a specific deployed contract.
func NewBrokerTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokerTransactor, error) {
	contract, err := bindBroker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerTransactor{contract: contract}, nil
}

// NewBrokerFilterer creates a new log filterer instance of Broker, bound to a specific deployed contract.
func NewBrokerFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokerFilterer, error) {
	contract, err := bindBroker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokerFilterer{contract: contract}, nil
}

// bindBroker binds a generic wrapper to an already deployed contract.
func bindBroker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.BrokerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transact(opts, method, params...)
}

// GetTime is a free data retrieval call binding the contract method 0x18c5e502.
//
// Solidity: function GetTime() view returns(uint256)
func (_Broker *BrokerCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTime is a free data retrieval call binding the contract method 0x18c5e502.
//
// Solidity: function GetTime() view returns(uint256)
func (_Broker *BrokerSession) GetTime() (*big.Int, error) {
	return _Broker.Contract.GetTime(&_Broker.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x18c5e502.
//
// Solidity: function GetTime() view returns(uint256)
func (_Broker *BrokerCallerSession) GetTime() (*big.Int, error) {
	return _Broker.Contract.GetTime(&_Broker.CallOpts)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCaller) FindBookingsByMiner(opts *bind.CallOpts, _miner common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "findBookingsByMiner", _miner)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCaller) FindBookingsByUser(opts *bind.CallOpts, _owner common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "findBookingsByUser", _owner)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// GetAvailableOffers is a free data retrieval call binding the contract method 0x7ee080c6.
//
// Solidity: function getAvailableOffers(uint256 _vmTypeId) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCaller) GetAvailableOffers(opts *bind.CallOpts, _vmTypeId *big.Int) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getAvailableOffers", _vmTypeId)

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetAvailableOffers is a free data retrieval call binding the contract method 0x7ee080c6.
//
// Solidity: function getAvailableOffers(uint256 _vmTypeId) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerSession) GetAvailableOffers(_vmTypeId *big.Int) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffers(&_Broker.CallOpts, _vmTypeId)
}

// GetAvailableOffers is a free data retrieval call binding the contract method 0x7ee080c6.
//
// Solidity: function getAvailableOffers(uint256 _vmTypeId) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetAvailableOffers(_vmTypeId *big.Int) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffers(&_Broker.CallOpts, _vmTypeId)
}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256,uint256) booking)
func (_Broker *BrokerCaller) GetBooking(opts *bind.CallOpts, index uint64) (BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getBooking", index)

	if err != nil {
		return *new(BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new(BrokerBooking)).(*BrokerBooking)

	return out0, err

}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256,uint256) booking)
func (_Broker *BrokerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256,uint256) booking)
func (_Broker *BrokerCallerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetCoinDecimals is a free data retrieval call binding the contract method 0xcf696d14.
//
// Solidity: function getCoinDecimals() view returns(uint8)
func (_Broker *BrokerCaller) GetCoinDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCoinDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCoinDecimals is a free data retrieval call binding the contract method 0xcf696d14.
//
// Solidity: function getCoinDecimals() view returns(uint8)
func (_Broker *BrokerSession) GetCoinDecimals() (uint8, error) {
	return _Broker.Contract.GetCoinDecimals(&_Broker.CallOpts)
}

// GetCoinDecimals is a free data retrieval call binding the contract method 0xcf696d14.
//
// Solidity: function getCoinDecimals() view returns(uint8)
func (_Broker *BrokerCallerSession) GetCoinDecimals() (uint8, error) {
	return _Broker.Contract.GetCoinDecimals(&_Broker.CallOpts)
}

// GetCommunityContract is a free data retrieval call binding the contract method 0x569e2e0f.
//
// Solidity: function getCommunityContract() view returns(address)
func (_Broker *BrokerCaller) GetCommunityContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCommunityContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCommunityContract is a free data retrieval call binding the contract method 0x569e2e0f.
//
// Solidity: function getCommunityContract() view returns(address)
func (_Broker *BrokerSession) GetCommunityContract() (common.Address, error) {
	return _Broker.Contract.GetCommunityContract(&_Broker.CallOpts)
}

// GetCommunityContract is a free data retrieval call binding the contract method 0x569e2e0f.
//
// Solidity: function getCommunityContract() view returns(address)
func (_Broker *BrokerCallerSession) GetCommunityContract() (common.Address, error) {
	return _Broker.Contract.GetCommunityContract(&_Broker.CallOpts)
}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint256)
func (_Broker *BrokerCaller) GetCommunityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCommunityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint256)
func (_Broker *BrokerSession) GetCommunityFee() (*big.Int, error) {
	return _Broker.Contract.GetCommunityFee(&_Broker.CallOpts)
}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint256)
func (_Broker *BrokerCallerSession) GetCommunityFee() (*big.Int, error) {
	return _Broker.Contract.GetCommunityFee(&_Broker.CallOpts)
}

// GetMinerUrl is a free data retrieval call binding the contract method 0xab49ec45.
//
// Solidity: function getMinerUrl(address _user) view returns(bytes32)
func (_Broker *BrokerCaller) GetMinerUrl(opts *bind.CallOpts, _user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getMinerUrl", _user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMinerUrl is a free data retrieval call binding the contract method 0xab49ec45.
//
// Solidity: function getMinerUrl(address _user) view returns(bytes32)
func (_Broker *BrokerSession) GetMinerUrl(_user common.Address) ([32]byte, error) {
	return _Broker.Contract.GetMinerUrl(&_Broker.CallOpts, _user)
}

// GetMinerUrl is a free data retrieval call binding the contract method 0xab49ec45.
//
// Solidity: function getMinerUrl(address _user) view returns(bytes32)
func (_Broker *BrokerCallerSession) GetMinerUrl(_user common.Address) ([32]byte, error) {
	return _Broker.Contract.GetMinerUrl(&_Broker.CallOpts, _user)
}

// GetMinersOffers is a free data retrieval call binding the contract method 0xc700ff0a.
//
// Solidity: function getMinersOffers(address miner) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCaller) GetMinersOffers(opts *bind.CallOpts, miner common.Address) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getMinersOffers", miner)

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetMinersOffers is a free data retrieval call binding the contract method 0xc700ff0a.
//
// Solidity: function getMinersOffers(address miner) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerSession) GetMinersOffers(miner common.Address) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetMinersOffers(&_Broker.CallOpts, miner)
}

// GetMinersOffers is a free data retrieval call binding the contract method 0xc700ff0a.
//
// Solidity: function getMinersOffers(address miner) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetMinersOffers(miner common.Address) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetMinersOffers(&_Broker.CallOpts, miner)
}

// GetMtlsHash is a free data retrieval call binding the contract method 0xea5f3984.
//
// Solidity: function getMtlsHash(address _user) view returns(bytes20)
func (_Broker *BrokerCaller) GetMtlsHash(opts *bind.CallOpts, _user common.Address) ([20]byte, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getMtlsHash", _user)

	if err != nil {
		return *new([20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([20]byte)).(*[20]byte)

	return out0, err

}

// GetMtlsHash is a free data retrieval call binding the contract method 0xea5f3984.
//
// Solidity: function getMtlsHash(address _user) view returns(bytes20)
func (_Broker *BrokerSession) GetMtlsHash(_user common.Address) ([20]byte, error) {
	return _Broker.Contract.GetMtlsHash(&_Broker.CallOpts, _user)
}

// GetMtlsHash is a free data retrieval call binding the contract method 0xea5f3984.
//
// Solidity: function getMtlsHash(address _user) view returns(bytes20)
func (_Broker *BrokerCallerSession) GetMtlsHash(_user common.Address) ([20]byte, error) {
	return _Broker.Contract.GetMtlsHash(&_Broker.CallOpts, _user)
}

// GetStablecoinAddress is a free data retrieval call binding the contract method 0x7bef4f54.
//
// Solidity: function getStablecoinAddress() view returns(address)
func (_Broker *BrokerCaller) GetStablecoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getStablecoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStablecoinAddress is a free data retrieval call binding the contract method 0x7bef4f54.
//
// Solidity: function getStablecoinAddress() view returns(address)
func (_Broker *BrokerSession) GetStablecoinAddress() (common.Address, error) {
	return _Broker.Contract.GetStablecoinAddress(&_Broker.CallOpts)
}

// GetStablecoinAddress is a free data retrieval call binding the contract method 0x7bef4f54.
//
// Solidity: function getStablecoinAddress() view returns(address)
func (_Broker *BrokerCallerSession) GetStablecoinAddress() (common.Address, error) {
	return _Broker.Contract.GetStablecoinAddress(&_Broker.CallOpts)
}

// GetUsersBookings is a free data retrieval call binding the contract method 0x626bbb27.
//
// Solidity: function getUsersBookings(address user) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCaller) GetUsersBookings(opts *bind.CallOpts, user common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getUsersBookings", user)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// GetUsersBookings is a free data retrieval call binding the contract method 0x626bbb27.
//
// Solidity: function getUsersBookings(address user) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// GetUsersBookings is a free data retrieval call binding the contract method 0x626bbb27.
//
// Solidity: function getUsersBookings(address user) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// UserAllowance is a free data retrieval call binding the contract method 0x20255c7e.
//
// Solidity: function userAllowance() view returns(uint256)
func (_Broker *BrokerCaller) UserAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAllowance is a free data retrieval call binding the contract method 0x20255c7e.
//
// Solidity: function userAllowance() view returns(uint256)
func (_Broker *BrokerSession) UserAllowance() (*big.Int, error) {
	return _Broker.Contract.UserAllowance(&_Broker.CallOpts)
}

// UserAllowance is a free data retrieval call binding the contract method 0x20255c7e.
//
// Solidity: function userAllowance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserAllowance() (*big.Int, error) {
	return _Broker.Contract.UserAllowance(&_Broker.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() view returns(uint256)
func (_Broker *BrokerCaller) UserBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() view returns(uint256)
func (_Broker *BrokerSession) UserBalance() (*big.Int, error) {
	return _Broker.Contract.UserBalance(&_Broker.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserBalance() (*big.Int, error) {
	return _Broker.Contract.UserBalance(&_Broker.CallOpts)
}

// UserDeposit is a free data retrieval call binding the contract method 0x48146113.
//
// Solidity: function userDeposit() view returns(uint256)
func (_Broker *BrokerCaller) UserDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDeposit is a free data retrieval call binding the contract method 0x48146113.
//
// Solidity: function userDeposit() view returns(uint256)
func (_Broker *BrokerSession) UserDeposit() (*big.Int, error) {
	return _Broker.Contract.UserDeposit(&_Broker.CallOpts)
}

// UserDeposit is a free data retrieval call binding the contract method 0x48146113.
//
// Solidity: function userDeposit() view returns(uint256)
func (_Broker *BrokerCallerSession) UserDeposit() (*big.Int, error) {
	return _Broker.Contract.UserDeposit(&_Broker.CallOpts)
}

// UserLockedBalance is a free data retrieval call binding the contract method 0xbb0090c9.
//
// Solidity: function userLockedBalance() view returns(uint256)
func (_Broker *BrokerCaller) UserLockedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userLockedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLockedBalance is a free data retrieval call binding the contract method 0xbb0090c9.
//
// Solidity: function userLockedBalance() view returns(uint256)
func (_Broker *BrokerSession) UserLockedBalance() (*big.Int, error) {
	return _Broker.Contract.UserLockedBalance(&_Broker.CallOpts)
}

// UserLockedBalance is a free data retrieval call binding the contract method 0xbb0090c9.
//
// Solidity: function userLockedBalance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserLockedBalance() (*big.Int, error) {
	return _Broker.Contract.UserLockedBalance(&_Broker.CallOpts)
}

// UserTokenBalance is a free data retrieval call binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() view returns(uint256)
func (_Broker *BrokerCaller) UserTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTokenBalance is a free data retrieval call binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() view returns(uint256)
func (_Broker *BrokerSession) UserTokenBalance() (*big.Int, error) {
	return _Broker.Contract.UserTokenBalance(&_Broker.CallOpts)
}

// UserTokenBalance is a free data retrieval call binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserTokenBalance() (*big.Int, error) {
	return _Broker.Contract.UserTokenBalance(&_Broker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Broker *BrokerCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Broker *BrokerSession) Version() (*big.Int, error) {
	return _Broker.Contract.Version(&_Broker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Broker *BrokerCallerSession) Version() (*big.Int, error) {
	return _Broker.Contract.Version(&_Broker.CallOpts)
}

// SetMtlsHash is a paid mutator transaction binding the contract method 0x117298e4.
//
// Solidity: function SetMtlsHash(bytes20 _signature) returns()
func (_Broker *BrokerTransactor) SetMtlsHash(opts *bind.TransactOpts, _signature [20]byte) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "SetMtlsHash", _signature)
}

// SetMtlsHash is a paid mutator transaction binding the contract method 0x117298e4.
//
// Solidity: function SetMtlsHash(bytes20 _signature) returns()
func (_Broker *BrokerSession) SetMtlsHash(_signature [20]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMtlsHash(&_Broker.TransactOpts, _signature)
}

// SetMtlsHash is a paid mutator transaction binding the contract method 0x117298e4.
//
// Solidity: function SetMtlsHash(bytes20 _signature) returns()
func (_Broker *BrokerTransactorSession) SetMtlsHash(_signature [20]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMtlsHash(&_Broker.TransactOpts, _signature)
}

// AbortBooking is a paid mutator transaction binding the contract method 0x20e56d93.
//
// Solidity: function abortBooking(uint64 index, uint8 abortType) returns()
func (_Broker *BrokerTransactor) AbortBooking(opts *bind.TransactOpts, index uint64, abortType uint8) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "abortBooking", index, abortType)
}

// AbortBooking is a paid mutator transaction binding the contract method 0x20e56d93.
//
// Solidity: function abortBooking(uint64 index, uint8 abortType) returns()
func (_Broker *BrokerSession) AbortBooking(index uint64, abortType uint8) (*types.Transaction, error) {
	return _Broker.Contract.AbortBooking(&_Broker.TransactOpts, index, abortType)
}

// AbortBooking is a paid mutator transaction binding the contract method 0x20e56d93.
//
// Solidity: function abortBooking(uint64 index, uint8 abortType) returns()
func (_Broker *BrokerTransactorSession) AbortBooking(index uint64, abortType uint8) (*types.Transaction, error) {
	return _Broker.Contract.AbortBooking(&_Broker.TransactOpts, index, abortType)
}

// AddOffer is a paid mutator transaction binding the contract method 0x3b3e328b.
//
// Solidity: function addOffer(uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns(uint256)
func (_Broker *BrokerTransactor) AddOffer(opts *bind.TransactOpts, pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "addOffer", pricePerSecond, vmTypeId, machinesAvailable)
}

// AddOffer is a paid mutator transaction binding the contract method 0x3b3e328b.
//
// Solidity: function addOffer(uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns(uint256)
func (_Broker *BrokerSession) AddOffer(pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.AddOffer(&_Broker.TransactOpts, pricePerSecond, vmTypeId, machinesAvailable)
}

// AddOffer is a paid mutator transaction binding the contract method 0x3b3e328b.
//
// Solidity: function addOffer(uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns(uint256)
func (_Broker *BrokerTransactorSession) AddOffer(pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.AddOffer(&_Broker.TransactOpts, pricePerSecond, vmTypeId, machinesAvailable)
}

// BookVM is a paid mutator transaction binding the contract method 0xcbb6bfbd.
//
// Solidity: function bookVM(uint256 offerIndex, uint256 secs) returns(uint256)
func (_Broker *BrokerTransactor) BookVM(opts *bind.TransactOpts, offerIndex *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "bookVM", offerIndex, secs)
}

// BookVM is a paid mutator transaction binding the contract method 0xcbb6bfbd.
//
// Solidity: function bookVM(uint256 offerIndex, uint256 secs) returns(uint256)
func (_Broker *BrokerSession) BookVM(offerIndex *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.BookVM(&_Broker.TransactOpts, offerIndex, secs)
}

// BookVM is a paid mutator transaction binding the contract method 0xcbb6bfbd.
//
// Solidity: function bookVM(uint256 offerIndex, uint256 secs) returns(uint256)
func (_Broker *BrokerTransactorSession) BookVM(offerIndex *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.BookVM(&_Broker.TransactOpts, offerIndex, secs)
}

// ClaimExpired is a paid mutator transaction binding the contract method 0x62c29617.
//
// Solidity: function claimExpired(uint64 index) returns()
func (_Broker *BrokerTransactor) ClaimExpired(opts *bind.TransactOpts, index uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "claimExpired", index)
}

// ClaimExpired is a paid mutator transaction binding the contract method 0x62c29617.
//
// Solidity: function claimExpired(uint64 index) returns()
func (_Broker *BrokerSession) ClaimExpired(index uint64) (*types.Transaction, error) {
	return _Broker.Contract.ClaimExpired(&_Broker.TransactOpts, index)
}

// ClaimExpired is a paid mutator transaction binding the contract method 0x62c29617.
//
// Solidity: function claimExpired(uint64 index) returns()
func (_Broker *BrokerTransactorSession) ClaimExpired(index uint64) (*types.Transaction, error) {
	return _Broker.Contract.ClaimExpired(&_Broker.TransactOpts, index)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xca3b3b4b.
//
// Solidity: function depositCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactor) DepositCoin(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "depositCoin", numTokens)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xca3b3b4b.
//
// Solidity: function depositCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerSession) DepositCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.DepositCoin(&_Broker.TransactOpts, numTokens)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xca3b3b4b.
//
// Solidity: function depositCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactorSession) DepositCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.DepositCoin(&_Broker.TransactOpts, numTokens)
}

// ExtendBooking is a paid mutator transaction binding the contract method 0x9f60c50c.
//
// Solidity: function extendBooking(uint64 index, uint256 secs) returns()
func (_Broker *BrokerTransactor) ExtendBooking(opts *bind.TransactOpts, index uint64, secs *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "extendBooking", index, secs)
}

// ExtendBooking is a paid mutator transaction binding the contract method 0x9f60c50c.
//
// Solidity: function extendBooking(uint64 index, uint256 secs) returns()
func (_Broker *BrokerSession) ExtendBooking(index uint64, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.ExtendBooking(&_Broker.TransactOpts, index, secs)
}

// ExtendBooking is a paid mutator transaction binding the contract method 0x9f60c50c.
//
// Solidity: function extendBooking(uint64 index, uint256 secs) returns()
func (_Broker *BrokerTransactorSession) ExtendBooking(index uint64, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.ExtendBooking(&_Broker.TransactOpts, index, secs)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address newCommunityAddress) returns(bool)
func (_Broker *BrokerTransactor) SetCommunityContract(opts *bind.TransactOpts, newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityContract", newCommunityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address newCommunityAddress) returns(bool)
func (_Broker *BrokerSession) SetCommunityContract(newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, newCommunityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address newCommunityAddress) returns(bool)
func (_Broker *BrokerTransactorSession) SetCommunityContract(newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, newCommunityAddress)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns(bool)
func (_Broker *BrokerTransactor) SetCommunityFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityFee", fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns(bool)
func (_Broker *BrokerSession) SetCommunityFee(fee *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns(bool)
func (_Broker *BrokerTransactorSession) SetCommunityFee(fee *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetMunerUrl is a paid mutator transaction binding the contract method 0xf44f11cd.
//
// Solidity: function setMunerUrl(bytes32 url) returns()
func (_Broker *BrokerTransactor) SetMunerUrl(opts *bind.TransactOpts, url [32]byte) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setMunerUrl", url)
}

// SetMunerUrl is a paid mutator transaction binding the contract method 0xf44f11cd.
//
// Solidity: function setMunerUrl(bytes32 url) returns()
func (_Broker *BrokerSession) SetMunerUrl(url [32]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMunerUrl(&_Broker.TransactOpts, url)
}

// SetMunerUrl is a paid mutator transaction binding the contract method 0xf44f11cd.
//
// Solidity: function setMunerUrl(bytes32 url) returns()
func (_Broker *BrokerTransactorSession) SetMunerUrl(url [32]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMunerUrl(&_Broker.TransactOpts, url)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0x9ecd0500.
//
// Solidity: function setStablecoinAddress(address t) returns(bool)
func (_Broker *BrokerTransactor) SetStablecoinAddress(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setStablecoinAddress", t)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0x9ecd0500.
//
// Solidity: function setStablecoinAddress(address t) returns(bool)
func (_Broker *BrokerSession) SetStablecoinAddress(t common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetStablecoinAddress(&_Broker.TransactOpts, t)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0x9ecd0500.
//
// Solidity: function setStablecoinAddress(address t) returns(bool)
func (_Broker *BrokerTransactorSession) SetStablecoinAddress(t common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetStablecoinAddress(&_Broker.TransactOpts, t)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0x1032c610.
//
// Solidity: function updateOffer(uint256 offerIndex, uint256 vmTypeId, uint256 machinesAvailable) returns()
func (_Broker *BrokerTransactor) UpdateOffer(opts *bind.TransactOpts, offerIndex *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "updateOffer", offerIndex, vmTypeId, machinesAvailable)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0x1032c610.
//
// Solidity: function updateOffer(uint256 offerIndex, uint256 vmTypeId, uint256 machinesAvailable) returns()
func (_Broker *BrokerSession) UpdateOffer(offerIndex *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, vmTypeId, machinesAvailable)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0x1032c610.
//
// Solidity: function updateOffer(uint256 offerIndex, uint256 vmTypeId, uint256 machinesAvailable) returns()
func (_Broker *BrokerTransactorSession) UpdateOffer(offerIndex *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, vmTypeId, machinesAvailable)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x4eb264e8.
//
// Solidity: function withdrawCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactor) WithdrawCoin(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "withdrawCoin", numTokens)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x4eb264e8.
//
// Solidity: function withdrawCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerSession) WithdrawCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.WithdrawCoin(&_Broker.TransactOpts, numTokens)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x4eb264e8.
//
// Solidity: function withdrawCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactorSession) WithdrawCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.WithdrawCoin(&_Broker.TransactOpts, numTokens)
}

// BrokerBookingClaimedIterator is returned from FilterBookingClaimed and is used to iterate over the raw logs and unpacked data for BookingClaimed events raised by the Broker contract.
type BrokerBookingClaimedIterator struct {
	Event *BrokerBookingClaimed // Event containing the contract specifics and raw log

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
func (it *BrokerBookingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingClaimed)
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
		it.Event = new(BrokerBookingClaimed)
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
func (it *BrokerBookingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingClaimed represents a BookingClaimed event raised by the Broker contract.
type BrokerBookingClaimed struct {
	MinerPayout *big.Int
	Index       *big.Int
	Miner       common.Address
	User        common.Address
	TimeUsed    *big.Int
	VmTypeId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingClaimed is a free log retrieval operation binding the contract event 0x75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c.
//
// Solidity: event BookingClaimed(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingClaimed(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingClaimedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingClaimed", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingClaimedIterator{contract: _Broker.contract, event: "BookingClaimed", logs: logs, sub: sub}, nil
}

// WatchBookingClaimed is a free log subscription operation binding the contract event 0x75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c.
//
// Solidity: event BookingClaimed(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingClaimed(opts *bind.WatchOpts, sink chan<- *BrokerBookingClaimed, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingClaimed", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingClaimed)
				if err := _Broker.contract.UnpackLog(event, "BookingClaimed", log); err != nil {
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

// ParseBookingClaimed is a log parse operation binding the contract event 0x75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c.
//
// Solidity: event BookingClaimed(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingClaimed(log types.Log) (*BrokerBookingClaimed, error) {
	event := new(BrokerBookingClaimed)
	if err := _Broker.contract.UnpackLog(event, "BookingClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingExtendedIterator is returned from FilterBookingExtended and is used to iterate over the raw logs and unpacked data for BookingExtended events raised by the Broker contract.
type BrokerBookingExtendedIterator struct {
	Event *BrokerBookingExtended // Event containing the contract specifics and raw log

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
func (it *BrokerBookingExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingExtended)
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
		it.Event = new(BrokerBookingExtended)
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
func (it *BrokerBookingExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingExtended represents a BookingExtended event raised by the Broker contract.
type BrokerBookingExtended struct {
	Index    *big.Int
	Miner    common.Address
	User     common.Address
	VmTypeId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBookingExtended is a free log retrieval operation binding the contract event 0x075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd90298147468.
//
// Solidity: event BookingExtended(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingExtended(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingExtendedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingExtended", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingExtendedIterator{contract: _Broker.contract, event: "BookingExtended", logs: logs, sub: sub}, nil
}

// WatchBookingExtended is a free log subscription operation binding the contract event 0x075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd90298147468.
//
// Solidity: event BookingExtended(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingExtended(opts *bind.WatchOpts, sink chan<- *BrokerBookingExtended, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingExtended", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingExtended)
				if err := _Broker.contract.UnpackLog(event, "BookingExtended", log); err != nil {
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

// ParseBookingExtended is a log parse operation binding the contract event 0x075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd90298147468.
//
// Solidity: event BookingExtended(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingExtended(log types.Log) (*BrokerBookingExtended, error) {
	event := new(BrokerBookingExtended)
	if err := _Broker.contract.UnpackLog(event, "BookingExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingReportedIterator is returned from FilterBookingReported and is used to iterate over the raw logs and unpacked data for BookingReported events raised by the Broker contract.
type BrokerBookingReportedIterator struct {
	Event *BrokerBookingReported // Event containing the contract specifics and raw log

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
func (it *BrokerBookingReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingReported)
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
		it.Event = new(BrokerBookingReported)
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
func (it *BrokerBookingReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingReported represents a BookingReported event raised by the Broker contract.
type BrokerBookingReported struct {
	MinerPayout *big.Int
	Index       *big.Int
	Miner       common.Address
	User        common.Address
	TimeUsed    *big.Int
	VmTypeId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingReported is a free log retrieval operation binding the contract event 0x26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41.
//
// Solidity: event BookingReported(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingReported(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingReportedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingReported", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingReportedIterator{contract: _Broker.contract, event: "BookingReported", logs: logs, sub: sub}, nil
}

// WatchBookingReported is a free log subscription operation binding the contract event 0x26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41.
//
// Solidity: event BookingReported(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingReported(opts *bind.WatchOpts, sink chan<- *BrokerBookingReported, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingReported", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingReported)
				if err := _Broker.contract.UnpackLog(event, "BookingReported", log); err != nil {
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

// ParseBookingReported is a log parse operation binding the contract event 0x26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41.
//
// Solidity: event BookingReported(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingReported(log types.Log) (*BrokerBookingReported, error) {
	event := new(BrokerBookingReported)
	if err := _Broker.contract.UnpackLog(event, "BookingReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingStartedIterator is returned from FilterBookingStarted and is used to iterate over the raw logs and unpacked data for BookingStarted events raised by the Broker contract.
type BrokerBookingStartedIterator struct {
	Event *BrokerBookingStarted // Event containing the contract specifics and raw log

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
func (it *BrokerBookingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingStarted)
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
		it.Event = new(BrokerBookingStarted)
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
func (it *BrokerBookingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingStarted represents a BookingStarted event raised by the Broker contract.
type BrokerBookingStarted struct {
	Index    *big.Int
	Miner    common.Address
	User     common.Address
	VmTypeId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBookingStarted is a free log retrieval operation binding the contract event 0x246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6.
//
// Solidity: event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingStarted(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingStartedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingStarted", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingStartedIterator{contract: _Broker.contract, event: "BookingStarted", logs: logs, sub: sub}, nil
}

// WatchBookingStarted is a free log subscription operation binding the contract event 0x246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6.
//
// Solidity: event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingStarted(opts *bind.WatchOpts, sink chan<- *BrokerBookingStarted, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingStarted", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingStarted)
				if err := _Broker.contract.UnpackLog(event, "BookingStarted", log); err != nil {
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

// ParseBookingStarted is a log parse operation binding the contract event 0x246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6.
//
// Solidity: event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingStarted(log types.Log) (*BrokerBookingStarted, error) {
	event := new(BrokerBookingStarted)
	if err := _Broker.contract.UnpackLog(event, "BookingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingStoppedIterator is returned from FilterBookingStopped and is used to iterate over the raw logs and unpacked data for BookingStopped events raised by the Broker contract.
type BrokerBookingStoppedIterator struct {
	Event *BrokerBookingStopped // Event containing the contract specifics and raw log

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
func (it *BrokerBookingStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingStopped)
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
		it.Event = new(BrokerBookingStopped)
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
func (it *BrokerBookingStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingStopped represents a BookingStopped event raised by the Broker contract.
type BrokerBookingStopped struct {
	MinerPayout *big.Int
	Index       *big.Int
	Miner       common.Address
	User        common.Address
	TimeUsed    *big.Int
	VmTypeId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingStopped is a free log retrieval operation binding the contract event 0x232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1.
//
// Solidity: event BookingStopped(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingStopped(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingStoppedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingStopped", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingStoppedIterator{contract: _Broker.contract, event: "BookingStopped", logs: logs, sub: sub}, nil
}

// WatchBookingStopped is a free log subscription operation binding the contract event 0x232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1.
//
// Solidity: event BookingStopped(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingStopped(opts *bind.WatchOpts, sink chan<- *BrokerBookingStopped, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingStopped", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingStopped)
				if err := _Broker.contract.UnpackLog(event, "BookingStopped", log); err != nil {
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

// ParseBookingStopped is a log parse operation binding the contract event 0x232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1.
//
// Solidity: event BookingStopped(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingStopped(log types.Log) (*BrokerBookingStopped, error) {
	event := new(BrokerBookingStopped)
	if err := _Broker.contract.UnpackLog(event, "BookingStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
