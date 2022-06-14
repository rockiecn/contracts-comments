package cmd

import (
	"fmt"
	"math/big"
	callconts "memoc/callcontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// ISGet some getter functions in Issuance-contract
// input of method set by param
var ISGet = &cli.Command{
	Name:  "isget",
	Usage: "Get specified info of issuance contract",
	Flags: []cli.Flag{
		//
		&cli.StringFlag{
			Name:    "issu",
			Aliases: []string{"is"},
			Value:   callconts.IssuanceAddr.Hex(), // default fs addr
			Usage:   "issu address",
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
		mlCmd,
		lmCmd,
		prCmd,
		szCmd,
		stCmd,
		totalPayCmd,
		totalPaidCmd,
		periodTargetCmd,
		periodTotalRewardCmd,
		issuRatioCmd,
		minRatioCmd,
	},
}

//
var mlCmd = &cli.Command{
	Name:  "ml",
	Usage: "get mintlevel. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		ml, err := e.MintLevel()
		if err != nil {
			return err
		}
		fmt.Printf("\nmint level: %v [0x%x]\n", ml, ml)

		return nil
	},
}

//
var lmCmd = &cli.Command{
	Name:  "lm",
	Usage: "get last mint. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		lm, err := e.LastMint()
		if err != nil {
			return err
		}
		fmt.Printf("\nlast mint: %v [0x%x]\n", lm, lm)

		return nil
	},
}

var prCmd = &cli.Command{
	Name:  "pr",
	Usage: "get price. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		pr, err := e.Price()
		if err != nil {
			return err
		}
		fmt.Printf("\nprice: %v [0x%x]\n", pr, pr)

		return nil
	},
}

var szCmd = &cli.Command{
	Name:  "sz",
	Usage: "get size. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		sz, err := e.Size()
		if err != nil {
			return err
		}
		fmt.Printf("\nsize: %v [0x%x]\n", sz, sz)

		return nil
	},
}

var stCmd = &cli.Command{
	Name:  "st",
	Usage: "get space time. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		st, err := e.SpaceTime()
		if err != nil {
			return err
		}
		fmt.Printf("\nspace time: %v [0x%x]\n", st, st)

		return nil
	},
}

var totalPayCmd = &cli.Command{
	Name:  "totalPay",
	Usage: "get total pay in MEMO. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		tp, err := e.TotalPay()
		if err != nil {
			return err
		}
		fmt.Printf("\ntotalPay: %v [0x%x]\n", formatWei(tp), tp)

		return nil
	},
}

var totalPaidCmd = &cli.Command{
	Name:  "totalPaid",
	Usage: "get total paid in MEMO. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		tp, err := e.TotalPaid()
		if err != nil {
			return err
		}
		fmt.Printf("\ntotalPaid: %v [0x%x]\n", formatWei(tp), tp)

		return nil
	},
}

var periodTargetCmd = &cli.Command{
	Name:  "periodTarget",
	Usage: "get period issuance target in MEMO. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		tp, err := e.PeriodTarget()
		if err != nil {
			return err
		}
		fmt.Printf("\nperiodTarget: %v [0x%x]\n", formatWei(tp), tp)

		return nil
	},
}

var periodTotalRewardCmd = &cli.Command{
	Name:  "periodTotalReward",
	Usage: "get period total reward in MEMO. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		tp, err := e.PeriodTotalReward()
		if err != nil {
			return err
		}
		fmt.Printf("\nperiodTotalReward: %v [0x%x]\n", formatWei(tp), tp)

		return nil
	},
}

var issuRatioCmd = &cli.Command{
	Name:  "issuRatio",
	Usage: "get issuance ratio in MEMO. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		tp, err := e.IssuRatio()
		if err != nil {
			return err
		}
		fmt.Printf("\nissuRatio: %v [0x%x]\n", tp, tp)

		return nil
	},
}

var minRatioCmd = &cli.Command{
	Name:  "minRatio",
	Usage: "get minimum issuance ratio in MEMO. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
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

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		tp, err := e.MinRatio()
		if err != nil {
			return err
		}
		fmt.Printf("\nminRatio: %v [0x%x]\n", tp, tp)

		return nil
	},
}
