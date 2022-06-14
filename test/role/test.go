package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	callconts "memoc/callcontracts"
	iface "memoc/interfaces"
	"memoc/test"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// package level vars
var (
	ethEndPoint string
	//qethEndPoint string

	// tx options
	txopts *callconts.TxOpts

	// addrs and sks of test accounts
	addrs []common.Address
	sks   []string

	// admin account addresses
	adminAddr common.Address
	userAddr  common.Address

	// contract addresses
	roleAddr       common.Address
	rtokenAddr     common.Address
	rolefsAddr     common.Address
	pledgePoolAddr common.Address
	issuAddr       common.Address
	// filesys
	fsAddr1 common.Address
	fsAddr2 common.Address

	// pledge values
	pledgeK *big.Int
	pledgeP *big.Int
	oneEth  *big.Int
	nineEth *big.Int

	fast bool

	// callers
	roleCaller  iface.RoleInfo
	ppCaller    iface.PledgePoolInfo
	rfsCaller   iface.RoleFSInfo
	issuCaller  iface.IssuanceInfo
	erc20Caller iface.ERC20Info

	// group index
	gIndex1 uint64
	gIndex2 uint64

	config map[string]string

	err error
)

// init package level vars
func init() {
	pledgeK = big.NewInt(1e18)
	pledgeP = big.NewInt(1e18)
	oneEth = big.NewInt(1e18)
	nineEth = big.NewInt((9 * 1e18))

	txopts = &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	addrs = []common.Address{
		common.HexToAddress(test.Acc1),
		common.HexToAddress(test.Acc2),
		common.HexToAddress(test.Acc3),
		common.HexToAddress(test.Acc4),
		common.HexToAddress(test.Acc5),
	}
	sks = []string{
		test.Sk1,
		test.Sk2,
		test.Sk3,
		test.Sk4,
		test.Sk5,
	}

	adminAddr = common.HexToAddress(test.AdminAddr)
	userAddr = common.HexToAddress(test.Acc1)

	status := make(chan error)

	// init callers before contracts deployed
	// rAdmin = callconts.NewR(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint)
	// rUser = callconts.NewR(common.Address{}, userAddr, test.Sk1, txopts, ethEndPoint)
	roleCaller = callconts.NewR(common.Address{}, common.Address{}, "", txopts, callconts.EndPoint, status)
	ppCaller = callconts.NewPledgePool(common.Address{}, common.Address{}, "", txopts, callconts.EndPoint, status)
	rfsCaller = callconts.NewRFS(common.Address{}, common.Address{}, "", txopts, callconts.EndPoint, status)
	issuCaller = callconts.NewIssu(common.Address{}, common.Address{}, "", txopts, callconts.EndPoint, status)
}

func main() {

	peth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;") //dev网
	//qeth = flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	pfast := flag.Bool("f", false, "flag for fast test")

	flag.Parse()

	//qethEndPoint = *qeth
	fast = *pfast
	ethEndPoint = *peth
	callconts.EndPoint = *peth

	CheckEthBalance()
	CheckERC20Balance()

	fmt.Println("")
	fmt.Println("============ 1. prepair for testing ============")

	err = PrePaire()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("============ 2. test GetAddrGindex ============")

	// get acc address and gIndex by rIndex
	fmt.Println("test role 2")
	status := make(chan error)
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	acc, gIndex, err := roleCaller.GetAddrGindex(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("acc: ", acc, " gIndex: ", gIndex)

	if acc.String() != test.Acc2 {
		log.Fatal("acc address error")
	}
	if gIndex != 1 {
		log.Fatal("gIndex error")
	}

	fmt.Println("test role 4")
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	acc, gIndex, err = roleCaller.GetAddrGindex(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("acc: ", acc, " gIndex: ", gIndex)

	if acc.String() != test.Acc4 {
		log.Fatal("acc address error")
	}
	if gIndex != 2 {
		log.Fatal("gIndex error")
	}

	fmt.Println("")
	fmt.Println("============ 3. test GetGroupsNum ============")

	// test GetGroupsNum
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	gNum, err := roleCaller.GetGroupsNum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("group number:", gNum)

	if gNum != 2 {
		log.Fatal("GetGroupsNum test failed, group number should be 2")
	}

	fmt.Println("")
	fmt.Println("============ 4. test PledgePool ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	pp, err := roleCaller.PledgePool()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pledge pool:", pp)
	if pp != pledgePoolAddr {
		log.Fatal("test PledgePool failed")
	}

	fmt.Println("")
	fmt.Println("============ 5. test Foundation ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	fd, err := roleCaller.Foundation()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("foundation:", fd)
	if fd != test.Foundation {
		log.Fatal("test Foundation failed")
	}

	fmt.Println("")
	fmt.Println("============ 6. test RToken ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	rt, err := roleCaller.RToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rt:", rt)
	if rt != rtokenAddr {
		log.Fatal("test RToken failed")
	}

	fmt.Println("")
	fmt.Println("============ 7. test Issuance ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	is, err := roleCaller.Issuance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is:", is)
	if is != issuAddr {
		log.Fatal("test Issuance failed")
	}

	fmt.Println("")
	fmt.Println("============ 8. test Rolefs ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	rfs, err := roleCaller.Rolefs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rfs:", rfs)
	if rfs != rolefsAddr {
		log.Fatal("test Rolefs failed")
	}

	fmt.Println("")
	fmt.Println("============ 9. test GetAddrsNum ============")

	//
	if fast {
		fmt.Println("@@ skipped for fast testing")
	} else {
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		an, err := roleCaller.GetAddrsNum()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("an:", an)
		if an != 5 {
			log.Fatal("test GetAddrsNum failed")
		}
	}

	fmt.Println("")
	fmt.Println("============ 10. test GetAddr ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	ga, err := roleCaller.GetAddr(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ga:", ga)
	if ga.String() != test.Acc1 {
		log.Fatal("test GetAddr failed")
	}

	fmt.Println("")
	fmt.Println("============ 11. test GetRoleInfo ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	isActive, isBanned, roleType, index, gIndex, extra, err := roleCaller.GetRoleInfo(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
	fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)

	fmt.Println("")
	fmt.Println("============ 12. test GetGroupsNum ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	gn, err := roleCaller.GetGroupsNum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gn:", gn)
	if gn != 2 {
		log.Fatal("test GetGroupsNum failed")
	}

	fmt.Println("")
	fmt.Println("============ 13. test GetGKNum ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	kNum, err := roleCaller.GetGKNum(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("kNum:", kNum)
	if kNum != 2 {
		log.Fatal("test GetGKNum failed")
	}

	fmt.Println("")
	fmt.Println("============ 14. test GetGUPNum ============")

	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	_, pNum, err := roleCaller.GetGUPNum(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pNum:", pNum)

	// add a provider only when no provider in group
	// add role index 5 into group 1 as a provider
	if pNum == 0 {
		fmt.Println("")
		fmt.Println("============ 15. test AddProviderToGroup ============")

		// caller must be provider himself to avoid sign
		rForProvider := callconts.NewR(roleAddr, common.HexToAddress(test.Acc5), test.Sk5, txopts, ethEndPoint, status)
		err := rForProvider.AddProviderToGroup(5, 1, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		// output:
		// 2021/12/16 14:14:40 begin AddProviderToGroup in Role contract...
		// 2021/12/16 14:14:41 Check Tx hash: 0x6e5b639d01fc6432d8daa59d0d149fc0d62a714bbb23d99cf39e48c31b544461 nonce: 84 gasPrice: 200
		// 2021/12/16 14:14:41 waiting for miner...
		// 2021/12/16 14:14:56 GasUsed: 79781 CumulativeGasUsed: 79781
		// 2021/12/16 14:14:56 AddProviderToGroup in Role has been successful!
	} else {
		fmt.Println(">>>> role index 5 already exists in group 1, skip AddProviderToGroup")
	}

	fmt.Println("")
	fmt.Println("============ 16. test GetGroupK ============")

	// group 1 , keeper[1] == 3
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	ggk, err := roleCaller.GetGroupK(1, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ggk:", ggk)
	if ggk != 3 {
		log.Fatal("test GetGroupK failed")
	}

	fmt.Println("")
	fmt.Println("============ 17. test GetGroupP ============")

	// group 1 , provider[0] == 5
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	ggp, err := roleCaller.GetGroupP(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ggp:", ggp)
	if ggp != 5 {
		log.Fatal("test GetGroupP failed")
	}

	fmt.Println("")
	fmt.Println("============ 18. test PledgeK,PledgeP,SetPledgeMoney ============")

	if fast {
		fmt.Println("@@ skipped for fast testing")
	} else {
		// show old pledge
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		oldPK, err := roleCaller.PledgeK()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledgeK before set: ", oldPK)
		oldPP, err := roleCaller.PledgeP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledgeP before set: ", oldPP)

		// plus pledge by 1 eth
		err = roleCaller.SetPledgeMoney(new(big.Int).Add(oldPK, oneEth), new(big.Int).Add(oldPP, oneEth))
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// show new
		newPK, err := roleCaller.PledgeK()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledgeK after set: ", newPK)
		newPP, err := roleCaller.PledgeP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledgeP after set: ", newPP)

		if new(big.Int).Sub(newPK, oldPK).Cmp(oneEth) != 0 {
			log.Fatal("test set PledgeK failed")
		}

		if new(big.Int).Sub(newPP, oldPP).Cmp(oneEth) != 0 {
			log.Fatal("test set PledgeP failed")
		}
	}

	fmt.Println("")
	fmt.Println("============ 19. test Recharge ============")

	if fast {
		fmt.Println("@@ skipped for fast testing")
	} else {
		// to assure balance is enough
		CheckERC20Balance()

		// check old balance of user in erc20
		erc20Caller = callconts.NewERC20(test.PrimaryToken, userAddr, test.Sk1, txopts, ethEndPoint, status)

		oldBalErc20, err := erc20Caller.BalanceOf(userAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> old balance in erc20 is: ", oldBalErc20)

		// check old balance of user in fs1
		fs1 := callconts.NewFileSys(fsAddr1, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		oldBalFs, tmp, err := fs1.GetBalance(1, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> old balance in fs is: ", oldBalFs, " tmp:", tmp)

		// user approve before recharge
		err = erc20Caller.Approve(fsAddr1, oneEth)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		// 查询user allowance
		allo, err := erc20Caller.Allowance(userAddr, fsAddr1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The allowance of ", userAddr, " to ", fsAddr1, " is ", allo)

		// user recharge 1 eth into fileSys
		roleCaller = callconts.NewR(roleAddr, userAddr, test.Sk1, txopts, ethEndPoint, status)
		usrIndex, err := roleCaller.GetRoleIndex(userAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("user index:", usrIndex)

		// query token index
		rtCaller := callconts.NewRT(rtokenAddr, userAddr, test.Sk1, txopts, ethEndPoint)
		tIndex, ok, err := rtCaller.GetTI(test.PrimaryToken)
		if err != nil {
			log.Fatal(err)
		}
		if !ok {
			log.Fatal("GetTI return invalid")
		}
		fmt.Println("token index:", tIndex)

		// user(acc1) recharge 1 eth into fs, with primary token
		roleCaller = callconts.NewR(roleAddr, userAddr, test.Sk1, txopts, ethEndPoint, status)
		err = roleCaller.Recharge(rtokenAddr, 1, 0, oneEth, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// check new balance of user in erc20
		newBalErc20, err := erc20Caller.BalanceOf(userAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> new balance in erc20 is: ", newBalErc20)

		// check new balance of user in fs
		newBalFs, tmp, err := fs1.GetBalance(1, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> new balance in fs is: ", newBalFs, " tmp:", tmp)

		// test
		if new(big.Int).Sub(oldBalErc20, newBalErc20).Cmp(oneEth) != 0 {
			log.Fatal("test Recharge failed, erc20 balance error")
		}
		if new(big.Int).Sub(newBalFs, oldBalFs).Cmp(oneEth) != 0 {
			log.Fatal("test Recharge failed, fs balance error")
		}
	}

	fmt.Println("")
	fmt.Println("============ 20. test WithdrawFromFs ============")

	if fast {
		fmt.Println("@@ skipped for fast testing")
	} else {
		// to assure balance is enough
		CheckERC20Balance()

		// check old balance of user in erc20
		userAddr = common.HexToAddress(test.Acc1)
		erc20Caller = callconts.NewERC20(test.PrimaryToken, userAddr, test.Sk1, txopts, ethEndPoint, status)
		oldBalErc20, err := erc20Caller.BalanceOf(userAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> old balance in erc20 is: ", oldBalErc20)

		// query user balance before withdraw
		fs1 := callconts.NewFileSys(fsAddr1, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		oldBalFs, tmp, err := fs1.GetBalance(1, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> old balance in fs1 is: ", oldBalFs, " tmp:", tmp)

		// withdraw
		roleCaller = callconts.NewR(roleAddr, userAddr, test.Sk1, txopts, ethEndPoint, status)
		err = roleCaller.WithdrawFromFs(rtokenAddr, 1, 0, oneEth, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// check new balance of user in erc20
		newBalErc20, err := erc20Caller.BalanceOf(userAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> new balance in erc20 is: ", newBalErc20)

		// query new balance in fs1 after withdraw
		newBalFs, tmp, err := fs1.GetBalance(1, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(">>>> old balance in erc20 is: ", newBalFs, " tmp:", tmp)

		// test
		if new(big.Int).Sub(newBalErc20, oldBalErc20).Cmp(oneEth) != 0 {
			log.Fatal("test Withdraw failed, erc20 balance error")
		}
		if new(big.Int).Sub(oldBalFs, newBalFs).Cmp(oneEth) != 0 {
			log.Fatal("test WithdrawFromFs failed")
		}
	}

	fmt.Println("")
	fmt.Println("============ 21. test SignForRegister ============")

	fmt.Println(">>>> test call register by other account")

	if fast {
		fmt.Println("@@ skipped for fast testing")
	} else {
		// signed by acc's sk
		sig, err := callconts.SignForRegister(adminAddr, test.Sk6)
		if err != nil {
			log.Fatal(err)
		}

		// other account call register for regist acc6
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		err = roleCaller.Register(common.HexToAddress(test.Acc6), sig)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		var rIndex uint64
		// check role index
		for retry := 5; retry > 0; retry-- {
			fmt.Println("call GetRoleIndex and store index ")

			roleCaller = callconts.NewR(roleAddr, common.HexToAddress(test.Acc6), test.Sk6, txopts, ethEndPoint, status)
			rIndex, err = roleCaller.GetRoleIndex(common.HexToAddress(test.Acc6))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rIndex: ", rIndex)

			if rIndex != 0 {
				if rIndex != 6 {
					log.Fatal("register acc7 failed, rIndex not 6")
				}
				fmt.Println("register acc7 succeed, rIndex: ", rIndex)
				break
			}

			fmt.Printf("GetRoleIndex return 0, retry %v times.\n", retry)
			// wait 2 seconds for Register to complete
			time.Sleep(time.Duration(2) * time.Second)
		}
		// failed
		if rIndex == 0 {
			log.Fatal("get role index failed after retry 5 times.")
		}
	}

	fmt.Println("")
	fmt.Println("============ 22. test SignForRegisterKeeper ============")

	if fast {
		fmt.Println("@@ skipped for fast test, because role can only register once.")
	} else {

		//--- register acc to get role index
		fmt.Println(">>>> begin register acc7 to get role index")

		acc7 := common.HexToAddress(test.Acc7)
		fmt.Printf("Role address: %s, acc7 address: %s\n", roleAddr.Hex(), test.Acc7)

		roleCaller = callconts.NewR(roleAddr, acc7, test.Sk7, txopts, ethEndPoint, status)
		err = roleCaller.Register(acc7, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		var rIndex uint64
		// check role index
		for retry := 5; retry > 0; retry-- {
			fmt.Println("call GetRoleIndex and store index ")

			roleCaller = callconts.NewR(roleAddr, common.HexToAddress(test.Acc7), test.Sk7, txopts, ethEndPoint, status)
			rIndex, err = roleCaller.GetRoleIndex(acc7)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rIndex: ", rIndex)

			if rIndex != 0 {
				if rIndex != 7 {
					log.Fatal("register acc7 failed, rIndex not 7")
				}

				fmt.Println("register acc7 succeed, rIndex: ", rIndex)
				break
			}

			fmt.Printf("GetRoleIndex return 0, retry %v times.\n", retry)
			// wait 2 seconds for Register to complete
			time.Sleep(time.Duration(2) * time.Second)
		}
		// failed
		if rIndex == 0 {
			log.Fatal("get role index failed after retry 5 times.")
		}

		//--- admin send Erc20 token to acc7
		erc20Caller = callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		bal, err := erc20Caller.BalanceOf(acc7)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc7: ", acc7, " Erc20 balance : ", bal)
		if bal.Cmp(oneEth) < 0 {
			err = erc20Caller.Transfer(acc7, oneEth) // admin给测试账户转账，用于测试（充值或质押）
			if err != nil {
				log.Fatal(err)
			}
			if err = <-status; err != nil {
				log.Fatal(err)
			}
			bal, err = erc20Caller.BalanceOf(acc7)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("after transfer, acc7 ", acc7, " Erc20 balance: ", bal)
		}

		//--- pledge money for register role
		fmt.Println(">>>> pledge acc7 as keeper")
		err = toPledge(acc7, roleAddr, pledgePoolAddr, test.Sk7, rIndex, pledgeK, txopts)
		if err != nil {
			log.Fatal(err)
		}

		//--- set pledge money to 1 eth
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		err = roleCaller.SetPledgeMoney(oneEth, oneEth)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		fmt.Println("plege money for keeper and provider have been set to 1 eth")

		//--- call SignForRegisterKeeper
		fmt.Println("call SignForRegisterKeeper")
		// get signature, call with admin, register acc7, blsKey is nil here
		sig, err := callconts.SignForRegisterKeeper(adminAddr, nil, test.Sk7)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sig: %x\n", sig)

		fmt.Println("call RegisterKeeper")
		// set plege pool addr to package level var for register keeper
		//callconts.PledgePoolAddr = pledgePoolAddr
		// other account call register for regist acc7
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		// rIndex = 7, blsKey = nil
		err = roleCaller.RegisterKeeper(pledgePoolAddr, 7, nil, sig)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// get role info
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		isActive, isBanned, roleType, index, gIndex, extra, err = roleCaller.GetRoleInfo(common.HexToAddress(test.Acc7))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
		fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)
	}

	fmt.Println("")
	fmt.Println("============ 23. test SignForRegisterProvider and SignForAddProviderToGroup ============")

	if fast {
		fmt.Println("@@ skipped for fast test, because role can only register once.")
	} else {
		//--- register acc to get role index
		fmt.Println(">>>> begin register acc8 to get role index")

		acc8 := common.HexToAddress(test.Acc8)
		fmt.Printf("Role address: %s, acc8 address: %s\n", roleAddr.Hex(), test.Acc8)

		roleCaller = callconts.NewR(roleAddr, acc8, test.Sk8, txopts, ethEndPoint, status)
		err = roleCaller.Register(acc8, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		var rIndex uint64
		// check role index
		for retry := 5; retry > 0; retry-- {
			fmt.Println("call GetRoleIndex and store index ")
			roleCaller = callconts.NewR(roleAddr, common.HexToAddress(test.Acc8), test.Sk8, txopts, ethEndPoint, status)
			rIndex, err = roleCaller.GetRoleIndex(common.HexToAddress(test.Acc8))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rIndex: ", rIndex)

			if rIndex != 0 {
				if rIndex != 8 {
					log.Fatal("register acc8 failed, rIndex not 8")
				}
				fmt.Println("register acc8 succeed, rIndex: ", rIndex)
				break
			}

			fmt.Printf("GetRoleIndex return 0, retry %v times.\n", retry)
			// wait 2 seconds for Register to complete
			time.Sleep(time.Duration(2) * time.Second)
		}
		// failed
		if rIndex == 0 {
			log.Fatal("get role index failed after retry 5 times.")
		}

		//--- admin send Erc20 token to acc8
		erc20Caller = callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		bal, err := erc20Caller.BalanceOf(acc8)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc8: ", acc8, " Erc20 balance : ", bal)
		if bal.Cmp(oneEth) < 0 {
			err = erc20Caller.Transfer(acc8, oneEth) // admin给测试账户转账，用于测试（充值或质押）
			if err != nil {
				log.Fatal(err)
			}
			if err = <-status; err != nil {
				log.Fatal(err)
			}
			bal, err = erc20Caller.BalanceOf(acc8)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("after transfer, acc8 ", acc8, " Erc20 balance: ", bal)
		}

		//--- pledge money for register role
		fmt.Println(">>>> pledge acc8 as provider")
		err = toPledge(acc8, roleAddr, pledgePoolAddr, test.Sk8, rIndex, pledgeK, txopts)
		if err != nil {
			log.Fatal(err)
		}

		//--- set pledge money to 1 eth
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		err = roleCaller.SetPledgeMoney(oneEth, oneEth)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		fmt.Println("plege money for keeper and provider has been set to 1 eth")

		//--- call SignForRegisterProvider
		fmt.Println("call SignForRegisterProvider")
		// get signature, call with admin, register acc8
		sig, err := callconts.SignForRegisterProvider(adminAddr, test.Sk8)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sig: %x\n", sig)

		fmt.Println("call RegisterProvider")
		// set plege pool addr to package level var for register keeper
		callconts.PledgePoolAddr = pledgePoolAddr
		// admin as caller
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		// rIndex = 8
		err = roleCaller.RegisterProvider(pledgePoolAddr, 8, sig)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// get role info
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		isActive, isBanned, roleType, index, gIndex, extra, err = roleCaller.GetRoleInfo(common.HexToAddress(test.Acc8))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
		fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)

		if roleType != 2 {
			log.Println("test SignForRegisterProvider failed, roleType not 2")
		}
		if index != 8 {
			log.Println("test SignForRegisterProvider failed, index not 8")
		}

		//--- begin test SignForAddProviderToGroup

		//--- call SignForAddProviderToGroup
		fmt.Println("call SignForAddProviderToGroup")
		// get signature, call with admin,  group 1, register acc8
		sig, err = callconts.SignForAddProviderToGroup(adminAddr, 1, test.Sk8)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sig: %x\n", sig)

		fmt.Println("call AddProviderToGroup")
		// admin as caller
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		// rIndex = 8, gIndex = 1
		err = roleCaller.AddProviderToGroup(8, 1, sig)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// get role info
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		isActive, isBanned, roleType, index, gIndex, extra, err = roleCaller.GetRoleInfo(common.HexToAddress(test.Acc8))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
		fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)

		if gIndex != 1 {
			log.Fatal("test SignForAddProviderToGroup failed, gIndex not 1")
		}
	}

	fmt.Println("")
	fmt.Println("============ 24. test SignForRegisterUser ============")

	if fast {
		fmt.Println("@@ skipped for fast test, because role can only register once.")
	} else {
		//--- register acc to get role index
		fmt.Println(">>>> begin register acc9 to get role index")

		acc9 := common.HexToAddress(test.Acc9)
		fmt.Printf("Role address: %s, acc9 address: %s\n", roleAddr.Hex(), test.Acc9)

		roleCaller = callconts.NewR(roleAddr, acc9, test.Sk9, txopts, ethEndPoint, status)
		err = roleCaller.Register(acc9, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		// check role index
		fmt.Println("call GetRoleIndex")
		rIndex, err := roleCaller.GetRoleIndex(acc9)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("registered rIndex:", rIndex)

		if rIndex != 9 {
			log.Fatal("register acc9 failed, rIndex not 9")
		}

		//--- call SignForRegisterUser
		fmt.Println("call SignForRegisterUser")
		// get signature, call with admin,  group 1, token 0, blsKey nil, register acc9
		sig, err := callconts.SignForRegisterUser(adminAddr, 1, nil, test.Sk9)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sig: %x\n", sig)

		fmt.Println("call RegisterUser")
		// set plege pool addr to package level var for register role
		callconts.PledgePoolAddr = pledgePoolAddr
		// admin as caller
		roleCaller := callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		// rIndex = 9
		//roleAddr, rTokenAddr common.Address, index uint64, gindex uint64, tindex uint32, blskey []byte, sign []byte
		err = roleCaller.RegisterUser(rtokenAddr, 9, 1, nil, sig)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}

		// get role info
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		isActive, isBanned, roleType, index, gIndex, extra, err = roleCaller.GetRoleInfo(common.HexToAddress(test.Acc9))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
		fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)

		if roleType != 1 {
			log.Fatal("test failed, roleType not 1")
		}
	}

	fmt.Println("")
	fmt.Println("============ 25. test SignForRecharge ============")

	// recharge only used by user

	// if fast {
	// 	fmt.Println("@@ skipped for fast test")
	// } else {

	// to assure balance is enough
	CheckERC20Balance()

	// check old balance of user in erc20
	userAddr = common.HexToAddress(test.Acc1)
	erc20Caller = callconts.NewERC20(test.PrimaryToken, userAddr, test.Sk1, txopts, ethEndPoint, status)
	oldBalErc20, err := erc20Caller.BalanceOf(userAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> old balance in erc20 is: ", oldBalErc20)

	// check old balance of user in fs
	fs1 := callconts.NewFileSys(fsAddr1, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	oldBalFs, tmp, err := fs1.GetBalance(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> old balance in fs is: ", oldBalFs, " tIndex:", tmp)

	// user approve before recharge
	err = erc20Caller.Approve(fsAddr1, oneEth)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	// 查询user allowance
	allo, err := erc20Caller.Allowance(userAddr, fsAddr1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The allowance of ", userAddr, " to ", fsAddr1, " is ", allo)

	// user recharge 1 eth into fileSys
	roleCaller = callconts.NewR(roleAddr, userAddr, test.Sk1, txopts, ethEndPoint, status)
	usrIndex, err := roleCaller.GetRoleIndex(userAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user index:", usrIndex)

	// query token index
	rtCaller := callconts.NewRT(rtokenAddr, userAddr, test.Sk1, txopts, ethEndPoint)
	tIndex, ok, err := rtCaller.GetTI(test.PrimaryToken)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		log.Fatal("GetTI return invalid")
	}
	fmt.Println("token index:", tIndex)

	//--- call SignForRecharge

	fmt.Println("call SignForRecharge")
	//caller common.Address, uIndex uint64, tIndex uint32, money *big.Int, accSk string
	sig, err := callconts.SignForRecharge(adminAddr, 1, 0, oneEth, test.Sk1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sig: %x\n", sig)

	fmt.Println("call Recharge with sig")

	// recharge

	// user(acc1) recharge 1 eth into fs, with primary token
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	err = roleCaller.Recharge(rtokenAddr, 1, 0, oneEth, sig)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	// check new balance of user in erc20
	newBalErc20, err := erc20Caller.BalanceOf(userAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> new balance in erc20 is: ", newBalErc20)

	// check new balance of user in fs
	newBalFs, tmp, err := fs1.GetBalance(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> new balance in fs is: ", newBalFs, " tIndex:", tmp)

	// test
	if new(big.Int).Sub(oldBalErc20, newBalErc20).Cmp(oneEth) != 0 {
		log.Fatal("test Recharge failed, erc20 balance error")
	}
	if new(big.Int).Sub(newBalFs, oldBalFs).Cmp(oneEth) != 0 {
		log.Fatal("test Recharge failed, fs balance error")
	}

	fmt.Println("")
	fmt.Println("============ 26. test SignForWithdrawFromFs ============")

	// if fast {
	// 	fmt.Println("@@ skipped for fast test")
	// } else {

	// to assure balance is enough
	CheckERC20Balance()

	// check old balance of user in erc20
	erc20Caller = callconts.NewERC20(test.PrimaryToken, userAddr, test.Sk1, txopts, ethEndPoint, status)
	oldBalErc20, err = erc20Caller.BalanceOf(userAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> old balance in erc20 is: ", oldBalErc20)

	// query user balance before withdraw
	fs1 = callconts.NewFileSys(fsAddr1, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	oldBalFs, tmp, err = fs1.GetBalance(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> old balance in fs is: ", oldBalFs, " tIndex:", tmp)

	//--- call SignForWithdrawFromFs

	fmt.Println("call SignForWithdrawFromFs")
	//caller common.Address, tIndex uint32, amount *big.Int, accSk string
	sig, err = callconts.SignForWithdrawFromFs(adminAddr, 0, oneEth, test.Sk1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sig: %x\n", sig)

	fmt.Println("call withdraw with sig")

	// withdraw
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	err = roleCaller.WithdrawFromFs(rtokenAddr, 1, 0, oneEth, sig)
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}

	// check new balance of user in erc20
	newBalErc20, err = erc20Caller.BalanceOf(userAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> new balance in erc20 is: ", newBalErc20)

	// query new balance in fs after withdraw
	newBalFs, tmp, err = fs1.GetBalance(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> new balance in fs is: ", newBalFs, " tIndex:", tmp)

	// test
	if new(big.Int).Sub(newBalErc20, oldBalErc20).Cmp(oneEth) != 0 {
		log.Fatal("test Withdraw failed, erc20 balance error")
	}
	if new(big.Int).Sub(oldBalFs, newBalFs).Cmp(oneEth) != 0 {
		log.Fatal("test WithdrawFromFs failed")
	}

	fmt.Println("")
	fmt.Println("============test success!============")

	// if success and not fast test, save new config into config file
	if !fast {
		fmt.Println("saving config..")
		cfg := fmt.Sprintf("role=%s\nrtoken=%s\nrolefs=%s\npledgePool=%s\nissuance=%s", roleAddr, rtokenAddr, rolefsAddr, pledgePoolAddr, issuAddr)
		SaveConfig(cfg)
	}
}

// approve before pledge
func toPledge(addr, roleAddr, pledgePoolAddr common.Address, sk string, rindex uint64, pledgek *big.Int, txopts *callconts.TxOpts) error {

	// 调用pledge前需要先approve
	status := make(chan error)
	erc20 := callconts.NewERC20(test.PrimaryToken, addr, sk, txopts, ethEndPoint, status)
	err := erc20.Approve(pledgePoolAddr, pledgek)
	if err != nil {
		return err
	}
	if err = <-status; err != nil {
		return err
	}
	// 查询allowance
	allow, err := erc20.Allowance(addr, pledgePoolAddr)
	if err != nil {
		return err
	}
	fmt.Println("The allowance of ", addr, " to ", pledgePoolAddr, " is ", allow)

	// 质押
	ppCaller = callconts.NewPledgePool(pledgePoolAddr, addr, sk, txopts, ethEndPoint, status)
	err = ppCaller.Pledge(test.PrimaryToken, roleAddr, rindex, pledgek, nil)
	if err != nil {
		return err
	}
	if err = <-status; err != nil {
		return err
	}
	return nil
}

// read contract addresses from config file
func ReadConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

// save contract addresses into config file if test success
func SaveConfig(cfg string) error {
	err := ioutil.WriteFile("contracts.ini", []byte(cfg), 0666)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// check accounts' eth balance, used for sending tx, recharge if not enough
func CheckEthBalance() {

	fmt.Println(">>>> checking eth balance")

	ethBal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Printf("admin balance: %x in Ethereum\n", ethBal)
	for i, acc := range addrs {
		ethBal = callconts.QueryEthBalance(acc.Hex(), ethEndPoint)
		fmt.Println("acc", i, " balance: ", ethBal, " in Ethereum")
	}
}

// check each account's balance, increase it if not enough
func CheckERC20Balance() {

	fmt.Println(">>>> checking balance of admin, 9 eth at least")

	// 查看测试账户的ERC20代币余额，不足时自动充值
	status := make(chan error)
	erc20 := callconts.NewERC20(test.PrimaryToken, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

	// 确保admin账户的ERC20代币余额充足（至少9 eth）
	bal, err := erc20.BalanceOf(adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("admin balance in primaryToken is ", bal)

	fmt.Println(">>>> checking balance of accounts, 1 eth each")
	// 确保每个测试账户的ERC20代币余额充足（不少于Pledge值，默认1 eth）
	tNum := 0
	var errT error
	for i, addr := range addrs {
		bal, err = erc20.BalanceOf(addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc", i, " balance in primaryToken is ", bal)
		if bal.Cmp(oneEth) < 0 {
			err = erc20.Transfer(addr, oneEth) // admin给测试账户转账，用于测试（充值或质押）
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
		log.Fatal(err)
	}
}

// prepare everything before starting test
func PrePaire() (err error) {

	status := make(chan error)
	// read contract addresses from config file for fast test
	if fast {
		fmt.Println(">>>> Fast test read contract addresses from config file")

		config = ReadConfig("./contracts.ini")

		// read contract address
		roleAddr = common.HexToAddress(config["role"])
		rtokenAddr = common.HexToAddress(config["rtoken"])
		rolefsAddr = common.HexToAddress(config["rolefs"])
		pledgePoolAddr = common.HexToAddress(config["pledgePool"])
		issuAddr = common.HexToAddress(config["issuance"])

		// update rfs caller after deployed
		// rfsCaller = callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint)
		// // update pledge pool caller after deployed
		// ppCaller = callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint)
		// // update issuance address after deployed
		// issuCaller = callconts.NewIssu(issuAddr, adminAddr, test.AdminSk, txopts, ethEndPoint)
	}

	// deploy: Role、RoleFS、PledgePool、CreateGroup(FileSys)

	// deploy Role
	if fast {
		fmt.Println("@@ fast test skip deploy Role, use existing address")
	} else {
		fmt.Println(">>>> begin deploy Role")

		// deploy Role
		roleCaller = callconts.NewR(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		roleAddr, _, err = roleCaller.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP, 1)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The Role contract address: ", roleAddr.Hex())

		// get and store RToken
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		rtokenAddr, err = roleCaller.RToken()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The RToken contract address: ", rtokenAddr.Hex())

		// deploy RoleFS
		rfsCaller = callconts.NewRFS(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		rolefsAddr, _, err = rfsCaller.DeployRoleFS()
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The RoleFS contract address: ", rolefsAddr.Hex())
		// update rfs caller after roleFS deployed
		rfsCaller = callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	}

	// register accounts： User=[acc1] Keeper=[acc2, acc3, acc4] Provider=[acc5]
	var rIndexes = make([]uint64, 5)
	if fast {
		fmt.Println("@@ fast test skip register accounts, use default role indexes")
		rIndexes = []uint64{1, 2, 3, 4, 5}
	} else {
		fmt.Println(">>>> begin register accounts to get role indexes")
		for i, addr := range addrs {
			fmt.Printf("Role address: %s, register acc: %s\n", roleAddr.Hex(), addr)
			roleCaller = callconts.NewR(roleAddr, addr, sks[i], txopts, ethEndPoint, status)
			err = roleCaller.Register(addr, nil)
			if err != nil {
				log.Fatal(err)
			}
			if err = <-status; err != nil {
				log.Fatal(err)
			}
			// get index
			for retry := 3; retry > 0; retry-- {
				fmt.Println("call GetRoleIndex and store index into rIndexes[]")
				rIndexes[i], err = roleCaller.GetRoleIndex(addr)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("rIndexes[i]: ", rIndexes[i])
				// got correct index value
				if rIndexes[i] != 0 {
					break
				}

				fmt.Printf("GetRoleIndex return 0, retry %v times.\n", retry)
				// wait 2 seconds for Register to complete
				time.Sleep(time.Duration(2) * time.Second)
			}
		}
	}

	fmt.Println("Role indexes is ", rIndexes)

	var p iface.PledgePoolInfo
	// deploy pledge pool
	if fast {
		fmt.Println("@@ fast test skip deloy pledge pool, use existing contract")
	} else {
		fmt.Println(">>>> begin deploy PledgePool")
		ppCaller = callconts.NewPledgePool(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		pledgePoolAddr, _, err = ppCaller.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The PledgePool contract address: ", pledgePoolAddr.Hex())
		// update pledge pool caller after deployed
		ppCaller = callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	}

	// get pledge
	var pledgek *big.Int
	var pledgep *big.Int
	if fast {
		fmt.Println("@@ fast test skip get Pledge, use default: 1 eth")
	} else {
		fmt.Println(">>>> Getting min pledge from Role")
		// get min Pledge for keeper and provider
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		pledgek, err = roleCaller.PledgeK() // 申请Keeper最少需质押的金额
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("min pledge for applying keeper: ", pledgek)
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		pledgep, err = roleCaller.PledgeP() // 申请Provider最少需质押的金额
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("min pledge for applying provider: ", pledgep)
	}

	// deploy Issuance
	if fast {
		fmt.Println("@@ fast test skip deploy Issuance")
	} else {
		fmt.Println(">>>> begin deploy Issuance")

		// deploy issuance contract
		issuCaller = callconts.NewIssu(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		issuAddr, _, err = issuCaller.DeployIssuance(rolefsAddr)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The Issuance contract address is: ", issuAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1
		// update issuance address after deployed
		issuCaller = callconts.NewIssu(issuAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

		// 给Role合约指定所有相关合约地址
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		err = roleCaller.SetPI(pledgePoolAddr, issuAddr, rolefsAddr)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
	}

	// pledge
	if fast {
		fmt.Println("@@ fast test skip Pledge, use existing")
	} else {
		fmt.Println(">>>> begin keepers pledge")
		for i, rindex := range rIndexes[1:4] {
			fmt.Println("params:")
			fmt.Println(
				"addr:", addrs[i+1],
				" roleAddr:", roleAddr,
				" pledgePoolAddr:", pledgePoolAddr,
				" sk:", sks[i+1],
				" rindex:", rindex,
				" pledgek:", pledgek,
				" txopts:", txopts,
			)
			err = toPledge(addrs[i+1], roleAddr, pledgePoolAddr, sks[i+1], rindex, pledgek, txopts)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(">>>> begin provider pledge")
		err = toPledge(addrs[4], roleAddr, pledgePoolAddr, sks[4], rIndexes[4], pledgep, txopts)
		if err != nil {
			log.Fatal(err)
		}
	}

	// keepers and providers register role
	if fast {
		fmt.Println("@@ fast test skip register roles, use existing roles")
	} else {
		fmt.Println(">>>> begin register keepers")

		callconts.PledgePoolAddr = pledgePoolAddr
		rkNum := 0
		var errT error
		for i, rindex := range rIndexes[1:4] {
			rKeeper := callconts.NewR(roleAddr, addrs[i+1], sks[i+1], txopts, ethEndPoint, status)
			fmt.Println(addrs[i+1].Hex(), " begin to register Keeper...")

			// query admin's ERC20 balance in pledgepool
			fmt.Println("pledgePoolAddr: ", pledgePoolAddr, " rindex:", rindex)
			p = callconts.NewPledgePool(pledgePoolAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
			bal, err := p.GetBalanceInPPool(rindex, 0)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rindex ", rindex, " balance in PledgePool is ", bal)

			// get keeper role info
			_, _, roleType, index, _, _, err := rKeeper.GetRoleInfo(addrs[i+1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The rindex", rindex, " information in Role contract, roleType:", roleType, " index:", index)

			// register keeper role
			if roleType == 0 {
				err = rKeeper.RegisterKeeper(pledgePoolAddr, rindex, []byte("Hello, test"), nil)
				if err != nil {
					for j := 0; j < rkNum; j++ {
						<-status
					}
					log.Fatal(err)
				}
				rkNum++
			}
		}
		for i := 0; i < rkNum; i++ {
			err = <-status
			if err != nil {
				errT = err
			}
		}
		if errT != nil {
			log.Fatal(errT)
		}

		fmt.Println(">>>> begin register provider ")

		rProvider := callconts.NewR(roleAddr, addrs[4], sks[4], txopts, ethEndPoint, status)
		// isActive, isBanned, roleType, index, gIndex, extra
		isActive, _, roleType, index, gIndex, _, err := rProvider.GetRoleInfo(addrs[4])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("provider info: ")
		fmt.Println("isActive:", isActive, " roleType:", roleType, " index:", index, " gIndex:", gIndex)

		// register provider
		if roleType == 0 {
			err = rProvider.RegisterProvider(pledgePoolAddr, rIndexes[4], nil)
			if err != nil {
				log.Fatal(err)
			}
			if err = <-status; err != nil {
				log.Fatal(err)
			}
		}
	}

	// create groups
	if fast {
		fmt.Println("@@ fast test with existing group index 1,2")
		gIndex1 = 1
		gIndex2 = 2
	} else {
		// create group, keeper1 keeper2 in group1, keeper3 in group2
		fmt.Println(">>>> begin create group 1")
		// 需要admin先调用CreateGroup,同时将部署FileSys合约
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		gIndex1, err = roleCaller.CreateGroup(rolefsAddr, uint64(0), rIndexes[1:3], 2)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		if gIndex1 != 1 {
			log.Fatal("gIndex1 should be 1")
		}

		// store fsAddr of group1
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		_, _, _, _, _, _, fsAddr1, err = roleCaller.GetGroupInfo(1)
		if err != nil {
			log.Fatal("get group1 info failed")
		}
		fmt.Println("fs1 addr: ", fsAddr1)

		// test create group 2 with inActive, level=2, but only 1 keeper involved
		fmt.Println(">>>> begin create group 2")
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		gIndex2, err = roleCaller.CreateGroup(rolefsAddr, uint64(0), rIndexes[3:4], 2)
		if err != nil {
			log.Fatal(err)
		}
		if err = <-status; err != nil {
			log.Fatal(err)
		}
		if gIndex2 != 2 {
			log.Fatal("gIndex2 should be 2")
		}
		// store fsAddr of group2
		roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
		_, _, _, _, _, _, fsAddr2, err = roleCaller.GetGroupInfo(2)
		if err != nil {
			log.Fatal("get group2 info failed")
		}
		fmt.Println("fs2 addr: ", fsAddr2)
	}

	// register user role
	if fast {
		fmt.Println("@@ fast test skip register user, use existing user")
	} else {
		fmt.Println(">>>> begin register user")
		// role caller for user
		rUser := callconts.NewR(roleAddr, addrs[0], sks[0], txopts, ethEndPoint, status)
		// isActive, isBanned, roleType, index, gIndex, extra
		isActive, _, roleType, index, gIndex, _, err := rUser.GetRoleInfo(addrs[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("user info: ")
		fmt.Println("isActive:", isActive, " roleType:", roleType, " index:", index, " gIndex:", gIndex)

		// register user
		if roleType == 0 {
			err = rUser.RegisterUser(rtokenAddr, rIndexes[0], 1, nil, nil)
			if err != nil {
				log.Fatal(err)
			}
			if err = <-status; err != nil {
				log.Fatal(err)
			}
		}
	}

	// show info
	{
		fmt.Println(">>>> Show Info:")
		fmt.Println("rIndexes: ", rIndexes)
		fmt.Println("addrs: ", addrs)
		fmt.Println("gIndex1: ", gIndex1)
		fmt.Println("gIndex2: ", gIndex2)
		fmt.Println("role: ", roleAddr)
		fmt.Println("rtoken: ", rtokenAddr)
		fmt.Println("rolefs: ", rolefsAddr)
		fmt.Println("pledgePool: ", pledgePoolAddr)
		fmt.Println("issuance: ", issuAddr)
	}

	var (
		//isActive, isBanned, isReady, level, _size, price, fsAddr1, err
		isActive bool
		isBanned bool
		isReady  bool
		level    uint16
		_size    *big.Int
		price    *big.Int
	)
	// get group1 info
	fmt.Println(">>>> getting group1 info")
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	isActive, isBanned, isReady, level, _size, price, fsAddr1, err = roleCaller.GetGroupInfo(gIndex1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr1.Hex())

	if !isActive {
		log.Fatal("group1 should be active")
	}

	// get group2 info
	fmt.Println(">>>> getting group2 info")
	roleCaller = callconts.NewR(roleAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	isActive, isBanned, isReady, level, _size, price, fsAddr2, err = roleCaller.GetGroupInfo(gIndex2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr2.Hex())

	if isActive {
		log.Fatal("group2 should be inactive")
	}

	return nil
}
