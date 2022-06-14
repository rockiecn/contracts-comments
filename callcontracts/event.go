package callconts

import (
	"log"
	"math/big"
	"memoc/contracts/pledgepool"
	"memoc/contracts/role"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/xerrors"
)

// get gIndex from logs in receipt
func getGIndexFromRLogs(endPoint string, hash common.Hash) (uint64, error) {
	receipt := getTransactionReceipt(endPoint, hash)

	if len(receipt.Logs) != 1 {
		return 0, xerrors.New("length of logs in receipt is error")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, err
	}

	intr, err := contractAbi.Events["CreateGroup"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, err
	}
	return intr[0].(uint64), nil
}

// get RUser event info from tx hash
func getRUserInfoFromRLogs(endPoint string, hash common.Hash) (uint64, common.Address, error) {
	receipt := getTransactionReceipt(endPoint, hash)

	if len(receipt.Logs) != 1 {
		return 0, common.Address{}, xerrors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, common.Address{}, err
	}

	// only 1 log in this tx
	intr, err := contractAbi.Events["RUser"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, common.Address{}, err
	}

	// index, address
	return intr[0].(uint64), intr[1].(common.Address), nil
}

// get RKeeper event info from tx hash
func getRKeeperInfoFromRLogs(endPoint string, hash common.Hash) (uint64, common.Address, error) {
	receipt := getTransactionReceipt(endPoint, hash)

	if len(receipt.Logs) != 1 {
		return 0, common.Address{}, xerrors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, common.Address{}, err
	}

	// only 1 log in this tx
	intr, err := contractAbi.Events["RKeeper"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, common.Address{}, err
	}

	// index, address
	return intr[0].(uint64), intr[1].(common.Address), nil
}

// get RProvider event info from tx hash
func getRProviderInfoFromRLogs(endPoint string, hash common.Hash) (uint64, common.Address, error) {
	receipt := getTransactionReceipt(endPoint, hash)

	if len(receipt.Logs) != 1 {
		return 0, common.Address{}, xerrors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, common.Address{}, err
	}

	// only 1 log in this tx
	intr, err := contractAbi.Events["RProvider"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, common.Address{}, err
	}

	// index, address
	return intr[0].(uint64), intr[1].(common.Address), nil
}

// get Pledge event info in PledgePool contract from tx hash
func getPledgeInfoFromRLogs(endPoint string, hash common.Hash) (common.Address, *big.Int, error) {
	receipt := getTransactionReceipt(endPoint, hash)

	log.Println("tx logs count: ", len(receipt.Logs))

	if len(receipt.Logs) != 3 {
		return common.Address{}, nil, xerrors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(pledgepool.PledgePoolABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return common.Address{}, nil, err
	}

	// 3 logs in this tx, event Pledge is the last one
	// get data
	d, err := contractAbi.Events["Pledge"].Inputs.UnpackValues(receipt.Logs[2].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return common.Address{}, nil, err
	}
	log.Println("data:", d)

	// actual topic start from index 1, and 0 is keccack256 of event signature: Pledge(address,uint256)
	t := receipt.Logs[2].Topics[1].String()
	log.Println("topic:", t)

	m, ok := d[0].(*big.Int)
	if !ok {
		log.Println("money type assertion failed")
	}

	// address, money
	return common.HexToAddress(t), m, nil
}

// get Withdraw event info in PledgePool contract from tx hash
func getWithdrawInfoFromRLogs(endPoint string, hash common.Hash) (common.Address, *big.Int, error) {
	receipt := getTransactionReceipt(endPoint, hash)

	log.Println("tx logs count: ", len(receipt.Logs))

	if len(receipt.Logs) != 2 {
		return common.Address{}, nil, xerrors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(pledgepool.PledgePoolABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return common.Address{}, nil, err
	}

	// 2 logs in this tx, event Withdraw is the last one
	// get data
	d, err := contractAbi.Events["Withdraw"].Inputs.UnpackValues(receipt.Logs[1].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return common.Address{}, nil, err
	}
	log.Println("data:", d)

	// actual topic start from index 1, and 0 is keccack256 of event signature: Pledge(address,uint256)
	t := receipt.Logs[1].Topics[1].String()
	log.Println("topic:", t)

	m, ok := d[0].(*big.Int)
	if !ok {
		log.Println("money type assertion failed")
	}

	// address, money
	return common.HexToAddress(t), m, nil
}
