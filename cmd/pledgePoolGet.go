package cmd

import (
	"fmt"
	"math/big"
	callconts "memoc/callcontracts"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// PPGet some getter function in PledgePool-contract
// input of method set by param
var PPGet = &cli.Command{
	Name:  "ppget",
	Usage: "Get specified info of pledge pool contract",
	Flags: []cli.Flag{
		// pp
		&cli.StringFlag{
			Name:    "ppaddr",
			Aliases: []string{"pp"},
			Value:   callconts.PledgePoolAddr.Hex(), // default rtoken addr
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
		ppbalCmd,
		plCmd,
		tpCmd,
	},
}

// get balance of acc
var ppbalCmd = &cli.Command{
	Name:  "bal",
	Usage: "get balance of acc. args0: acc index, arg1: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		ppaddr := common.HexToAddress(cctx.String("ppaddr"))
		fmt.Println("pledge pool:", ppaddr)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		rindex := cctx.Args().Get(0)
		i, _ := strconv.Atoi(rindex)
		r64 := uint64(i)
		// // default index = 1
		// if t32 == 0 {
		// 	t32 = 1
		// }
		fmt.Println("rindex:", r64)

		tindex := cctx.Args().Get(1)
		i, _ = strconv.Atoi(tindex)
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

		// pp caller
		pp := callconts.NewPledgePool(ppaddr, caller, "", txopts, endPoint, make(chan error))

		// call contract
		b, err := pp.GetBalanceInPPool(r64, t32)
		if err != nil {
			return err
		}
		fmt.Printf("\nbalance: %v\n", formatWei(b))

		return nil
	},
}

// GetPledge Get all pledge amount in specified token.
var plCmd = &cli.Command{
	Name:  "pl",
	Usage: "GetPledge Get all pledge amount in specified token. arg0: token index",
	Action: func(cctx *cli.Context) error {
		// parse flags
		ppaddr := common.HexToAddress(cctx.String("ppaddr"))
		fmt.Println("pledge pool:", ppaddr)
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

		// pp caller
		pp := callconts.NewPledgePool(ppaddr, caller, "", txopts, endPoint, make(chan error))

		// call contract
		b, err := pp.GetPledge(t32)
		if err != nil {
			return err
		}
		fmt.Printf("\npledge: %v\n", formatWei(b))

		return nil
	},
}

// TotalPledge Get total pledge amount in pledge pool.

var tpCmd = &cli.Command{
	Name:  "tp",
	Usage: "TotalPledge Get total pledge amount in pledge pool",
	Action: func(cctx *cli.Context) error {

		// parse flags
		ppaddr := common.HexToAddress(cctx.String("ppaddr"))
		fmt.Println("pledge pool:", ppaddr)
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

		// pp caller
		pp := callconts.NewPledgePool(ppaddr, caller, "", txopts, endPoint, make(chan error))

		// call contract
		tt, err := pp.TotalPledge()
		if err != nil {
			return err
		}

		fmt.Printf("\ntotal pledge: %v\n", formatWei(tt))

		return nil
	},
}
