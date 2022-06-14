package cmd

import (
	"fmt"
	"math/big"
	callconts "memoc/callcontracts"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// FSGet some getter functions in FileSys-contract
// input of method set by param
var FSGet = &cli.Command{
	Name:  "fsget",
	Usage: "Get specified info of fs contract",
	Flags: []cli.Flag{
		// fs
		&cli.StringFlag{
			Name:    "fs",
			Aliases: []string{"f"},
			Value:   callconts.FileSysAddr.Hex(), // default fs addr
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
		infoCmd,
		pnCmd,
		proCmd,
		aoCmd,
		siCmd,
		ciCmd,
		setiCmd,
		balCmd,
	},
}

// get fs info
var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "get fs info. arg0: user index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		user := cctx.Args().Get(0)
		i, _ := strconv.Atoi(user)
		u64 := uint64(i)
		fmt.Println("user:", u64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		e := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))
		a, tIndex, err := e.GetFsInfo(u64)
		if err != nil {
			return err
		}
		fmt.Printf("\nisActive: %v token Index: %v\n", a, tIndex)

		return nil
	},
}

// get fs provider number
var pnCmd = &cli.Command{
	Name:  "pn",
	Usage: "get fs provider number. arg0: user index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		user := cctx.Args().Get(0)
		i, _ := strconv.Atoi(user)
		u64 := uint64(i)
		fmt.Println("user:", u64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		e := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		pn, err := e.GetFsProviderSum(u64)
		if err != nil {
			return err
		}
		fmt.Printf("\nprovider number: %v\n", pn)

		return nil
	},
}

// get fs provider index by array index
var proCmd = &cli.Command{
	Name:  "pro",
	Usage: "get fs provider index. arg0: user index, arg1: provider index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		user := cctx.Args().Get(0)
		i, _ := strconv.Atoi(user)
		u64 := uint64(i)
		fmt.Println("user:", u64)

		pindex := cctx.Args().Get(1)
		i, _ = strconv.Atoi(pindex)
		p64 := uint64(i)
		fmt.Println("pindex:", p64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		f := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		// call contract
		pIndex, err := f.GetFsProvider(u64, p64)
		if err != nil {
			return err
		}
		fmt.Printf("\nprovider index: %v\n", pIndex)

		return nil
	},
}

// get provider's aggregate order in FsInfo
var aoCmd = &cli.Command{
	Name:  "ao",
	Usage: "get provider's aggregate order in FsInfo. arg0: user index, arg1: provider index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		user := cctx.Args().Get(0)
		i, _ := strconv.Atoi(user)
		u64 := uint64(i)
		fmt.Println("user:", u64)

		pindex := cctx.Args().Get(1)
		i, _ = strconv.Atoi(pindex)
		p64 := uint64(i)
		fmt.Println("pindex:", p64)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		f := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		// call contract
		nonce, subnonce, err := f.GetFsInfoAggOrder(u64, p64)
		if err != nil {
			return err
		}
		fmt.Printf("\nnonce: %v, subnonce: %v\n", nonce, subnonce)

		return nil
	},
}

// get storeInfo in fs
var siCmd = &cli.Command{
	Name:  "si",
	Usage: "get storeInfo in fs. arg0: user index, arg1: provider index, arg2: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		token := cctx.String("token")
		fmt.Println("token:", token)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		user := cctx.Args().Get(0)
		i, _ := strconv.Atoi(user)
		u64 := uint64(i)
		fmt.Println("user:", u64)

		index := cctx.Args().Get(1)
		i, _ = strconv.Atoi(index)
		p64 := uint64(i)
		fmt.Println("index:", p64)

		t := cctx.Args().Get(2)
		i, _ = strconv.Atoi(t)
		t32 := uint32(i)
		fmt.Println("token:", t32)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		f := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		// call contract
		time, size, price, err := f.GetStoreInfo(u64, p64, t32)
		if err != nil {
			return err
		}
		fmt.Printf("\ntime: %v, size: %v, price: %v\n", time, size, price)

		return nil
	},
}

// get channelInfo in fs
var ciCmd = &cli.Command{
	Name:  "ci",
	Usage: "get channelInfo in fs. arg0: user index, arg1: provider index, arg2: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
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

		// parse args
		user := cctx.Args().Get(0)
		i, _ := strconv.Atoi(user)
		u64 := uint64(i)
		fmt.Println("user:", u64)

		index := cctx.Args().Get(1)
		i, _ = strconv.Atoi(index)
		p64 := uint64(i)
		fmt.Println("index:", p64)

		t := cctx.Args().Get(2)
		i, _ = strconv.Atoi(t)
		t32 := uint32(i)
		fmt.Println("token:", t32)

		// fs caller
		f := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		// call contract
		amount, nonce, expire, err := f.GetChannelInfo(u64, p64, t32)
		if err != nil {
			return err
		}
		fmt.Printf("\namount: %v, nonce: %v, expire: %v\n", amount, nonce, expire)

		return nil
	},
}

// 获得支付计费相关的信息
var setiCmd = &cli.Command{
	Name:  "seti",
	Usage: "get settle info in fs. arg0: provider index, arg1: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		p := cctx.Args().Get(0)
		i, _ := strconv.Atoi(p)
		p64 := uint64(i)
		fmt.Println("pindex:", p64)

		t := cctx.Args().Get(1)
		i, _ = strconv.Atoi(t)
		t32 := uint32(i)
		fmt.Println("token:", t32)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		f := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		// call contract
		// se.time,se.size,se.price,se.maxPay,se.hasPaid,se.canPay,se.lost,se.lostPaid,se.managePay,se.endPaid,se.linearPaid
		time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err := f.GetSettleInfo(p64, t32)
		if err != nil {
			return err
		}
		fmt.Println("\n", time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid)

		return nil
	},
}

// 获得账户收益余额
var balCmd = &cli.Command{
	Name:  "bal",
	Usage: "get balance of role in fs. arg0: role index, arg1: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		fs := common.HexToAddress(cctx.String("fs"))
		fmt.Println("fs:", fs)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		index := cctx.Args().Get(0)
		i, _ := strconv.Atoi(index)
		r64 := uint64(i)
		fmt.Println("index:", r64)

		t := cctx.Args().Get(1)
		i, _ = strconv.Atoi(t)
		t32 := uint32(i)
		fmt.Println("token:", t32)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// fs caller
		f := callconts.NewFileSys(fs, caller, "", txopts, endPoint, make(chan error))

		// call contract
		avail, tmp, err := f.GetBalance(r64, t32)
		if err != nil {
			return err
		}
		fmt.Printf("\navail: %v, tmp: %v\n", avail, tmp)

		return nil
	},
}
