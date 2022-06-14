package cmd

import (
	"fmt"
	"math/big"
	callconts "memoc/callcontracts"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// RGet some getter functions in Role-contract
// input of method set by param
var RGet = &cli.Command{
	Name:  "rget",
	Usage: "Get specified info of role contract",
	Flags: []cli.Flag{
		// role
		&cli.StringFlag{
			Name:    "role",
			Aliases: []string{"r"},
			Value:   callconts.RoleAddr.Hex(), // default role addr
			Usage:   "fs address",
		},
		// caller
		&cli.StringFlag{
			Name:    "caller",
			Aliases: []string{"c"},
			Value:   callconts.AdminAddr.Hex(), // default caller = admin
			Usage:   "tx caller",
		},
		// end point
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   callconts.EndPoint, //默认值为common.go中的endPoint
			Usage:   "the geth endPoint",
		},
	},
	Subcommands: []*cli.Command{
		anCmd,
		addrCmd,
		riCmd,
		rinfoCmd,
		gnCmd,
		ginfoCmd,
		faCmd,
		knCmd,
		upCmd,
		gkCmd,
		gpCmd,
		guCmd,
		ownerCmd,
		ppCmd,
		pkCmd,
		rvCmd,
	},
}

// get addr num
var anCmd = &cli.Command{
	Name:  "an",
	Usage: "get addr num",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		n, err := r.GetAddrsNum()
		if err != nil {
			return err
		}
		fmt.Printf("\naddr num: %v\n", n)

		return nil
	},
}

// get addr
var addrCmd = &cli.Command{
	Name:  "addr",
	Usage: "get acc addr. args0: acc index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		index := cctx.Args().Get(0)
		i, _ := strconv.Atoi(index)
		i64 := uint64(i)
		fmt.Println("index:", i64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		addr, err := r.GetAddr(i64)
		if err != nil {
			return err
		}
		fmt.Printf("\naddr: %v\n", addr)

		return nil
	},
}

// get role index
var riCmd = &cli.Command{
	Name:  "ri",
	Usage: "get role index. args0: role address",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		addr := cctx.Args().Get(0)
		if len(addr) == 0 {
			// use admin as default
			addr = callconts.AdminAddr.Hex()
		} else if len(addr) != 42 || addr == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("addr:", addr)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		index, err := r.GetRoleIndex(common.HexToAddress(addr))
		if err != nil {
			return err
		}
		fmt.Printf("\nindex: %v\n", index)

		return nil
	},
}

// get role info
var rinfoCmd = &cli.Command{
	Name:  "rinfo",
	Usage: "get role info. args0: role address",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		addr := cctx.Args().Get(0)
		if len(addr) == 0 {
			// use admin as default
			addr = callconts.AdminAddr.Hex()
		} else if len(addr) != 42 || addr == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("addr:", addr)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		// info[acc].isActive, info[acc].isBanned, info[acc].roleType, info[acc].index, info[acc].gIndex, info[acc].extra
		ia, ib, rt, index, gindex, extra, err := r.GetRoleInfo(common.HexToAddress(addr))
		if err != nil {
			return err
		}
		fmt.Println("\nisActive, isBanned, roleType, index, gIndex, extra")
		fmt.Println("info: ", ia, ib, rt, index, gindex, extra)

		return nil
	},
}

// get group num
var gnCmd = &cli.Command{
	Name:  "gn",
	Usage: "get group num",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		n, err := r.GetGroupsNum()
		if err != nil {
			return err
		}
		fmt.Printf("\ngroup num: %v\n", n)

		return nil
	},
}

// get group info
var ginfoCmd = &cli.Command{
	Name:  "ginfo",
	Usage: "get group info. args0: group index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		gindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(gindex)
		g64 := uint64(i)
		fmt.Println("gindex:", g64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		// groups[i].isActive, groups[i].isBanned, groups[i].isReady, groups[i].level, groups[i].size, groups[i].price, groups[i].fsAddr
		ia, ib, ir, lvl, size, price, fsaddr, err := r.GetGroupInfo(g64)
		if err != nil {
			return err
		}
		fmt.Println("\nisActive, isBanned, isReady, level, size, price, fsAddr")
		fmt.Println("info: ", ia, ib, ir, lvl, size, price, fsaddr)

		return nil
	},
}

// get addr and gindex by acc index
var faCmd = &cli.Command{
	Name:  "ag",
	Usage: "get addr and gindex. args0: acc index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		index := cctx.Args().Get(0)
		i, _ := strconv.Atoi(index)
		i64 := uint64(i)
		// default acc index = 1
		if i64 == 0 {
			i64 = 1
		}
		fmt.Println("index:", i64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		addr, gi, err := r.GetAddrGindex(i64)
		if err != nil {
			return err
		}
		fmt.Printf("\nacc addr: %v, gindex: %v\n", addr, gi)

		return nil
	},
}

// get group's keeper num by index
var knCmd = &cli.Command{
	Name:  "kn",
	Usage: "get group's keeper num by index. args0: group index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		index := cctx.Args().Get(0)
		i, _ := strconv.Atoi(index)
		i64 := uint64(i)
		// default index = 1
		if i64 == 0 {
			i64 = 1
		}
		fmt.Println("index:", i64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		n, err := r.GetGKNum(i64)
		if err != nil {
			return err
		}
		fmt.Printf("\nkeeper num: %v\n", n)

		return nil
	},
}

// get group's user and provider num by index
var upCmd = &cli.Command{
	Name:  "up",
	Usage: "get group's user and provider num by index. args0: group index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		index := cctx.Args().Get(0)
		i, _ := strconv.Atoi(index)
		i64 := uint64(i)
		// default index = 1
		if i64 == 0 {
			i64 = 1
		}
		fmt.Println("index:", i64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		u, p, err := r.GetGUPNum(i64)
		if err != nil {
			return err
		}
		fmt.Printf("\nuser num: %v, provider num: %v\n", u, p)

		return nil
	},
}

// get group's keeper
var gkCmd = &cli.Command{
	Name:  "gk",
	Usage: "get group's keeper. args0: group index, arg1: keeper index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		gindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(gindex)
		g64 := uint64(i)
		// default index = 1
		if g64 == 0 {
			g64 = 1
		}
		fmt.Println("gindex:", g64)

		kindex := cctx.Args().Get(0)
		i, _ = strconv.Atoi(kindex)
		k64 := uint64(i)
		// default index = 1
		if k64 == 0 {
			k64 = 1
		}
		fmt.Println("kindex:", k64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		kIndex, err := r.GetGroupK(g64, k64)
		if err != nil {
			return err
		}
		fmt.Printf("\nkeeper index: %v\n", kIndex)

		return nil
	},
}

// get group's provider
var gpCmd = &cli.Command{
	Name:  "gp",
	Usage: "get group's provider. args0: group index, arg1: provider index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		gindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(gindex)
		g64 := uint64(i)
		// default index = 1
		if g64 == 0 {
			g64 = 1
		}
		fmt.Println("gindex:", g64)

		pindex := cctx.Args().Get(0)
		i, _ = strconv.Atoi(pindex)
		p64 := uint64(i)
		// default index = 1
		if p64 == 0 {
			p64 = 1
		}
		fmt.Println("pindex:", p64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		pIndex, err := r.GetGroupP(g64, p64)
		if err != nil {
			return err
		}
		fmt.Printf("\nprovider index: %v\n", pIndex)

		return nil
	},
}

// get group's provider
var guCmd = &cli.Command{
	Name:  "gu",
	Usage: "get group's user. args0: group index, arg1: user index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		gindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(gindex)
		g64 := uint64(i)
		// default index = 1
		if g64 == 0 {
			g64 = 1
		}
		fmt.Println("gindex:", g64)

		uindex := cctx.Args().Get(0)
		i, _ = strconv.Atoi(uindex)
		u64 := uint64(i)
		// default index = 1
		if u64 == 0 {
			u64 = 1
		}
		fmt.Println("uindex:", u64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role caller
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		uIndex, err := r.GetGroupU(g64, u64)
		if err != nil {
			return err
		}
		fmt.Printf("\nuser index: %v\n", uIndex)

		return nil
	},
}

// get owner
var ownerCmd = &cli.Command{
	Name:  "owner",
	Usage: "get owner",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// owner use the same address with role
		o := callconts.NewOwn(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		owner, err := o.GetOwner()
		if err != nil {
			return err
		}
		fmt.Printf("\nowner: %v\n", owner)

		return nil
	},
}

// get money when pledge Provider
var ppCmd = &cli.Command{
	Name:  "pp",
	Usage: "get pledge money for Provider",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// owner use the same address with role
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		pledge, err := r.PledgeP()
		if err != nil {
			return err
		}
		fmt.Printf("\npledgeProvider: %v\n", formatWei(pledge))

		return nil
	},
}

// get money when pledge Keeper
var pkCmd = &cli.Command{
	Name:  "pk",
	Usage: "get pledge money for Keeper",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// owner use the same address with role
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))

		// call contract
		pledge, err := r.PledgeK()
		if err != nil {
			return err
		}
		fmt.Printf("\npledgekeeper: %v\n", formatWei(pledge))

		return nil
	},
}

// get Role version
var rvCmd = &cli.Command{
	Name:  "v",
	Usage: "get Role version. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		role := common.HexToAddress(cctx.String("role"))
		fmt.Println("role:", role)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// owner use the same address with role
		r := callconts.NewR(role, caller, "", txopts, endPoint, make(chan error))
		n, err := r.GetRVersion()
		if err != nil {
			return err
		}
		fmt.Println("Role version:", n)

		return nil
	},
}
