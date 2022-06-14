package main

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	callconts "memoc/callcontracts"
	"memoc/cmd"
	test "memoc/test"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

var errHexskFormat = errors.New("the hexsk'format is wrong")

// for test
func main() {
	fmt.Println("welcome to test contract!")
	fmt.Println("curent endPoint: ", callconts.EndPoint)

	commands := []*cli.Command{
		cmd.AdminCmd,
		cmd.MoneyCmd,
		cmd.GetERC20Cmd,
		cmd.FSGet,
		cmd.RGet,
		cmd.RTGet,
		cmd.PPGet,
		cmd.ISGet,
	}

	app := &cli.App{
		Name:                 "memoc",
		Usage:                "Tool for test memo contracts",
		EnableBashCompletion: true,
		Commands:             commands,
	}

	app.Setup()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

func testHash() {
	encodePackedHex := "00000000000000030000000061bc37780000000061bc3779000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000561"
	fmt.Println("len:", len(encodePackedHex))

	encodePacked := hexToByte(encodePackedHex)
	solh := hexToByte("85c6ceefb89fb4e3e80dd6db29861de012e185df27d365278b5187e990088736")
	fmt.Println("encodePacked:", encodePacked)
	fmt.Println("sol-hash:", solh)
	fmt.Println("sol-hash:", bytesToHex(solh))

	h, err := getHash(3, 1639724920, 1639724921, 10, 0, 0, big.NewInt(5), "a")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("h:", h)

	sign, err := callconts.SignForRepair(test.AdminSk, 3, 1639724920, 1639724921, 10, 0, 0, big.NewInt(5), "a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signature:", sign)
	fmt.Println("signature:", bytesToHex(sign))
}

func hexToByte(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func bytesToHex(b []byte) string {
	s := hex.EncodeToString(b)
	return s
}

func getHash(pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, label string) ([]byte, error) {
	by := make([]byte, 77)
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, pIndex)
	for i, b := range tmp {
		by[i] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, start)
	for i, b := range tmp {
		by[i+8] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, end)
	for i, b := range tmp {
		by[i+16] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, size)
	for i, b := range tmp {
		by[i+24] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, nonce)
	for i, b := range tmp {
		by[i+32] = byte(b)
	}
	t := make([]byte, 4)
	binary.BigEndian.PutUint32(t, tIndex)
	for i, b := range tmp {
		by[i+40] = byte(b)
	}
	spriceNew := common.LeftPadBytes(sprice.Bytes(), 32)
	for i, b := range spriceNew {
		by[i+44] = byte(b)
	}
	labelNew := []byte(label)
	by[76] = byte(labelNew[0])

	hash := crypto.Keccak256(by)

	return hash, nil
}
