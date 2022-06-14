package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	callconts "memoc/callcontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
)

const (
	MoneyTo   = 1e18
	WaitTime  = 3 * time.Second
	AdminAddr = "0x6916c30dD1E2BC8d2550AC206c03C56834751e41"
	AdminSk   = "85b7f07c7c0d84292ca565a623ca30c80c48c8225b6b43330a9ad071972bac49"
	Acc1      = "0x9011B1c901d330A63d029B6B325EdE69aeEe11d4"
	Sk1       = "7000cd6cee7cdb6bcc7eda212d1c5aea1d8d35321895f4d601fdd49be96fbc7b"
	Acc2      = "0x7B024b830B0a9a315Baf9C76c0E53ec2dD19cf85"
	Sk2       = "a6d933d11c64217989e8a9449c2ad498ff33df51825a425db71c220e6158ee87"
	Acc3      = "0xeA1652B7a432C3A5607C13aEcDa44BE0Ab24692C"
	Sk3       = "376187ebad236c9e45827d6fee5f00573c6d690f3600f3b3cb4159b9a89446c9"
	Acc4      = "0xC7DFcac0e19e67561e80aEc20474cC3D5eC88ad0"
	Sk4       = "c46a4316171ee0070b4f7c8e08229c63d9f5e683b2ca895d38f2e25e0a27ea4f"
	Acc5      = "0xe23627B7c85A55afA2dF7689A10e7a880723ab8D"
	Sk5       = "457a3b8a46f0990b8a75eaa1bf7a157813dd0908a106b216a1af8de6e3a2881f"
	Acc6      = "0xff055f0a5df3e1e52d1110e072316e9f90a56d7c"
	Sk6       = "26c330ac5d67890ee770ec58763579f2e7e0f891bf3244304dab49b2c737ecae"
	Acc7      = "0xa1776d1aa3b039e0fc91a05bceb2ef42785e7598"
	Sk7       = "1e2b78621a49ade3e1eb70c691802cd37bee078a0f10f690870c5e5b3dee8c3b"
	Acc8      = "0x1ca59b8200e96ba6c7718ba1278b66e141383a49"
	Sk8       = "b80a63a52061dc9c1624edd76945b295d16487a68accf7d55c56c2ed11bc2bbe"
	Acc9      = "0x97041772c3bd9b97af8615fcf04812db9f81ee74"
	Sk9       = "561904230d600c7b9b9842a57aff2b8a754693ce927d27c9f2e00b251a4a0480"
	Acc10     = "0x4242c00fea991f432ae2ffb7ae88b8b353739a13"
	Sk10      = "ed319477dc5cad5478954ac9751cfa094193451c0de161dedd75b515c12e9cf4"
)

var (
	Foundation = common.HexToAddress("0x30F6551c2F970b21C1A9426aeb289c4ED6d570Fd")

	PrimaryToken = common.HexToAddress("0xe4Fe7d73F7f7593f41E06FffF550732B2E25C7eD")
	RTokenAddr   = common.HexToAddress("0xB4888a5a29C3a33dA7c23c1d25375e693b4562f1")
)

// TransferTo trans eth to addr for test
func TransferTo(value *big.Int, addr common.Address, eth, qeth string) error {
	client, err := ethclient.Dial(eth)
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("aca26228a9ed5ca4da2dd08d225b1b1e049d80e1b126c0d7e644d04d0fb910a3")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasLimit := uint64(23000)           // in units
	gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("client.NetworkID error,use the default chainID")
		chainID = big.NewInt(666)
	}

	retry := 0
	for {
		if retry > 10 {
			return xerrors.New("fail to transfer")
		}
		retry++
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			continue
		}

		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			continue
		}

		tx := types.NewTransaction(nonce, addr, value, gasLimit, gasPrice, nil)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			continue
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Println("trans transcation fail:", err)
			continue
		}

		qCount := 0
		for ; qCount < 10; qCount++ {
			balance := callconts.QueryEthBalance(addr.Hex(), qeth)
			if balance.Cmp(value) >= 0 {
				break
			}
			fmt.Println(addr, "'s Balance now:", balance.String(), ", waiting for transfer success")
			t := 20 * (qCount + 1)
			time.Sleep(time.Duration(t) * time.Second)
		}

		if qCount < 10 {
			break
		}
	}

	fmt.Println("transfer ", value.String(), "to", addr)
	return nil
}
