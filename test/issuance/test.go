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

const (
	PeriodTargetAmount = "50000000000000000000000000000"
	IssuRatio          = 50
	MinRatio           = 3
)

// 仍然需要通过调用rolefs合约中的addOrder等函数，从而触发代币发行、触发Issuance合约中的参数被更改，再次测试getter类函数
func main() {
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	rolefsAddr := common.HexToAddress("0xAAaC6D27153BF52d66Eed127e0321372B2FFF67C") // 测试期间部署的RoleFS合约地址

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: nil,
		GasLimit: callconts.DefaultGasLimit,
	}

	fmt.Println("============1. begin test deploy Issuance contract============")
	status := make(chan error)
	issu := callconts.NewIssu(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	issuAddr, _, err := issu.DeployIssuance(rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Issuance contract address is ", issuAddr.Hex())

	fmt.Println("============2. begin test MintLevel============")
	issu = callconts.NewIssu(issuAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	mintLevel, err := issu.MintLevel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The mintLevel is ", mintLevel)
	if mintLevel.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("mintlevel should be 0")
	}

	fmt.Println("============3. begin test LastMint============")
	lastMint, err := issu.LastMint()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The lastMint is ", lastMint)

	fmt.Println("============4. begin test Price============")
	price, err := issu.Price()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The price is ", price)
	if price.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("price should be 0")
	}

	fmt.Println("============5. begin test Size============")
	size, err := issu.Size()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The size is ", size)
	if size.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("size should be 0")
	}

	fmt.Println("============6. begin test SpaceTime============")
	spaceTime, err := issu.SpaceTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The spaceTime is ", spaceTime)
	if spaceTime.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("spacetime should be 0")
	}

	fmt.Println("============7. begin test TotalPay============")
	totalPay, err := issu.TotalPay()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The totalPay is ", totalPay)
	if totalPay.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("totalPay should be 0")
	}

	fmt.Println("============8. begin test TotalPaid============")
	totalPaid, err := issu.TotalPaid()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The totalPaid is ", totalPaid)
	if totalPaid.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("totalPaid should be 0")
	}

	fmt.Println("============9. begin test PeriodTarget============")
	periodTarget, err := issu.PeriodTarget()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The periodTarget is ", periodTarget)
	targetAmount, ok := new(big.Int).SetString(PeriodTargetAmount, 10)
	if !ok {
		log.Fatal("transform string to bigInt fails")
	}
	if periodTarget.Cmp(targetAmount) != 0 {
		log.Fatal("periodTarget should be ", targetAmount)
	}

	fmt.Println("============10. begin test PeriodTotalReward============")
	totalReward, err := issu.PeriodTotalReward()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The periodTotalReward is ", totalReward)
	if totalReward.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("periodTotalReward should be 0")
	}

	fmt.Println("============11. begin test IssuRatio============")
	ir, err := issu.IssuRatio()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The issuRatio is ", ir)
	if ir != IssuRatio {
		log.Fatal("issuRatio should be ", IssuRatio)
	}

	fmt.Println("============12. begin test MinRatio============")
	mr, err := issu.MinRatio()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The minRatio is ", mr)
	if mr != MinRatio {
		log.Fatal("minRatio should be ", MinRatio)
	}

	fmt.Println("============test success!============")
}
