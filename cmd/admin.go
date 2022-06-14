package cmd

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	callconts "memoc/callcontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	errSStr = errors.New("trans string amount to bigInt error")
)

// AdminCmd is about contracts functions called by admin
// admin value is optional. [completed]
var AdminCmd = &cli.Command{
	Name:  "admin",
	Usage: "Admin deploy/call contracts",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "adminAddr",
			Aliases: []string{"aa"},
			Value:   callconts.AdminAddr.Hex(), //默认值为common.go中的admin账户地址
			Usage:   "the admin account's address",
		},
		&cli.StringFlag{
			Name:    "adminSk",
			Aliases: []string{"as"},
			Value:   "", //默认值为空，需手动输入
			Usage:   "the admin account's secretkey",
		},
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   callconts.EndPoint, //默认值为common.go中的endPoint
			Usage:   "the geth endPoint",
		},
	},
	Subcommands: []*cli.Command{
		// erc20
		deployERC20Cmd,
		mintCmd,
		burnCmd,
		setUpRoleCmd,
		revokeRoleCmd,
		pauseCmd,
		unpauseCmd,
		// role
		deployRoleCmd,
		setPICmd,
		registerTokenCmd,
		createGroupCmd, // 同时会deployFileSys并且setGF
		addKeeperToGroupCmd,
		setPledgeMoneyCmd,
		alterOwnerCmd,
		// issuance
		deployIssuanceCmd,
		// rolefs
		deployRolefsCmd,
		setAddrCmd,
		// fileSys
		deployFileSysCmd,
		// pledgePool
		deployPledgePoolCmd,
	},
}

var deployERC20Cmd = &cli.Command{
	Name:      "deperc",
	Usage:     "ERC20 - Admin deploy ERC20 contract for primary token.",
	ArgsUsage: "<name symbol>",
	Description: `
A function of ERC20 contract.
Admin deploy an ERC20 standard contract for the Primary Token used in memo.
Index of Primary Token is 0, other erc20 tokens start from 1, and increases by one after each.

Arguments:
name - the name of newly deployed token
symbol - the symbol of newly deployed token
version - the version of ERC20 contract
addrs - the addrs that used to mint with multi-signature, input 5 addresses separated by commas
	`,
	Action: func(cctx *cli.Context) error {
		name := cctx.Args().Get(0)
		symbol := cctx.Args().Get(1)
		version := cctx.Args().Get(2)
		addrs := cctx.Args().Get(3)
		fmt.Println("name:", name, " symbol:", symbol, " version:", version, " addrs:", addrs)

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		// erc20 caller
		e := callconts.NewERC20(addr, addr, sk, txopts, endPoint, status)
		ver, err := strconv.Atoi(version)
		if err != nil {
			return err
		}
		as := strings.Split(addrs, ",")
		asAddr := [5]common.Address{}
		for i := 0; i < 5; i++ {
			asAddr[i] = common.HexToAddress(as[i])
		}
		erc20Addr, _, err := e.DeployERC20(name, symbol, uint16(ver), asAddr)
		if err != nil {
			return err
		}

		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("\nERC20 contract address:", erc20Addr.Hex())

		return nil
	},
}

var mintCmd = &cli.Command{
	Name:      "mint",
	Usage:     "ERC20 - Admin mint ERC20 token. ",
	ArgsUsage: "<target amount>",
	Description: `
Mint is a function in ERC20 contract.
Admin mint some token for target.

Arguments:
target - the account address to receive minted token
amount - the number of token to be minted
signs - signature information for 5 accounts, separated by commas
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hex(), //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		target := cctx.Args().Get(0)
		amount := cctx.Args().Get(1)
		signs := cctx.Args().Get(2)
		fmt.Println("amount:", amount, ", target:", target, " signs:", signs)

		// TODO:check target
		if len(target) != 42 || target == callconts.InvalidAddr {
			fmt.Println("target should be with prefix 0x and shouldn't be 0x0")
			return nil
		}

		mintValue := big.NewInt(0)
		mintValue, ok := mintValue.SetString(amount, 10)
		if !ok {
			return errSStr
		}
		if mintValue.Cmp(big.NewInt(0)) <= 0 {
			fmt.Println("mintValue should be more than 0")
			return nil
		}

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		status := make(chan error)
		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, status)
		ss := strings.Split(signs, ",")
		signByte := [5][]byte{}
		for i, s := range ss {
			signByte[i] = []byte(s)
		}
		err := e.MintToken(common.HexToAddress(target), mintValue, signByte)
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}

		return nil
	},
}

var burnCmd = &cli.Command{
	Name:      "burn",
	Usage:     "ERC20 - Admin burn ERC20 token.",
	ArgsUsage: "<amount>",
	Description: `
Burn is a function in ERC20 contract.
Admin burn some ERC20 token, the number of burned is specifiied by arg amount.

Arguments:
amount - the number of token to be burned.
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hex(), //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		amount := cctx.Args().Get(0)
		fmt.Println("amount:", amount)

		burnValue := big.NewInt(0)
		burnValue, ok := burnValue.SetString(amount, 10)
		if !ok {
			return errSStr
		}
		if burnValue.Cmp(big.NewInt(0)) <= 0 {
			fmt.Println("burnValue should be more than 0")
			return nil
		}

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		status := make(chan error)
		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, status)
		err := e.Burn(burnValue)
		if err != nil {
			return err
		}

		err = <-status
		if err != nil {
			return err
		}

		return nil
	},
}

var setUpRoleCmd = &cli.Command{
	Name:      "setrole",
	Usage:     "ERC20 - Admin setup Role in ERC20 token. ",
	ArgsUsage: "<rfsAddr role>",
	Description: `
SetupRole is a function in ERC20 contract.
Admin setup role for a target in the primary token. 

Arguments:
target - the address of the account to be setup.
role - specifys the role of target, 0|1|2 stand for DEFAULT_ADMIN_ROLE|MINTER_ROLE|PAUSER_ROLE respectively.
	`,
	Action: func(cctx *cli.Context) error {
		target := cctx.Args().Get(0)
		role := cctx.Args().Get(1)
		r, err := strconv.Atoi(role)
		if err != nil {
			fmt.Println("trans string", role, " to int error")
			return err
		}
		if len(target) != 42 || target == callconts.InvalidAddr {
			fmt.Println("target should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		if r != 1 && r != 2 && r != 0 {
			return errors.New("role should be 0(DEFAULT_ADMIN_ROLE)、1(MINTER_ROLE)、2(PAUSER_ROLE)")
		}
		fmt.Println("target:", target, ", role:", r)
		txopts := &callconts.TxOpts{Nonce: nil, GasPrice: big.NewInt(callconts.DefaultGasPrice), GasLimit: callconts.DefaultGasLimit}
		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)
		status := make(chan error)
		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, status)
		err = e.SetUpRole(uint8(r), common.HexToAddress(target))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
	Flags:                  []cli.Flag{&cli.StringFlag{Name: "erc20Addr", Aliases: []string{"e"}, Value: callconts.ERC20Addr.Hex(), Usage: "the ERC20 contract address"}},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

var revokeRoleCmd = &cli.Command{
	Name:      "revrole",
	Usage:     "ERC20 - Admin revoke target's Role in ERC20 token. ",
	ArgsUsage: "<target, role>",
	Description: `
RevokeRole is a function in ERC20 contract.
Revoke target's role in primary token by admin. 

Arguments:
target - the account address to be revoked.
ole - the role code to be set. 0|1|2 stands for DEFAULT_ADMIN_ROLE|MINTER_ROLE|PAUSER_ROLE respectively.
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hex(), //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		target := cctx.Args().Get(0)
		role := cctx.Args().Get(1)
		r, err := strconv.Atoi(role)
		if err != nil {
			fmt.Println("trans string", role, " to int error")
			return err
		}
		if len(target) != 42 || target == callconts.InvalidAddr {
			fmt.Println("target should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		if r != 1 && r != 2 && r != 0 {
			return errors.New("role should be 0(DEFAULT_ADMIN_ROLE)、1(MINTER_ROLE)、2(PAUSER_ROLE)")
		}
		fmt.Println("target:", target, ", role:", r)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		status := make(chan error)
		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, status)
		err = e.RevokeRole(uint8(r), common.HexToAddress(target))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}

		return nil
	},
}

var pauseCmd = &cli.Command{
	Name:      "pause",
	Usage:     "ERC20 - Admin set true to prohibit transfer operation in erc20. ",
	ArgsUsage: " ",
	Description: `
Prohibit the transfer operation in primary token temporaly, the caller must have PAUSER_ROLE right.

Arguments:
no arguments
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hex(), //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, status)
		err := e.Pause()
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var unpauseCmd = &cli.Command{
	Name:      "unpause",
	Usage:     "ERC20 - Admin set false to allow transfer operation in erc20. ",
	ArgsUsage: " ",
	Description: `
Allow transfer operation in primary token. Called by who has PAUSER_ROLE

Arguments:
no arguments
`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hex(), //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, status)
		err := e.Unpause()
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var deployRoleCmd = &cli.Command{
	Name:      "deprole",
	Usage:     "Role - Admin deploy Role contract. ",
	ArgsUsage: "<foundation primaryToken pledgeK pledgeP>",
	Description: `
Deploy a Role contract by admin.

Arguments:
foundation - the address of foundation
primaryToken - the address of erc20 contract
pledgeK - the minimum amount to pledge in Role when apply for the Keeper role.
pledgeP - the minimum anount to pledge in Role when apply for the Provider role. 
version - the version of Role contract
	`,
	Action: func(cctx *cli.Context) error {
		foundation := cctx.Args().Get(0)
		primaryToken := cctx.Args().Get(1)
		pledgeKeeper := cctx.Args().Get(2)
		pledgeProvider := cctx.Args().Get(3)
		version := cctx.Args().Get(4)

		// 输入值判断并格式转换
		if len(foundation) != 42 || foundation == callconts.InvalidAddr {
			fmt.Println("foundation should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("foundation:", foundation)
		if len(primaryToken) != 42 || primaryToken == callconts.InvalidAddr {
			fmt.Println("primaryToken should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("primaryToken:", primaryToken)
		pledgeK := big.NewInt(0)
		pledgeK, ok := pledgeK.SetString(pledgeKeeper, 10)
		if !ok {
			return errSStr
		}
		fmt.Println("pledgeKeeper:", pledgeK)
		pledgeP := big.NewInt(0)
		pledgeP, ok = pledgeP.SetString(pledgeProvider, 10)
		if !ok {
			return errSStr
		}
		fmt.Println("pledgeProvider:", pledgeP)

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		r := callconts.NewR(callconts.AdminAddr, addr, sk, txopts, endPoint, status)
		v, err := strconv.Atoi(version)
		if err != nil {
			return err
		}
		roleAddr, _, err := r.DeployRole(common.HexToAddress(foundation), common.HexToAddress(primaryToken), pledgeK, pledgeP, uint16(v))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("Role contract address:", roleAddr.Hex())
		r = callconts.NewR(roleAddr, addr, sk, txopts, endPoint, status)
		rTokenAddr, err := r.RToken()
		if err != nil {
			return err
		}
		fmt.Println("RToken contract address:", rTokenAddr.Hex())
		return nil
	},
}

var setPICmd = &cli.Command{
	Name:      "setPI",
	Usage:     "Role - Admin set PledgePool-address,Issuance-address and RoleFS-address to Role contract. ",
	ArgsUsage: "<pledgePoolAddr issuAddr rolefsAddr>",
	Description: `
SetPI is a function in Role contract.
Set the pledgePool contract address, issuance contract address and rolefs contract address for role contract by admin.

Arguments:
pledgePoolAddr - address of pledge pool contract
issuAddr - address of issuance contract
rolefsAddr - address of rolefs contract
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "roleAddr",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), //默认值为common.go中的role合约地址
			Usage:   "the Role contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		pledgePoolAddr := cctx.Args().Get(0)
		issuAddr := cctx.Args().Get(1)
		rolefsAddr := cctx.Args().Get(2)

		// 输入值判断并格式转换
		if len(pledgePoolAddr) != 42 || pledgePoolAddr == callconts.InvalidAddr {
			fmt.Println("pledgePoolAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("pledgePool address:", pledgePoolAddr)
		if len(issuAddr) != 42 || issuAddr == callconts.InvalidAddr {
			fmt.Println("issuAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("issuance address:", issuAddr)
		if len(rolefsAddr) != 42 || rolefsAddr == callconts.InvalidAddr {
			fmt.Println("rolefsAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("roleFS address:", rolefsAddr)

		roleAddr := common.HexToAddress(cctx.String("roleAddr"))
		fmt.Println("roleAddr:", roleAddr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		r := callconts.NewR(roleAddr, addr, sk, txopts, endPoint, status)
		err := r.SetPI(common.HexToAddress(pledgePoolAddr), common.HexToAddress(issuAddr), common.HexToAddress(rolefsAddr))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var registerTokenCmd = &cli.Command{
	Name:      "regtk",
	Usage:     "Role - Admin registerToken in Role contract. ",
	ArgsUsage: "<tokenAddr>",
	Description: `
RegisterToken is a function in Role contract.
Register a token into Role contract by admin.

Arguments:
tokenAddr - the address of the token to be registered.
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "roleAddr",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), //默认值为common.go中的role合约地址
			Usage:   "the Role contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		tokenAddr := cctx.Args().Get(0)

		// 输入值判断并格式转换
		if len(tokenAddr) != 42 || tokenAddr == callconts.InvalidAddr {
			fmt.Println("tokenAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("token address:", tokenAddr)

		roleAddr := common.HexToAddress(cctx.String("roleAddr"))
		fmt.Println("roleAddr:", roleAddr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		r := callconts.NewR(roleAddr, addr, sk, txopts, endPoint, status)
		err := r.RegisterToken(common.HexToAddress(tokenAddr))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var createGroupCmd = &cli.Command{
	Name:      "cg",
	Usage:     "Role - Admin create group in Role contract. ",
	ArgsUsage: "<rfsAddr level>",
	Description: `
CreateGroup is a function in Role contract.
Create a group in Role contract, and deploy a FileSys contract for it. 

Arguments:
rfsAddr - the address of rolefs contract
level - the minimum number of members needed by the group for it to be active
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "roleAddr",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), //默认值为common.go中的role合约地址
			Usage:   "the Role contract address",
		},
		&cli.StringFlag{
			Name:        "kindexes",
			Aliases:     []string{"ks"},
			Value:       "", //默认值为nil
			Usage:       "keeper indexes included when creating group",
			DefaultText: "; kindexes such as 1,2,3",
		},
	},
	Action: func(cctx *cli.Context) error {
		rfsAddr := cctx.Args().Get(0)
		founder := uint64(0)
		level := cctx.Args().Get(1)

		// 输入值判断并格式转换
		if len(rfsAddr) != 42 || rfsAddr == callconts.InvalidAddr {
			fmt.Println("rfsAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("rfs address:", rfsAddr)
		le, err := strconv.Atoi(level)
		if err != nil {
			fmt.Println("trans string", level, " to int error")
			return err
		}
		fmt.Println("level:", le)

		// 获取command options
		roleAddr := common.HexToAddress(cctx.String("roleAddr"))
		fmt.Println("roleAddr:", roleAddr.Hex())
		kIndexes := cctx.String("kindexes")
		ks := strings.Split(kIndexes, ",")
		fmt.Println("kIndexes:", ks)
		var keepers []uint64
		if ks[0] == "" {
			keepers = nil
		} else {
			for _, tmp := range ks {
				intNum, err := strconv.Atoi(tmp)
				if err != nil {
					fmt.Println("trans string", tmp, " to int error")
					return err
				}
				keepers = append(keepers, uint64(intNum))
			}
		}
		fmt.Println("keepers:", keepers)
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		r := callconts.NewR(roleAddr, addr, sk, txopts, endPoint, status)
		gIndex, err := r.CreateGroup(common.HexToAddress(rfsAddr), founder, keepers, uint16(le))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("gIndex:", gIndex)
		time.Sleep(2 * time.Second)
		_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
		if err != nil {
			return err
		}
		fmt.Println("FileSys contract address:", fsAddr.Hex())
		return nil
	},
}

var addKeeperToGroupCmd = &cli.Command{
	Name:      "addk",
	Usage:     "Role - Admin add keeper to group in Role contract. ",
	ArgsUsage: "<kIndex gIndex>",
	Description: `
AddKeeper is a function in Role contract.
Add a keeper into the group with group index specified in argument.

Arguments:
kIndex - the index of the keeper to be added
gIndex - the index of the group
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "roleAddr",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), //默认值为common.go中的role合约地址
			Usage:   "the Role contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return errors.New("should have 2 arguments. Args0:kIndex, Args1:gIndex")
		}
		kIndex := cctx.Args().Get(0)
		gIndex := cctx.Args().Get(1)

		// 输入值判断并格式转换
		ki, err := strconv.Atoi(kIndex)
		if err != nil {
			fmt.Println("trans string", kIndex, " to int error")
			return err
		}
		gi, err := strconv.Atoi(gIndex)
		if err != nil {
			fmt.Println("trans string", gIndex, " to int error")
			return err
		}
		fmt.Println("kIndex:", ki, ", gIndex:", gi)

		roleAddr := common.HexToAddress(cctx.String("roleAddr"))
		fmt.Println("roleAddr:", roleAddr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		r := callconts.NewR(roleAddr, addr, sk, txopts, endPoint, status)
		err = r.AddKeeperToGroup(uint64(ki), uint64(gi))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var setPledgeMoneyCmd = &cli.Command{
	Name:      "setp",
	Usage:     "Role - Admin set pledgeK and pledgeP in Role contract. ",
	ArgsUsage: "<pledgeKeeper pledgeProvider>",
	Description: `
SetPledgeMoney is a function in role contract.
Set the value of pledgeK and pledgeP variables in the Role contract.

Arguments:
pledgeK - the minimum amount to pledge in Role when apply for the Keeper role.
pledgeP - the minimum anount to pledge in Role when apply for the Provider role. 
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "roleAddr",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), //默认值为common.go中的role合约地址
			Usage:   "the Role contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {
			return errors.New("should have 2 arguments. Args0:pledgeKeeper, Args1:pledgeProvider")
		}
		pledgeK := cctx.Args().Get(0)
		pledgeP := cctx.Args().Get(1)

		// 输入值判断并格式转换
		pk := big.NewInt(0)
		pk, ok := pk.SetString(pledgeK, 10)
		if !ok {
			fmt.Println("trans string", pledgeK, " to big.Int error")
			return errSStr
		}
		pp := big.NewInt(0)
		pp, ok = pp.SetString(pledgeP, 10)
		if !ok {
			fmt.Println("trans string", pledgeP, " to big.Int error")
			return errSStr
		}
		fmt.Println("pledgeKeeper:", pk, ", pledgeProvider:", pp)

		roleAddr := common.HexToAddress(cctx.String("roleAddr"))
		fmt.Println("roleAddr:", roleAddr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		r := callconts.NewR(roleAddr, addr, sk, txopts, endPoint, status)
		err := r.SetPledgeMoney(pk, pp)
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var alterOwnerCmd = &cli.Command{
	Name:      "ao",
	Usage:     "Role - Admin alter the owner of Role contract. ",
	ArgsUsage: "<newOwner>",
	Description: `
AlterOwner is a funtion in Role contract.
Alter the owner of Role contract by admin.

Arguments:
newOwner - specify the new owner of Role Contract.
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "roleAddr",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), //默认值为common.go中的role合约地址
			Usage:   "the Role contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return errors.New("should have 1 arguments. Args0:newOwner")
		}
		newOwner := cctx.Args().Get(0)

		// 输入值判断并格式转换
		if len(newOwner) != 42 || newOwner == callconts.InvalidAddr {
			fmt.Println("new owner should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("new owner address:", newOwner)

		roleAddr := common.HexToAddress(cctx.String("roleAddr"))
		fmt.Println("roleAddr:", roleAddr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		own := callconts.NewOwn(roleAddr, addr, sk, txopts, endPoint, status)
		err := own.AlterOwner(common.HexToAddress(newOwner))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var deployIssuanceCmd = &cli.Command{
	Name:      "depissu",
	Usage:     "Issuance - Admin deploy Issuance contract. ",
	ArgsUsage: "<rfsAddr>",
	Description: `
Deploy an Issuance contract by admin. 
The rolefs contract need to be deploy first.

Arguments:
rfsAddr - the address of rolefs contract. 
	`,
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return errors.New("should have 1 arguments. Args0:rfsAddr")
		}
		rfsAddr := cctx.Args().Get(0)

		// 输入值判断并格式转换
		if len(rfsAddr) != 42 || rfsAddr == callconts.InvalidAddr {
			fmt.Println("rfsAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("rolefs address:", rfsAddr)

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		issu := callconts.NewIssu(callconts.AdminAddr, addr, sk, txopts, endPoint, status)
		issuAddr, _, err := issu.DeployIssuance(common.HexToAddress(rfsAddr))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("\nIssuance contract address:", issuAddr.Hex())
		return nil
	},
}

var deployRolefsCmd = &cli.Command{
	Name:      "deprfs",
	Usage:     "RoleFS - Admin deploy RoleFS contract.",
	ArgsUsage: " ",
	Description: `
Deploy a RoleFS contract by admin.

Arguments:
none
	`,
	Action: func(cctx *cli.Context) error {
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		rfs := callconts.NewRFS(callconts.AdminAddr, addr, sk, txopts, endPoint, status)
		rfsAddr, _, err := rfs.DeployRoleFS()
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("\nRoleFS contract address:", rfsAddr.Hex())
		return nil
	},
}

var setAddrCmd = &cli.Command{
	Name:      "setAddr",
	Usage:     "RoleFS - Admin set issuan, role, rtoken address to RoleFS contract.",
	ArgsUsage: "<issuAddr roleAddr rtokenAddr>",
	Description: `
SetAddr is a function in RoleFS contract.
Set the address of issuance contract, role contract, rtoken contract for RoleFS contract.

Arguments:
issuaAddr - the address of issuance contract
roleAddr - the address of role contract
rtokenAddr - the address of rtoken contract
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "rolefsAddr",
			Aliases: []string{"rfs"},
			Value:   callconts.RoleFSAddr.Hex(), //默认值为common.go中的rolefs合约地址
			Usage:   "the RoleFS contract address",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {
			return errors.New("should have 3 arguments. Args0:issuAddr, Args1:roleAddr, Args2:rtokenAddr")
		}
		issuAddr := cctx.Args().Get(0)
		roleAddr := cctx.Args().Get(1)
		rtAddr := cctx.Args().Get(2)

		// 输入值判断并格式转换
		if len(issuAddr) != 42 || issuAddr == callconts.InvalidAddr {
			fmt.Println("issuAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("issuance address:", issuAddr)
		if len(roleAddr) != 42 || roleAddr == callconts.InvalidAddr {
			fmt.Println("roleAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("role address:", roleAddr)
		if len(rtAddr) != 42 || rtAddr == callconts.InvalidAddr {
			fmt.Println("rtAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("rToken address:", rtAddr)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		rfsAddr := common.HexToAddress(cctx.String("rolefsAddr"))
		fmt.Println("rfsAddr:", rfsAddr.Hex())
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)
		status := make(chan error)
		rfs := callconts.NewRFS(rfsAddr, addr, sk, txopts, endPoint, status)
		err := rfs.SetAddr(common.HexToAddress(issuAddr), common.HexToAddress(roleAddr), common.HexToAddress(rtAddr))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		return nil
	},
}

var deployFileSysCmd = &cli.Command{
	Name:      "depfs",
	Usage:     "FileSys - Admin deploy FileSys contract. ",
	ArgsUsage: "<roleAddr rfsAddr gIndex>",
	Description: `
Deploy a fileSys contract for a specified group index by admin. Usually called after a group is created.

Arguments:
roleAddr - the address of role contract
rfsAddr - the address of rolefs contract
gIndex - the index of the group which the newly deployed fileSys contract belongs to.
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "kindexes",
			Aliases:     []string{"ks"},
			Value:       "", //默认值为nil
			Usage:       "keeper indexes included when deploy FileSys",
			DefaultText: "kindexes such as 1,2,3",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {
			return errors.New("should have 3 arguments. Args0:roleAddr, Args1:rfsAddr, Args2:gIndex")
		}

		roleAddr := cctx.Args().Get(0)
		foundation := uint64(0)
		rfsAddr := cctx.Args().Get(1)
		gIndex := cctx.Args().Get(2)

		// 输入值判断并格式转换
		if len(roleAddr) != 42 || roleAddr == callconts.InvalidAddr {
			fmt.Println("roleAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("role address:", roleAddr)
		if len(rfsAddr) != 42 || rfsAddr == callconts.InvalidAddr {
			fmt.Println("rfsAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("rfs address:", rfsAddr)
		gi, err := strconv.Atoi(gIndex)
		if err != nil {
			fmt.Println("trans string", gIndex, " to int error")
			return err
		}
		fmt.Println("gIndex:", gi)

		// 获取command options
		kIndexes := cctx.String("kindexes")
		ks := strings.Split(kIndexes, ",")
		fmt.Println("kIndexes:", ks)
		var keepers []uint64
		for _, tmp := range ks {
			intNum, err := strconv.Atoi(tmp)
			if err != nil {
				fmt.Println("trans string", tmp, " to int error")
				return err
			}
			keepers = append(keepers, uint64(intNum))
		}
		fmt.Println("keepers:", keepers)
		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		fs := callconts.NewFileSys(callconts.AdminAddr, addr, sk, txopts, endPoint, status)
		fsAddr, _, err := fs.DeployFileSys(foundation, uint64(gi), common.HexToAddress(roleAddr), common.HexToAddress(rfsAddr), keepers)
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("\nFileSys contract address:", fsAddr.Hex())
		return nil
	},
}

var deployPledgePoolCmd = &cli.Command{
	Name:      "deppp",
	Usage:     "PledgePool - Admin deploy PledgePool contract. ",
	ArgsUsage: "<primeTokenAddr rTokenAddr roleAddr>",
	Description: `
Deploy a pledgePool contract by admin.

Arguments:
primeTokenAddr - the address of the primary token
rTokenAddr - the address of rToken contract
roleAddr - the address of role contract
	`,
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {
			return errors.New("should have 3 arguments. Args0:primeTokenAddr, Args1:rTokenAddr, Args2:roleAddr")
		}

		primeTokenAddr := cctx.Args().Get(0)
		rTokenAddr := cctx.Args().Get(1)
		roleAddr := cctx.Args().Get(2)

		// 输入值判断并格式转换
		if len(roleAddr) != 42 || roleAddr == callconts.InvalidAddr {
			fmt.Println("roleAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("role address:", roleAddr)
		if len(primeTokenAddr) != 42 || primeTokenAddr == callconts.InvalidAddr {
			fmt.Println("primeTokenAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("primeToken address:", primeTokenAddr)
		if len(rTokenAddr) != 42 || rTokenAddr == callconts.InvalidAddr {
			fmt.Println("rTokenAddr should be with prefix 0x and shouldn't be 0x0")
			return nil
		}
		fmt.Println("rToken address:", rTokenAddr)

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		status := make(chan error)
		pp := callconts.NewPledgePool(callconts.AdminAddr, addr, sk, txopts, endPoint, status)
		pledgePoolAddr, _, err := pp.DeployPledgePool(common.HexToAddress(primeTokenAddr), common.HexToAddress(rTokenAddr), common.HexToAddress(roleAddr))
		if err != nil {
			return err
		}
		err = <-status
		if err != nil {
			return err
		}
		fmt.Println("\nPledgePool contract address:", pledgePoolAddr.Hex())
		return nil
	},
}
