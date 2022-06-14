// only prepare to complete calling basic functions of contract now, after that, will judge the input parameters of functions
// And also need to add Getter functions.

package callconts

import (
	"fmt"
	"log"
	"math/big"
	"memoc/contracts/role"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
)

// NewR new a instance of ContractModule. 'roleAddr' indicates Role contract address
func NewR(roleAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.RoleInfo {
	r := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: roleAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return r
}

// DeployRole deploy a Role contract, called by admin, specify foundation、primaryToken、pledgeK、pledgeP
func (r *ContractModule) DeployRole(foundation, primaryToken common.Address, pledgeKeeper, pledgeProvider *big.Int, version uint16) (common.Address, *role.Role, error) {
	var roleAddr common.Address
	var roleIns *role.Role

	log.Println("begin deploy Role contract...")
	client := getClient(r.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return roleAddr, roleIns, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	roleAddr, tx, roleIns, err := role.DeployRole(auth, client, foundation, primaryToken, pledgeKeeper, pledgeProvider, version)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployRole Err:", err)
		return roleAddr, roleIns, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "DeployRole")

	log.Println("Role address is ", roleAddr.Hex())
	return roleAddr, roleIns, nil
}

// newRole new an instance of Role contract, 'roleAddr' indicates Role contract address
func newRole(roleAddr common.Address, client *ethclient.Client) (*role.Role, error) {
	roleIns, err := role.NewRole(roleAddr, client)
	if err != nil {
		return nil, err
	}
	return roleIns, nil
}

// SetPI callled by admin, set pledgePool-address、 issuance-address and rolefs-address
func (r *ContractModule) SetPI(pledgePoolAddr, issuAddr, rolefsAddr common.Address) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	owner, err := r.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != r.addr.Hex() {
		log.Println("owner of Role-contract is", owner.Hex(), ", but caller is", r.addr.Hex())
		return errNotOwner
	}

	log.Println("begin SetPI in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.SetPI(auth, pledgePoolAddr, issuAddr, rolefsAddr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SetPI Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "SetPI")

	return nil
}

// Register called by anyone to get index
func (r *ContractModule) Register(addr common.Address, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check if addr has registered
	_, _, _, index, _, _, err := r.GetRoleInfo(addr)
	if index > 0 { // has registered already
		log.Println("Has registered, index is ", index)
		go func() {
			r.Status <- nil
		}()
		return nil
	}

	log.Println("begin Register in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.Register(auth, addr, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Register Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "Register")

	return nil
}

// RegisterKeeper called by anyone to register Keeper role, befor this, you should pledge in PledgePool
func (r *ContractModule) RegisterKeeper(pledgePoolAddr common.Address, index uint64, blskey []byte, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	addr, err := r.GetAddr(index)
	if err != nil {
		return err
	}
	log.Println("account address get by rIndex", index, "is:", addr.Hex())
	_, _, roleType, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if roleType != 0 { // role already registered
		return ErrRoleReg
	}
	pp := NewPledgePool(pledgePoolAddr, r.addr, r.hexSk, r.txopts, r.endPoint, r.Status)
	pledge, err := pp.GetBalanceInPPool(index, 0) // tindex:0 表示主代币
	if err != nil {
		return err
	}
	pledgek, err := r.PledgeK()
	if err != nil {
		return err
	}
	if pledge.Cmp(pledgek) < 0 {
		log.Println("the rindex:", index, ", addr:", addr.Hex(), ", pledgeMoney:", pledge, " is not enough, shouldn't less than ", pledgek)
		return errPledgeNE
	}

	log.Println("begin RegisterKeeper in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.RegisterKeeper(auth, index, blskey, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("RegisterKeeper Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "RegisterKeeper")

	return nil
}

// RegisterProvider called by anyone to register Provider role, befor this, you should pledge in PledgePool
func (r *ContractModule) RegisterProvider(pledgePoolAddr common.Address, index uint64, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	addr, err := r.GetAddr(index)
	if err != nil {
		return err
	}
	log.Println("account address get by rIndex", index, "is:", addr.Hex())
	_, _, roleType, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if roleType != 0 { // role already registered
		log.Println("account roleType is:", roleType)
		return ErrRoleReg
	}
	pp := NewPledgePool(pledgePoolAddr, r.addr, r.hexSk, r.txopts, r.endPoint, r.Status)
	pledge, err := pp.GetBalanceInPPool(index, 0) // tindex:0 表示主代币
	if err != nil {
		return err
	}
	log.Println("account pledge value is:", pledge)
	pledgep, err := r.PledgeP()
	if err != nil {
		return err
	}
	log.Println("register provider need pledge value:", pledgep)
	if pledge.Cmp(pledgep) < 0 {
		log.Println("the rindex ", index, " addr:", addr.Hex(), " pledgeMoney:", pledge, " is not enough, shouldn't less than ", pledgep)
		return errPledgeNE
	}

	log.Println("begin RegisterProvider in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.RegisterProvider(auth, index, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("RegisterProvider Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "RegisterProvider")

	return nil
}

// RegisterUser called by anyone to register User role
func (r *ContractModule) RegisterUser(rTokenAddr common.Address, index uint64, gindex uint64, blskey []byte, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check index
	addr, err := r.GetAddr(index)
	if err != nil {
		return err
	}
	log.Println("account address get by rIndex", index, "is:", addr.Hex())
	_, _, roleType, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if roleType != 0 { // role already registered
		log.Println("account roleType is:", roleType)
		return ErrRoleReg
	}

	// check gindex
	isActive, isBanned, _, _, _, _, _, err := r.GetGroupInfo(gindex)
	if err != nil {
		return err
	}
	if !isActive || isBanned {
		log.Println("group ", gindex, " isActive:", isActive, " isBanned:", isBanned)
		return ErrInvalidG
	}

	// don't need to check fs

	log.Println("begin RegisterUser in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.RegisterUser(auth, index, gindex, blskey, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("RegisterUser Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "RegisterUser")

	return nil
}

// RegisterToken called by admin to register token. Once token is registered, it is supported by memo.
func (r *ContractModule) RegisterToken(tokenAddr common.Address) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check whether pledgePool address in Role contract is valid
	pledgePool, err := r.PledgePool()
	if err != nil {
		return err
	}
	if pledgePool.Hex() == InvalidAddr {
		return ErrNotSetPP
	}

	log.Println("begin RegisterToken in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.RegisterToken(auth, tokenAddr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("RegisterToken Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "RegisterToken")

	return nil
}

// CreateGroup called by admin to create a group and then deploy FileSys contract and then set FileSys-address to group.
// founder default is 0, len(keepers)>=level then group is active
func (r *ContractModule) CreateGroup(rfsAddr common.Address, founder uint64, kindexes []uint64, level uint16) (uint64, error) {
	gIndex, err := r.createGroup(kindexes, level)
	if err != nil {
		return gIndex, err
	}

	// deploy FileSys
	fsAddr, _, err := r.DeployFileSys(founder, gIndex, r.contractAddress, rfsAddr, kindexes)
	if err != nil {
		return gIndex, err
	}

	err = r.SetGF(fsAddr, gIndex)
	return gIndex, err
}

// CreateGroup called by admin to create a group.
func (r *ContractModule) createGroup(kindexes []uint64, level uint16) (uint64, error) {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return 0, err
	}

	// check caller
	owner, err := r.GetOwner()
	if err != nil {
		return 0, err
	}
	if owner.Hex() != r.addr.Hex() {
		log.Println("owner is", owner.Hex(), "but caller is", r.addr.Hex())
		return 0, errNotOwner
	}

	// check kindexes
	var tmpAddr common.Address
	for _, kindex := range kindexes {
		tmpAddr, err = r.GetAddr(kindex)
		isActive, isBanned, roleType, _, _, _, err := r.GetRoleInfo(tmpAddr)
		if err != nil {
			return 0, err
		}
		if roleType != KeeperRoleType || isActive || isBanned {
			log.Println("rindex ", kindex, " in kindexes is invalid, the address is", tmpAddr)
			log.Println("its roleType:", roleType, " isActive:", isActive, "isBanned: ", isBanned)
			return 0, ErrIndex
		}
	}

	log.Println("begin CreateGroup in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return 0, errMA
	}
	tx, err := roleIns.CreateGroup(auth, kindexes, level)
	if err != nil {
		log.Println("CreateGroup Err:", err)
		return 0, err
	}

	cg := make(chan error)
	// NOTE： 此处需等待checkTx执行完毕,从而获取gIndex
	go checkTx(r.endPoint, tx, cg, "CreateGroup")
	err = <-cg
	if err != nil {
		return 0, err
	}

	gIndex, err := getGIndexFromRLogs(r.endPoint, tx.Hash())
	if err != nil {
		return 0, err
	}
	log.Println("CreateGroup in Role, the gIndex is", gIndex)
	return gIndex, nil
}

// SetGF called by admin to set fsAddress for group after CreateGroup and deployFileSys.
func (r *ContractModule) SetGF(fsAddr common.Address, gIndex uint64) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check gIndex
	num, err := r.GetGroupsNum()
	if err != nil {
		return err
	}
	if gIndex == 0 || gIndex > num {
		return ErrInvalidG
	}

	log.Println("begin SetGF in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.SetGF(auth, gIndex, fsAddr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SetGF Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "SetGF")

	return nil
}

// AddKeeperToGroup called by admin.
func (r *ContractModule) AddKeeperToGroup(kIndex uint64, gIndex uint64) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check caller
	owner, err := r.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != r.addr.Hex() {
		log.Println("owner is", owner.Hex(), "but caller is", r.addr.Hex())
		return errNotOwner
	}

	// check gIndex
	num, err := r.GetGroupsNum()
	if err != nil {
		return err
	}
	if gIndex == 0 || gIndex > num {
		log.Println("the gIndex", gIndex, "shouldn't be zero or more than groupsNum", num)
		return ErrInvalidG
	}
	_, isBanned, _, _, _, _, _, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	if isBanned {
		log.Println("the group represented by gIndex is banned")
		return ErrInvalidG
	}

	// check kIndex
	addr, err := r.GetAddr(kIndex)
	if err != nil {
		return err
	}
	log.Println("account address get by rIndex", kIndex, "is:", addr.Hex())
	isActive, isBanned, roleType, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if isActive || isBanned || roleType != KeeperRoleType {
		log.Println("The account represented by kIndex", kIndex, " isActive:", isActive, " isBanned:", isBanned, " roleType:", roleType)
		return ErrIndex
	}

	log.Println("begin AddKeeperToGroup in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.AddKeeperToGroup(auth, kIndex, gIndex)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("AddKeeperToGroup Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "AddKeeperToGroup")

	return nil
}

// AddProviderToGroup called by provider or called by others.
func (r *ContractModule) AddProviderToGroup(pIndex uint64, gIndex uint64, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex
	addr, err := r.GetAddr(pIndex)
	if err != nil {
		return err
	}
	log.Println("account address get by rIndex", pIndex, "is:", addr.Hex())
	isActive, isBanned, roleType, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}

	if isActive || isBanned || roleType != ProviderRoleType {
		log.Println("The account represented by pIndex", pIndex, " isActive:", isActive, " isBanned:", isBanned, " roleType:", roleType)
		return ErrIndex
	}

	// check gIndex
	num, err := r.GetGroupsNum()
	if err != nil {
		return err
	}
	if gIndex == 0 || gIndex > num {
		log.Println("gIndex shouldn't be zero or more than groupsNum", num)
		return ErrInvalidG
	}
	kNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	fmt.Println("keeper num of group: ", kNum)

	isActive, isBanned, _, _, _, _, _, err = r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	if !isActive {
		log.Println("the group is not active.")
		return ErrNotActive
	}
	if isBanned {
		log.Println("the group represented by gIndex is banned")
		return ErrInvalidG
	}

	log.Println("begin AddProviderToGroup in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.AddProviderToGroup(auth, pIndex, gIndex, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("AddProviderToGroup Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "AddProviderToGroup")

	return nil
}

// SetPledgeMoney called by admin to set the amount that the keeper and provider needs to pledge.
func (r *ContractModule) SetPledgeMoney(kpledge *big.Int, ppledge *big.Int) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check caller
	owner, err := r.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != r.addr.Hex() {
		log.Println("owner is", owner.Hex(), "but caller is", r.addr.Hex())
		return errNotOwner
	}

	log.Println("begin SetPledgeMoney in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.SetPledgeMoney(auth, kpledge, ppledge)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SetPledgeMoney Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "SetPledgeMoney")

	return nil
}

// Recharge called by user or called by others.
func (r *ContractModule) Recharge(rTokenAddr common.Address, uIndex uint64, tIndex uint32, money *big.Int, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// uIndex need to be user
	addr, err := r.GetAddr(uIndex)
	if err != nil {
		return err
	}
	log.Println("account address get by rIndex", uIndex, "is:", addr.Hex())
	_, isBanned, roleType, _, gIndex, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if isBanned || roleType != UserRoleType {
		log.Println("The uIndex", uIndex, " isBanned:", isBanned, " roleType:", roleType)
		return ErrIndex
	}

	// check tindex
	rt := NewRT(rTokenAddr, r.addr, r.hexSk, r.txopts, r.endPoint)
	isValid, err := rt.IsValid(tIndex)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrTIndex
	}

	// check allowance
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	tAddr, err := rt.GetTA(tIndex)
	if err != nil {
		return err
	}
	erc20 := NewERC20(tAddr, r.addr, r.hexSk, r.txopts, r.endPoint, r.Status)
	allo, err := erc20.Allowance(addr, fsAddr)
	if err != nil {
		return err
	}
	if allo.Cmp(money) < 0 {
		log.Println("uIndex", uIndex, " addr:", addr.Hex(), " allowance to fsAddr", fsAddr.Hex(), "is", allo, "less than recharge money", money)
		return ErrAlloNotE
	}

	log.Println("begin Recharge in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.Recharge(auth, uIndex, tIndex, money, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Recharge Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "Recharge")

	return nil
}

// WithdrawFromFs called by memo-role or called by others.
// foundation、user、keeper withdraw money from filesystem
func (r *ContractModule) WithdrawFromFs(rTokenAddr common.Address, rIndex uint64, tIndex uint32, amount *big.Int, sign []byte) error {
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return err
	}

	// check amount
	if amount.Cmp(big.NewInt(0)) <= 0 {
		return xerrors.New("amount shouldn't be 0")
	}

	// check tindex
	rt := NewRT(rTokenAddr, r.addr, r.hexSk, r.txopts, r.endPoint)
	isValid, err := rt.IsValid(tIndex)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrTIndex
	}

	// check rIndex
	addr, err := r.GetAddr(rIndex)
	if err != nil {
		return err
	}
	_, isBanned, rtype, _, gIndex, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	// 需要存在有效的gIndex
	if gIndex == 0 {
		log.Println("rIndex:", rIndex, " addr:", addr.Hex(), " gIndex:", gIndex)
		return ErrInvalidG
	}
	// 非 foundation 取回余额
	if r.addr.Hex() != Foundation.Hex() {
		// 账户不能被禁止
		if isBanned {
			log.Println("rIndex:", rIndex, " addr:", addr.Hex(), " isBanned:", isBanned)
			return ErrIsBanned
		}
		// 账户不能是provider
		if rtype == ProviderRoleType {
			log.Println("rIndex:", rIndex, " addr:", addr.Hex(), " roleType:", rtype, " is a provider, please call proWithdraw rather than this")
			return ErrIndex
		}
	}

	log.Println("begin WithdrawFromFs in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(r.endPoint, r.hexSk, nil, r.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.WithdrawFromFs(auth, rIndex, tIndex, amount, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("WithdrawFromFs Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(r.endPoint, tx, r.Status, "WithdrawFromFs")

	return nil
}

// PledgePool get pledgepool
func (r *ContractModule) PledgePool() (common.Address, error) {
	var pp common.Address
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return pp, err
	}

	retryCount := 0
	for {
		retryCount++
		pp, err = roleIns.PledgePool(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pp, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pp, nil
	}
}

// Foundation get foundation address
func (r *ContractModule) Foundation() (common.Address, error) {
	var f common.Address
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return f, err
	}

	retryCount := 0
	for {
		retryCount++
		f, err = roleIns.Foundation(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return f, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return f, nil
	}
}

// PledgeK get the pledgeAmount that the account need to pledge when it register Keeper
func (r *ContractModule) PledgeK() (*big.Int, error) {
	pk := big.NewInt(0)
	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return pk, err
	}

	retryCount := 0
	for {
		retryCount++
		pk, err = roleIns.PledgeK(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pk, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pk, nil
	}
}

// PledgeP get the pledgeAmount that the account need to pledge when it register Provider
func (r *ContractModule) PledgeP() (*big.Int, error) {
	pp := big.NewInt(0)

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return pp, err
	}

	retryCount := 0
	for {
		retryCount++
		pp, err = roleIns.PledgeP(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pp, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pp, nil
	}
}

// RToken get RToken contract address
func (r *ContractModule) RToken() (common.Address, error) {
	var rt common.Address

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return rt, err
	}

	retryCount := 0
	for {
		retryCount++
		rt, err = roleIns.RToken(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return rt, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return rt, nil
	}
}

// Issuance get Issuance contract address
func (r *ContractModule) Issuance() (common.Address, error) {
	var is common.Address

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return is, err
	}

	retryCount := 0
	for {
		retryCount++
		is, err = roleIns.Issuance(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return is, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return is, nil
	}
}

// Rolefs get RoleFS contract address
func (r *ContractModule) Rolefs() (common.Address, error) {
	var rfs common.Address

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return rfs, err
	}

	retryCount := 0
	for {
		retryCount++
		rfs, err = roleIns.Rolefs(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return rfs, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return rfs, nil
	}
}

// GetAddrsNum get the number of registered addresses.
func (r *ContractModule) GetAddrsNum() (uint64, error) {
	var anum uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return anum, err
	}

	retryCount := 0
	for {
		retryCount++
		anum, err = roleIns.GetAddrsNum(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return anum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return anum, nil
	}
}

// GetAddr get address by role index.
func (r *ContractModule) GetAddr(rIndex uint64) (common.Address, error) {
	var addr common.Address

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return addr, err
	}

	retryCount := 0
	if rIndex == 0 {
		return addr, ErrIndexZero
	}
	sum, err := r.GetAddrsNum()
	if err != nil {
		return addr, err
	}
	log.Println("addrs total:", sum)
	if rIndex > sum {
		return addr, ErrOARange
	}

	rIndex-- // get address by array index actually in contract, which is rIndex minus 1

	for {
		retryCount++
		addr, err = roleIns.GetAddr(&bind.CallOpts{
			From: r.addr,
		}, rIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return addr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return addr, nil
	}
}

// GetRoleIndex get the account role index by address.
func (r *ContractModule) GetRoleIndex(addr common.Address) (uint64, error) {
	var rIndex uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return rIndex, err
	}

	retryCount := 0
	for {
		retryCount++
		rIndex, err = roleIns.GetRoleIndex(&bind.CallOpts{
			From: r.addr,
		}, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return rIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return rIndex, nil
	}
}

// GetRoleInfo get account information by address.
func (r *ContractModule) GetRoleInfo(addr common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	var isActive, isBanned bool
	var roleType uint8
	var index, gIndex uint64
	var extra []byte

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return isActive, isBanned, roleType, index, gIndex, extra, err
	}

	retryCount := 0
	for {
		retryCount++
		isActive, isBanned, roleType, index, gIndex, extra, err = roleIns.GetRoleInfo(&bind.CallOpts{
			From: r.addr,
		}, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isActive, isBanned, roleType, index, gIndex, extra, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isActive, isBanned, roleType, index, gIndex, extra, nil
	}
}

// GetGroupsNum get the number of groups.
func (r *ContractModule) GetGroupsNum() (uint64, error) {
	var gnum uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return gnum, err
	}

	retryCount := 0
	for {
		retryCount++
		gnum, err = roleIns.GetGroupsNum(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return gnum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return gnum, nil
	}
}

// GetGroupInfo get group information by gIndex.
func (r *ContractModule) GetGroupInfo(gIndex uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	var isActive, isBanned, isReady bool
	var level uint16
	var size, price *big.Int
	var fsAddr common.Address

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return isActive, isBanned, isReady, level, size, price, fsAddr, err
	}

	retryCount := 0
	if gIndex == 0 {
		return isActive, isBanned, isReady, level, size, price, fsAddr, ErrIndexZero
	}
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		isActive, isBanned, isReady, level, size, price, fsAddr, err = roleIns.GetGroupInfo(&bind.CallOpts{
			From: r.addr,
		}, gIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isActive, isBanned, isReady, level, size, price, fsAddr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isActive, isBanned, isReady, level, size, price, fsAddr, nil
	}
}

// GetAddrGindex get account's address and its gIndex by rIndex.
func (r *ContractModule) GetAddrGindex(rIndex uint64) (common.Address, uint64, error) {
	var addr common.Address
	var gIndex uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return addr, gIndex, err
	}

	if rIndex == 0 {
		return addr, 0, ErrIndex
	}

	retryCount := 0
	for {
		retryCount++
		addr, gIndex, err = roleIns.GetAddrGindex(&bind.CallOpts{
			From: r.addr,
		}, rIndex-1)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return addr, gIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return addr, gIndex, nil
	}
}

// GetGKNum get the number of keepers in the group.
func (r *ContractModule) GetGKNum(gIndex uint64) (uint64, error) {
	var gkNum uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return gkNum, err
	}

	retryCount := 0
	if gIndex == 0 {
		return gkNum, ErrIndexZero
	}
	gIndex--
	for {
		retryCount++
		gkNum, err = roleIns.GetGKNum(&bind.CallOpts{
			From: r.addr,
		}, gIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return gkNum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return gkNum, nil
	}
}

// GetGUPNum get the number of user、providers in the group.
func (r *ContractModule) GetGUPNum(gIndex uint64) (uint64, uint64, error) {
	var gpNum, guNum uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return guNum, gpNum, err
	}

	retryCount := 0
	if gIndex == 0 {
		return guNum, gpNum, ErrIndexZero
	}
	gIndex--
	for {
		retryCount++
		guNum, gpNum, err = roleIns.GetGUPNum(&bind.CallOpts{
			From: r.addr,
		}, gIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return guNum, gpNum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return guNum, gpNum, nil
	}
}

// GetGroupK get keeper role index by gIndex and keeper array index.
func (r *ContractModule) GetGroupK(gIndex uint64, index uint64) (uint64, error) {
	var kIndex uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return kIndex, err
	}

	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return kIndex, err
	}
	if index >= gkNum {
		fmt.Println("the array range is", gkNum)
		return kIndex, ErrOARange
	}

	retryCount := 0
	if gIndex == 0 {
		return kIndex, ErrIndexZero
	}
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		kIndex, err = roleIns.GetGroupK(&bind.CallOpts{
			From: r.addr,
		}, gIndex, index)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return kIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return kIndex, nil
	}
}

// GetGroupP get provider role index by gIndex and provider array index.
func (r *ContractModule) GetGroupP(gIndex uint64, index uint64) (uint64, error) {
	var pIndex uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return pIndex, err
	}

	_, gpNum, err := r.GetGUPNum(gIndex)
	if err != nil {
		return pIndex, err
	}
	if index >= gpNum {
		fmt.Println("the array range is", gpNum)
		return pIndex, ErrOARange
	}

	retryCount := 0
	if gIndex == 0 {
		return pIndex, ErrIndexZero
	}
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		pIndex, err = roleIns.GetGroupP(&bind.CallOpts{
			From: r.addr,
		}, gIndex, index)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pIndex, nil
	}
}

// GetGroupU get user role index by gIndex and user array index.
func (r *ContractModule) GetGroupU(gIndex uint64, index uint64) (uint64, error) {
	var uIndex uint64

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return uIndex, err
	}

	guNum, _, err := r.GetGUPNum(gIndex)
	if err != nil {
		return uIndex, err
	}
	if index >= guNum {
		fmt.Println("the array range is", guNum)
		return uIndex, ErrOARange
	}

	retryCount := 0
	if gIndex == 0 {
		return uIndex, ErrIndexZero
	}
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		uIndex, err = roleIns.GetGU(&bind.CallOpts{
			From: r.addr,
		}, gIndex, index)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return uIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return uIndex, nil
	}
}

// GetRVersion get the version of Role-contract.
func (r *ContractModule) GetRVersion() (uint16, error) {
	var v uint16

	client := getClient(r.endPoint)
	defer client.Close()
	roleIns, err := newRole(r.contractAddress, client)
	if err != nil {
		return v, err
	}

	retryCount := 0
	for {
		retryCount++
		v, err = roleIns.Version(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return v, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return v, nil
	}
}
