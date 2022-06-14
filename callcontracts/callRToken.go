package callconts

import (
	"memoc/contracts/role"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NewRT new a instance of ContractModule
func NewRT(rTokenAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string) iface.RTokenInfo {
	rt := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: rTokenAddr,
		endPoint:        endPoint,
	}

	return rt
}

// IsValid check whether the tokenIndex is valid, rTokenAddr indicates RToken contract address, get it by RToken() in callRole.go
func (rt *ContractModule) IsValid(tIndex uint32) (bool, error) {
	var isvalid bool
	client := getClient(rt.endPoint)
	defer client.Close()
	rToken, err := role.NewRToken(rt.contractAddress, client)
	if err != nil {
		return false, err
	}

	retryCount := 0
	for {
		retryCount++
		isvalid, err = rToken.IsValid(&bind.CallOpts{
			From: rt.addr,
		}, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return false, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isvalid, nil
	}
}

// GetTA get the address of tokenIndex. Return (0x0, nil) if the tIndex is invalid
func (rt *ContractModule) GetTA(tIndex uint32) (common.Address, error) {
	var taddr common.Address

	client := getClient(rt.endPoint)
	defer client.Close()
	rToken, err := role.NewRToken(rt.contractAddress, client)
	if err != nil {
		return taddr, err
	}

	retryCount := 0
	for {
		retryCount++
		taddr, err = rToken.GetTA(&bind.CallOpts{
			From: rt.addr,
		}, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return taddr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return taddr, nil
	}
}

// GetTI get the tokenIndex by token address and if it is valid
func (rt *ContractModule) GetTI(taddr common.Address) (uint32, bool, error) {
	var tindex uint32
	var isvalid bool

	client := getClient(rt.endPoint)
	defer client.Close()
	rToken, err := role.NewRToken(rt.contractAddress, client)
	if err != nil {
		return tindex, isvalid, err
	}

	retryCount := 0
	for {
		retryCount++
		tindex, isvalid, err = rToken.GetTI(&bind.CallOpts{
			From: rt.addr,
		}, taddr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return tindex, isvalid, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return tindex, isvalid, nil
	}
}

// GetTNum get the number of tokens that memo supports
func (rt *ContractModule) GetTNum() (uint32, error) {
	var tnum uint32

	client := getClient(rt.endPoint)
	defer client.Close()
	rToken, err := role.NewRToken(rt.contractAddress, client)
	if err != nil {
		return tnum, err
	}

	retryCount := 0
	for {
		retryCount++
		tnum, err = rToken.GetTNum(&bind.CallOpts{
			From: rt.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return tnum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return tnum, nil
	}
}
