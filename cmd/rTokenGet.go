package cmd

import (
	"fmt"
	"math/big"
	callconts "memoc/callcontracts"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// RTGet rtoken address set by flags
// input of method set by param
var RTGet = &cli.Command{
	Name:  "rtget",
	Usage: "Get specified info of rtoken contract",
	Flags: []cli.Flag{
		// role
		&cli.StringFlag{
			Name:    "rtoken",
			Aliases: []string{"rt"},
			Value:   callconts.RTokenAddr.Hex(), // default rtoken addr
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
		ivCmd,
		taCmd,
		tiCmd,
		tnCmd,
	},
}

// get is valid
var ivCmd = &cli.Command{
	Name:  "iv",
	Usage: "get is valid. arg0: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		rtoken := common.HexToAddress(cctx.String("rtoken"))
		fmt.Println("rtoken:", rtoken)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		tindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(tindex)
		t32 := uint32(i)
		// // default index = 1
		// if t32 == 0 {
		// 	t32 = 1
		// }
		fmt.Println("tindex:", t32)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role token caller
		rt := callconts.NewRT(rtoken, caller, "", txopts, endPoint)

		// call contract
		b, err := rt.IsValid(t32)
		if err != nil {
			return err
		}
		fmt.Printf("\nisValid: %v\n", b)

		return nil
	},
}

// get token addr
var taCmd = &cli.Command{
	Name:  "ta",
	Usage: "get token addr. arg0: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		rtoken := common.HexToAddress(cctx.String("rtoken"))
		fmt.Println("rtoken:", rtoken)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		tindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(tindex)
		t32 := uint32(i)
		// // default index = 1
		// if t32 == 0 {
		// 	t32 = 1
		// }
		fmt.Println("tindex:", t32)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// role token caller
		rt := callconts.NewRT(rtoken, caller, "", txopts, endPoint)

		// call contract
		addr, err := rt.GetTA(t32)
		if err != nil {
			return err
		}
		fmt.Printf("\ntoken addr: %v\n", addr)

		return nil
	},
}

// get token index from addr
var tiCmd = &cli.Command{
	Name:  "ti",
	Usage: "get token index from addr. arg0: token address",
	Action: func(cctx *cli.Context) error {
		// parse flags
		rtoken := common.HexToAddress(cctx.String("rtoken"))
		fmt.Println("rtoken:", rtoken)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		addr := cctx.Args().Get(0)
		if len(addr) == 0 {
			// use erc20 as default
			addr = callconts.ERC20Addr.Hex()
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

		// role token caller
		rt := callconts.NewRT(rtoken, caller, "", txopts, endPoint)

		// call contract
		index, b, err := rt.GetTI(common.HexToAddress(addr))
		if err != nil {
			return err
		}
		fmt.Printf("\ntoken index: %v registered: %v\n", index, b)

		return nil
	},
}

// get token num
var tnCmd = &cli.Command{
	Name:  "tn",
	Usage: "get token num",
	Action: func(cctx *cli.Context) error {
		// parse flags
		rtoken := common.HexToAddress(cctx.String("rtoken"))
		fmt.Println("rtoken:", rtoken)
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

		// role token caller
		rt := callconts.NewRT(rtoken, caller, "", txopts, endPoint)

		// call contract
		n, err := rt.GetTNum()
		if err != nil {
			return err
		}
		fmt.Printf("\ntoken number: %v\n", n)

		return nil
	},
}
