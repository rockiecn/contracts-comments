package cmd

import (
	"fmt"
	"math/big"

	callconts "memoc/callcontracts"
	"memoc/test"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	transferMoney = big.NewInt(1e9)
	transferEth   = big.NewInt(2e18)
	token         = float64(1e18)
	nano          = float64(1e9)
)

// MoneyCmd is about transfer eth or ERC20-token, and get balance
var MoneyCmd = &cli.Command{
	Name:  "money",
	Usage: "Transfer Eth/ERC20-token",
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
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hex(), //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   callconts.EndPoint, //默认值为common.go中的endPoint
			Usage:   "the geth endPoint",
		},
	},
	Subcommands: []*cli.Command{
		balanceCmd,
		balanceEthCmd,
		transferCmd,
		transferEthCmd,
	},
}

var balanceCmd = &cli.Command{
	Name:      "balance",
	Usage:     "Get balance in ERC20 contract. ",
	ArgsUsage: "<target>",
	Description: `
This command get the balance of a specified account in ERC20 contract.

Arguments:
target - the address of the target account to get balance for.
	`,
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

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

		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, make(chan error))
		bal, err := e.BalanceOf(common.HexToAddress(acc))
		if err != nil {
			return err
		}
		fmt.Println("\nerc20 balance:", formatWei(bal))
		return nil
	},
}

var balanceEthCmd = &cli.Command{
	Name:  "balanceEth",
	Usage: "Get Eth balance. ",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   callconts.EndPoint, //默认值为common.go中的endPoint
			Usage:   "the geth endPoint",
		},
	},
	ArgsUsage: "<target>",
	Description: `
This command gets the Eth balance of a specified target account.

Arguments:
target - the target account address
	`,
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		bal := callconts.QueryEthBalance(acc, endPoint)
		fmt.Println("\neth balance:", bal)

		return nil
	},
}

var transferCmd = &cli.Command{
	Name:      "transfer",
	Usage:     "Transfer erc20-token to target address. ",
	ArgsUsage: "<target>",
	Description: `
Transfer is a function in ERC20 contract.
This command call transfer function of ERC20 contract to transfer erc20 tokens from the caller of the function to a specified target address.

Arguments:
target - the address which the erc20 tokens to be transfered to.
	`,
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

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

		e := callconts.NewERC20(erc20Addr, addr, sk, txopts, endPoint, make(chan error))
		err := e.Transfer(common.HexToAddress(acc), transferMoney)
		if err != nil {
			return err
		}
		fmt.Println("\nadmin transfer", transferMoney, " to", acc, " successfully!")
		return nil
	},
}

var transferEthCmd = &cli.Command{
	Name:      "transferEth",
	Usage:     "Transfer eth to target account. ",
	ArgsUsage: "<target>",
	Description: `
This command transfers Eth from the caller to a specified target address.

Arguments:
target - the target address which the Eth is transfered to.
	`,
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println()

		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		err := test.TransferTo(transferEth, common.HexToAddress(acc), endPoint, endPoint)
		if err != nil {
			return err
		}

		return nil
	},
}

func formatWei(i *big.Int) string {
	f := new(big.Float).SetInt(i)
	res, _ := f.Float64()
	switch {
	case res >= token:
		return fmt.Sprintf("%.02f MEMO", res/token)
	case res >= nano:
		return fmt.Sprintf("%.02f NAmo", res/nano)
	default:
		return fmt.Sprintf("%d ATmo", i.Int64())
	}
}
