package callconts

import (
	"log"
	"math/big"
	rolefs "memoc/contracts/rolefs"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
)

// NewRFS new a instance of ContractModule
func NewRFS(roleFSAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.RoleFSInfo {
	rfs := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: roleFSAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return rfs
}

// DeployRoleFS deploy an RoleFS contract, called by admin
func (rfs *ContractModule) DeployRoleFS() (common.Address, *rolefs.RoleFS, error) {
	var roleFSAddr common.Address
	var roleFSIns *rolefs.RoleFS

	log.Println("begin deploy RoleFS contract...")
	client := getClient(rfs.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return roleFSAddr, roleFSIns, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	roleFSAddr, tx, roleFSIns, err := rolefs.DeployRoleFS(auth, client)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployRoleFS Err:", err)
		return roleFSAddr, roleFSIns, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "DeployRoleFS")

	log.Println("RoleFS address is ", roleFSAddr.Hex())
	return roleFSAddr, roleFSIns, nil
}

func newRoleFS(roleFSAddr common.Address, client *ethclient.Client) (*rolefs.RoleFS, error) {
	roleFSIns, err := rolefs.NewRoleFS(roleFSAddr, client)
	if err != nil {
		return nil, err
	}
	return roleFSIns, nil
}

func (rfs *ContractModule) checkParam(uIndex, pIndex uint64, uRoleType, pRoleType uint8, tIndex uint32, roleAddr, rTokenAddr common.Address, label uint8) (uint64, error) {
	// check whether uIndex is user
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	uaddr, err := r.GetAddr(uIndex)
	if err != nil {
		return 0, err
	}
	isActive, isBanned, roleType, _, ugIndex, _, err := r.GetRoleInfo(uaddr)
	if err != nil {
		return 0, err
	}
	if roleType != uRoleType || !isActive || isBanned {
		log.Println("uIndex ", uIndex, " roleType:", roleType, " isBanned:", isBanned, " isActive:", isActive)
		return 0, ErrIndex
	}
	// check whether pIndex is provider
	paddr, err := r.GetAddr(pIndex)
	if err != nil {
		return 0, err
	}
	isActive, isBanned, roleType, _, pgIndex, _, err := r.GetRoleInfo(paddr)
	if err != nil {
		return 0, err
	}
	if roleType != pRoleType || !isActive || isBanned {
		log.Println("pIndex ", pIndex, " roleType:", roleType, " isBanned:", isBanned, " isActive:", isActive)
		return 0, ErrIndex
	}
	// check whether their gIndex is same
	isActive, isBanned, roleType, cIndex, cgIndex, _, err := r.GetRoleInfo(rfs.addr)
	if err != nil {
		return 0, err
	}
	if label == 1 { // addOrder, caller is user
		if cIndex != uIndex {
			log.Println("addOrder can only be called by user, but the caller index:", cIndex, " uIndex:", uIndex)
			return 0, errCaller
		}
	} else if label == 2 { // subOrder, caller is user or keeper
		if cIndex != uIndex {
			if roleType != KeeperRoleType || !isActive || isBanned {
				log.Println("SubOrder: caller:", rfs.addr.Hex(), " roleType:", roleType, "(should be keeper) isBanned:", isBanned, "(should not be banned) isActive:", isActive, "(should be active)")
				return 0, ErrIndex
			}
		}
	} else { // addReapir, subRepair
		if roleType != KeeperRoleType || !isActive || isBanned {
			log.Println("caller ", rfs.addr.Hex(), " roleType:", roleType, "(should be keeper) isBanned:", isBanned, "(should not be banned) isActive:", isActive, "(should be active)")
			return 0, ErrIndex
		}
	}
	if ugIndex != pgIndex || ugIndex != cgIndex {
		log.Println("uIndex's gIndex:", ugIndex, " pIndex's gIndex:", pgIndex, " caller's gIndex:", cgIndex)
		return 0, ErrIndex
	}
	// check tindex
	rt := NewRT(rTokenAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	isValid, err := rt.IsValid(tIndex)
	if err != nil {
		return 0, err
	}
	if !isValid {
		return 0, ErrTIndex
	}
	return ugIndex, nil
}

// SetAddr called by admin, which is the deployer. Set some address type variables.
func (rfs *ContractModule) SetAddr(issuan, role, rtoken common.Address) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check caller
	r := NewOwn(role, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	owner, err := r.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != rfs.addr.Hex() {
		log.Println("owner is", owner.Hex(), " but caller is", rfs.addr.Hex())
		return errNotOwner
	}

	log.Println("begin SetAddr in RoleFS contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleFSIns.SetAddr(auth, issuan, role, rtoken)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SetAddr Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "SetAddr")

	return nil
}

// AddOrder called by keeper? Add the storage order in the FileSys.
// hash(uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice)?
// 目前合约中还未对签名进行判断处理
// nonce需要从0开始依次累加
// 调用该函数前，需要admin为RoleFS合约账户赋予MINTER_ROLE权限
func (rfs *ContractModule) AddOrder(roleAddr, rTokenAddr common.Address, uIndex, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, usign, psign []byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()

	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check start,end,size
	if size == 0 {
		return errSize
	}
	if end <= start {
		log.Println("start:", start, " end:", end)
		return errEnd
	}
	if (end/86400)*86400 != end {
		log.Println("end:", end)
		return xerrors.New("end should be divisible by 86400(one day)")
	}
	// check uIndex,pIndex,gIndex,tIndex
	gIndex, err := rfs.checkParam(uIndex, pIndex, UserRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr, 1)
	if err != nil {
		return err
	}

	// check balance
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	pay := big.NewInt(0).Mul(sprice, new(big.Int).SetUint64(end-start))
	manageAndTax := big.NewInt(0).Div(pay, big.NewInt(20)) // pay/100*4 + pay/100*1
	payAndTax := big.NewInt(0).Add(pay, manageAndTax)
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	avail, _, err := fs.GetBalance(uIndex, tIndex)
	if err != nil {
		return err
	}
	if avail.Cmp(payAndTax) < 0 {
		log.Println("payAndTax is", payAndTax, " but avail is", avail)
		return ErrBalNotE
	}
	// check nonce
	_nonce, _, err := fs.GetFsInfoAggOrder(uIndex, pIndex)
	if err != nil {
		return err
	}
	if _nonce != nonce {
		log.Println("nonce:", nonce, " should be", _nonce)
		return errNonce
	}
	// check start
	_time, _, _, err := fs.GetStoreInfo(uIndex, pIndex, tIndex)
	if err != nil {
		return err
	}
	if end < _time {
		log.Println("end:", start, " should be more than time:", _time)
		return xerrors.New("end error")
	}
	// check whether rolefsAddr has Minter-Role
	if tIndex == 0 {
		erc20 := NewERC20(ERC20Addr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
		has, err := erc20.HasRole(MinterRole, rfs.contractAddress)
		if err != nil {
			return err
		}
		if !has {
			log.Println("rolefsAddr:", rfs.contractAddress.Hex(), " hasn't MinterRole, please setUpRole first")
			return xerrors.New("rolefsAddr has not MinterRole")
		}
	}

	log.Println("begin AddOrder in RoleFS contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return errMA
	}

	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	// tx, err := roleFSIns.AddOrder(auth, uIndex, pIndex, start, end, size, nonce, tIndex, sprice, usign, psign, ksigns)
	// use struct to call addOder
	ps := rolefs.AOParams{
		UIndex: uIndex,
		PIndex: pIndex,
		Start:  start,
		End:    end,
		Size:   size,
		Nonce:  nonce,
		TIndex: tIndex,
		SPrice: sprice,
		Usign:  usign,
		Psign:  psign,
		//Ksigns: ksigns,
	}
	// call with struct param
	tx, err := roleFSIns.AddOrder(auth, ps)

	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("AddOrder Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "AddOrder")

	return nil
}

// SubOrder called by keeper? Reduce the storage order in the FileSys.
// hash(uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice)?
// 目前合约中还未对签名信息做判断处理
func (rfs *ContractModule) SubOrder(roleAddr, rTokenAddr common.Address, uIndex, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, usign, psign []byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check size,start.end
	if size <= 0 {
		return errSize
	}
	now := uint64(time.Now().Unix())
	if end <= start || end > now {
		log.Println("end:", end, " start:", start, " now:", now)
		return errEndNow
	}
	// check uIndex,pIndex,gIndex,tIndex
	gIndex, err := rfs.checkParam(uIndex, pIndex, UserRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr, 2)
	if err != nil {
		return err
	}

	/*
		// check ksigns's length
		r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
		gkNum, err := r.GetGKNum(gIndex)
		if err != nil {
			return err
		}
		if len(ksigns) < int(gkNum*2/3) {
			return ErrKSignsNE
		}
	*/

	// check nonce
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	_nonce, _subNonce, err := fs.GetFsInfoAggOrder(uIndex, pIndex)
	if err != nil {
		return err
	}
	if _nonce <= nonce {
		log.Println("nonce:", nonce, " should less than addNonce:", _nonce, ", you should call addOrder first")
		return errNonce
	}
	if _subNonce != nonce {
		log.Println("nonce:", nonce, " should be", _subNonce)
		return errNonce
	}
	// check size
	_, _size, _price, err := fs.GetStoreInfo(uIndex, pIndex, tIndex)
	if err != nil {
		return err
	}
	if size > _size {
		log.Println("size:", size, " shouldn't be more than store.size:", _size)
		return errSize
	}
	if sprice.Cmp(_price) > 0 {
		log.Println("sprice:", sprice, " shouldn't be more than store.price:", _price)
		return errSprice
	}

	log.Println("begin SubOrder in RoleFS contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return errMA
	}

	// prepair params for subOrder
	ps := rolefs.SOParams{
		KIndex: 0,
		UIndex: uIndex,
		PIndex: pIndex,
		Start:  start,
		End:    end,
		Size:   size,
		Nonce:  nonce,
		TIndex: tIndex,
		SPrice: sprice,
		Usign:  usign,
		Psign:  psign,
		//Ksigns: ksigns,
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	//tx, err := roleFSIns.SubOrder(auth, uIndex, pIndex, start, end, size, nonce, tIndex, sprice, usign, psign, ksigns)
	tx, err := roleFSIns.SubOrder(auth, ps)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SubOrder Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "SubOrder")

	return nil
}

// AddRepair called by keeper. Add the repair order in the FileSys.
// hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "a"), signed by newProvider
func (rfs *ContractModule) AddRepair(roleAddr, rTokenAddr common.Address, pIndex, nPIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, psign []byte, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex, nPIndex,tIndex,gIndex
	gIndex, err := rfs.checkParam(pIndex, nPIndex, ProviderRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr, 3)
	if err != nil {
		return err
	}
	// check ksigns's length
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	if len(ksigns) < int(gkNum*2/3) {
		return ErrKSignsNE
	}
	// check start,end,size
	if size <= 0 {
		return errSize
	}
	if end <= start {
		log.Println("end:", end, " start:", start)
		return errEnd
	}
	// check lost,lostPaid
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	_, _, _, _, _, _, _lost, _lostPaid, _, _, _, err := fs.GetSettleInfo(pIndex, tIndex)
	if err != nil {
		return err
	}
	if _lost.Cmp(_lostPaid) < 0 {
		log.Println("pIndex:", pIndex, " lost:", _lost, " lostPaid:", _lostPaid, ", lost shouldn't less than lostPaid")
		return xerrors.New("lost error")
	}
	bal := big.NewInt(0).Sub(_lost, _lostPaid)
	pay := big.NewInt(0).Mul(sprice, new(big.Int).SetUint64(end-start))
	if bal.Cmp(pay) < 0 {
		log.Println("pIndex:", pIndex, " bal:", bal, " pay:", pay, ", bal shouldn't be less than pay")
		return xerrors.New("pay error")
	}
	// check nonce
	_nonce, _, err := fs.GetFsInfoAggOrder(0, nPIndex)
	if err != nil {
		return err
	}
	if _nonce != nonce {
		log.Println("newPro:", nPIndex, " repairFs.nonce:", _nonce, " nonce:", nonce, ", they should be same")
		return errNonce
	}

	log.Println("begin AddRepair in RoleFS contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleFSIns.AddRepair(auth, pIndex, nPIndex, start, end, size, nonce, tIndex, sprice, psign, ksigns)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("AddRepair Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "AddRepair")

	return nil
}

// SubRepair called by keeper. Reduce the repair order in the FileSys.
// hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "s"), signed by newProvider
func (rfs *ContractModule) SubRepair(roleAddr, rTokenAddr common.Address, pIndex, nPIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, psign []byte, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex,npIndex,gIndex,tIndex
	gIndex, err := rfs.checkParam(pIndex, nPIndex, ProviderRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr, 3)
	if err != nil {
		return err
	}
	// check ksigns's length
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	if len(ksigns) < int(gkNum*2/3) {
		return ErrKSignsNE
	}
	// check start,end,size
	if size <= 0 {
		return errSize
	}
	now := uint64(time.Now().Unix())
	if end <= start || end > now {
		log.Println("end:", end, " start:", start, " now:", now)
		return errEndNow
	}
	// check nonce
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	_, _subNonce, err := fs.GetFsInfoAggOrder(0, nPIndex)
	if err != nil {
		return err
	}
	if _subNonce != nonce {
		log.Println("nonce:", nonce, " should be", _subNonce)
		return errNonce
	}
	// check size
	_, _size, _price, err := fs.GetStoreInfo(0, nPIndex, tIndex)
	if err != nil {
		return err
	}
	if size > _size {
		log.Println("size:", size, " shouldn't be more than store.size:", _size)
		return errSize
	}
	if sprice.Cmp(_price) > 0 {
		log.Println("sprice:", sprice, " shouldn't be more than store.price:", _price)
		return errSprice
	}

	log.Println("begin SubRepair in RoleFS contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleFSIns.SubRepair(auth, pIndex, nPIndex, start, end, size, nonce, tIndex, sprice, psign, ksigns)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SubRepair Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "SubRepair")

	return nil
}

// ProWithdraw called by keeper? Retrieve the Provider's balance in FileSys.
// hash(pIndex, tIndex, pay, lost)?
func (rfs *ContractModule) ProWithdraw(roleAddr, rTokenAddr common.Address, pIndex uint64, tIndex uint32, pay, lost *big.Int, kIndexes []uint64, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	addr, err := r.GetAddr(pIndex)
	if err != nil {
		return err
	}
	isActive, isBanned, roleType, _, gIndex, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if !isActive || isBanned || roleType != ProviderRoleType {
		log.Println("pIndex isActive:", isActive, " isBanned:", isBanned, " roleType:", roleType, ", should be active,not be banned,roleType should be 2")
		return ErrIndex
	}

	// check ksigns's length
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	l := int(gkNum * 2 / 3)
	le := len(ksigns)
	if le < l {
		log.Println("ksigns length", le, " shouldn't be less than", l)
		return ErrKSignsNE
	}

	log.Println("begin call ProWithdraw in RoleFS contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(rfs.endPoint, rfs.hexSk, nil, rfs.txopts)
	if errMA != nil {
		return errMA
	}

	// get provider address for calling proWithdraw
	r = NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint, rfs.Status)
	proAddr, gIndex, err := r.GetAddrGindex(pIndex)
	if err != nil {
		return err
	}
	// get token address from tIndex for calling proWithdraw
	rt := NewRT(rTokenAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	tAddr, err := rt.GetTA(tIndex)
	if err != nil {
		return err
	}

	// prepair params for subOrder
	ps := rolefs.PWParams{
		PIndex:   pIndex,
		TIndex:   tIndex,
		PAddr:    proAddr,
		TAddr:    tAddr,
		Pay:      pay,
		Lost:     lost,
		KIndexes: kIndexes,
		Ksigns:   ksigns,
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	//tx, err := roleFSIns.ProWithdraw(auth, pIndex, tIndex, pay, lost, ksigns)
	tx, err := roleFSIns.ProWithdraw(auth, ps)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("ProWithdraw Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(rfs.endPoint, tx, rfs.Status, "ProWithdraw")

	return nil
}
