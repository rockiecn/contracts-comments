package cmd

import (
	"fmt"
	"math/big"
	callconts "memoc/callcontracts"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// GetERC20Cmd erc20 address and caller address set by flags
// input of method set by param
var GetERC20Cmd = &cli.Command{
	Name:  "eget",
	Usage: "Get specified info of the ERC20-contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hash().Hex(), // default caller = admin
			Usage:   "erc20 address",
		},
		&cli.StringFlag{
			Name:    "caller",
			Aliases: []string{"c"},
			Value:   callconts.AdminAddr.Hex(), // default caller = admin
			Usage:   "tx caller",
		},
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   callconts.EndPoint, //默认值为common.go中的endPoint
			Usage:   "the geth endPoint",
		},
	},
	Subcommands: []*cli.Command{
		nameCmd,
		symbolCmd,
		decCmd,
		tsCmd,
		boCmd,
		alCmd,
		hrCmd,
		psCmd,
		msCmd,
		vCmd,
		msaCmd,
	},
}

// get erc20 name
var nameCmd = &cli.Command{
	Name:  "name",
	Usage: "get erc20 token name. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		n, err := e.GetName()
		if err != nil {
			return err
		}
		fmt.Println("erc20 name:", n)

		return nil
	},
}

// get erc20 symbol
var symbolCmd = &cli.Command{
	Name:  "sym",
	Usage: "get erc20 symbol.",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		n, err := e.GetSymbol()
		if err != nil {
			return err
		}
		fmt.Println("erc20 symbol:", n)

		return nil
	},
}

// get erc20 decimal
var decCmd = &cli.Command{
	Name:  "dec",
	Usage: "get erc20 decimal.",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		dec, err := e.GetDecimals()
		if err != nil {
			return err
		}
		fmt.Println("erc20 decimal:", dec)

		return nil
	},
}

// get erc20 total supply
var tsCmd = &cli.Command{
	Name:  "ts",
	Usage: "get erc20 total supply.",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		ts, err := e.GetTotalSupply()
		if err != nil {
			return err
		}
		fmt.Println("erc20 total supply:", formatWei(ts))

		return nil
	},
}

// get balance of acc
var boCmd = &cli.Command{
	Name:  "bo",
	Usage: "get balance of an acc. arg0: acc address",
	Action: func(cctx *cli.Context) error {
		// parse args
		acc := cctx.Args().Get(0)
		// use primary token as default
		if acc == "" {
			acc = callconts.AdminAddr.Hex()
		} else {
			// check addr
			if len(acc) != 42 || acc == callconts.InvalidAddr {
				fmt.Println("acc should be with prefix 0x, and shouldn't be 0x0")
				return nil
			}
		}
		fmt.Println("acc address:", acc)

		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20 address:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller address:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		b, err := e.BalanceOf(common.HexToAddress(acc))
		if err != nil {
			return err
		}
		fmt.Println("balance of acc: ", formatWei(b))

		return nil
	},
}

// get allowance of an acc pair
var alCmd = &cli.Command{
	Name:  "al",
	Usage: "get allowance of an acc pair.  arg0: owner address, arg1: spender address",
	Action: func(cctx *cli.Context) error {
		// parse args
		owner := cctx.Args().Get(0)
		// check addr
		if len(owner) != 42 || owner == callconts.InvalidAddr {
			fmt.Println("owner should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("owner address:", owner)

		spender := cctx.Args().Get(1)
		// check addr
		if len(spender) != 42 || spender == callconts.InvalidAddr {
			fmt.Println("spender should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("spender address:", spender)

		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20 address:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller address:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		al, err := e.Allowance(common.HexToAddress(owner), common.HexToAddress(spender))
		if err != nil {
			return err
		}
		fmt.Printf("allowance of %s to %s is : %v\n", owner, spender, formatWei(al))

		return nil
	},
}

// hasrole
var hrCmd = &cli.Command{
	Name:  "hr",
	Usage: "call hasrole method of accesscontrol. arg0: role type, arg1: role address",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20 address:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller address:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// parse args
		rType := cctx.Args().Get(0)
		rt, _ := strconv.Atoi(rType)
		r8 := uint8(rt)
		// type 1 as default
		if r8 == 0 {
			r8 = 1
		}
		fmt.Println("role type:", r8)

		acc := cctx.Args().Get(1)
		if len(acc) == 0 {
			acc = "0x9011B1c901d330A63d029B6B325EdE69aeEe11d4" // acc1 as default
		} else if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("acc should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("acc address:", acc)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		b, err := e.HasRole(r8, common.HexToAddress(acc))
		if err != nil {
			return err
		}
		fmt.Printf("\nhas role: %v\n", b)

		return nil
	},
}

// get paused
var psCmd = &cli.Command{
	Name:  "ps",
	Usage: "call getpaused method of accesscontrol",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20 address:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller address:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		b, err := e.GetPaused()
		if err != nil {
			return err
		}
		fmt.Printf("\nget paused: %v\n", b)

		return nil
	},
}

// get erc20 maxSupply
var msCmd = &cli.Command{
	Name:  "ms",
	Usage: "get erc20 maxSupply. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		n, err := e.GetMaxSupply()
		if err != nil {
			return err
		}
		fmt.Println("erc20 maxSupply:", formatWei(n))

		return nil
	},
}

// get erc20 version
var vCmd = &cli.Command{
	Name:  "v",
	Usage: "get erc20 version. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		n, err := e.GetVersion()
		if err != nil {
			return err
		}
		fmt.Println("erc20 version:", n)

		return nil
	},
}

// get erc20 multiSigAddrs
var msaCmd = &cli.Command{
	Name:  "msa",
	Usage: "get erc20 multiSigAddrs. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
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
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts, endPoint, make(chan error))
		n, err := e.MultiSigAddrs()
		if err != nil {
			return err
		}
		output := [5]string{}
		for i, a := range n {
			output[i] = a.Hex()
		}
		fmt.Println("erc20 multiSigAddrs:", output)

		return nil
	},
}
