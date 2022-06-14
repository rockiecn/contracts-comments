package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	callconts "memoc/callcontracts"
	"memoc/test"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ethEndPoint  string
	qethEndPoint string
)

// 仍然需要通过调用rolefs合约中的addOrder等函数去触发代币发行，从而触发PledgePool合约中的分润机制，之后再次测试getter类函数
func main() {
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	_ = qethEndPoint
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	roleAddr := common.HexToAddress("0xA2f493B23a23A39d3f1072f3961dbAe53EAe384a")
	pledgeMoney := big.NewInt(1e6)

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: nil,
		GasLimit: callconts.DefaultGasLimit,
	}

	// 查询在erc20代币上的余额
	status := make(chan error)
	erc20 := callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	bal, err := erc20.BalanceOf(adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("admin balance in primaryToken is ", bal)

	pp := callconts.NewPledgePool(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

	fmt.Println("============1. begin test deploy PledgePool contract============")
	ppAddr, _, err := pp.DeployPledgePool(test.PrimaryToken, test.RTokenAddr, roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PledgePool contract address is ", ppAddr.Hex())

	fmt.Println("============2. begin test Pledge============")
	// 首先需approve
	err = erc20.Approve(ppAddr, pledgeMoney)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 然后质押账户需要在Role合约中Register从而获得rindex
	r := callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	err = r.Register(adminAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	rIndex, err := r.GetRoleIndex(adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The account's role index is ", rIndex)
	
	// 开始pledge
	pp = callconts.NewPledgePool(ppAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	err = pp.Pledge(test.PrimaryToken, roleAddr, rIndex, pledgeMoney, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	fmt.Println("============3. begin test GetPledge============")
	p, err := pp.GetPledge(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total pledge value of token0 is ", p)
	if p.Cmp(pledgeMoney) != 0 {
		log.Fatal("pledge value should be ", pledgeMoney)
	}

	fmt.Println("============4. begin test GetBalanceInPPool============")
	p, err = pp.GetBalanceInPPool(rIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The balance value of token0 in PledgePool contract is ", p)
	if p.Cmp(pledgeMoney) != 0 {
		log.Fatal("balance value in PledgePool should be ", pledgeMoney)
	}

	fmt.Println("============5. begin test TotalPledge============")
	p, err = pp.TotalPledge()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total pledge of primaryToken is ", p)
	if p.Cmp(pledgeMoney) != 0 {
		log.Fatal("totalPledge should be ", pledgeMoney)
	}

	fmt.Println("============6. begin test Withdraw============")
	_addrNum, err := r.GetAddrsNum()
	if err != nil {
		log.Fatal("get addrNum failed: ", err)
	}
	fmt.Println("addrNum: ", _addrNum)

	err = pp.Withdraw(roleAddr, test.RTokenAddr, rIndex, 0, pledgeMoney, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	p, err = pp.GetBalanceInPPool(rIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After Withdraw, the balance value of token0 in PledgePool contract is ", p)
	if p.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("After Withdraw, the balance in PledgePool should be 0")
	}

	fmt.Println("============test success!============")
}
