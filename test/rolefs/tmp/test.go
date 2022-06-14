package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"math/big"
// 	callconts "memoc/callcontracts"
// 	"memoc/test"
// 	"time"

// 	"github.com/ethereum/go-ethereum/common"
// )

// var (
// 	ethEndPoint  string
// 	qethEndPoint string
// )

// //
// func main() {
// 	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
// 	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
// 	flag.Parse()
// 	ethEndPoint = *eth
// 	qethEndPoint = *qeth
// 	callconts.EndPoint = ethEndPoint

// 	// 用于测试的一些参数
// 	adminAddr := common.HexToAddress(test.AdminAddr)
// 	acc1Addr := common.HexToAddress(test.Acc1) // 用于注册User
// 	acc2Addr := common.HexToAddress(test.Acc2) // 用于注册provider
// 	acc3Addr := common.HexToAddress(test.Acc3) // 用于注册keeper1
// 	acc4Addr := common.HexToAddress(test.Acc4) // 用于注册keeper2
// 	accs := []common.Address{acc1Addr, acc2Addr, acc3Addr, acc4Addr}
// 	pledgeK := big.NewInt(1e6)
// 	pledgeP := big.NewInt(1e6)
// 	start := uint64(time.Now().Unix())
// 	fmt.Println("start:", start)
// 	end := 86400 - start%86400 + start + 100*86400
// 	size := uint64(1000000)
// 	nonce := uint64(0)
// 	sprice := big.NewInt(1e11) // 1e11 1e6 1e2 1e7
// 	addOrderpay := big.NewInt(0).Mul(big.NewInt(int64(end-start)), sprice)
// 	fmt.Println("addOrderPay:", addOrderpay)
// 	rechargeValue := big.NewInt(0).Mul(addOrderpay, big.NewInt(2)) // need not less than (end - start)*sprice + managePay + taxPay
// 	fmt.Println("rechargeValue:", rechargeValue)
// 	zero := big.NewInt(0)
// 	managePay := big.NewInt(1).Div(addOrderpay, big.NewInt(25))
// 	fmt.Println("managePay:", managePay)
// 	reduceBal := big.NewInt(1).Add(addOrderpay, big.NewInt(1).Div(addOrderpay, big.NewInt(20)))

// 	fmt.Println("============ 1. check acc balance of eth ============")

// 	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
// 	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
// 	fmt.Println("admin balance: ", bal, " in Ethereum")
// 	for i, acc := range accs {
// 		bal = callconts.QueryEthBalance(acc.Hex(), ethEndPoint)
// 		fmt.Println("acc", i, " balance: ", bal, " in Ethereum")
// 	}

// 	txopts := &callconts.TxOpts{
// 		Nonce:    nil,
// 		GasPrice: nil,
// 		GasLimit: callconts.DefaultGasLimit,
// 	}

// 	status := make(chan error)

// 	fmt.Println("============ 2. check acc balance of erc20 ============")

// 	// 查询在erc20代币上的余额
// 	erc20 := callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	bal, err := erc20.BalanceOf(adminAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("admin balance in primaryToken is ", bal)
// 	// 余额不足，自动充值
// 	mintValue := big.NewInt(1).Add(big.NewInt(test.MoneyTo), rechargeValue)
// 	if bal.Cmp(mintValue) < 0 {
// 		// mintToken
// 		err = erc20.MintToken(adminAddr, mintValue)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if err = <-status; err != nil {
// 			log.Fatal(err)
// 		}
// 		bal, err = erc20.BalanceOf(adminAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("after mint, admin balance in primaryToken is ", bal)
// 	}
// 	tNum := 0
// 	var errT error
// 	for i, acc := range accs {
// 		bal, err = erc20.BalanceOf(acc)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("acc", i, " balance in primaryToken is ", bal)
// 		if i == 0 {
// 			if bal.Cmp(rechargeValue) < 0 {
// 				err = erc20.Transfer(acc, rechargeValue)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				tNum++
// 			}
// 		} else if bal.Cmp(pledgeK) < 0 {
// 			err = erc20.Transfer(acc, pledgeK)
// 			if err != nil {
// 				for j := 0; j < tNum; j++ {
// 					<-status
// 				}
// 				log.Fatal(err)
// 			}
// 			tNum++
// 		}
// 	}
// 	for i := 0; i < tNum; i++ {
// 		err = <-status
// 		if err != nil {
// 			errT = err
// 		}
// 	}
// 	if errT != nil {
// 		log.Fatal(errT)
// 	}

// 	fmt.Println("============ 3. deploy RoleFS contract ============")

// 	// rfs caller for deploy
// 	rfs := callconts.NewRFS(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	rolefsAddr, _, err := rfs.DeployRoleFS()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("RoleFS contract Address:", rolefsAddr.Hex())

// 	fmt.Println("============ 4. deploy Role contract ============")

// 	// 部署Role合约
// 	r := callconts.NewR(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeK)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The Role contract address is ", roleAddr.Hex())

// 	// 获得RToken合约地址
// 	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	rTokenAddr, err := r.RToken()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The RToken contract address is ", rTokenAddr.Hex())

// 	fmt.Println("============ 5. deploy PledgePool contract ============")

// 	// 部署PledgePool合约
// 	pp := callconts.NewPledgePool(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	pledgePoolAddr, _, err := pp.DeployPledgePool(test.PrimaryToken, rTokenAddr, roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The PledgePool contract address is ", pledgePoolAddr.Hex())
// 	//pp = callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

// 	fmt.Println("============ 6. deploy Issuance contract ============")

// 	// 部署Issuance合约
// 	issu := callconts.NewIssu(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The Issuance contract address is ", issuanceAddr.Hex())

// 	fmt.Println("============ 7. call SetPI for Role ============")

// 	// role caller
// 	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	// setPI
// 	err = r.SetPI(pledgePoolAddr, issuanceAddr, rolefsAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 8. register acc for user and provider ============")

// 	// register user with acc1
// 	r = callconts.NewR(roleAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
// 	err = r.Register(acc1Addr, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	uIndex, err := r.GetRoleIndex(acc1Addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("uIndex ", uIndex)

// 	// register provider with acc2
// 	r = callconts.NewR(roleAddr, acc2Addr, test.Sk2, txopts, ethEndPoint, status)
// 	err = r.Register(acc2Addr, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	pIndex, err := r.GetRoleIndex(acc2Addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("pIndex ", pIndex)

// 	fmt.Println("============ 9. register acc for keepers ============")

// 	// register acc3
// 	r = callconts.NewR(roleAddr, acc3Addr, test.Sk3, txopts, ethEndPoint, status)
// 	err = r.Register(acc3Addr, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// register acc4
// 	r = callconts.NewR(roleAddr, acc4Addr, test.Sk4, txopts, ethEndPoint, status)
// 	err = r.Register(acc4Addr, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 10. keepers plege and register role ============")

// 	// acc3 pledge
// 	r = callconts.NewR(roleAddr, acc3Addr, test.Sk3, txopts, ethEndPoint, status)
// 	kIndex3, err := r.GetRoleIndex(acc3Addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("kIndex3 ", kIndex3)
// 	// pp caller
// 	pp = callconts.NewPledgePool(pledgePoolAddr, acc3Addr, test.Sk3, txopts, ethEndPoint, status)
// 	err = pp.Pledge(test.PrimaryToken, roleAddr, kIndex3, pledgeK, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// acc3 register role
// 	err = r.RegisterKeeper(pledgePoolAddr, kIndex3, []byte("test"), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// acc4 pledge
// 	r = callconts.NewR(roleAddr, acc4Addr, test.Sk4, txopts, ethEndPoint, status)
// 	kIndex4, err := r.GetRoleIndex(acc4Addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("kIndex4 ", kIndex4)
// 	// pp caller
// 	pp = callconts.NewPledgePool(pledgePoolAddr, acc4Addr, test.Sk4, txopts, ethEndPoint, status)
// 	err = pp.Pledge(test.PrimaryToken, roleAddr, kIndex4, pledgeK, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// acc4 register role
// 	err = r.RegisterKeeper(pledgePoolAddr, kIndex4, []byte("test"), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 11. createGroup and get fsAddr ============")

// 	// createGroup, and deploy FileSys contract
// 	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	gIndex, err := r.CreateGroup(rolefsAddr, 0, []uint64{kIndex3, kIndex4}, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("gIndex ", gIndex)

// 	// 获取FileSys合约地址
// 	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The FileSys contract address is ", fsAddr.Hex())
// 	fmt.Println("The issuance contract address is ", issuanceAddr.Hex())
// 	fmt.Println("The role contract address is ", roleAddr.Hex())
// 	fmt.Println("The rToken contract address is ", rTokenAddr.Hex())

// 	fmt.Println("============ 12. call setAddr for roleFS ============")

// 	rfs = callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	err = rfs.SetAddr(issuanceAddr, roleAddr, fsAddr, rTokenAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 13. user register role and recharge ============")

// 	r = callconts.NewR(roleAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
// 	err = r.RegisterUser(rTokenAddr, uIndex, gIndex, []byte("test"), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// user approve
// 	erc20 = callconts.NewERC20(test.PrimaryToken, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
// 	err = erc20.Approve(fsAddr, rechargeValue)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// 查询该User对FileSys合约的allowance
// 	allo, err := erc20.Allowance(acc1Addr, fsAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The allowance of User to FileSys contract is ", allo)
// 	if allo.Cmp(rechargeValue) != 0 {
// 		log.Fatal("allowance should be ", rechargeValue)
// 	}

// 	// user往FileSys中充值rechargeValue，用于存储服务付费
// 	fmt.Println("rechargeValue: ", rechargeValue)
// 	err = r.Recharge(rTokenAddr, uIndex, 0, rechargeValue, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// to assure recharge complete.
// 	time.Sleep(1 * time.Second)

// 	// user获取其在FileSys中的balance
// 	fs := callconts.NewFileSys(fsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	avail, _, err := fs.GetBalance(uIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("avail:", avail)
// 	if avail.Cmp(rechargeValue) < 0 {
// 		log.Fatal("user avail should not less than ", rechargeValue)
// 	}

// 	fmt.Println("============ 14. provider pledge and register role ============")

// 	// acc2 pledge
// 	pp = callconts.NewPledgePool(pledgePoolAddr, acc2Addr, test.Sk2, txopts, ethEndPoint, status)
// 	err = pp.Pledge(test.PrimaryToken, roleAddr, pIndex, pledgeP, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// provider register
// 	r = callconts.NewR(roleAddr, acc2Addr, test.Sk2, txopts, ethEndPoint, status)
// 	err = r.RegisterProvider(pledgePoolAddr, pIndex, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	// for activate provider
// 	err = r.AddProviderToGroup(pIndex, gIndex, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 15. grant minter role to rfs ============")

// 	//admin为RoleFS赋予Minter_Role权限
// 	erc20 = callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	err = erc20.SetUpRole(callconts.MinterRole, rolefsAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	time.Sleep(1 * time.Second)

// 	fmt.Println("============ 16. calc signatures for addOrder ============")

// 	// use primary token for testing
// 	tIndex := uint32(0)

// 	// user = acc1
// 	usig, err := callconts.SignForOrder(1, 2, nonce, start, end, size, tIndex, sprice, test.Sk1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// provide = acc2
// 	psig, err := callconts.SignForOrder(1, 2, nonce, start, end, size, tIndex, sprice, test.Sk2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 17. call addOrder ============")

// 	// rfs caller, called by user
// 	rfs = callconts.NewRFS(rolefsAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
// 	// ERC20Addr's value is used in checkParam
// 	callconts.ERC20Addr = test.PrimaryToken
// 	// call AddOrder
// 	err = rfs.AddOrder(
// 		roleAddr,   // role address
// 		rTokenAddr, // rtoken address
// 		1,          // uID
// 		2,          // pID
// 		start,      // start
// 		end,        // end
// 		size,       // size
// 		nonce,      // nonce
// 		0,          // tIndex
// 		sprice,     // price
// 		usig,       // usig
// 		psig,       // psig
// 		//ksigs,      // ksigs
// 	)
// 	if err != nil {
// 		log.Print("call add order err:", err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	time.Sleep(3 * time.Second)

// 	fmt.Println("============ 18. test addOrder ============")

// 	fmt.Println("check provider sum")
// 	// check provider sum
// 	pSum, err := fs.GetFsProviderSum(uIndex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if pSum != 1 {
// 		fmt.Println("pSum ", pSum)
// 		log.Fatal("The pSum should be 1")
// 	}
// 	pIndexGotten, err := fs.GetFsProvider(uIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if pIndexGotten != pIndex {
// 		fmt.Println("pIndexGotten ", pIndexGotten)
// 		log.Fatal("The pIndex gotten should be ", pIndex)
// 	}
// 	fmt.Println("check getChannelInfo")
// 	// check getChannelInfo
// 	_amount, _nonce, _expire, err := fs.GetChannelInfo(uIndex, pIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: channel amount", _amount, ",nonce", _nonce, ",expire", _expire)
// 	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != end {
// 		log.Fatal("result error")
// 	}
// 	fmt.Println("check getStoreInfo")
// 	// check getStoreInfo
// 	_time, _size, _price, err := fs.GetStoreInfo(uIndex, pIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: storeinfo time", _time, ",size", _size, ",price", _price)
// 	if _price.Cmp(sprice) != 0 || _size != size || _time != end {
// 		log.Fatal("price in storeInfo after addOrder should be ", sprice, ", size in storeInfo should be ", size, ", time should be ", end)
// 	}
// 	// check getSettleInfo
// 	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(pIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: settleinfo time", _time, ",size", _size, ",price", _price, ",maxPay", _maxPay, ",hasPaid", _hasPaid, ",canPay", _canPay, ",lost", _lost, ",lostPaid", _lostPaid, ",managePay", _managePay, ",endPaid", _endPaid, ",linearPaid", _linearPaid)
// 	if _time != start {
// 		fmt.Println("time:", _time)
// 		log.Fatal("time should be ", start)
// 	}
// 	if _size != size {
// 		fmt.Println("size:", _size)
// 		log.Fatal("sieze should be ", size)
// 	}
// 	if _price.Cmp(sprice) != 0 {
// 		fmt.Println("price:", _price)
// 		log.Fatal("price should be ", sprice)
// 	}
// 	if _maxPay.Cmp(addOrderpay) != 0 {
// 		fmt.Println("maxPay:", _maxPay)
// 		log.Fatal("maxPay should be ", addOrderpay)
// 	}
// 	if _managePay.Cmp(managePay) != 0 {
// 		fmt.Println("managePay:", _managePay)
// 		log.Fatal("managePay should be ", managePay)
// 	}
// 	fmt.Println("check getFsInfoAggOrder")
// 	// check getFsInfoAggOrder
// 	_nonce, _subNonce, err := fs.GetFsInfoAggOrder(uIndex, pIndex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: aggOrderInfo nonce ", _nonce, ",subNonce:", _subNonce)
// 	if _nonce != 1 || _subNonce != 0 {
// 		log.Fatal("nonce in aggOrder should be ", 1)
// 	}
// 	fmt.Println("check balance of foundation")
// 	// check balance of foundation
// 	_availF, _tmpF, err := fs.GetBalance(0, 0) // 查询基金会在FileSys中的balance
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: foundation avail ", _availF, ",tmp ", _tmpF)
// 	if _availF.Cmp(big.NewInt(1).Div(addOrderpay, big.NewInt(100))) != 0 || _tmpF.Cmp(zero) != 0 {
// 		log.Fatal("availF should be ", 1, " tmpF should be ", _tmpF)
// 	}
// 	fmt.Println("check balance of user")
// 	// check balance of user
// 	_avail, _tmp, err := fs.GetBalance(uIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: user avail ", _avail, ",tmp ", _tmp)
// 	if _avail.Sub(avail, _avail).Cmp(reduceBal) != 0 {
// 		log.Fatal("the new avail should be", reduceBal, "less than avail")
// 	}
// 	fmt.Println("check group info")
// 	// check group info
// 	_isActive, _isBanned, _isReady, _level, sizeG, priceG, _fsAddr, err := r.GetGroupInfo(gIndex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("After addOrder: groupInfo isActive ", _isActive, ",isBanned ", _isBanned, ",isReady ", _isReady, ",level ", _level, ",size ", sizeG, ",price ", priceG, ",fsAddr ", _fsAddr.Hex())
// 	if sizeG.Cmp(big.NewInt(int64(size))) != 0 || priceG.Cmp(sprice) != 0 {
// 		log.Fatal("the sizeG should be ", size, " priceG should be ", sprice)
// 	}
// 	fmt.Println("check mint info")
// 	// check mint info (初次addOrder，mintToken值为0)
// 	issu = callconts.NewIssu(issuanceAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
// 	mintLevel, err := issu.MintLevel()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("mintLevel:", mintLevel)
// 	if mintLevel.Cmp(zero) != 0 {
// 		log.Fatal("mintLevel should be 0")
// 	}
// 	st, err := issu.SpaceTime()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("spaceTime:", st)
// 	if st.Cmp(big.NewInt(0).Mul(new(big.Int).SetUint64(size), new(big.Int).SetUint64(end-start))) != 0 {
// 		log.Fatal("st should be ", big.NewInt(0).Mul(new(big.Int).SetUint64(size), new(big.Int).SetUint64(end-start)))
// 	}
// 	//size:10, price:10, totalPay:100,上面测试通过的话，这里就不需要再重复测试了

// 	fmt.Println("============ 19. calc signatures for sub Order ============")

// 	// user = acc1
// 	usig, err = callconts.SignForOrder(1, 2, nonce, start+5, start+10, size-5, tIndex, big.NewInt(0).Sub(sprice, big.NewInt(5)), test.Sk1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// provide = acc2
// 	psig, err = callconts.SignForOrder(1, 2, nonce, start+5, start+10, size-5, tIndex, big.NewInt(0).Sub(sprice, big.NewInt(5)), test.Sk2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	/*
// 		// keeper = acc3, acc4
// 		ksig3, err = callconts.SignForAddOrder(1, 2, nonce, start+5, start+10, size-5, big.NewInt(0).Sub(sprice, big.NewInt(5)), test.Sk3)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		ksig4, err = callconts.SignForAddOrder(1, 2, nonce, start+5, start+10, size-5, big.NewInt(0).Sub(sprice, big.NewInt(5)), test.Sk4)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		ksigs = [][]byte{ksig3, ksig4}
// 	*/

// 	fmt.Println("============20. call SubOrder============")

// 	// rfs caller, called by user
// 	rfs = callconts.NewRFS(rolefsAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
// 	err = rfs.SubOrder(
// 		roleAddr,
// 		rTokenAddr,
// 		1, // uIndex
// 		2, // pIndex
// 		start+5,
// 		start+10,
// 		size-5,
// 		nonce,
// 		0,                                        // tIndex
// 		big.NewInt(0).Sub(sprice, big.NewInt(5)), // sprice
// 		usig,                                     // usig
// 		psig,                                     // psig
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============21. test SubOrder============")

// 	// 这里不重复测试了，仅抽取部分信息进行测试
// 	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(pIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("after SubOrder: time", _time, ", size", _size, ", price", _price, ", maxPay", _maxPay, ", hasPaid", _hasPaid, ", canPay", _canPay, ", lost", _lost, ", lostPaid", _lostPaid, ", managePay", _managePay, ", endPaid", _endPaid, ", linearPaid", _linearPaid)
// 	if _time != start+10 {
// 		log.Fatal("time should be ", start+10)
// 	}
// 	if _size != 5 {
// 		log.Fatal("size should be ", 5)
// 	}
// 	if _price.Cmp(big.NewInt(5)) != 0 {
// 		log.Fatal("price should be 5")
// 	}
// 	if _maxPay.Cmp(addOrderpay) != 0 {
// 		log.Fatal("managePay should be", addOrderpay)
// 	}
// 	tmp := big.NewInt(0).Mul(sprice, big.NewInt(10)) //sprice*(newEnd - oldStart), newEnd是上面调用subOrder的start+10
// 	if _canPay.Cmp(tmp) != 0 {
// 		log.Fatal("canPay should be ", tmp)
// 	}

// 	fmt.Println("============ 22. calc signatures for proWithdraw ============")

// 	pay := big.NewInt(20)
// 	lost := big.NewInt(10)

// 	// keeper = acc3, acc4
// 	ksig3, err := callconts.SignForProWithdraw(2, 0, pay, lost, test.Sk3)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ksig4, err := callconts.SignForProWithdraw(2, 0, pay, lost, test.Sk4)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ksigs := [][]byte{ksig3, ksig4}
// 	kIndexes := []uint64{3, 4}

// 	fmt.Println("============ 23. call proWithdraw============")

// 	fmt.Println("pay: ", pay.String())
// 	fmt.Println("lost: ", lost.String())

// 	// rfs caller, called by user
// 	rfs = callconts.NewRFS(rolefsAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
// 	err = rfs.ProWithdraw(
// 		roleAddr,
// 		rTokenAddr,
// 		2,        // pIndex
// 		0,        // tIndex
// 		pay,      // pay
// 		lost,     // lost
// 		kIndexes, // keeper indexes
// 		ksigs,    // keeper signatures
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = <-status; err != nil {
// 		log.Fatal(err)
// 	}
// 	time.Sleep(3 * time.Second)

// 	fmt.Println("============ 24. test proWithdraw============")

// 	fmt.Println("pIndex: ", pIndex)
// 	// 获取settlement信息，并判断正确性
// 	_time2, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(pIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("after ProWithdraw: time", _time2, ", size", _size, ", price", _price, ", maxPay", _maxPay, ", hasPaid", _hasPaid, ", canPay", _canPay, ", lost", _lost, ", lostPaid", _lostPaid, ", managePay", _managePay, ", endPaid", _endPaid, ", linearPaid", _linearPaid)
// 	if _size != 5 {
// 		log.Fatal("size should be ", 5)
// 	}
// 	if _hasPaid.Cmp(pay) != 0 {
// 		log.Fatal("hasPaid should be ", pay)
// 	}
// 	thisCanPay := big.NewInt(0).Mul(new(big.Int).SetUint64(_time2-_time), _price)
// 	tmp2 := big.NewInt(0).Add(tmp, thisCanPay)
// 	if _canPay.Cmp(tmp2) != 0 {
// 		log.Fatal("canPay should be ", tmp2)
// 	}
// 	if _lost.Cmp(lost) != 0 {
// 		log.Fatal("lost should be ", lost)
// 	}
// 	// 获取balance信息
// 	_avail, _tmp, err = fs.GetBalance(pIndex, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("after ProWithdraw, provider avail", _avail, ", tmp", _tmp)
// 	if _avail.Cmp(zero) != 0 {
// 		log.Fatal("avail should be 0")
// 	}
// 	if _tmp.Cmp(big.NewInt(0).Sub(_canPay, _hasPaid)) < 0 {
// 		log.Fatal("tmp should not be less than", big.NewInt(0).Sub(_canPay, _hasPaid))
// 	}
// 	// 获取provider在tIndex上的代币余额
// 	bal, err = erc20.BalanceOf(acc3Addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("after ProWithdraw, provider balance is ", bal)

// 	fmt.Println("============ All tests passed ============")

// }
