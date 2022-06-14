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
	rechargeMoney := big.NewInt(1e8)
	start := uint64(time.Now().Unix()) // 当前时间的时间戳
	end := 86400 - start%86400 + start // end需要与天对齐
	size := uint64(10)
	sprice := big.NewInt(100)
	// 部署Role的参数
	pledgeK := big.NewInt(1e18)
	pledgeP := big.NewInt(1e18)
	version := uint16(1)
	var addrs []common.Address = []common.Address{common.HexToAddress(test.Acc1), common.HexToAddress(test.Acc2), common.HexToAddress(test.Acc3), common.HexToAddress(test.Acc4), common.HexToAddress(test.Acc5)}
	var sks []string = []string{test.Sk1, test.Sk2, test.Sk3, test.Sk4, test.Sk5}

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")
	for i, addr := range addrs {
		bal = callconts.QueryEthBalance(addr.Hex(), ethEndPoint)
		fmt.Println("acc", i, " balance: ", bal, " in Ethereum")
	}

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: nil,
		GasLimit: callconts.DefaultGasLimit,
	}

	// 查看测试账户在ERC20代币上的余额，不足时，自动充值
	statusTran := make(chan error)
	erc20 := callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	bal, err := erc20.BalanceOf(adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("admin balance in primaryToken is ", bal)

	tranNum := 0 // 表示转账checkTx的goroutine个数
	for i, addr := range addrs {
		bal, err = erc20.BalanceOf(addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc", i, " balance in primaryToken is ", bal)
		if bal.Cmp(pledgeK) < 0 {
			err = erc20.Transfer(addr, pledgeK) // admin给测试账户转账，用于测试（充值或质押）
			if err != nil {
				log.Fatal(err)
			}
			// 将开启一个goroutine
			tranNum++
		}
	}
	var errT error
	for i := 0; i < tranNum; i++ {
		err = <-statusTran
		if err != nil {
			errT = err // 避免goroutine泄漏
		}
	}
	if errT != nil {
		log.Fatal(errT)
	}

	fmt.Println("============1. begin test deploy FileSys contract============")
	// 部署FileSys合约前准备工作：部署Role合约、部署RoleFS合约、账户注册角色、调用Role合约中的CreateGroup函数
	// 顺序：Role、RoleFS、PledgePool、Issuance、CreateGroup(FileSys)、
	var roleAddr common.Address
	tranNum = 0
	r := callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	roleAddr, _, err = r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP, version)
	if err != nil {
		log.Fatal(err)
	}
	tranNum++
	fmt.Println("The Role contract address: ", roleAddr.Hex())
	//roleAddr := common.HexToAddress("0xe26F77e3268ae064514b18211Ed46DC0460197FE")

	var rolefsAddr common.Address
	rfs := callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	rolefsAddr, _, err = rfs.DeployRoleFS()
	if err != nil {
		errT = <-statusTran
		fmt.Println("DeployRole err:", errT)
		log.Fatal(err)
	}
	tranNum++
	fmt.Println("The RoleFS contract address: ", rolefsAddr.Hex())
	//rolefsAddr := common.HexToAddress("0xAAaC6D27153BF52d66Eed127e0321372B2FFF67C")
	for i := 0; i < tranNum; i++ {
		err = <-statusTran
		if err != nil {
			errT = err
		}
	}
	if errT != nil {
		log.Fatal(errT)
	}

	// 账户注册, acc1:User; acc2、acc3、acc4:Keeper; acc5:Provider
	rIndexes := make([]uint64, 5)
	regNum := 0
	for i, addr := range addrs {
		r = callconts.NewR(roleAddr, addr, sks[i], txopts, ethEndPoint, statusTran)
		err = r.Register(addr, nil)
		if err != nil {
			for j := 0; j < regNum; j++ {
				<-statusTran
			}
			log.Fatal(err)
		}
		regNum++
	}
	for i := 0; i < regNum; i++ {
		err = <-statusTran
		if err != nil {
			errT = err
		}
	}
	if errT != nil {
		log.Fatal(err)
	}
	for i, addr := range addrs {
		rIndexes[i], err = r.GetRoleIndex(addr)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Role indexes is ", rIndexes)
	//rIndexes := []uint64{1, 2, 3, 4, 5}

	// 部署PledgePool合约用于账户质押，之后才可以申请Keeper、Provider角色
	rtokenAddr, err := r.RToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RToken contract address: ", rtokenAddr.Hex())
	//rtokenAddr := common.HexToAddress("0x081458b892fb2caEb3e4a6234F9183594531b715")
	var pledgePoolAddr common.Address
	p := callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	pledgePoolAddr, _, err = p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	err = <-statusTran
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PledgePool contract address: ", pledgePoolAddr.Hex())
	//pledgePoolAddr := common.HexToAddress("0xfB279b29E437cEE73ac2F2423a3Cfb22060Eb9d7")
	pledgek, err := r.PledgeK() // 申请Keeper最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Apply Keeper need to pledge ", pledgek)
	pledgep, err := r.PledgeP() // 申请Provider最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Apply Provider need to pledge ", pledgep)
	// Keepers质押
	for i, rindex := range rIndexes[1:4] {
		err = toPledge(addrs[i+1], roleAddr, pledgePoolAddr, sks[i+1], rindex, pledgek, txopts)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Provider质押
	err = toPledge(addrs[4], roleAddr, pledgePoolAddr, sks[4], rIndexes[4], pledgep, txopts)
	if err != nil {
		log.Fatal(err)
	}

	var issuanceAddr common.Address
	issu := callconts.NewIssu(issuanceAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	issuanceAddr, _, err = issu.DeployIssuance(rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Issuance contract address is: ", issuanceAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1

	// 顺便给Role合约赋值
	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	err = r.SetPI(pledgePoolAddr, issuanceAddr, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}

	// 申请Keeper
	callconts.PledgePoolAddr = pledgePoolAddr
	rkNum := 0
	p = callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	for i, rindex := range rIndexes[1:4] {
		r := callconts.NewR(roleAddr, addrs[i+1], sks[i+1], txopts, ethEndPoint, statusTran)
		fmt.Println(addrs[i+1].Hex(), " begin to register Keeper...")
		// 获取PledgePool合约中的balance
		bal, err := p.GetBalanceInPPool(rindex, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("rindex ", rindex, " balance in PledgePool is ", bal)
		// 获取账户信息
		_, _, roleType, index, _, _, err := r.GetRoleInfo(addrs[i+1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The rindex", rindex, " information in Role contract, roleType:", roleType, " index:", index)

		err = r.RegisterKeeper(pledgePoolAddr, rindex, []byte("Hello, test"), nil)
		if err != nil {
			for j := 0; j < rkNum; j++ {
				<-statusTran
			}
			log.Fatal(err)
		}
		rkNum++
	}
	for i := 0; i < rkNum; i++ {
		err = <-statusTran
		if err != nil {
			errT = err
		}
	}
	if errT != nil {
		log.Fatal(err)
	}
	// 申请Provider
	r = callconts.NewR(roleAddr, addrs[4], sks[4], txopts, ethEndPoint, statusTran)
	err = r.RegisterProvider(pledgePoolAddr, rIndexes[4], nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 需要admin先调用CreateGroup,同时将部署FileSys合约
	r = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	gIndex, err := r.CreateGroup(rolefsAddr, uint64(0), rIndexes[1:3], 2)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 获取group信息
	isActive, isBanned, isReady, level, _size, price, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info: isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr.Hex())
	if !isActive {
		log.Fatal("group should be active")
	}
	if isBanned {
		log.Fatal("group shouldn't be banned")
	}
	if isReady {
		log.Fatal("group shouldn't be ready")
	}
	if level != 2 {
		log.Fatal("level shoule be 2")
	}

	fmt.Println("============2. begin test GetFsInfo============")
	// 先注册User，与group绑定
	r = callconts.NewR(roleAddr, addrs[0], sks[0], txopts, ethEndPoint, statusTran)
	err = r.RegisterUser(rtokenAddr, rIndexes[0], gIndex, []byte("Hello,test user"), nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 根据uIndex获取filesys信息
	fs := callconts.NewFileSys(fsAddr, addrs[0], sks[0], txopts, ethEndPoint, statusTran)
	isActive, tokenIndex, err := fs.GetFsInfo(rIndexes[0])
	if err != nil {
		log.Fatal(err)
	}
	if !isActive || tokenIndex != 0 {
		log.Fatal("The filesys with uIndex: ", rIndexes[0], " isActive: ", isActive, " tokenIndex: ", tokenIndex)
	}

	fmt.Println("============3. begin test GetBalance============")
	// Provider申请加入group
	r = callconts.NewR(roleAddr, addrs[4], sks[4], txopts, ethEndPoint, statusTran)
	err = r.AddProviderToGroup(rIndexes[4], gIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 获得FileSys中的AggOrder信息，则需要先提前调用RoleFS合约中的AddOrder函数，这样FileSys中才会有Order信息
	// 而调用RoleFS合约中的AddOrder函数前，需要先给RoleFS合约赋Issuance合约值
	// 给RoleFS合约赋值
	rfs = callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	err = rfs.SetAddr(issuanceAddr, roleAddr, rtokenAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// User给FileSys充值，用于存储服务支付
	fmt.Println("User approve FileSys contract account")
	erc20 = callconts.NewERC20(test.PrimaryToken, addrs[0], sks[0], txopts, ethEndPoint, statusTran)
	err = erc20.Approve(fsAddr, rechargeMoney)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	fmt.Println("User recharge to FileSys contract")
	// 查询该User对应的fsAddr,判断是否和上述相同
	r = callconts.NewR(roleAddr, addrs[0], sks[0], txopts, ethEndPoint, statusTran)
	_, _, _, _, _gIndex, _, err := r.GetRoleInfo(addrs[0])
	if err != nil {
		log.Fatal(err)
	}
	_, _, _, _, _, _, _fsAddr, err := r.GetGroupInfo(_gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The fsAddr of user is ", _fsAddr.Hex())
	fmt.Println("The user address is ", addrs[0].Hex()) // 用来判断是否和rIndexes[0]指代的地址相同
	// 查询该User对FileSys合约的allowance
	allo, err := erc20.Allowance(addrs[0], fsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The allowance of User to FileSys contract is ", allo)
	// 查询fs是否isActive
	_isActive, _, err := fs.GetFsInfo(rIndexes[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The fs of user isActive: ", _isActive)
	err = r.Recharge(rtokenAddr, rIndexes[0], 0, rechargeMoney, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 查询User在FileSys合约中的可用balance
	avail, tmp, err := fs.GetBalance(rIndexes[0], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The User's avail balance in FileSys contract is ", avail, " tmp is ", tmp)

	fmt.Println("============4. begin test GetFsInfoAggOrder============")
	// 调用AddOrder函数前，需要为RoleFS赋予MINTER_ROLE权限
	erc20 = callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, statusTran)
	err = erc20.SetUpRole(callconts.MinterRole, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 调用RoleFS合约中的AddOrder函数,keeper调用
	callconts.ERC20Addr = test.PrimaryToken
	rfs = callconts.NewRFS(rolefsAddr, addrs[0], sks[0], txopts, ethEndPoint, statusTran)
	fmt.Println("============ calc signatures for addOrder ============")
	// use primary token for testing
	tIndex := uint32(0)
	// user = acc1
	usig, err := callconts.SignForOrder(rIndexes[0], rIndexes[4], 0, start, end, size, tIndex, sprice, test.Sk1)
	if err != nil {
		log.Fatal(err)
	}
	// provide = acc5
	psig, err := callconts.SignForOrder(rIndexes[0], rIndexes[4], 0, start, end, size, tIndex, sprice, test.Sk5)
	if err != nil {
		log.Fatal(err)
	}
	err = rfs.AddOrder(roleAddr, rtokenAddr, rIndexes[0], rIndexes[4], start, end, size, 0, 0, sprice, usig, psig)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 获取order信息
	_nonce, subNonce, err := fs.GetFsInfoAggOrder(rIndexes[0], rIndexes[4])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The order info nonce is ", _nonce, " subNonce is ", subNonce)
	if _nonce != 1 || subNonce != 0 {
		log.Fatal("nonce should be 1, subNonce should be 0")
	}

	fmt.Println("============5. begin test GetFsProviderSum============")
	pSum, err := fs.GetFsProviderSum(rIndexes[0])
	if err != nil {
		log.Fatal(err)
	}
	if pSum != 1 {
		fmt.Println("The pSum is ", pSum)
		log.Fatal("The gotten pSum should be 1")
	}

	fmt.Println("============6. begin test GetFsProvider============")
	pIndexGotten, err := fs.GetFsProvider(rIndexes[0], 0)
	if err != nil {
		log.Fatal(err)
	}
	if pIndexGotten != rIndexes[4] {
		fmt.Println("The pIndex gotten is ", pIndexGotten)
		log.Fatal("The pIndex gotten should be ", rIndexes[4])
	}

	fmt.Println("============7. begin test GetStoreInfo============")
	_t, _s, _p, err := fs.GetStoreInfo(rIndexes[0], rIndexes[4], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The store information time is ", _t, ", size is ", _s, ", price is ", _p) // 应该和上述的time、size、price相同
	if _s != size || _p.Cmp(sprice) != 0 {
		log.Fatal("size shoule be ", size, ", price should be", sprice)
	}

	fmt.Println("============8. begin test GetChannelInfo============")
	amount, _n, expire, err := fs.GetChannelInfo(rIndexes[0], rIndexes[4], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The channel info amount is ", amount, ", nonce is ", _n, ", expire is ", expire)
	if expire != end || amount.Cmp(big.NewInt(0)) != 0 || _n != 0 {
		log.Fatal("expire should be ", end, ", amount shoule be 0, nonce should be 0")
	}

	fmt.Println("============9. begin test GetSettleInfo============")
	setime, sesize, seprice, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err := fs.GetSettleInfo(rIndexes[4], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The settlement info time is ", setime, ", size is ", sesize, ", price is ", seprice, ", maxPay is ", maxPay, ", hasPaid is ", hasPaid, ", canPay is ", canPay, ", lost is ", lost, ", lostPaid is ", lostPaid, ", managePay is ", managePay, ", endPaid is ", endPaid, ", linearPaid is ", linearPaid)
	if setime != start {
		log.Fatal("time shoule be ", start)
	}
	if sesize != size {
		log.Fatal("size shoule be ", size)
	}
	if seprice.Cmp(sprice) != 0 {
		log.Fatal("price shoule be ", sprice)
	}
	_maxPay := big.NewInt(0).Mul(new(big.Int).SetUint64(end-start), sprice)
	if maxPay.Cmp(_maxPay) != 0 {
		log.Fatal("maxPay should be ", _maxPay)
	}
	zero := big.NewInt(0)
	if hasPaid.Cmp(zero) != 0 || canPay.Cmp(zero) != 0 || lost.Cmp(zero) != 0 || lostPaid.Cmp(zero) != 0 || endPaid.Cmp(zero) != 0 || linearPaid.Cmp(zero) != 0 {
		log.Fatal("hasPaid,canPay,lost,lostPaid,endPaid,linearPaid shoule be 0")
	}
	_managePay := big.NewInt(0).Div(_maxPay, big.NewInt(25))
	if managePay.Cmp(_managePay) != 0 {
		log.Fatal("managePay should be ", _managePay)
	}

	fmt.Println("============10. begin test AddOrder again============")
	// 调用RoleFS合约中的AddOrder函数,keeper调用
	rfs = callconts.NewRFS(rolefsAddr, addrs[0], sks[0], txopts, ethEndPoint, statusTran)
	fmt.Println("============ calc signatures for addOrder ============")
	// user = acc1
	usig, err = callconts.SignForOrder(rIndexes[0], rIndexes[4], 1, end, end+86400, size, tIndex, sprice, test.Sk1)
	if err != nil {
		log.Fatal(err)
	}
	// provide = acc5
	psig, err = callconts.SignForOrder(rIndexes[0], rIndexes[4], 1, end, end+86400, size, tIndex, sprice, test.Sk5)
	if err != nil {
		log.Fatal(err)
	}
	err = rfs.AddOrder(roleAddr, rtokenAddr, rIndexes[0], rIndexes[4], end, end+86400, size, 1, 0, sprice, usig, psig)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-statusTran; err != nil {
		log.Fatal(err)
	}
	// 获取order信息
	_nonce, subNonce, err = fs.GetFsInfoAggOrder(rIndexes[0], rIndexes[4])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The order info nonce is ", _nonce, " subNonce is ", subNonce)
	if _nonce != 2 || subNonce != 0 {
		log.Fatal("nonce should be 2, subNonce should be 0")
	}

	fmt.Println("============test success!============")
}

func toPledge(addr, roleAddr, pledgePoolAddr common.Address, sk string, rindex uint64, pledgek *big.Int, txopts *callconts.TxOpts) error {
	status := make(chan error)
	p := callconts.NewPledgePool(pledgePoolAddr, addr, sk, txopts, ethEndPoint, status)
	// 调用pledge前需要先approve
	erc20 := callconts.NewERC20(test.PrimaryToken, addr, sk, txopts, ethEndPoint, status)
	err := erc20.Approve(pledgePoolAddr, pledgek)
	if err != nil {
		return err
	}
	err = <-status
	if err != nil {
		return err
	}
	// 查询allowance
	allo, err := erc20.Allowance(addr, pledgePoolAddr)
	if err != nil {
		return err
	}
	fmt.Println("The allowance of ", addr, " to ", pledgePoolAddr, " is ", allo)
	// 质押
	err = p.Pledge(test.PrimaryToken, roleAddr, rindex, pledgek, nil)
	if err != nil {
		return err
	}
	err = <-status
	if err != nil {
		return err
	}
	return nil
}
