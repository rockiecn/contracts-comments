package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	callconts "memoc/callcontracts"
	"memoc/test"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ethEndPoint string
	//qethEndPoint string
)

// 该测试需花费约14分钟
func main() {
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;") //dev网
	flag.Parse()
	ethEndPoint = *eth
	//qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	acc1Addr := common.HexToAddress(test.Acc1)   // User
	acc2Addr := common.HexToAddress(test.Acc2)   // keeper
	acc3Addr := common.HexToAddress(test.Acc3)   // provider
	acc4Addr := common.HexToAddress(test.Acc4)   // provider
	acc5Addr := common.HexToAddress(test.Acc5)   // keeper
	acc6Addr := common.HexToAddress(test.Acc6)   // keeper
	acc7Addr := common.HexToAddress(test.Acc7)   // keeper
	acc8Addr := common.HexToAddress(test.Acc8)   // keeper
	acc9Addr := common.HexToAddress(test.Acc9)   // keeper
	acc10Addr := common.HexToAddress(test.Acc10) // keeper
	// put all accounts into slice
	accs := []common.Address{
		acc1Addr,  // user
		acc2Addr,  // keeper
		acc3Addr,  // provider
		acc4Addr,  // provider
		acc5Addr,  // keeper
		acc6Addr,  // keeper
		acc7Addr,  // keeper
		acc8Addr,  // keeper
		acc9Addr,  // keeper
		acc10Addr, // keeper
	}

	pledgeK := big.NewInt(1e6)
	start := uint64(time.Now().Unix())
	fmt.Println("start:", start)
	end := 86400 - start%86400 + start + 100*86400
	fmt.Println("end:", end)
	size := uint64(1000000)
	nonce := uint64(0)
	sprice := big.NewInt(1e11) // 1e11 1e6 1e2 1e7
	addOrderpay := big.NewInt(0).Mul(big.NewInt(int64(end-start)), sprice)
	fmt.Println("addOrderPay:", addOrderpay)
	rechargeValue := big.NewInt(0).Mul(addOrderpay, big.NewInt(2)) // need not less than (end - start)*sprice + managePay + taxPay
	fmt.Println("rechargeValue:", rechargeValue)
	zero := big.NewInt(0)
	managePay := big.NewInt(1).Div(addOrderpay, big.NewInt(25))
	fmt.Println("managePay:", managePay)
	reduceBal := big.NewInt(1).Add(addOrderpay, big.NewInt(1).Div(addOrderpay, big.NewInt(20)))

	//start := 1639911209
	//uIndex := uint64(2)
	//gIndex := uint64(1)
	//kIndex := uint64(1)
	//pIndex := uint64(3)
	//p2Index := uint64(4)
	//rolefsAddr := common.HexToAddress("0x506703d5e2126003944B707bDb10c1030F891e15")
	//roleAddr := common.HexToAddress("0x05B2A72D0045aF7A5dD574c138a28528f46A6E33")
	//rTokenAddr := common.HexToAddress("0xAEAFe15CBd5e6f7788779330Ad4A152E8ddB06eF")
	//pledgePoolAddr := common.HexToAddress("0xcfcBCac22D55e1FFeD9887158E642Fde09924071")
	//issuanceAddr := common.HexToAddress("0x5Be65871560ee3038aE99c9D9d2b0c23AEe381ab")
	//fsAddr := common.HexToAddress("0x44f813AFaAA5e65c8D06Be806C2DeB2E2335Bd76")

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")
	for i, acc := range accs {
		bal = callconts.QueryEthBalance(acc.Hex(), ethEndPoint)
		fmt.Println("acc", i, " balance: ", bal, " in Ethereum")
	}

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: nil,
		GasLimit: callconts.DefaultGasLimit,
	}

	status := make(chan error)

	// 查询在erc20代币上的余额
	erc20 := callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	bal, err := erc20.BalanceOf(adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("admin balance in primaryToken is ", bal)

	tNum := 0
	var errT error
	for i, acc := range accs {
		bal, err = erc20.BalanceOf(acc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc", i, " balance in primaryToken is ", bal)
		if i == 0 {
			if bal.Cmp(new(big.Int).Mul(rechargeValue, big.NewInt(2))) < 0 {
				err = erc20.Transfer(acc, new(big.Int).Mul(rechargeValue, big.NewInt(2)))
				if err != nil {
					log.Fatal(err)
				}
				tNum++
			}
		} else if bal.Cmp(pledgeK) < 0 {
			err = erc20.Transfer(acc, pledgeK)
			if err != nil {
				for j := 0; j < tNum; j++ {
					<-status
				}
				log.Fatal(err)
			}
			tNum++
		}
	}
	for i := 0; i < tNum; i++ {
		err = <-status
		if err != nil {
			errT = err
		}
	}
	if errT != nil {
		log.Fatal(errT)
	}

	rfs := callconts.NewRFS(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

	fmt.Println("============1. begin test deploy RoleFS contract============")
	rolefsAddr, _, err := rfs.DeployRoleFS()
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("RoleFS contract Address:", rolefsAddr.Hex())
	rfs = callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

	// 部署RoleFS合约之后，需要调用SetAddr函数赋予合约需要的状态变量信息
	fmt.Println("============2. begin test SetAddr============")
	// 部署Role合约
	r := callconts.NewR(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeK, 1)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Role contract address is ", roleAddr.Hex())
	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	// 获得RToken合约地址
	rTokenAddr, err := r.RToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RToken contract address is ", rTokenAddr.Hex())
	// 部署PledgePool合约
	pp := callconts.NewPledgePool(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	pledgePoolAddr, _, err := pp.DeployPledgePool(test.PrimaryToken, rTokenAddr, roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PledgePool contract address is ", pledgePoolAddr.Hex())
	//pp = callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	// 部署Issuance合约
	issu := callconts.NewIssu(adminAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Issuance contract address is ", issuanceAddr.Hex())
	issu = callconts.NewIssu(issuanceAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	// setPI
	err = r.SetPI(pledgePoolAddr, issuanceAddr, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	// keeper注册、质押、申请角色

	fmt.Println("=== begin register acc for keepers ===")

	// acc2
	registerAcc(roleAddr, acc2Addr, test.Sk2, txopts, ethEndPoint, status)
	// acc5
	registerAcc(roleAddr, acc5Addr, test.Sk5, txopts, ethEndPoint, status)
	// acc6
	registerAcc(roleAddr, acc6Addr, test.Sk6, txopts, ethEndPoint, status)
	// acc7
	registerAcc(roleAddr, acc7Addr, test.Sk7, txopts, ethEndPoint, status)
	// acc8
	registerAcc(roleAddr, acc8Addr, test.Sk8, txopts, ethEndPoint, status)
	// acc9
	registerAcc(roleAddr, acc9Addr, test.Sk9, txopts, ethEndPoint, status)
	// acc10
	registerAcc(roleAddr, acc10Addr, test.Sk10, txopts, ethEndPoint, status)

	time.Sleep(1 * time.Second)

	// register keepers

	fmt.Println("=== begin register keepers ===")

	// register keeper for acc2
	kIndex2 := registerKeeper(roleAddr, acc2Addr, test.Sk2, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)
	// register keeper for acc5
	kIndex5 := registerKeeper(roleAddr, acc5Addr, test.Sk5, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)
	// register keeper for acc6
	kIndex6 := registerKeeper(roleAddr, acc6Addr, test.Sk6, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)
	// register keeper for acc7
	kIndex7 := registerKeeper(roleAddr, acc7Addr, test.Sk7, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)
	// register keeper for acc8
	kIndex8 := registerKeeper(roleAddr, acc8Addr, test.Sk8, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)
	// register keeper for acc9
	kIndex9 := registerKeeper(roleAddr, acc9Addr, test.Sk9, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)
	// register keeper for acc10
	kIndex10 := registerKeeper(roleAddr, acc10Addr, test.Sk10, txopts, ethEndPoint, status, pledgePoolAddr, pledgeK)

	fmt.Println("begin deploy fileSys contract")

	// createGroup, and deploy FileSys contract
	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	gIndex, err := r.CreateGroup(rolefsAddr, 0, []uint64{kIndex2}, 1)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("gIndex ", gIndex)
	// 获取FileSys合约地址
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The FileSys contract address is ", fsAddr.Hex())
	// 给RoleFS合约赋予状态变量值
	err = rfs.SetAddr(issuanceAddr, roleAddr, rTokenAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	fmt.Println("============3. begin test AddOrder============")
	// user注册、充值
	r = callconts.NewR(roleAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
	err = r.Register(acc1Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	uIndex, err := r.GetRoleIndex(acc1Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uIndex ", uIndex)
	err = r.RegisterUser(rTokenAddr, uIndex, gIndex, []byte("test"), nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// user approve
	erc20 = callconts.NewERC20(test.PrimaryToken, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
	err = erc20.Approve(fsAddr, rechargeValue)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 查询该User对FileSys合约的allowance
	allo, err := erc20.Allowance(acc1Addr, fsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The allowance of User to FileSys contract is ", allo)
	if allo.Cmp(rechargeValue) != 0 {
		log.Fatal("allowance should be ", rechargeValue)
	}
	// user往FileSys中充值rechargeValue，用于存储服务付费
	err = r.Recharge(rTokenAddr, uIndex, 0, rechargeValue, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// user获取其在FileSys中的balance
	fs := callconts.NewFileSys(fsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	avail, _, err := fs.GetBalance(uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user avail in fs:", avail)
	if avail.Cmp(rechargeValue) < 0 {
		log.Fatal("user avail should not less than ", rechargeValue)
	}
	// provider注册、质押、申请角色、加入group
	r = callconts.NewR(roleAddr, acc3Addr, test.Sk3, txopts, ethEndPoint, status)
	err = r.Register(acc3Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	pIndex, err := r.GetRoleIndex(acc3Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pIndex ", pIndex)
	pp = callconts.NewPledgePool(pledgePoolAddr, acc3Addr, test.Sk3, txopts, ethEndPoint, status)
	err = pp.Pledge(test.PrimaryToken, roleAddr, pIndex, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	err = r.RegisterProvider(pledgePoolAddr, pIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	err = r.AddProviderToGroup(pIndex, gIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	//admin为RoleFS赋予Minter_Role权限
	erc20 = callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	err = erc20.SetUpRole(callconts.MinterRole, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	fmt.Println("============ calc signatures for addOrder ============")

	// use primary token for testing
	tIndex := uint32(0)

	// user = acc1
	usig, err := callconts.SignForOrder(uIndex, pIndex, nonce, start, end, size, tIndex, sprice, test.Sk1)
	if err != nil {
		log.Fatal(err)
	}
	// provide = acc3
	psig, err := callconts.SignForOrder(uIndex, pIndex, nonce, start, end, size, tIndex, sprice, test.Sk3)
	if err != nil {
		log.Fatal(err)
	}

	// keeper调用addOrder
	time.Sleep(test.WaitTime)
	rfs = callconts.NewRFS(rolefsAddr, acc1Addr, test.Sk1, txopts, ethEndPoint, status)
	callconts.ERC20Addr = test.PrimaryToken
	err = rfs.AddOrder(roleAddr, rTokenAddr, uIndex, pIndex, start, end, size, nonce, 0, sprice, usig, psig)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 获取addOrder后fs的信息并测试正确性
	pSum, err := fs.GetFsProviderSum(uIndex)
	if err != nil {
		log.Fatal(err)
	}
	if pSum != 1 {
		fmt.Println("pSum ", pSum)
		log.Fatal("The pSum should be 1")
	}
	pIndexGotten, err := fs.GetFsProvider(uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	if pIndexGotten != pIndex {
		fmt.Println("pIndexGotten ", pIndexGotten)
		log.Fatal("The pIndex gotten should be ", pIndex)
	}
	// 获取addOrder后channel的信息并测试正确性
	_amount, _nonce, _expire, err := fs.GetChannelInfo(uIndex, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: channel amount", _amount, ",nonce", _nonce, ",expire", _expire)
	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != end {
		log.Fatal("result error")
	}
	// 获取addOrder后store的信息并测试正确性
	_time, _size, _price, err := fs.GetStoreInfo(uIndex, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: storeinfo time", _time, ",size", _size, ",price", _price)
	if _price.Cmp(sprice) != 0 || _size != size || _time != end {
		log.Fatal("price in storeInfo after addOrder should be ", sprice, ", size in storeInfo should be ", size, ", time should be ", end)
	}
	// 获取addOrder后proInfo中的Settlement的信息，并测试正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: settleinfo time", _time, ",size", _size, ",price", _price, ",maxPay", _maxPay, ",hasPaid", _hasPaid, ",canPay", _canPay, ",lost", _lost, ",lostPaid", _lostPaid, ",managePay", _managePay, ",endPaid", _endPaid, ",linearPaid", _linearPaid)
	if _time != start {
		fmt.Println("time:", _time)
		log.Fatal("time should be ", start)
	}
	if _size != size {
		fmt.Println("size:", _size)
		log.Fatal("sieze should be ", size)
	}
	if _price.Cmp(sprice) != 0 {
		fmt.Println("price:", _price)
		log.Fatal("price should be ", sprice)
	}
	if _maxPay.Cmp(addOrderpay) != 0 {
		fmt.Println("maxPay:", _maxPay)
		log.Fatal("maxPay should be ", addOrderpay)
	}
	if _managePay.Cmp(managePay) != 0 {
		fmt.Println("managePay:", _managePay)
		log.Fatal("managePay should be ", managePay)
	}
	// 获取addOrder后proInfo中的aggOrder的信息，并测试正确性
	_nonce, _subNonce, err := fs.GetFsInfoAggOrder(uIndex, pIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: aggOrderInfo nonce ", _nonce, ",subNonce:", _subNonce)
	if _nonce != 1 || _subNonce != 0 {
		log.Fatal("nonce in aggOrder should be ", 1)
	}
	// 获取addOrder后foundation的balance，并测试正确性
	_availF, _tmpF, err := fs.GetBalance(0, 0) // 查询基金会在FileSys中的balance
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: foundation avail ", _availF, ",tmp ", _tmpF)
	if _availF.Cmp(big.NewInt(1).Div(addOrderpay, big.NewInt(100))) != 0 || _tmpF.Cmp(zero) != 0 {
		log.Fatal("availF should be ", 1, " tmpF should be ", _tmpF)
	}
	// 获取addOrder后user的balance，并测试正确性
	_avail, _tmp, err := fs.GetBalance(uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: user avail ", _avail, ",tmp ", _tmp)
	if _avail.Sub(avail, _avail).Cmp(reduceBal) != 0 {
		log.Fatal("the new avail should be", reduceBal, "less than avail")
	}
	// 获取addOrder后Role合约里的group信息，并测试正确性
	_isActive, _isBanned, _isReady, _level, sizeG, priceG, _fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: groupInfo isActive ", _isActive, ",isBanned ", _isBanned, ",isReady ", _isReady, ",level ", _level, ",size ", sizeG, ",price ", priceG, ",fsAddr ", _fsAddr.Hex())
	if sizeG.Cmp(big.NewInt(int64(size))) != 0 || priceG.Cmp(sprice) != 0 {
		log.Fatal("the sizeG should be ", size, " priceG should be ", sprice)
	}
	// 获取addOrder后代币发行的相关情况，并测试正确性(初次addOrder，mintToken值为0)
	mintLevel, err := issu.MintLevel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mintLevel:", mintLevel)
	if mintLevel.Cmp(zero) != 0 {
		log.Fatal("mintLevel should be 0")
	}
	st, err := issu.SpaceTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("spaceTime:", st)
	if st.Cmp(big.NewInt(0).Mul(new(big.Int).SetUint64(size), new(big.Int).SetUint64(end-start))) != 0 {
		log.Fatal("st should be ", big.NewInt(0).Mul(new(big.Int).SetUint64(size), new(big.Int).SetUint64(end-start)))
	}
	//size:10, price:10, totalPay:100,上面测试通过的话，这里就不需要再重复测试了

	fmt.Println("============ calc signatures for sub Order ============")

	// user = acc1
	usig, err = callconts.SignForOrder(uIndex, pIndex, nonce, start+5, start+10, size-5, tIndex, big.NewInt(0).Sub(sprice, big.NewInt(5)), test.Sk1)
	if err != nil {
		log.Fatal(err)
	}
	// provide = acc3
	psig, err = callconts.SignForOrder(uIndex, pIndex, nonce, start+5, start+10, size-5, tIndex, big.NewInt(0).Sub(sprice, big.NewInt(5)), test.Sk3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("============4. begin test SubOrder============")
	err = rfs.SubOrder(roleAddr, rTokenAddr, uIndex, pIndex, start+5, start+10, size-5, nonce, 0, big.NewInt(0).Sub(sprice, big.NewInt(5)), usig, psig)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 这里不重复测试了，仅抽取部分信息进行测试
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubOrder: time", _time, ", size", _size, ", price", _price, ", maxPay", _maxPay, ", hasPaid", _hasPaid, ", canPay", _canPay, ", lost", _lost, ", lostPaid", _lostPaid, ", managePay", _managePay, ", endPaid", _endPaid, ", linearPaid", _linearPaid)
	if _time != start+10 {
		log.Fatal("time should be ", start+10)
	}
	if _size != 5 {
		log.Fatal("size should be ", 5)
	}
	if _price.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("price should be 5")
	}
	if _maxPay.Cmp(addOrderpay) != 0 {
		log.Fatal("managePay should be", addOrderpay)
	}
	tmp := big.NewInt(0).Mul(sprice, big.NewInt(10)) //sprice*(newEnd - oldStart), newEnd是上面调用subOrder的start+10
	if _canPay.Cmp(tmp) != 0 {
		log.Fatal("canPay should be ", tmp)
	}

	fmt.Println("============ calc keeper signatures for proWithdraw ============")

	pay := big.NewInt(20)
	lost := big.NewInt(10)

	// keeper = acc2
	ksig2, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk2)
	if err != nil {
		log.Fatal(err)
	}
	// keeper = acc5
	ksig5, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk5)
	if err != nil {
		log.Fatal(err)
	}
	// keeper = acc6
	ksig6, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk6)
	if err != nil {
		log.Fatal(err)
	}
	// keeper = acc7
	ksig7, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk7)
	if err != nil {
		log.Fatal(err)
	}
	// keeper = acc8
	ksig8, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk8)
	if err != nil {
		log.Fatal(err)
	}
	// keeper = acc9
	ksig9, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk9)
	if err != nil {
		log.Fatal(err)
	}
	// keeper = acc10
	ksig10, err := callconts.SignForProWithdraw(pIndex, 0, pay, lost, test.Sk10)
	if err != nil {
		log.Fatal(err)
	}

	ksigs := [][]byte{ksig2, ksig5, ksig6, ksig7, ksig8, ksig9, ksig10}                  // indexes for keeeprs
	kIndexes := []uint64{kIndex2, kIndex5, kIndex6, kIndex7, kIndex8, kIndex9, kIndex10} // keepers

	fmt.Println("============5. begin test ProWithdraw============")

	err = rfs.ProWithdraw(roleAddr, rTokenAddr, pIndex, 0, pay, lost, kIndexes, ksigs)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time2, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw: time", _time2, ", size", _size, ", price", _price, ", maxPay", _maxPay, ", hasPaid", _hasPaid, ", canPay", _canPay, ", lost", _lost, ", lostPaid", _lostPaid, ", managePay", _managePay, ", endPaid", _endPaid, ", linearPaid", _linearPaid)
	if _size != 5 {
		log.Fatal("size should be ", 5)
	}
	if _hasPaid.Cmp(pay) != 0 {
		log.Fatal("hasPaid should be ", pay)
	}
	thisCanPay := big.NewInt(0).Mul(new(big.Int).SetUint64(_time2-_time), _price)
	tmp2 := big.NewInt(0).Add(tmp, thisCanPay)
	if _canPay.Cmp(tmp2) != 0 {
		log.Fatal("canPay should be ", tmp2)
	}
	if _lost.Cmp(lost) != 0 {
		log.Fatal("lost should be ", lost)
	}
	// 获取balance信息
	_avail, _tmp, err = fs.GetBalance(pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw, provider avail", _avail, ", tmp", _tmp)
	if _avail.Cmp(zero) != 0 {
		log.Fatal("avail should be 0")
	}
	if _tmp.Cmp(big.NewInt(0).Sub(_canPay, _hasPaid)) < 0 {
		log.Fatal("tmp should not be less than", big.NewInt(0).Sub(_canPay, _hasPaid))
	}
	// 获取provider在tIndex上的代币余额
	bal, err = erc20.BalanceOf(acc3Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw, provider balance is ", bal)

	fmt.Println("============6. begin test AddRepair============")
	// 获取Issuance中的totalPaid
	_totalPaid, err := issu.TotalPaid()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issuance totalPaid:", _totalPaid)
	if _totalPaid.Cmp(zero) != 0 {
		log.Fatal("result wrong")
	}
	// 调用AddRepair前，需要先调用ProWithdraw、指定lost值
	// 新注册一个provider
	r = callconts.NewR(roleAddr, acc4Addr, test.Sk4, txopts, ethEndPoint, status)
	err = r.Register(acc4Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	p2Index, err := r.GetRoleIndex(acc4Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p2Index ", p2Index)
	if p2Index != 10 {
		log.Fatal("p2Index should be 10")
	}
	pp = callconts.NewPledgePool(pledgePoolAddr, acc4Addr, test.Sk4, txopts, ethEndPoint, status)
	err = pp.Pledge(test.PrimaryToken, roleAddr, p2Index, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	err = r.RegisterProvider(pledgePoolAddr, p2Index, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	err = r.AddProviderToGroup(p2Index, gIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 先构造签名信息
	//bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sprice));
	npSig, err := callconts.SignForRepair(test.Sk4, pIndex, start+9, start+10, size, nonce, 0, big.NewInt(5), "a") // new provider sign，此处需要确保sprice*(end-start)<=lost
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("params:", pIndex, start+9, start+10, size, nonce, 0, big.NewInt(5), "a")
	fmt.Println("npSign:", npSig)
	// 调用AddRepair
	time.Sleep(test.WaitTime)
	rfs = callconts.NewRFS(rolefsAddr, acc2Addr, test.Sk2, txopts, ethEndPoint, status)
	err = rfs.AddRepair(roleAddr, rTokenAddr, pIndex, p2Index, start+9, start+10, size, nonce, 0, big.NewInt(5), npSig, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair: time", _time, ", size", _size, ", price", _price, ", maxPay", _maxPay, ", hasPaid", _hasPaid, ", canPay", _canPay, ", lost", _lost, ", lostPaid", _lostPaid, ", managePay", _managePay, ", endPaid", _endPaid, ", linearPaid", _linearPaid)
	if _time != (start + 9) {
		log.Fatal("time should be ", start+9)
	}
	if _size != size {
		log.Fatal("size should be ", size)
	}
	if _price.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("price should be 5")
	}
	// if _max.Cmp()
	// 获取balance信息
	_avail, _tmp, err = fs.GetBalance(p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, provider avail", _avail, ", tmp", _tmp)
	if _avail.Cmp(zero) != 0 || _tmp.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("result wrong")
	}
	// 获取repairFs信息
	pSum, err = fs.GetFsProviderSum(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs providerSum:", pSum)
	if pSum != 1 {
		log.Fatal("pSum should be 1")
	}
	_nonce, _subNonce, err = fs.GetFsInfoAggOrder(0, p2Index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider nonce:", _nonce, ", subNonce:", _subNonce)
	if _nonce != 1 || _subNonce != 0 {
		log.Fatal("result wrong")
	}
	_time, _size, _price, err = fs.GetStoreInfo(0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider storeInfo: time", _time, ", size", _size, ", price", _price)
	if _time != 0 || _size != size || _price.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("result wrong")
	}
	_amount, _nonce, _expire, err = fs.GetChannelInfo(0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider channelInfo: amount", _amount, ", nonce", _nonce, ", expire", _expire)
	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != 0 {
		log.Fatal("result wrong")
	}

	fmt.Println("============7. begin test SubRepair============")
	// 先构造签名信息
	//bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sprice));
	npSig, err = callconts.SignForRepair(test.Sk4, pIndex, start+9, start+10, size, nonce, 0, big.NewInt(5), "s") // new provider sign
	if err != nil {
		log.Fatal(err)
	}
	// 调用SubRepair
	startSubRepair := start + 9
	endSubRepair := start + 10
	spriceSubRepair := big.NewInt(5)
	err = rfs.SubRepair(roleAddr, rTokenAddr, pIndex, p2Index, startSubRepair, endSubRepair, size, nonce, 0, spriceSubRepair, npSig, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair: time", _time, ", size", _size, ", price", _price, ", maxPay", _maxPay, ", hasPaid", _hasPaid, ", canPay", _canPay, ", lost", _lost, ", lostPaid", _lostPaid, ", managePay", _managePay, ", endPaid", _endPaid, ", linearPaid", _linearPaid)
	if _time != endSubRepair {
		log.Fatal("time should be ", endSubRepair)
	}
	if _price.Cmp(zero) != 0 { // spriceAddRepair - spriceSubRepair
		log.Fatal("price should be 0")
	}
	maxPaySubRepair := big.NewInt(0).Mul(spriceSubRepair, new(big.Int).SetUint64(endSubRepair-startSubRepair))
	if _maxPay.Cmp(maxPaySubRepair) != 0 {
		log.Fatal("maxPay should be ", maxPaySubRepair)
	}
	// 获取balance信息
	_avail, _tmp, err = fs.GetBalance(p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, provider avail", _avail, ", tmp", _tmp)
	if _avail.Cmp(zero) != 0 || _tmp.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("result wrong")
	}
	// 获取repairFs信息
	pSum, err = fs.GetFsProviderSum(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs providerSum:", pSum)
	if pSum != 1 {
		log.Fatal("pSum should be 1")
	}
	_nonce, _subNonce, err = fs.GetFsInfoAggOrder(0, p2Index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider nonce:", _nonce, ", subNonce:", _subNonce)
	if _nonce != 1 || _subNonce != 1 {
		log.Fatal("result wrong")
	}
	_time, _size, _price, err = fs.GetStoreInfo(0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider storeInfo: time", _time, ", size", _size, ", price", _price)
	if _time != 0 || _size != 0 || _price.Cmp(zero) != 0 {
		log.Fatal("result wrong")
	}
	_amount, _nonce, _expire, err = fs.GetChannelInfo(0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider channelInfo: amount", _amount, ", nonce", _nonce, ", expire", _expire)
	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != 0 {
		log.Fatal("result wrong")
	}

	fmt.Println("============test success!============")
}

// regitst an account with address and sk
func registerAcc(
	roleAddr common.Address,
	accAddr common.Address,
	accSk string,
	txopts *callconts.TxOpts,
	ethEndPoint string,
	status chan error,
) {
	r := callconts.NewR(roleAddr, accAddr, accSk, txopts, ethEndPoint, status)
	err := r.Register(accAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
}

// register keeper role for an account
func registerKeeper(
	roleAddr common.Address,
	accAddr common.Address,
	accSk string,
	txopts *callconts.TxOpts,
	ethEndPoint string,
	status chan error,
	pledgePoolAddr common.Address,
	pledgeK *big.Int,
) uint64 {
	// register keeper for acc2
	r := callconts.NewR(roleAddr, accAddr, accSk, txopts, ethEndPoint, status)
	kIndex, err := r.GetRoleIndex(accAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("kIndex ", kIndex)
	pp := callconts.NewPledgePool(pledgePoolAddr, accAddr, accSk, txopts, ethEndPoint, status)
	err = pp.Pledge(test.PrimaryToken, roleAddr, kIndex, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	err = r.RegisterKeeper(pledgePoolAddr, kIndex, []byte("test"), nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	return kIndex
}
