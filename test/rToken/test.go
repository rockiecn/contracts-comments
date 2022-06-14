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

func main() {
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	pledgeKP := big.NewInt(1e6)

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	status := make(chan error)

	fmt.Println("============1. begin test deploy RToken contract ============")
	// 注意：RToken合约是由Role合约部署的.在Role合约被admin部署时，Role合约通过create2创建RToken合约
	r := callconts.NewR(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeKP, pledgeKP, 1)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Role contract address is ", roleAddr)
	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	rTokenAddr, err := r.RToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RToken contract address is ", rTokenAddr.Hex())

	fmt.Println("============2. begin test GetTA ============")
	rt := callconts.NewRT(rTokenAddr, adminAddr, test.AdminSk, txopts, ethEndPoint)
	// RToken合约在被部署时，primaryToken地址就被记录到了RToken合约中，并且其tokenIndex为0
	tAddr, err := rt.GetTA(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The primaryToken address is ", tAddr.Hex())
	if tAddr.Hex() != test.PrimaryToken.Hex() {
		log.Fatal("gotten primaryToken address is wrong")
	}
	// 使用无效的tokenIndex进行测试
	tAddr, err = rt.GetTA(1)
	if err != nil {
		log.Fatal(err)
	}
	if tAddr.Hex() != callconts.InvalidAddr {
		log.Fatal("gotten token address should be zero by invalid tokenIndex")
	}

	fmt.Println("============3. begin test IsValid ============")
	isValid, err := rt.IsValid(0)
	if err != nil {
		log.Fatal(err)
	}
	if !isValid {
		log.Fatal("the tokenIndex 0 should be valid")
	}
	isValid, err = rt.IsValid(1)
	if err != nil {
		log.Fatal(err)
	}
	if isValid {
		log.Fatal("The tokenIndex 1 should be invalid")
	}

	fmt.Println("============4. begin test GetTI ============")
	tIndex, isValid, err := rt.GetTI(test.PrimaryToken)
	if err != nil {
		log.Fatal(err)
	}
	if tIndex != 0 {
		log.Fatal("gotten tokenIndex is wrong")
	}
	if !isValid {
		log.Fatal("primaryToken should be valid")
	}
	// 错误场景测试
	tIndex, isValid, err = rt.GetTI(adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	if tIndex != 0 || isValid {
		log.Fatal("the test tokenAddr should be invalid and tIndex should be zero")
	}

	fmt.Println("============5. begin test GetTNum ============")
	tNum, err := rt.GetTNum()
	if err != nil {
		log.Fatal(err)
	}
	if tNum != 1 {
		log.Fatal("tNum should be 1")
	}

	fmt.Println("============test success!============")
}
