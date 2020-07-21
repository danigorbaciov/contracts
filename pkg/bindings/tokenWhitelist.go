// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenWhitelistABI is the input ABI used to generate the binding from.
<<<<<<< HEAD
const TokenWhitelistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_oracleNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_stablecoinAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ensRegistry\",\"type\":\"address\"}],\"name\":\"ENSSetRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenLoadable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenRedeemable\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"_redeemable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getERC20RecipientAndAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"redeemableCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"redeemableTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"setTokenLoadable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"setTokenRedeemable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
||||||| parent of d04e1094... Upgrade to solc 0.6.11
const TokenWhitelistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_oracleNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_stablecoinAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenLoadable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenRedeemable\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"_redeemable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getERC20RecipientAndAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"redeemableCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"redeemableTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"setTokenLoadable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"setTokenRedeemable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
=======
const TokenWhitelistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ens_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_oracleNode_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_controllerNode_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_stablecoinAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"AddedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedExclusiveMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"RemovedMethodId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenLoadable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"UpdatedTokenRedeemable\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"_redeemable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getERC20RecipientAndAmount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"isERC20MethodWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemableCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemableTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_loadable\",\"type\":\"bool\"}],\"name\":\"setTokenLoadable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_redeemable\",\"type\":\"bool\"}],\"name\":\"setTokenRedeemable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
>>>>>>> d04e1094... Upgrade to solc 0.6.11

// TokenWhitelistBin is the compiled bytecode used for deploying new contracts.
<<<<<<< HEAD
<<<<<<< HEAD
var TokenWhitelistBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034553480156200005657600080fd5b506040516200272c3803806200272c833981810160405260808110156200007c57600080fd5b5080516020820151604083015160609093015191929091620000a7846001600160e01b03620001a616565b620000bb826001600160e01b036200028616565b603a9290925550603980546001600160a01b0319166001600160a01b039092169190911790555060366020527f36aaa0d2a881446a5d284caf55fe779596f3dd15619675545d6094c8909139c2805460ff1990811660019081179092557fbdd45af1473e0dd27b6e035ccafb153373c61968247ba4b7c0e6c1c96496792f80548216831790557f877171c55a6bbc4aa0ea34a24de07dd3c5952b4a7d66df2b953f79c7955a122780548216831790556323b872dd60e01b6000527f04324ec27bf37ba85969616ff61c6c0e3939d6b141697b749b0be22765b29870805490911690911790556200034b565b600054610100900460ff1680620001cb5750620001cb6001600160e01b036200034516565b80620001da575060005460ff16155b620002175760405162461bcd60e51b815260040180806020018281038252602e815260200180620026fe602e913960400191505060405180910390fd5b600054610100900460ff1615801562000243576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200026f57603380546001600160a01b0319166001600160a01b0384161790555b801562000282576000805461ff00191690555b5050565b600054610100900460ff1680620002ab5750620002ab6001600160e01b036200034516565b80620002ba575060005460ff16155b620002f75760405162461bcd60e51b815260040180806020018281038252602e815260200180620026fe602e913960400191505060405180910390fd5b600054610100900460ff1615801562000323576000805460ff1961ff0019909116610100171660011790555b81156200026f576034829055801562000282576000805461ff00191690555050565b303b1590565b6123a3806200035b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80636c3824ef116100a2578063996cba6811610071578063996cba6814610584578063afc72e93146105ba578063d545782e1461065b578063e2b4ce971461068d578063e9cbd8221461069557610116565b80636c3824ef146104bc5780637d73b2311461052a57806380cc0dda1461054e578063872773061461055657610116565b80633efec5e9116100e95780633efec5e9146103f0578063443dd2a4146103f857806344b049bc146104505780635d793a7d146104585780636a1744dc1461048657610116565b806313d5e8461461011b5780631d3a069f146101355780631f69565f1461017057806334c73edc14610240575b600080fd5b61012361069d565b60408051918252519081900360200190f35b61015c6004803603602081101561014b57600080fd5b50356001600160e01b0319166106a4565b604080519115158252519081900360200190f35b6101966004803603602081101561018657600080fd5b50356001600160a01b03166106c7565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b838110156101ff5781810151838201526020016101e7565b50505050905090810190601f16801561022c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103ee600480360360c081101561025657600080fd5b810190602081018135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460208302840111600160201b831117156102a357600080fd5b919390929091602081019035600160201b8111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460208302840111600160201b831117156102f357600080fd5b919390929091602081019035600160201b81111561031057600080fd5b82018360208201111561032257600080fd5b803590602001918460208302840111600160201b8311171561034357600080fd5b919390929091602081019035600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460208302840111600160201b8311171561039357600080fd5b919390929091602081019035600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460208302840111600160201b831117156103e357600080fd5b9193509150356107c5565b005b610196610c57565b610400610d55565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561043c578181015183820152602001610424565b505050509050019250505060405180910390f35b610400610db7565b6103ee6004803603604081101561046e57600080fd5b506001600160a01b0381351690602001351515610e85565b61015c6004803603604081101561049c57600080fd5b5080356001600160a01b031690602001356001600160e01b03191661107b565b6103ee600480360360208110156104d257600080fd5b810190602081018135600160201b8111156104ec57600080fd5b8201836020820111156104fe57600080fd5b803590602001918460208302840111600160201b8311171561051f57600080fd5b509092509050611102565b6105326113a5565b604080516001600160a01b039092168252519081900360200190f35b6101236113b4565b6103ee6004803603604081101561056c57600080fd5b506001600160a01b03813516906020013515156113ba565b6103ee6004803603606081101561059a57600080fd5b506001600160a01b038135811691602081013590911690604001356114f6565b610638600480360360408110156105d057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b5090925090506115a3565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103ee6004803603606081101561067157600080fd5b506001600160a01b0381351690602081013590604001356118e7565b610123611a31565b610532611a37565b6038545b90565b6001600160e01b0319811660009081526036602052604090205460ff165b919050565b6001600160a01b038116600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff8084169694840481169562010000909404169391929189918301828280156107a45780601f10610779576101008083540402835291602001916107a4565b820191906000526020600020905b81548152906001019060200180831161078757829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107ce33611a46565b610818576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561082657508986145b801561083157508984145b801561083c57508984145b61088d576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4957603560008d8d848181106108a857fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610923576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606109486109438c8c8581811061093757fe5b90506020020135611ada565b611b06565b90506040518060e001604052808281526020018a8a8581811061096757fe5b9050602002013581526020016000815260200160011515815260200188888581811061098f57fe5b905060200201351515151581526020018686858181106109ab57fe5b9050602002013515151515815260200184815250603560008f8f868181106109cf57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000019080519060200190610a1b929190612202565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560378d8d84818110610a9257fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b039590920293909301359390931692909217905550848483818110610ada57fe5b9050602002013515610afe57603854610afa90600163ffffffff611b5d16565b6038555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2c57fe5b905060200201356001600160a01b0316838c8c87818110610b4957fe5b905060200201358b8b88818110610b5c57fe5b9050602002013515158a8a89818110610b7157fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610c01578181015183820152602001610be9565b50505050905090810190601f168015610c2e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610890565b505050505050505050505050565b6039546001600160a01b0316600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff808416969484048116956201000090940416939192918991830182828015610d365780601f10610d0b57610100808354040283529160200191610d36565b820191906000526020600020905b815481529060010190602001808311610d1957829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606037805480602002602001604051908101604052809291908181526020018280548015610dad57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d8f575b5050505050905090565b606080603854604051908082528060200260200182016040528015610de6578160200160208202803883390190505b5090506000805b603754811015610e7d57600060378281548110610e0657fe5b60009182526020808320909101546001600160a01b0316808352603590915260409091206003015490915060ff620100009091041615610e745780848481518110610e4d57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610ded565b509091505090565b610e8e33611a46565b610ed8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff16610f48576040805162461bcd60e51b815260206004820152601f60248201527f72656465656d61626c653a20746f6b656e206e6f7420617661696c61626c6500604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff620100009091041615158115151415610fc6576040805162461bcd60e51b815260206004820152601b60248201527f72656465656d61626c653a206e6f207374617465206368616e67650000000000604482015290519081900360640190fd5b8015610fe857603854610fe090600163ffffffff611b5d16565b603855611000565b603854610ffc90600163ffffffff611bbe16565b6038555b6001600160a01b038216600081815260356020908152604091829020600301805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526035602052604081206003015460ff166110e0576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526036602052604090205460ff16919050565b61110b33611a46565b611155576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156113a057600083838381811061116e57fe5b602090810292909201356001600160a01b0316600081815260359093526040909220600301549192505060ff166111e5576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03811660009081526035602052604090206003015462010000900460ff16156112275760385461122390600163ffffffff611bbe16565b6038555b6001600160a01b03811660009081526035602052604081209061124a8282612280565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b60375461128690600163ffffffff611bbe16565b81101561133f57816001600160a01b0316603782815481106112a457fe5b6000918252602090912001546001600160a01b0316141561133757603780546112d490600163ffffffff611bbe16565b815481106112de57fe5b600091825260209091200154603780546001600160a01b03909216918390811061130457fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061133f565b600101611272565b5060378054906113539060001983016122c7565b50604080513381526001600160a01b038316602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a150600101611158565b505050565b6033546001600160a01b031690565b603a5490565b6113c333611a46565b61140d576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff1661147d576040805162461bcd60e51b815260206004820181905260248201527f6c6f616461626c653a20746f6b656e206973206e6f7420617661696c61626c65604482015290519081900360640190fd5b6001600160a01b0382166000818152603560209081526040918290206003018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b6114ff33611a46565b611549576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b611554838383611c1b565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b60008060248310156115fc576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b6000611648600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611ce5169050565b9050611654868261107b565b61169a576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b141561170857856116fe600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d64169050565b92509250506118df565b6001600160e01b031981166323b872dd60e01b141561180b576064841015611777576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b6117c1603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d77169050565b6116fe604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d64169050565b604484101561184b5760405162461bcd60e51b81526004018080602001828103825260258152602001806123206025913960400191505060405180910390fd5b611895601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d77169050565b6116fe602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d64169050565b935093915050565b60006118f4603a54611df9565b90506118ff33611a46565b806119125750336001600160a01b038216145b61195c576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03841660009081526035602052604090206003015460ff166119c5576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03841660008181526035602090815260409182902060028101879055600401859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60345490565b6039546001600160a01b031690565b6000611a53603454611df9565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611aa857600080fd5b505afa158015611abc573d6000803e3d6000fd5b505050506040513d6020811015611ad257600080fd5b505192915050565b611ae26122eb565b6040516020810160405282815280602083015250611aff82611f1a565b8152919050565b60608082600001516040519080825280601f01601f191660200182016040528015611b38576020820181803883390190505b5090506000602082019050611b568185602001518660000151611fb3565b5092915050565b600082820183811015611bb7576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600082821115611c15576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b038216611ccb576040516000906001600160a01b0385169083908381818185875af1925050503d8060008114611c74576040519150601f19603f3d011682016040523d82523d6000602084013e611c79565b606091505b5050905080611cc5576040805162461bcd60e51b81526020600482015260136024820152721cd85999551c985b9cd9995c8819985a5b1959606a1b604482015290519081900360640190fd5b506113a0565b6113a06001600160a01b038316848363ffffffff611ff116565b6000611cf882600463ffffffff611b5d16565b83511015611d44576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d5884602063ffffffff611b5d16565b90940151949350505050565b6000611cf882602063ffffffff611b5d16565b6000611d8a82601463ffffffff611b5d16565b83511015611dd6576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611dea84602063ffffffff611b5d16565b9094015160601c949350505050565b6033546000906001600160a01b0316611e59576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611ea557600080fd5b505afa158015611eb9573d6000803e3d6000fd5b505050506040513d6020811015611ecf57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015611aa857600080fd5b60008082611f2c5760009150506106c2565b6fffffffffffffffffffffffffffffffff8316611f5057601001600160801b830492505b67ffffffffffffffff8316611f715760080168010000000000000000830492505b63ffffffff8316611f8957600401600160201b830492505b61ffff8316611f9e5760020162010000830492505b60ff8316611faa576001015b60200392915050565b5b60208110611fd3578151835260209283019290910190601f1901611fb4565b905182516020929092036101000a6000190180199091169116179052565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526113a0908490612050826001600160a01b03166121fc565b6120a1576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b602083106120df5780518252601f1990920191602091820191016120c0565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612141576040519150601f19603f3d011682016040523d82523d6000602084013e612146565b606091505b50915091508161219d576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b8051156121f6578080602001905160208110156121b957600080fd5b50516121f65760405162461bcd60e51b815260040180806020018281038252602a815260200180612345602a913960400191505060405180910390fd5b50505050565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061224357805160ff1916838001178555612270565b82800160010185558215612270579182015b82811115612270578251825591602001919060010190612255565b5061227c929150612305565b5090565b50805460018160011615610100020316600290046000825580601f106122a657506122c4565b601f0160209004906000526020600020908101906122c49190612305565b50565b8154818355818111156113a0576000838152602090206113a0918101908301612305565b604051806040016040528060008152602001600081525090565b6106a191905b8082111561227c576000815560010161230b56fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f76655361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a723158201875979fbabae5249adb55cd5fdfd0ce5855f6ea91c233c71ad3a738da7b091064736f6c63430005110032436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"
||||||| parent of d04e1094... Upgrade to solc 0.6.11
var TokenWhitelistBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034553480156200005657600080fd5b506040516200272c3803806200272c833981810160405260808110156200007c57600080fd5b5080516020820151604083015160609093015191929091620000a7846001600160e01b03620001a616565b620000bb826001600160e01b036200028616565b603a9290925550603980546001600160a01b0319166001600160a01b039092169190911790555060366020527f36aaa0d2a881446a5d284caf55fe779596f3dd15619675545d6094c8909139c2805460ff1990811660019081179092557fbdd45af1473e0dd27b6e035ccafb153373c61968247ba4b7c0e6c1c96496792f80548216831790557f877171c55a6bbc4aa0ea34a24de07dd3c5952b4a7d66df2b953f79c7955a122780548216831790556323b872dd60e01b6000527f04324ec27bf37ba85969616ff61c6c0e3939d6b141697b749b0be22765b29870805490911690911790556200034b565b600054610100900460ff1680620001cb5750620001cb6001600160e01b036200034516565b80620001da575060005460ff16155b620002175760405162461bcd60e51b815260040180806020018281038252602e815260200180620026fe602e913960400191505060405180910390fd5b600054610100900460ff1615801562000243576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200026f57603380546001600160a01b0319166001600160a01b0384161790555b801562000282576000805461ff00191690555b5050565b600054610100900460ff1680620002ab5750620002ab6001600160e01b036200034516565b80620002ba575060005460ff16155b620002f75760405162461bcd60e51b815260040180806020018281038252602e815260200180620026fe602e913960400191505060405180910390fd5b600054610100900460ff1615801562000323576000805460ff1961ff0019909116610100171660011790555b81156200026f576034829055801562000282576000805461ff00191690555050565b303b1590565b6123a3806200035b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80636c3824ef116100a2578063996cba6811610071578063996cba6814610584578063afc72e93146105ba578063d545782e1461065b578063e2b4ce971461068d578063e9cbd8221461069557610116565b80636c3824ef146104bc5780637d73b2311461052a57806380cc0dda1461054e578063872773061461055657610116565b80633efec5e9116100e95780633efec5e9146103f0578063443dd2a4146103f857806344b049bc146104505780635d793a7d146104585780636a1744dc1461048657610116565b806313d5e8461461011b5780631d3a069f146101355780631f69565f1461017057806334c73edc14610240575b600080fd5b61012361069d565b60408051918252519081900360200190f35b61015c6004803603602081101561014b57600080fd5b50356001600160e01b0319166106a4565b604080519115158252519081900360200190f35b6101966004803603602081101561018657600080fd5b50356001600160a01b03166106c7565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b838110156101ff5781810151838201526020016101e7565b50505050905090810190601f16801561022c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103ee600480360360c081101561025657600080fd5b810190602081018135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460208302840111600160201b831117156102a357600080fd5b919390929091602081019035600160201b8111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460208302840111600160201b831117156102f357600080fd5b919390929091602081019035600160201b81111561031057600080fd5b82018360208201111561032257600080fd5b803590602001918460208302840111600160201b8311171561034357600080fd5b919390929091602081019035600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460208302840111600160201b8311171561039357600080fd5b919390929091602081019035600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460208302840111600160201b831117156103e357600080fd5b9193509150356107c5565b005b610196610c57565b610400610d55565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561043c578181015183820152602001610424565b505050509050019250505060405180910390f35b610400610db7565b6103ee6004803603604081101561046e57600080fd5b506001600160a01b0381351690602001351515610e85565b61015c6004803603604081101561049c57600080fd5b5080356001600160a01b031690602001356001600160e01b03191661107b565b6103ee600480360360208110156104d257600080fd5b810190602081018135600160201b8111156104ec57600080fd5b8201836020820111156104fe57600080fd5b803590602001918460208302840111600160201b8311171561051f57600080fd5b509092509050611102565b6105326113a5565b604080516001600160a01b039092168252519081900360200190f35b6101236113b4565b6103ee6004803603604081101561056c57600080fd5b506001600160a01b03813516906020013515156113ba565b6103ee6004803603606081101561059a57600080fd5b506001600160a01b038135811691602081013590911690604001356114f6565b610638600480360360408110156105d057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b5090925090506115a3565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103ee6004803603606081101561067157600080fd5b506001600160a01b0381351690602081013590604001356118e7565b610123611a31565b610532611a37565b6038545b90565b6001600160e01b0319811660009081526036602052604090205460ff165b919050565b6001600160a01b038116600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff8084169694840481169562010000909404169391929189918301828280156107a45780601f10610779576101008083540402835291602001916107a4565b820191906000526020600020905b81548152906001019060200180831161078757829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107ce33611a46565b610818576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561082657508986145b801561083157508984145b801561083c57508984145b61088d576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4957603560008d8d848181106108a857fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610923576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606109486109438c8c8581811061093757fe5b90506020020135611ada565b611b06565b90506040518060e001604052808281526020018a8a8581811061096757fe5b9050602002013581526020016000815260200160011515815260200188888581811061098f57fe5b905060200201351515151581526020018686858181106109ab57fe5b9050602002013515151515815260200184815250603560008f8f868181106109cf57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000019080519060200190610a1b929190612202565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560378d8d84818110610a9257fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b039590920293909301359390931692909217905550848483818110610ada57fe5b9050602002013515610afe57603854610afa90600163ffffffff611b5d16565b6038555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2c57fe5b905060200201356001600160a01b0316838c8c87818110610b4957fe5b905060200201358b8b88818110610b5c57fe5b9050602002013515158a8a89818110610b7157fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610c01578181015183820152602001610be9565b50505050905090810190601f168015610c2e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610890565b505050505050505050505050565b6039546001600160a01b0316600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff808416969484048116956201000090940416939192918991830182828015610d365780601f10610d0b57610100808354040283529160200191610d36565b820191906000526020600020905b815481529060010190602001808311610d1957829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606037805480602002602001604051908101604052809291908181526020018280548015610dad57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d8f575b5050505050905090565b606080603854604051908082528060200260200182016040528015610de6578160200160208202803883390190505b5090506000805b603754811015610e7d57600060378281548110610e0657fe5b60009182526020808320909101546001600160a01b0316808352603590915260409091206003015490915060ff620100009091041615610e745780848481518110610e4d57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610ded565b509091505090565b610e8e33611a46565b610ed8576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff16610f48576040805162461bcd60e51b815260206004820152601f60248201527f72656465656d61626c653a20746f6b656e206e6f7420617661696c61626c6500604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff620100009091041615158115151415610fc6576040805162461bcd60e51b815260206004820152601b60248201527f72656465656d61626c653a206e6f207374617465206368616e67650000000000604482015290519081900360640190fd5b8015610fe857603854610fe090600163ffffffff611b5d16565b603855611000565b603854610ffc90600163ffffffff611bbe16565b6038555b6001600160a01b038216600081815260356020908152604091829020600301805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526035602052604081206003015460ff166110e0576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526036602052604090205460ff16919050565b61110b33611a46565b611155576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156113a057600083838381811061116e57fe5b602090810292909201356001600160a01b0316600081815260359093526040909220600301549192505060ff166111e5576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03811660009081526035602052604090206003015462010000900460ff16156112275760385461122390600163ffffffff611bbe16565b6038555b6001600160a01b03811660009081526035602052604081209061124a8282612280565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b60375461128690600163ffffffff611bbe16565b81101561133f57816001600160a01b0316603782815481106112a457fe5b6000918252602090912001546001600160a01b0316141561133757603780546112d490600163ffffffff611bbe16565b815481106112de57fe5b600091825260209091200154603780546001600160a01b03909216918390811061130457fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061133f565b600101611272565b5060378054906113539060001983016122c7565b50604080513381526001600160a01b038316602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a150600101611158565b505050565b6033546001600160a01b031690565b603a5490565b6113c333611a46565b61140d576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff1661147d576040805162461bcd60e51b815260206004820181905260248201527f6c6f616461626c653a20746f6b656e206973206e6f7420617661696c61626c65604482015290519081900360640190fd5b6001600160a01b0382166000818152603560209081526040918290206003018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b6114ff33611a46565b611549576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b611554838383611c1b565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b60008060248310156115fc576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b6000611648600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611ce5169050565b9050611654868261107b565b61169a576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b141561170857856116fe600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d64169050565b92509250506118df565b6001600160e01b031981166323b872dd60e01b141561180b576064841015611777576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b6117c1603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d77169050565b6116fe604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d64169050565b604484101561184b5760405162461bcd60e51b81526004018080602001828103825260258152602001806123206025913960400191505060405180910390fd5b611895601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d77169050565b6116fe602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d64169050565b935093915050565b60006118f4603a54611df9565b90506118ff33611a46565b806119125750336001600160a01b038216145b61195c576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03841660009081526035602052604090206003015460ff166119c5576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03841660008181526035602090815260409182902060028101879055600401859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60345490565b6039546001600160a01b031690565b6000611a53603454611df9565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611aa857600080fd5b505afa158015611abc573d6000803e3d6000fd5b505050506040513d6020811015611ad257600080fd5b505192915050565b611ae26122eb565b6040516020810160405282815280602083015250611aff82611f1a565b8152919050565b60608082600001516040519080825280601f01601f191660200182016040528015611b38576020820181803883390190505b5090506000602082019050611b568185602001518660000151611fb3565b5092915050565b600082820183811015611bb7576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600082821115611c15576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6001600160a01b038216611ccb576040516000906001600160a01b0385169083908381818185875af1925050503d8060008114611c74576040519150601f19603f3d011682016040523d82523d6000602084013e611c79565b606091505b5050905080611cc5576040805162461bcd60e51b81526020600482015260136024820152721cd85999551c985b9cd9995c8819985a5b1959606a1b604482015290519081900360640190fd5b506113a0565b6113a06001600160a01b038316848363ffffffff611ff116565b6000611cf882600463ffffffff611b5d16565b83511015611d44576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d5884602063ffffffff611b5d16565b90940151949350505050565b6000611cf882602063ffffffff611b5d16565b6000611d8a82601463ffffffff611b5d16565b83511015611dd6576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611dea84602063ffffffff611b5d16565b9094015160601c949350505050565b6033546000906001600160a01b0316611e59576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611ea557600080fd5b505afa158015611eb9573d6000803e3d6000fd5b505050506040513d6020811015611ecf57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015611aa857600080fd5b60008082611f2c5760009150506106c2565b6fffffffffffffffffffffffffffffffff8316611f5057601001600160801b830492505b67ffffffffffffffff8316611f715760080168010000000000000000830492505b63ffffffff8316611f8957600401600160201b830492505b61ffff8316611f9e5760020162010000830492505b60ff8316611faa576001015b60200392915050565b5b60208110611fd3578151835260209283019290910190601f1901611fb4565b905182516020929092036101000a6000190180199091169116179052565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526113a0908490612050826001600160a01b03166121fc565b6120a1576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b602083106120df5780518252601f1990920191602091820191016120c0565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612141576040519150601f19603f3d011682016040523d82523d6000602084013e612146565b606091505b50915091508161219d576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b8051156121f6578080602001905160208110156121b957600080fd5b50516121f65760405162461bcd60e51b815260040180806020018281038252602a815260200180612345602a913960400191505060405180910390fd5b50505050565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061224357805160ff1916838001178555612270565b82800160010185558215612270579182015b82811115612270578251825591602001919060010190612255565b5061227c929150612305565b5090565b50805460018160011615610100020316600290046000825580601f106122a657506122c4565b601f0160209004906000526020600020908101906122c49190612305565b50565b8154818355818111156113a0576000838152602090206113a0918101908301612305565b604051806040016040528060008152602001600081525090565b6106a191905b8082111561227c576000815560010161230b56fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f76655361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a7231582092b5c5c651728c56367f69e37ffc68e2c78ef720a003b6496ae94f4b393934b264736f6c63430005110032436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"
=======
var TokenWhitelistBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034553480156200005657600080fd5b50604051620028e4380380620028e4833981810160405260808110156200007c57600080fd5b5080516020820151604083015160609093015191929091620000a7846001600160e01b03620001a616565b620000bb826001600160e01b036200028616565b603a9290925550603980546001600160a01b0319166001600160a01b039092169190911790555060366020527f36aaa0d2a881446a5d284caf55fe779596f3dd15619675545d6094c8909139c2805460ff1990811660019081179092557fbdd45af1473e0dd27b6e035ccafb153373c61968247ba4b7c0e6c1c96496792f80548216831790557f877171c55a6bbc4aa0ea34a24de07dd3c5952b4a7d66df2b953f79c7955a122780548216831790556323b872dd60e01b6000527f04324ec27bf37ba85969616ff61c6c0e3939d6b141697b749b0be22765b29870805490911690911790556200034b565b600054610100900460ff1680620001cb5750620001cb6001600160e01b036200034516565b80620001da575060005460ff16155b620002175760405162461bcd60e51b815260040180806020018281038252602e815260200180620028b6602e913960400191505060405180910390fd5b600054610100900460ff1615801562000243576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200026f57603380546001600160a01b0319166001600160a01b0384161790555b801562000282576000805461ff00191690555b5050565b600054610100900460ff1680620002ab5750620002ab6001600160e01b036200034516565b80620002ba575060005460ff16155b620002f75760405162461bcd60e51b815260040180806020018281038252602e815260200180620028b6602e913960400191505060405180910390fd5b600054610100900460ff1615801562000323576000805460ff1961ff0019909116610100171660011790555b81156200026f576034829055801562000282576000805461ff00191690555050565b303b1590565b61255b806200035b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80636c3824ef116100a2578063996cba6811610071578063996cba6814610584578063afc72e93146105ba578063d545782e1461065b578063e2b4ce971461068d578063e9cbd8221461069557610116565b80636c3824ef146104bc5780637d73b2311461052a57806380cc0dda1461054e578063872773061461055657610116565b80633efec5e9116100e95780633efec5e9146103f0578063443dd2a4146103f857806344b049bc146104505780635d793a7d146104585780636a1744dc1461048657610116565b806313d5e8461461011b5780631d3a069f146101355780631f69565f1461017057806334c73edc14610240575b600080fd5b61012361069d565b60408051918252519081900360200190f35b61015c6004803603602081101561014b57600080fd5b50356001600160e01b0319166106a4565b604080519115158252519081900360200190f35b6101966004803603602081101561018657600080fd5b50356001600160a01b03166106c7565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b838110156101ff5781810151838201526020016101e7565b50505050905090810190601f16801561022c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103ee600480360360c081101561025657600080fd5b810190602081018135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460208302840111600160201b831117156102a357600080fd5b919390929091602081019035600160201b8111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460208302840111600160201b831117156102f357600080fd5b919390929091602081019035600160201b81111561031057600080fd5b82018360208201111561032257600080fd5b803590602001918460208302840111600160201b8311171561034357600080fd5b919390929091602081019035600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460208302840111600160201b8311171561039357600080fd5b919390929091602081019035600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460208302840111600160201b831117156103e357600080fd5b9193509150356107c5565b005b610196610c57565b610400610d55565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561043c578181015183820152602001610424565b505050509050019250505060405180910390f35b610400610db7565b6103ee6004803603604081101561046e57600080fd5b506001600160a01b0381351690602001351515610e9c565b61015c6004803603604081101561049c57600080fd5b5080356001600160a01b031690602001356001600160e01b031916611092565b6103ee600480360360208110156104d257600080fd5b810190602081018135600160201b8111156104ec57600080fd5b8201836020820111156104fe57600080fd5b803590602001918460208302840111600160201b8311171561051f57600080fd5b509092509050611119565b6105326113db565b604080516001600160a01b039092168252519081900360200190f35b6101236113ea565b6103ee6004803603604081101561056c57600080fd5b506001600160a01b03813516906020013515156113f0565b6103ee6004803603606081101561059a57600080fd5b506001600160a01b0381358116916020810135909116906040013561152c565b610638600480360360408110156105d057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b5090925090506115d9565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103ee6004803603606081101561067157600080fd5b506001600160a01b03813516906020810135906040013561191d565b610123611a67565b610532611a6d565b6038545b90565b6001600160e01b0319811660009081526036602052604090205460ff165b919050565b6001600160a01b038116600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff8084169694840481169562010000909404169391929189918301828280156107a45780601f10610779576101008083540402835291602001916107a4565b820191906000526020600020905b81548152906001019060200180831161078757829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107ce33611a7c565b610818576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561082657508986145b801561083157508984145b801561083c57508984145b61088d576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4957603560008d8d848181106108a857fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610923576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606109486109438c8c8581811061093757fe5b90506020020135611b10565b611b3c565b90506040518060e001604052808281526020018a8a8581811061096757fe5b9050602002013581526020016000815260200160011515815260200188888581811061098f57fe5b905060200201351515151581526020018686858181106109ab57fe5b9050602002013515151515815260200184815250603560008f8f868181106109cf57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000019080519060200190610a1b9291906123a3565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560378d8d84818110610a9257fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b039590920293909301359390931692909217905550848483818110610ada57fe5b9050602002013515610afe57603854610afa90600163ffffffff611baa16565b6038555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2c57fe5b905060200201356001600160a01b0316838c8c87818110610b4957fe5b905060200201358b8b88818110610b5c57fe5b9050602002013515158a8a89818110610b7157fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610c01578181015183820152602001610be9565b50505050905090810190601f168015610c2e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610890565b505050505050505050505050565b6039546001600160a01b0316600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff808416969484048116956201000090940416939192918991830182828015610d365780601f10610d0b57610100808354040283529160200191610d36565b820191906000526020600020905b815481529060010190602001808311610d1957829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606037805480602002602001604051908101604052809291908181526020018280548015610dad57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d8f575b5050505050905090565b60608060385467ffffffffffffffff81118015610dd357600080fd5b50604051908082528060200260200182016040528015610dfd578160200160208202803683370190505b5090506000805b603754811015610e9457600060378281548110610e1d57fe5b60009182526020808320909101546001600160a01b0316808352603590915260409091206003015490915060ff620100009091041615610e8b5780848481518110610e6457fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610e04565b509091505090565b610ea533611a7c565b610eef576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff16610f5f576040805162461bcd60e51b815260206004820152601f60248201527f72656465656d61626c653a20746f6b656e206e6f7420617661696c61626c6500604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff620100009091041615158115151415610fdd576040805162461bcd60e51b815260206004820152601b60248201527f72656465656d61626c653a206e6f207374617465206368616e67650000000000604482015290519081900360640190fd5b8015610fff57603854610ff790600163ffffffff611baa16565b603855611017565b60385461101390600163ffffffff611c0b16565b6038555b6001600160a01b038216600081815260356020908152604091829020600301805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526035602052604081206003015460ff166110f7576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526036602052604090205460ff16919050565b61112233611a7c565b61116c576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156113d657600083838381811061118557fe5b602090810292909201356001600160a01b0316600081815260359093526040909220600301549192505060ff166111fc576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03811660009081526035602052604090206003015462010000900460ff161561123e5760385461123a90600163ffffffff611c0b16565b6038555b6001600160a01b0381166000908152603560205260408120906112618282612421565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b60375461129d90600163ffffffff611c0b16565b81101561135657816001600160a01b0316603782815481106112bb57fe5b6000918252602090912001546001600160a01b0316141561134e57603780546112eb90600163ffffffff611c0b16565b815481106112f557fe5b600091825260209091200154603780546001600160a01b03909216918390811061131b57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550611356565b600101611289565b50603780548061136257fe5b6000828152602090819020820160001990810180546001600160a01b0319169055909101909155604080513381526001600160a01b0384169281019290925280517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad09281900390910190a15060010161116f565b505050565b6033546001600160a01b031690565b603a5490565b6113f933611a7c565b611443576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff166114b3576040805162461bcd60e51b815260206004820181905260248201527f6c6f616461626c653a20746f6b656e206973206e6f7420617661696c61626c65604482015290519081900360640190fd5b6001600160a01b0382166000818152603560209081526040918290206003018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b61153533611a7c565b61157f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b61158a838383611c4d565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6000806024831015611632576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b600061167e600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c93169050565b905061168a8682611092565b6116d0576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b141561173e5785611734600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b9250925050611915565b6001600160e01b031981166323b872dd60e01b14156118415760648410156117ad576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b6117f7603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d25169050565b611734604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b60448410156118815760405162461bcd60e51b815260040180806020018281038252602581526020018061249d6025913960400191505060405180910390fd5b6118cb601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d25169050565b611734602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b935093915050565b600061192a603a54611da7565b905061193533611a7c565b806119485750336001600160a01b038216145b611992576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03841660009081526035602052604090206003015460ff166119fb576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03841660008181526035602090815260409182902060028101879055600401859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60345490565b6039546001600160a01b031690565b6000611a89603454611da7565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611ade57600080fd5b505afa158015611af2573d6000803e3d6000fd5b505050506040513d6020811015611b0857600080fd5b505192915050565b611b18612468565b6040516020810160405282815280602083015250611b3582611ec8565b8152919050565b606080826000015167ffffffffffffffff81118015611b5a57600080fd5b506040519080825280601f01601f191660200182016040528015611b85576020820181803683370190505b5090506000602082019050611ba38185602001518660000151611f61565b5092915050565b600082820183811015611c04576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000611c0483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611f9f565b6001600160a01b038216611c7957611c746001600160a01b0384168263ffffffff61203616565b6113d6565b6113d66001600160a01b038316848363ffffffff61211b16565b6000611ca682600463ffffffff611baa16565b83511015611cf2576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d0684602063ffffffff611baa16565b90940151949350505050565b6000611ca682602063ffffffff611baa16565b6000611d3882601463ffffffff611baa16565b83511015611d84576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d9884602063ffffffff611baa16565b9094015160601c949350505050565b6033546000906001600160a01b0316611e07576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611e5357600080fd5b505afa158015611e67573d6000803e3d6000fd5b505050506040513d6020811015611e7d57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015611ade57600080fd5b60008082611eda5760009150506106c2565b6fffffffffffffffffffffffffffffffff8316611efe57601001600160801b830492505b67ffffffffffffffff8316611f1f5760080168010000000000000000830492505b63ffffffff8316611f3757600401600160201b830492505b61ffff8316611f4c5760020162010000830492505b60ff8316611f58576001015b60200392915050565b5b60208110611f81578151835260209283019290910190601f1901611f62565b905182516020929092036101000a6000190180199091169116179052565b6000818484111561202e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ff3578181015183820152602001611fdb565b50505050905090810190601f1680156120205780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b8047101561208b576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d80600081146120d6576040519150601f19603f3d011682016040523d82523d6000602084013e6120db565b606091505b50509050806113d65760405162461bcd60e51b815260040180806020018281038252603a8152602001806124c2603a913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526113d690849060606121bd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166122199092919063ffffffff16565b8051909150156113d6578080602001905160208110156121dc57600080fd5b50516113d65760405162461bcd60e51b815260040180806020018281038252602a8152602001806124fc602a913960400191505060405180910390fd5b60606122288484600085612230565b949350505050565b606061223b8561239d565b61228c576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106122cb5780518252601f1990920191602091820191016122ac565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461232d576040519150601f19603f3d011682016040523d82523d6000602084013e612332565b606091505b509150915081156123465791506122289050565b8051156123565780518082602001fd5b60405162461bcd60e51b8152602060048201818152865160248401528651879391928392604401919085019080838360008315611ff3578181015183820152602001611fdb565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106123e457805160ff1916838001178555612411565b82800160010185558215612411579182015b828111156124115782518255916020019190600101906123f6565b5061241d929150612482565b5090565b50805460018160011615610100020316600290046000825580601f106124475750612465565b601f0160209004906000526020600020908101906124659190612482565b50565b604051806040016040528060008152602001600081525090565b6106a191905b8082111561241d576000815560010161248856fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f7665416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d617920686176652072657665727465645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a264697066735822122042b0948b0099f07c159a5b491735eb0cd85fc9e33b8288d8b3e618ad9d5e980564736f6c634300060b0033436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"
>>>>>>> d04e1094... Upgrade to solc 0.6.11
||||||| parent of 3f937509... Update prettier plugin to support 0.6 syntax
var TokenWhitelistBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034553480156200005657600080fd5b50604051620028e4380380620028e4833981810160405260808110156200007c57600080fd5b5080516020820151604083015160609093015191929091620000a7846001600160e01b03620001a616565b620000bb826001600160e01b036200028616565b603a9290925550603980546001600160a01b0319166001600160a01b039092169190911790555060366020527f36aaa0d2a881446a5d284caf55fe779596f3dd15619675545d6094c8909139c2805460ff1990811660019081179092557fbdd45af1473e0dd27b6e035ccafb153373c61968247ba4b7c0e6c1c96496792f80548216831790557f877171c55a6bbc4aa0ea34a24de07dd3c5952b4a7d66df2b953f79c7955a122780548216831790556323b872dd60e01b6000527f04324ec27bf37ba85969616ff61c6c0e3939d6b141697b749b0be22765b29870805490911690911790556200034b565b600054610100900460ff1680620001cb5750620001cb6001600160e01b036200034516565b80620001da575060005460ff16155b620002175760405162461bcd60e51b815260040180806020018281038252602e815260200180620028b6602e913960400191505060405180910390fd5b600054610100900460ff1615801562000243576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200026f57603380546001600160a01b0319166001600160a01b0384161790555b801562000282576000805461ff00191690555b5050565b600054610100900460ff1680620002ab5750620002ab6001600160e01b036200034516565b80620002ba575060005460ff16155b620002f75760405162461bcd60e51b815260040180806020018281038252602e815260200180620028b6602e913960400191505060405180910390fd5b600054610100900460ff1615801562000323576000805460ff1961ff0019909116610100171660011790555b81156200026f576034829055801562000282576000805461ff00191690555050565b303b1590565b61255b806200035b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80636c3824ef116100a2578063996cba6811610071578063996cba6814610584578063afc72e93146105ba578063d545782e1461065b578063e2b4ce971461068d578063e9cbd8221461069557610116565b80636c3824ef146104bc5780637d73b2311461052a57806380cc0dda1461054e578063872773061461055657610116565b80633efec5e9116100e95780633efec5e9146103f0578063443dd2a4146103f857806344b049bc146104505780635d793a7d146104585780636a1744dc1461048657610116565b806313d5e8461461011b5780631d3a069f146101355780631f69565f1461017057806334c73edc14610240575b600080fd5b61012361069d565b60408051918252519081900360200190f35b61015c6004803603602081101561014b57600080fd5b50356001600160e01b0319166106a4565b604080519115158252519081900360200190f35b6101966004803603602081101561018657600080fd5b50356001600160a01b03166106c7565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b838110156101ff5781810151838201526020016101e7565b50505050905090810190601f16801561022c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103ee600480360360c081101561025657600080fd5b810190602081018135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460208302840111600160201b831117156102a357600080fd5b919390929091602081019035600160201b8111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460208302840111600160201b831117156102f357600080fd5b919390929091602081019035600160201b81111561031057600080fd5b82018360208201111561032257600080fd5b803590602001918460208302840111600160201b8311171561034357600080fd5b919390929091602081019035600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460208302840111600160201b8311171561039357600080fd5b919390929091602081019035600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460208302840111600160201b831117156103e357600080fd5b9193509150356107c5565b005b610196610c57565b610400610d55565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561043c578181015183820152602001610424565b505050509050019250505060405180910390f35b610400610db7565b6103ee6004803603604081101561046e57600080fd5b506001600160a01b0381351690602001351515610e9c565b61015c6004803603604081101561049c57600080fd5b5080356001600160a01b031690602001356001600160e01b031916611092565b6103ee600480360360208110156104d257600080fd5b810190602081018135600160201b8111156104ec57600080fd5b8201836020820111156104fe57600080fd5b803590602001918460208302840111600160201b8311171561051f57600080fd5b509092509050611119565b6105326113db565b604080516001600160a01b039092168252519081900360200190f35b6101236113ea565b6103ee6004803603604081101561056c57600080fd5b506001600160a01b03813516906020013515156113f0565b6103ee6004803603606081101561059a57600080fd5b506001600160a01b0381358116916020810135909116906040013561152c565b610638600480360360408110156105d057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b5090925090506115d9565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103ee6004803603606081101561067157600080fd5b506001600160a01b03813516906020810135906040013561191d565b610123611a67565b610532611a6d565b6038545b90565b6001600160e01b0319811660009081526036602052604090205460ff165b919050565b6001600160a01b038116600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff8084169694840481169562010000909404169391929189918301828280156107a45780601f10610779576101008083540402835291602001916107a4565b820191906000526020600020905b81548152906001019060200180831161078757829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107ce33611a7c565b610818576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561082657508986145b801561083157508984145b801561083c57508984145b61088d576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4957603560008d8d848181106108a857fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610923576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606109486109438c8c8581811061093757fe5b90506020020135611b10565b611b3c565b90506040518060e001604052808281526020018a8a8581811061096757fe5b9050602002013581526020016000815260200160011515815260200188888581811061098f57fe5b905060200201351515151581526020018686858181106109ab57fe5b9050602002013515151515815260200184815250603560008f8f868181106109cf57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000019080519060200190610a1b9291906123a3565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560378d8d84818110610a9257fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b039590920293909301359390931692909217905550848483818110610ada57fe5b9050602002013515610afe57603854610afa90600163ffffffff611baa16565b6038555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2c57fe5b905060200201356001600160a01b0316838c8c87818110610b4957fe5b905060200201358b8b88818110610b5c57fe5b9050602002013515158a8a89818110610b7157fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610c01578181015183820152602001610be9565b50505050905090810190601f168015610c2e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610890565b505050505050505050505050565b6039546001600160a01b0316600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff808416969484048116956201000090940416939192918991830182828015610d365780601f10610d0b57610100808354040283529160200191610d36565b820191906000526020600020905b815481529060010190602001808311610d1957829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606037805480602002602001604051908101604052809291908181526020018280548015610dad57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d8f575b5050505050905090565b60608060385467ffffffffffffffff81118015610dd357600080fd5b50604051908082528060200260200182016040528015610dfd578160200160208202803683370190505b5090506000805b603754811015610e9457600060378281548110610e1d57fe5b60009182526020808320909101546001600160a01b0316808352603590915260409091206003015490915060ff620100009091041615610e8b5780848481518110610e6457fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610e04565b509091505090565b610ea533611a7c565b610eef576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff16610f5f576040805162461bcd60e51b815260206004820152601f60248201527f72656465656d61626c653a20746f6b656e206e6f7420617661696c61626c6500604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff620100009091041615158115151415610fdd576040805162461bcd60e51b815260206004820152601b60248201527f72656465656d61626c653a206e6f207374617465206368616e67650000000000604482015290519081900360640190fd5b8015610fff57603854610ff790600163ffffffff611baa16565b603855611017565b60385461101390600163ffffffff611c0b16565b6038555b6001600160a01b038216600081815260356020908152604091829020600301805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526035602052604081206003015460ff166110f7576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526036602052604090205460ff16919050565b61112233611a7c565b61116c576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156113d657600083838381811061118557fe5b602090810292909201356001600160a01b0316600081815260359093526040909220600301549192505060ff166111fc576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03811660009081526035602052604090206003015462010000900460ff161561123e5760385461123a90600163ffffffff611c0b16565b6038555b6001600160a01b0381166000908152603560205260408120906112618282612421565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b60375461129d90600163ffffffff611c0b16565b81101561135657816001600160a01b0316603782815481106112bb57fe5b6000918252602090912001546001600160a01b0316141561134e57603780546112eb90600163ffffffff611c0b16565b815481106112f557fe5b600091825260209091200154603780546001600160a01b03909216918390811061131b57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550611356565b600101611289565b50603780548061136257fe5b6000828152602090819020820160001990810180546001600160a01b0319169055909101909155604080513381526001600160a01b0384169281019290925280517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad09281900390910190a15060010161116f565b505050565b6033546001600160a01b031690565b603a5490565b6113f933611a7c565b611443576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff166114b3576040805162461bcd60e51b815260206004820181905260248201527f6c6f616461626c653a20746f6b656e206973206e6f7420617661696c61626c65604482015290519081900360640190fd5b6001600160a01b0382166000818152603560209081526040918290206003018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b61153533611a7c565b61157f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b61158a838383611c4d565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6000806024831015611632576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b600061167e600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c93169050565b905061168a8682611092565b6116d0576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b141561173e5785611734600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b9250925050611915565b6001600160e01b031981166323b872dd60e01b14156118415760648410156117ad576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b6117f7603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d25169050565b611734604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b60448410156118815760405162461bcd60e51b815260040180806020018281038252602581526020018061249d6025913960400191505060405180910390fd5b6118cb601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d25169050565b611734602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b935093915050565b600061192a603a54611da7565b905061193533611a7c565b806119485750336001600160a01b038216145b611992576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03841660009081526035602052604090206003015460ff166119fb576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03841660008181526035602090815260409182902060028101879055600401859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60345490565b6039546001600160a01b031690565b6000611a89603454611da7565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611ade57600080fd5b505afa158015611af2573d6000803e3d6000fd5b505050506040513d6020811015611b0857600080fd5b505192915050565b611b18612468565b6040516020810160405282815280602083015250611b3582611ec8565b8152919050565b606080826000015167ffffffffffffffff81118015611b5a57600080fd5b506040519080825280601f01601f191660200182016040528015611b85576020820181803683370190505b5090506000602082019050611ba38185602001518660000151611f61565b5092915050565b600082820183811015611c04576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000611c0483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611f9f565b6001600160a01b038216611c7957611c746001600160a01b0384168263ffffffff61203616565b6113d6565b6113d66001600160a01b038316848363ffffffff61211b16565b6000611ca682600463ffffffff611baa16565b83511015611cf2576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d0684602063ffffffff611baa16565b90940151949350505050565b6000611ca682602063ffffffff611baa16565b6000611d3882601463ffffffff611baa16565b83511015611d84576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d9884602063ffffffff611baa16565b9094015160601c949350505050565b6033546000906001600160a01b0316611e07576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611e5357600080fd5b505afa158015611e67573d6000803e3d6000fd5b505050506040513d6020811015611e7d57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015611ade57600080fd5b60008082611eda5760009150506106c2565b6fffffffffffffffffffffffffffffffff8316611efe57601001600160801b830492505b67ffffffffffffffff8316611f1f5760080168010000000000000000830492505b63ffffffff8316611f3757600401600160201b830492505b61ffff8316611f4c5760020162010000830492505b60ff8316611f58576001015b60200392915050565b5b60208110611f81578151835260209283019290910190601f1901611f62565b905182516020929092036101000a6000190180199091169116179052565b6000818484111561202e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ff3578181015183820152602001611fdb565b50505050905090810190601f1680156120205780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b8047101561208b576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d80600081146120d6576040519150601f19603f3d011682016040523d82523d6000602084013e6120db565b606091505b50509050806113d65760405162461bcd60e51b815260040180806020018281038252603a8152602001806124c2603a913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526113d690849060606121bd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166122199092919063ffffffff16565b8051909150156113d6578080602001905160208110156121dc57600080fd5b50516113d65760405162461bcd60e51b815260040180806020018281038252602a8152602001806124fc602a913960400191505060405180910390fd5b60606122288484600085612230565b949350505050565b606061223b8561239d565b61228c576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106122cb5780518252601f1990920191602091820191016122ac565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461232d576040519150601f19603f3d011682016040523d82523d6000602084013e612332565b606091505b509150915081156123465791506122289050565b8051156123565780518082602001fd5b60405162461bcd60e51b8152602060048201818152865160248401528651879391928392604401919085019080838360008315611ff3578181015183820152602001611fdb565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106123e457805160ff1916838001178555612411565b82800160010185558215612411579182015b828111156124115782518255916020019190600101906123f6565b5061241d929150612482565b5090565b50805460018160011615610100020316600290046000825580601f106124475750612465565b601f0160209004906000526020600020908101906124659190612482565b50565b604051806040016040528060008152602001600081525090565b6106a191905b8082111561241d576000815560010161248856fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f7665416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d617920686176652072657665727465645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a264697066735822122042b0948b0099f07c159a5b491735eb0cd85fc9e33b8288d8b3e618ad9d5e980564736f6c634300060b0033436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"
=======
var TokenWhitelistBin = "0x6080604052603380546001600160a01b0319166e0c2e074ec69a0dfb2997ba6c7d2e1e1790557f7f2ce995617d2816b426c5c8698c5ec2952f7a34bb10f38326f74933d58936976034553480156200005657600080fd5b50604051620028e4380380620028e4833981810160405260808110156200007c57600080fd5b5080516020820151604083015160609093015191929091620000a7846001600160e01b03620001a616565b620000bb826001600160e01b036200028616565b603a9290925550603980546001600160a01b0319166001600160a01b039092169190911790555060366020527f36aaa0d2a881446a5d284caf55fe779596f3dd15619675545d6094c8909139c2805460ff1990811660019081179092557fbdd45af1473e0dd27b6e035ccafb153373c61968247ba4b7c0e6c1c96496792f80548216831790557f877171c55a6bbc4aa0ea34a24de07dd3c5952b4a7d66df2b953f79c7955a122780548216831790556323b872dd60e01b6000527f04324ec27bf37ba85969616ff61c6c0e3939d6b141697b749b0be22765b29870805490911690911790556200034b565b600054610100900460ff1680620001cb5750620001cb6001600160e01b036200034516565b80620001da575060005460ff16155b620002175760405162461bcd60e51b815260040180806020018281038252602e815260200180620028b6602e913960400191505060405180910390fd5b600054610100900460ff1615801562000243576000805460ff1961ff0019909116610100171660011790555b6001600160a01b038216156200026f57603380546001600160a01b0319166001600160a01b0384161790555b801562000282576000805461ff00191690555b5050565b600054610100900460ff1680620002ab5750620002ab6001600160e01b036200034516565b80620002ba575060005460ff16155b620002f75760405162461bcd60e51b815260040180806020018281038252602e815260200180620028b6602e913960400191505060405180910390fd5b600054610100900460ff1615801562000323576000805460ff1961ff0019909116610100171660011790555b81156200026f576034829055801562000282576000805461ff00191690555050565b303b1590565b61255b806200035b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80636c3824ef116100a2578063996cba6811610071578063996cba6814610584578063afc72e93146105ba578063d545782e1461065b578063e2b4ce971461068d578063e9cbd8221461069557610116565b80636c3824ef146104bc5780637d73b2311461052a57806380cc0dda1461054e578063872773061461055657610116565b80633efec5e9116100e95780633efec5e9146103f0578063443dd2a4146103f857806344b049bc146104505780635d793a7d146104585780636a1744dc1461048657610116565b806313d5e8461461011b5780631d3a069f146101355780631f69565f1461017057806334c73edc14610240575b600080fd5b61012361069d565b60408051918252519081900360200190f35b61015c6004803603602081101561014b57600080fd5b50356001600160e01b0319166106a4565b604080519115158252519081900360200190f35b6101966004803603602081101561018657600080fd5b50356001600160a01b03166106c7565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b838110156101ff5781810151838201526020016101e7565b50505050905090810190601f16801561022c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103ee600480360360c081101561025657600080fd5b810190602081018135600160201b81111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460208302840111600160201b831117156102a357600080fd5b919390929091602081019035600160201b8111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460208302840111600160201b831117156102f357600080fd5b919390929091602081019035600160201b81111561031057600080fd5b82018360208201111561032257600080fd5b803590602001918460208302840111600160201b8311171561034357600080fd5b919390929091602081019035600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460208302840111600160201b8311171561039357600080fd5b919390929091602081019035600160201b8111156103b057600080fd5b8201836020820111156103c257600080fd5b803590602001918460208302840111600160201b831117156103e357600080fd5b9193509150356107c5565b005b610196610c57565b610400610d55565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561043c578181015183820152602001610424565b505050509050019250505060405180910390f35b610400610db7565b6103ee6004803603604081101561046e57600080fd5b506001600160a01b0381351690602001351515610e9c565b61015c6004803603604081101561049c57600080fd5b5080356001600160a01b031690602001356001600160e01b031916611092565b6103ee600480360360208110156104d257600080fd5b810190602081018135600160201b8111156104ec57600080fd5b8201836020820111156104fe57600080fd5b803590602001918460208302840111600160201b8311171561051f57600080fd5b509092509050611119565b6105326113db565b604080516001600160a01b039092168252519081900360200190f35b6101236113ea565b6103ee6004803603604081101561056c57600080fd5b506001600160a01b03813516906020013515156113f0565b6103ee6004803603606081101561059a57600080fd5b506001600160a01b0381358116916020810135909116906040013561152c565b610638600480360360408110156105d057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b5090925090506115d9565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6103ee6004803603606081101561067157600080fd5b506001600160a01b03813516906020810135906040013561191d565b610123611a67565b610532611a6d565b6038545b90565b6001600160e01b0319811660009081526036602052604090205460ff165b919050565b6001600160a01b038116600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff8084169694840481169562010000909404169391929189918301828280156107a45780601f10610779576101008083540402835291602001916107a4565b820191906000526020600020905b81548152906001019060200180831161078757829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b6107ce33611a7c565b610818576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b898814801561082657508986145b801561083157508984145b801561083c57508984145b61088d576040805162461bcd60e51b815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610c4957603560008d8d848181106108a857fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff1615610923576040805162461bcd60e51b815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b60606109486109438c8c8581811061093757fe5b90506020020135611b10565b611b3c565b90506040518060e001604052808281526020018a8a8581811061096757fe5b9050602002013581526020016000815260200160011515815260200188888581811061098f57fe5b905060200201351515151581526020018686858181106109ab57fe5b9050602002013515151515815260200184815250603560008f8f868181106109cf57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000019080519060200190610a1b9291906123a3565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560378d8d84818110610a9257fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b039590920293909301359390931692909217905550848483818110610ada57fe5b9050602002013515610afe57603854610afa90600163ffffffff611baa16565b6038555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610b2c57fe5b905060200201356001600160a01b0316838c8c87818110610b4957fe5b905060200201358b8b88818110610b5c57fe5b9050602002013515158a8a89818110610b7157fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610c01578181015183820152602001610be9565b50505050905090810190601f168015610c2e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610890565b505050505050505050505050565b6039546001600160a01b0316600090815260356020908152604080832060018082015460028084015460038501546004860154865488516101009782161588026000190190911694909404601f81018a90048a0285018a019098528784526060999889988998899889988998929788979196909560ff808416969484048116956201000090940416939192918991830182828015610d365780601f10610d0b57610100808354040283529160200191610d36565b820191906000526020600020905b815481529060010190602001808311610d1957829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606037805480602002602001604051908101604052809291908181526020018280548015610dad57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d8f575b5050505050905090565b60608060385467ffffffffffffffff81118015610dd357600080fd5b50604051908082528060200260200182016040528015610dfd578160200160208202803683370190505b5090506000805b603754811015610e9457600060378281548110610e1d57fe5b60009182526020808320909101546001600160a01b0316808352603590915260409091206003015490915060ff620100009091041615610e8b5780848481518110610e6457fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b50600101610e04565b509091505090565b610ea533611a7c565b610eef576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff16610f5f576040805162461bcd60e51b815260206004820152601f60248201527f72656465656d61626c653a20746f6b656e206e6f7420617661696c61626c6500604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff620100009091041615158115151415610fdd576040805162461bcd60e51b815260206004820152601b60248201527f72656465656d61626c653a206e6f207374617465206368616e67650000000000604482015290519081900360640190fd5b8015610fff57603854610ff790600163ffffffff611baa16565b603855611017565b60385461101390600163ffffffff611c0b16565b6038555b6001600160a01b038216600081815260356020908152604091829020600301805485151562010000810262ff0000199092169190911790915582513381529182019390935280820192909252517fcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d9181900360600190a15050565b6001600160a01b03821660009081526035602052604081206003015460ff166110f7576040805162461bcd60e51b81526020600482015260126024820152713737b716b2bc34b9ba34b733903a37b5b2b760711b604482015290519081900360640190fd5b506001600160e01b03191660009081526036602052604090205460ff16919050565b61112233611a7c565b61116c576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b60005b818110156113d657600083838381811061118557fe5b602090810292909201356001600160a01b0316600081815260359093526040909220600301549192505060ff166111fc576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03811660009081526035602052604090206003015462010000900460ff161561123e5760385461123a90600163ffffffff611c0b16565b6038555b6001600160a01b0381166000908152603560205260408120906112618282612421565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b60375461129d90600163ffffffff611c0b16565b81101561135657816001600160a01b0316603782815481106112bb57fe5b6000918252602090912001546001600160a01b0316141561134e57603780546112eb90600163ffffffff611c0b16565b815481106112f557fe5b600091825260209091200154603780546001600160a01b03909216918390811061131b57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550611356565b600101611289565b50603780548061136257fe5b6000828152602090819020820160001990810180546001600160a01b0319169055909101909155604080513381526001600160a01b0384169281019290925280517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad09281900390910190a15060010161116f565b505050565b6033546001600160a01b031690565b603a5490565b6113f933611a7c565b611443576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03821660009081526035602052604090206003015460ff166114b3576040805162461bcd60e51b815260206004820181905260248201527f6c6f616461626c653a20746f6b656e206973206e6f7420617661696c61626c65604482015290519081900360640190fd5b6001600160a01b0382166000818152603560209081526040918290206003018054851515610100810261ff00199092169190911790915582513381529182019390935280820192909252517f0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec9181900360600190a15050565b61153533611a7c565b61157f576040805162461bcd60e51b815260206004820152601660248201527539b2b73232b91034b9903737ba1030b71030b236b4b760511b604482015290519081900360640190fd5b61158a838383611c4d565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6000806024831015611632576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206d6574686f642d656e636f64696e67206279746573604482015290519081900360640190fd5b600061167e600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611c93169050565b905061168a8682611092565b6116d0576040805162461bcd60e51b81526020600482015260126024820152711d5b9cdd5c1c1bdc9d1959081b595d1a1bd960721b604482015290519081900360640190fd5b6001600160e01b03198116630852cd8d60e31b141561173e5785611734600487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b9250925050611915565b6001600160e01b031981166323b872dd60e01b14156118415760648410156117ad576040805162461bcd60e51b815260206004820181905260248201527f6e6f7420656e6f756768206461746120666f72207472616e7366657246726f6d604482015290519081900360640190fd5b6117f7603086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d25169050565b611734604487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b60448410156118815760405162461bcd60e51b815260040180806020018281038252602581526020018061249d6025913960400191505060405180910390fd5b6118cb601086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d25169050565b611734602487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff611d12169050565b935093915050565b600061192a603a54611da7565b905061193533611a7c565b806119485750336001600160a01b038216145b611992576040805162461bcd60e51b815260206004820152601660248201527532b4ba3432b91037b930b1b6329037b91030b236b4b760511b604482015290519081900360640190fd5b6001600160a01b03841660009081526035602052604090206003015460ff166119fb576040805162461bcd60e51b8152602060048201526016602482015275746f6b656e206973206e6f7420617661696c61626c6560501b604482015290519081900360640190fd5b6001600160a01b03841660008181526035602090815260409182902060028101879055600401859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60345490565b6039546001600160a01b031690565b6000611a89603454611da7565b6001600160a01b03166324d7806c836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611ade57600080fd5b505afa158015611af2573d6000803e3d6000fd5b505050506040513d6020811015611b0857600080fd5b505192915050565b611b18612468565b6040516020810160405282815280602083015250611b3582611ec8565b8152919050565b606080826000015167ffffffffffffffff81118015611b5a57600080fd5b506040519080825280601f01601f191660200182016040528015611b85576020820181803683370190505b5090506000602082019050611ba38185602001518660000151611f61565b5092915050565b600082820183811015611c04576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000611c0483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611f9f565b6001600160a01b038216611c7957611c746001600160a01b0384168263ffffffff61203616565b6113d6565b6113d66001600160a01b038316848363ffffffff61211b16565b6000611ca682600463ffffffff611baa16565b83511015611cf2576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d0684602063ffffffff611baa16565b90940151949350505050565b6000611ca682602063ffffffff611baa16565b6000611d3882601463ffffffff611baa16565b83511015611d84576040805162461bcd60e51b8152602060048201526014602482015273736c6963696e67206f7574206f662072616e676560601b604482015290519081900360640190fd5b600080611d9884602063ffffffff611baa16565b9094015160601c949350505050565b6033546000906001600160a01b0316611e07576040805162461bcd60e51b815260206004820152601d60248201527f454e535265736f6c7661626c65206e6f7420696e697469616c697a6564000000604482015290519081900360640190fd5b60335460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b158015611e5357600080fd5b505afa158015611e67573d6000803e3d6000fd5b505050506040513d6020811015611e7d57600080fd5b505160408051631d9dabef60e11b81526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b158015611ade57600080fd5b60008082611eda5760009150506106c2565b6fffffffffffffffffffffffffffffffff8316611efe57601001600160801b830492505b67ffffffffffffffff8316611f1f5760080168010000000000000000830492505b63ffffffff8316611f3757600401600160201b830492505b61ffff8316611f4c5760020162010000830492505b60ff8316611f58576001015b60200392915050565b5b60208110611f81578151835260209283019290910190601f1901611f62565b905182516020929092036101000a6000190180199091169116179052565b6000818484111561202e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ff3578181015183820152602001611fdb565b50505050905090810190601f1680156120205780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b8047101561208b576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b6040516000906001600160a01b0384169083908381818185875af1925050503d80600081146120d6576040519150601f19603f3d011682016040523d82523d6000602084013e6120db565b606091505b50509050806113d65760405162461bcd60e51b815260040180806020018281038252603a8152602001806124c2603a913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526113d690849060606121bd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166122199092919063ffffffff16565b8051909150156113d6578080602001905160208110156121dc57600080fd5b50516113d65760405162461bcd60e51b815260040180806020018281038252602a8152602001806124fc602a913960400191505060405180910390fd5b60606122288484600085612230565b949350505050565b606061223b8561239d565b61228c576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106122cb5780518252601f1990920191602091820191016122ac565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461232d576040519150601f19603f3d011682016040523d82523d6000602084013e612332565b606091505b509150915081156123465791506122289050565b8051156123565780518082602001fd5b60405162461bcd60e51b8152602060048201818152865160248401528651879391928392604401919085019080838360008315611ff3578181015183820152602001611fdb565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106123e457805160ff1916838001178555612411565b82800160010185558215612411579182015b828111156124115782518255916020019190600101906123f6565b5061241d929150612482565b5090565b50805460018160011615610100020316600290046000825580601f106124475750612465565b601f0160209004906000526020600020908101906124659190612482565b50565b604051806040016040528060008152602001600081525090565b6106a191905b8082111561241d576000815560010161248856fe6e6f7420656e6f756768206461746120666f72207472616e736665722f61707070726f7665416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d617920686176652072657665727465645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a264697066735822122094c7cee700372938833e0bfd608db66486a9dbf61eb42a45a0ac9fc235e256df64736f6c634300060b0033436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564"
>>>>>>> 3f937509... Update prettier plugin to support 0.6 syntax

// DeployTokenWhitelist deploys a new Ethereum contract, binding an instance of TokenWhitelist to it.
func DeployTokenWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _oracleNode_ [32]byte, _controllerNode_ [32]byte, _stablecoinAddress_ common.Address) (common.Address, *types.Transaction, *TokenWhitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistBin), backend, _ens_, _oracleNode_, _controllerNode_, _stablecoinAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenWhitelist{TokenWhitelistCaller: TokenWhitelistCaller{contract: contract}, TokenWhitelistTransactor: TokenWhitelistTransactor{contract: contract}, TokenWhitelistFilterer: TokenWhitelistFilterer{contract: contract}}, nil
}

// TokenWhitelist is an auto generated Go binding around an Ethereum contract.
type TokenWhitelist struct {
	TokenWhitelistCaller     // Read-only binding to the contract
	TokenWhitelistTransactor // Write-only binding to the contract
	TokenWhitelistFilterer   // Log filterer for contract events
}

// TokenWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWhitelistSession struct {
	Contract     *TokenWhitelist   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWhitelistCallerSession struct {
	Contract *TokenWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWhitelistTransactorSession struct {
	Contract     *TokenWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWhitelistRaw struct {
	Contract *TokenWhitelist // Generic contract binding to access the raw methods on
}

// TokenWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWhitelistCallerRaw struct {
	Contract *TokenWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWhitelistTransactorRaw struct {
	Contract *TokenWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWhitelist creates a new instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelist(address common.Address, backend bind.ContractBackend) (*TokenWhitelist, error) {
	contract, err := bindTokenWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelist{TokenWhitelistCaller: TokenWhitelistCaller{contract: contract}, TokenWhitelistTransactor: TokenWhitelistTransactor{contract: contract}, TokenWhitelistFilterer: TokenWhitelistFilterer{contract: contract}}, nil
}

// NewTokenWhitelistCaller creates a new read-only instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistCaller(address common.Address, caller bind.ContractCaller) (*TokenWhitelistCaller, error) {
	contract, err := bindTokenWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistCaller{contract: contract}, nil
}

// NewTokenWhitelistTransactor creates a new write-only instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWhitelistTransactor, error) {
	contract, err := bindTokenWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistTransactor{contract: contract}, nil
}

// NewTokenWhitelistFilterer creates a new log filterer instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWhitelistFilterer, error) {
	contract, err := bindTokenWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistFilterer{contract: contract}, nil
}

// bindTokenWhitelist binds a generic wrapper to an already deployed contract.
func bindTokenWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelist *TokenWhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelist.Contract.TokenWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelist *TokenWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelist *TokenWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelist *TokenWhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelist *TokenWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelist *TokenWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.contract.Transact(opts, method, params...)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) ControllerNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCallerSession) ControllerNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.CallOpts)
}

// GetERC20RecipientAndAmount is a free data retrieval call binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) constant returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetERC20RecipientAndAmount(opts *bind.CallOpts, _token common.Address, _data []byte) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getERC20RecipientAndAmount", _token, _data)
	return *ret0, *ret1, err
}

// GetERC20RecipientAndAmount is a free data retrieval call binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) constant returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetERC20RecipientAndAmount(_token common.Address, _data []byte) (common.Address, *big.Int, error) {
	return _TokenWhitelist.Contract.GetERC20RecipientAndAmount(&_TokenWhitelist.CallOpts, _token, _data)
}

// GetERC20RecipientAndAmount is a free data retrieval call binding the contract method 0xafc72e93.
//
// Solidity: function getERC20RecipientAndAmount(address _token, bytes _data) constant returns(address, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetERC20RecipientAndAmount(_token common.Address, _data []byte) (common.Address, *big.Int, error) {
	return _TokenWhitelist.Contract.GetERC20RecipientAndAmount(&_TokenWhitelist.CallOpts, _token, _data)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetStablecoinInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getStablecoinInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetTokenInfo(opts *bind.CallOpts, _a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getTokenInfo", _a)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.CallOpts, _a)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.CallOpts, _a)
}

// IsERC20MethodSupported is a free data retrieval call binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCaller) IsERC20MethodSupported(opts *bind.CallOpts, _token common.Address, _methodId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "isERC20MethodSupported", _token, _methodId)
	return *ret0, err
}

// IsERC20MethodSupported is a free data retrieval call binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsERC20MethodSupported(_token common.Address, _methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodSupported(&_TokenWhitelist.CallOpts, _token, _methodId)
}

// IsERC20MethodSupported is a free data retrieval call binding the contract method 0x6a1744dc.
//
// Solidity: function isERC20MethodSupported(address _token, bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCallerSession) IsERC20MethodSupported(_token common.Address, _methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodSupported(&_TokenWhitelist.CallOpts, _token, _methodId)
}

// IsERC20MethodWhitelisted is a free data retrieval call binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCaller) IsERC20MethodWhitelisted(opts *bind.CallOpts, _methodId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "isERC20MethodWhitelisted", _methodId)
	return *ret0, err
}

// IsERC20MethodWhitelisted is a free data retrieval call binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsERC20MethodWhitelisted(_methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodWhitelisted(&_TokenWhitelist.CallOpts, _methodId)
}

// IsERC20MethodWhitelisted is a free data retrieval call binding the contract method 0x1d3a069f.
//
// Solidity: function isERC20MethodWhitelisted(bytes4 _methodId) constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCallerSession) IsERC20MethodWhitelisted(_methodId [4]byte) (bool, error) {
	return _TokenWhitelist.Contract.IsERC20MethodWhitelisted(&_TokenWhitelist.CallOpts, _methodId)
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCaller) OracleNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "oracleNode")
	return *ret0, err
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) OracleNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.CallOpts)
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCallerSession) OracleNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.CallOpts)
}

// RedeemableCounter is a free data retrieval call binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCaller) RedeemableCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "redeemableCounter")
	return *ret0, err
}

// RedeemableCounter is a free data retrieval call binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistSession) RedeemableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.RedeemableCounter(&_TokenWhitelist.CallOpts)
}

// RedeemableCounter is a free data retrieval call binding the contract method 0x13d5e846.
//
// Solidity: function redeemableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) RedeemableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.RedeemableCounter(&_TokenWhitelist.CallOpts)
}

// RedeemableTokens is a free data retrieval call binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) RedeemableTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "redeemableTokens")
	return *ret0, err
}

// RedeemableTokens is a free data retrieval call binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) RedeemableTokens() ([]common.Address, error) {
	return _TokenWhitelist.Contract.RedeemableTokens(&_TokenWhitelist.CallOpts)
}

// RedeemableTokens is a free data retrieval call binding the contract method 0x44b049bc.
//
// Solidity: function redeemableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) RedeemableTokens() ([]common.Address, error) {
	return _TokenWhitelist.Contract.RedeemableTokens(&_TokenWhitelist.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) Stablecoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "stablecoin")
	return *ret0, err
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) Stablecoin() (common.Address, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) Stablecoin() (common.Address, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.CallOpts)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) TokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "tokenAddressArray")
	return *ret0, err
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.CallOpts)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _redeemable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _redeemable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "addTokens", _tokens, _symbols, _magnitude, _loadable, _redeemable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _redeemable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _redeemable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _redeemable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _redeemable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _redeemable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _redeemable, _lastUpdate)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "claim", _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Claim(&_TokenWhitelist.TransactOpts, _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Claim(&_TokenWhitelist.TransactOpts, _to, _asset, _amount)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) RemoveTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "removeTokens", _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RemoveTokens(&_TokenWhitelist.TransactOpts, _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RemoveTokens(&_TokenWhitelist.TransactOpts, _tokens)
}

// SetTokenLoadable is a paid mutator transaction binding the contract method 0x87277306.
//
// Solidity: function setTokenLoadable(address _token, bool _loadable) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) SetTokenLoadable(opts *bind.TransactOpts, _token common.Address, _loadable bool) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "setTokenLoadable", _token, _loadable)
}

// SetTokenLoadable is a paid mutator transaction binding the contract method 0x87277306.
//
// Solidity: function setTokenLoadable(address _token, bool _loadable) returns()
func (_TokenWhitelist *TokenWhitelistSession) SetTokenLoadable(_token common.Address, _loadable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenLoadable(&_TokenWhitelist.TransactOpts, _token, _loadable)
}

// SetTokenLoadable is a paid mutator transaction binding the contract method 0x87277306.
//
// Solidity: function setTokenLoadable(address _token, bool _loadable) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) SetTokenLoadable(_token common.Address, _loadable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenLoadable(&_TokenWhitelist.TransactOpts, _token, _loadable)
}

// SetTokenRedeemable is a paid mutator transaction binding the contract method 0x5d793a7d.
//
// Solidity: function setTokenRedeemable(address _token, bool _redeemable) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) SetTokenRedeemable(opts *bind.TransactOpts, _token common.Address, _redeemable bool) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "setTokenRedeemable", _token, _redeemable)
}

// SetTokenRedeemable is a paid mutator transaction binding the contract method 0x5d793a7d.
//
// Solidity: function setTokenRedeemable(address _token, bool _redeemable) returns()
func (_TokenWhitelist *TokenWhitelistSession) SetTokenRedeemable(_token common.Address, _redeemable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenRedeemable(&_TokenWhitelist.TransactOpts, _token, _redeemable)
}

// SetTokenRedeemable is a paid mutator transaction binding the contract method 0x5d793a7d.
//
// Solidity: function setTokenRedeemable(address _token, bool _redeemable) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) SetTokenRedeemable(_token common.Address, _redeemable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.SetTokenRedeemable(&_TokenWhitelist.TransactOpts, _token, _redeemable)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) UpdateTokenRate(opts *bind.TransactOpts, _token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "updateTokenRate", _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.UpdateTokenRate(&_TokenWhitelist.TransactOpts, _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.UpdateTokenRate(&_TokenWhitelist.TransactOpts, _token, _rate, _updateDate)
}

// TokenWhitelistAddedExclusiveMethodIterator is returned from FilterAddedExclusiveMethod and is used to iterate over the raw logs and unpacked data for AddedExclusiveMethod events raised by the TokenWhitelist contract.
type TokenWhitelistAddedExclusiveMethodIterator struct {
	Event *TokenWhitelistAddedExclusiveMethod // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAddedExclusiveMethodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedExclusiveMethod)
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
		it.Event = new(TokenWhitelistAddedExclusiveMethod)
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
func (it *TokenWhitelistAddedExclusiveMethodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedExclusiveMethodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedExclusiveMethod represents a AddedExclusiveMethod event raised by the TokenWhitelist contract.
type TokenWhitelistAddedExclusiveMethod struct {
	Token    common.Address
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddedExclusiveMethod is a free log retrieval operation binding the contract event 0xfb181256b03ef9051c59b29b98e8ef8dc1161e61d9062e1192ddd073806b0876.
//
// Solidity: event AddedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedExclusiveMethod(opts *bind.FilterOpts) (*TokenWhitelistAddedExclusiveMethodIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedExclusiveMethodIterator{contract: _TokenWhitelist.contract, event: "AddedExclusiveMethod", logs: logs, sub: sub}, nil
}

// WatchAddedExclusiveMethod is a free log subscription operation binding the contract event 0xfb181256b03ef9051c59b29b98e8ef8dc1161e61d9062e1192ddd073806b0876.
//
// Solidity: event AddedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedExclusiveMethod(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedExclusiveMethod) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedExclusiveMethod)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedExclusiveMethod", log); err != nil {
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

// ParseAddedExclusiveMethod is a log parse operation binding the contract event 0xfb181256b03ef9051c59b29b98e8ef8dc1161e61d9062e1192ddd073806b0876.
//
// Solidity: event AddedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseAddedExclusiveMethod(log types.Log) (*TokenWhitelistAddedExclusiveMethod, error) {
	event := new(TokenWhitelistAddedExclusiveMethod)
	if err := _TokenWhitelist.contract.UnpackLog(event, "AddedExclusiveMethod", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistAddedMethodIdIterator is returned from FilterAddedMethodId and is used to iterate over the raw logs and unpacked data for AddedMethodId events raised by the TokenWhitelist contract.
type TokenWhitelistAddedMethodIdIterator struct {
	Event *TokenWhitelistAddedMethodId // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAddedMethodIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedMethodId)
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
		it.Event = new(TokenWhitelistAddedMethodId)
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
func (it *TokenWhitelistAddedMethodIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedMethodIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedMethodId represents a AddedMethodId event raised by the TokenWhitelist contract.
type TokenWhitelistAddedMethodId struct {
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddedMethodId is a free log retrieval operation binding the contract event 0xcad8cc4e064e022264c8f21f5293f8b3c267eaa6895ee7c9e0b34689726eae71.
//
// Solidity: event AddedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedMethodId(opts *bind.FilterOpts) (*TokenWhitelistAddedMethodIdIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedMethodId")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedMethodIdIterator{contract: _TokenWhitelist.contract, event: "AddedMethodId", logs: logs, sub: sub}, nil
}

// WatchAddedMethodId is a free log subscription operation binding the contract event 0xcad8cc4e064e022264c8f21f5293f8b3c267eaa6895ee7c9e0b34689726eae71.
//
// Solidity: event AddedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedMethodId(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedMethodId) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedMethodId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedMethodId)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedMethodId", log); err != nil {
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

// ParseAddedMethodId is a log parse operation binding the contract event 0xcad8cc4e064e022264c8f21f5293f8b3c267eaa6895ee7c9e0b34689726eae71.
//
// Solidity: event AddedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseAddedMethodId(log types.Log) (*TokenWhitelistAddedMethodId, error) {
	event := new(TokenWhitelistAddedMethodId)
	if err := _TokenWhitelist.contract.UnpackLog(event, "AddedMethodId", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistAddedTokenIterator is returned from FilterAddedToken and is used to iterate over the raw logs and unpacked data for AddedToken events raised by the TokenWhitelist contract.
type TokenWhitelistAddedTokenIterator struct {
	Event *TokenWhitelistAddedToken // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAddedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedToken)
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
		it.Event = new(TokenWhitelistAddedToken)
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
func (it *TokenWhitelistAddedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedToken represents a AddedToken event raised by the TokenWhitelist contract.
type TokenWhitelistAddedToken struct {
	Sender     common.Address
	Token      common.Address
	Symbol     string
	Magnitude  *big.Int
	Loadable   bool
	Redeemable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedToken(opts *bind.FilterOpts) (*TokenWhitelistAddedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedTokenIterator{contract: _TokenWhitelist.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedToken(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedToken) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedToken)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedToken", log); err != nil {
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

// ParseAddedToken is a log parse operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseAddedToken(log types.Log) (*TokenWhitelistAddedToken, error) {
	event := new(TokenWhitelistAddedToken)
	if err := _TokenWhitelist.contract.UnpackLog(event, "AddedToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the TokenWhitelist contract.
type TokenWhitelistClaimedIterator struct {
	Event *TokenWhitelistClaimed // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistClaimed)
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
		it.Event = new(TokenWhitelistClaimed)
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
func (it *TokenWhitelistClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistClaimed represents a Claimed event raised by the TokenWhitelist contract.
type TokenWhitelistClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterClaimed(opts *bind.FilterOpts) (*TokenWhitelistClaimedIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistClaimedIterator{contract: _TokenWhitelist.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *TokenWhitelistClaimed) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistClaimed)
				if err := _TokenWhitelist.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseClaimed(log types.Log) (*TokenWhitelistClaimed, error) {
	event := new(TokenWhitelistClaimed)
	if err := _TokenWhitelist.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistENSSetRegistryIterator is returned from FilterENSSetRegistry and is used to iterate over the raw logs and unpacked data for ENSSetRegistry events raised by the TokenWhitelist contract.
type TokenWhitelistENSSetRegistryIterator struct {
	Event *TokenWhitelistENSSetRegistry // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistENSSetRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistENSSetRegistry)
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
		it.Event = new(TokenWhitelistENSSetRegistry)
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
func (it *TokenWhitelistENSSetRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistENSSetRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistENSSetRegistry represents a ENSSetRegistry event raised by the TokenWhitelist contract.
type TokenWhitelistENSSetRegistry struct {
	EnsRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterENSSetRegistry is a free log retrieval operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterENSSetRegistry(opts *bind.FilterOpts) (*TokenWhitelistENSSetRegistryIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistENSSetRegistryIterator{contract: _TokenWhitelist.contract, event: "ENSSetRegistry", logs: logs, sub: sub}, nil
}

// WatchENSSetRegistry is a free log subscription operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchENSSetRegistry(opts *bind.WatchOpts, sink chan<- *TokenWhitelistENSSetRegistry) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "ENSSetRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistENSSetRegistry)
				if err := _TokenWhitelist.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
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

// ParseENSSetRegistry is a log parse operation binding the contract event 0x186a5507043ec82d7cac3080e4bf22b3f781d96b198adf165a2df507bda90773.
//
// Solidity: event ENSSetRegistry(address _ensRegistry)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseENSSetRegistry(log types.Log) (*TokenWhitelistENSSetRegistry, error) {
	event := new(TokenWhitelistENSSetRegistry)
	if err := _TokenWhitelist.contract.UnpackLog(event, "ENSSetRegistry", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistRemovedExclusiveMethodIterator is returned from FilterRemovedExclusiveMethod and is used to iterate over the raw logs and unpacked data for RemovedExclusiveMethod events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedExclusiveMethodIterator struct {
	Event *TokenWhitelistRemovedExclusiveMethod // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistRemovedExclusiveMethodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedExclusiveMethod)
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
		it.Event = new(TokenWhitelistRemovedExclusiveMethod)
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
func (it *TokenWhitelistRemovedExclusiveMethodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedExclusiveMethodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedExclusiveMethod represents a RemovedExclusiveMethod event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedExclusiveMethod struct {
	Token    common.Address
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovedExclusiveMethod is a free log retrieval operation binding the contract event 0xe01bc5ecc4d7ff06fdb26bad9a3601ef089d9e5aa6f7dd03dc713b468eec117a.
//
// Solidity: event RemovedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedExclusiveMethod(opts *bind.FilterOpts) (*TokenWhitelistRemovedExclusiveMethodIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedExclusiveMethodIterator{contract: _TokenWhitelist.contract, event: "RemovedExclusiveMethod", logs: logs, sub: sub}, nil
}

// WatchRemovedExclusiveMethod is a free log subscription operation binding the contract event 0xe01bc5ecc4d7ff06fdb26bad9a3601ef089d9e5aa6f7dd03dc713b468eec117a.
//
// Solidity: event RemovedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedExclusiveMethod(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedExclusiveMethod) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedExclusiveMethod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedExclusiveMethod)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedExclusiveMethod", log); err != nil {
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

// ParseRemovedExclusiveMethod is a log parse operation binding the contract event 0xe01bc5ecc4d7ff06fdb26bad9a3601ef089d9e5aa6f7dd03dc713b468eec117a.
//
// Solidity: event RemovedExclusiveMethod(address _token, bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseRemovedExclusiveMethod(log types.Log) (*TokenWhitelistRemovedExclusiveMethod, error) {
	event := new(TokenWhitelistRemovedExclusiveMethod)
	if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedExclusiveMethod", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistRemovedMethodIdIterator is returned from FilterRemovedMethodId and is used to iterate over the raw logs and unpacked data for RemovedMethodId events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedMethodIdIterator struct {
	Event *TokenWhitelistRemovedMethodId // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistRemovedMethodIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedMethodId)
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
		it.Event = new(TokenWhitelistRemovedMethodId)
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
func (it *TokenWhitelistRemovedMethodIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedMethodIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedMethodId represents a RemovedMethodId event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedMethodId struct {
	MethodId [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovedMethodId is a free log retrieval operation binding the contract event 0x006dd38caa262b48ea0824b897ee1c4f238521632ad2c5d12f3f0225a1378d1d.
//
// Solidity: event RemovedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedMethodId(opts *bind.FilterOpts) (*TokenWhitelistRemovedMethodIdIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedMethodId")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedMethodIdIterator{contract: _TokenWhitelist.contract, event: "RemovedMethodId", logs: logs, sub: sub}, nil
}

// WatchRemovedMethodId is a free log subscription operation binding the contract event 0x006dd38caa262b48ea0824b897ee1c4f238521632ad2c5d12f3f0225a1378d1d.
//
// Solidity: event RemovedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedMethodId(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedMethodId) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedMethodId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedMethodId)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedMethodId", log); err != nil {
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

// ParseRemovedMethodId is a log parse operation binding the contract event 0x006dd38caa262b48ea0824b897ee1c4f238521632ad2c5d12f3f0225a1378d1d.
//
// Solidity: event RemovedMethodId(bytes4 _methodId)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseRemovedMethodId(log types.Log) (*TokenWhitelistRemovedMethodId, error) {
	event := new(TokenWhitelistRemovedMethodId)
	if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedMethodId", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistRemovedTokenIterator is returned from FilterRemovedToken and is used to iterate over the raw logs and unpacked data for RemovedToken events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedTokenIterator struct {
	Event *TokenWhitelistRemovedToken // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistRemovedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedToken)
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
		it.Event = new(TokenWhitelistRemovedToken)
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
func (it *TokenWhitelistRemovedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedToken represents a RemovedToken event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedToken struct {
	Sender common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedToken is a free log retrieval operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedToken(opts *bind.FilterOpts) (*TokenWhitelistRemovedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedTokenIterator{contract: _TokenWhitelist.contract, event: "RemovedToken", logs: logs, sub: sub}, nil
}

// WatchRemovedToken is a free log subscription operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedToken(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedToken) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedToken)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedToken", log); err != nil {
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

// ParseRemovedToken is a log parse operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseRemovedToken(log types.Log) (*TokenWhitelistRemovedToken, error) {
	event := new(TokenWhitelistRemovedToken)
	if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistUpdatedTokenLoadableIterator is returned from FilterUpdatedTokenLoadable and is used to iterate over the raw logs and unpacked data for UpdatedTokenLoadable events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenLoadableIterator struct {
	Event *TokenWhitelistUpdatedTokenLoadable // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistUpdatedTokenLoadableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenLoadable)
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
		it.Event = new(TokenWhitelistUpdatedTokenLoadable)
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
func (it *TokenWhitelistUpdatedTokenLoadableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenLoadableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenLoadable represents a UpdatedTokenLoadable event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenLoadable struct {
	Sender   common.Address
	Token    common.Address
	Loadable bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenLoadable is a free log retrieval operation binding the contract event 0x0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec.
//
// Solidity: event UpdatedTokenLoadable(address _sender, address _token, bool _loadable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenLoadable(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenLoadableIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenLoadable")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenLoadableIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenLoadable", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenLoadable is a free log subscription operation binding the contract event 0x0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec.
//
// Solidity: event UpdatedTokenLoadable(address _sender, address _token, bool _loadable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenLoadable(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenLoadable) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenLoadable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenLoadable)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenLoadable", log); err != nil {
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

// ParseUpdatedTokenLoadable is a log parse operation binding the contract event 0x0e086282e8e406857ef1dce65e04a192ad8405e48484524cb2ddbf28e5d84eec.
//
// Solidity: event UpdatedTokenLoadable(address _sender, address _token, bool _loadable)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseUpdatedTokenLoadable(log types.Log) (*TokenWhitelistUpdatedTokenLoadable, error) {
	event := new(TokenWhitelistUpdatedTokenLoadable)
	if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenLoadable", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistUpdatedTokenRateIterator is returned from FilterUpdatedTokenRate and is used to iterate over the raw logs and unpacked data for UpdatedTokenRate events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRateIterator struct {
	Event *TokenWhitelistUpdatedTokenRate // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistUpdatedTokenRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenRate)
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
		it.Event = new(TokenWhitelistUpdatedTokenRate)
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
func (it *TokenWhitelistUpdatedTokenRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenRate represents a UpdatedTokenRate event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRate struct {
	Sender common.Address
	Token  common.Address
	Rate   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenRate is a free log retrieval operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenRate(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenRateIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenRate")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenRateIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenRate is a free log subscription operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenRate(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenRate) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenRate)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRate", log); err != nil {
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

// ParseUpdatedTokenRate is a log parse operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseUpdatedTokenRate(log types.Log) (*TokenWhitelistUpdatedTokenRate, error) {
	event := new(TokenWhitelistUpdatedTokenRate)
	if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistUpdatedTokenRedeemableIterator is returned from FilterUpdatedTokenRedeemable and is used to iterate over the raw logs and unpacked data for UpdatedTokenRedeemable events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRedeemableIterator struct {
	Event *TokenWhitelistUpdatedTokenRedeemable // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistUpdatedTokenRedeemableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenRedeemable)
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
		it.Event = new(TokenWhitelistUpdatedTokenRedeemable)
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
func (it *TokenWhitelistUpdatedTokenRedeemableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenRedeemableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenRedeemable represents a UpdatedTokenRedeemable event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRedeemable struct {
	Sender     common.Address
	Token      common.Address
	Redeemable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenRedeemable is a free log retrieval operation binding the contract event 0xcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d.
//
// Solidity: event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenRedeemable(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenRedeemableIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenRedeemable")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenRedeemableIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenRedeemable", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenRedeemable is a free log subscription operation binding the contract event 0xcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d.
//
// Solidity: event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenRedeemable(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenRedeemable) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenRedeemable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenRedeemable)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRedeemable", log); err != nil {
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

// ParseUpdatedTokenRedeemable is a log parse operation binding the contract event 0xcaa111d70d53608b9c8e3278c634595491de54f572a17a297dedad20f517039d.
//
// Solidity: event UpdatedTokenRedeemable(address _sender, address _token, bool _redeemable)
func (_TokenWhitelist *TokenWhitelistFilterer) ParseUpdatedTokenRedeemable(log types.Log) (*TokenWhitelistUpdatedTokenRedeemable, error) {
	event := new(TokenWhitelistUpdatedTokenRedeemable)
	if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRedeemable", log); err != nil {
		return nil, err
	}
	return event, nil
}
