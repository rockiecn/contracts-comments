package callconts

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	//"github.com/memoio/go-mefs-v2/lib/utils"
)

// SignForRegister Used to call Register on behalf of other accounts
// hash(caller, register)
func SignForRegister(caller common.Address, regAccSK string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(regAccSK)
	if err != nil {
		log.Fatal(err)
	}
	//(caller, register)的哈希值
	//label := common.LeftPadBytes([]byte(register), 32)
	label := []byte(register)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//私钥对上述哈希值签名
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,_blsKey,"keeper")
func SignForRegisterKeeper(caller common.Address, _blsKey []byte, regAccSk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,_blsKey,"keeper")
	label := []byte(labelKeeper)
	hash := crypto.Keccak256(caller.Bytes(), _blsKey, label)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,"provider")
func SignForRegisterProvider(caller common.Address, regAccSk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,"provider")
	label := []byte(labelProvider)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,gIndex,payToken,blsKey)
func SignForRegisterUser(caller common.Address, gIndex uint64, _blsKey []byte, regAccSk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,gIndex,payToken,blsKey)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	// append caller
	b = append(b, caller.Bytes()...)
	// append gIndex
	binary.BigEndian.PutUint64(tmp8, gIndex)
	b = append(b, tmp8...)
	// append blsKey
	b = append(b, _blsKey...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,gIndex)
func SignForAddProviderToGroup(caller common.Address, gIndex uint64, accSk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(accSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,gIndex,payToken,blsKey)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	// append caller
	b = append(b, caller.Bytes()...)
	// append gIndex
	binary.BigEndian.PutUint64(tmp8, gIndex)
	b = append(b, tmp8...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller, uIndex, tIndex, money)
func SignForRecharge(caller common.Address, uIndex uint64, tIndex uint32, money *big.Int, accSk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(caller, uIndex, tIndex, money)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	tmp4 := make([]byte, 4)

	// append caller
	b = append(b, caller.Bytes()...)
	// append uIndex
	binary.BigEndian.PutUint64(tmp8, uIndex)
	b = append(b, tmp8...)
	// append tIndex
	binary.BigEndian.PutUint32(tmp4, tIndex)
	b = append(b, tmp4...)
	// append money
	m := common.LeftPadBytes(money.Bytes(), 32)
	b = append(b, m...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller, tIndex, amount)
func SignForWithdrawFromFs(caller common.Address, tIndex uint32, amount *big.Int, accSk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(caller, tIndex, amount)
	b := make([]byte, 0)
	tmp4 := make([]byte, 4)

	// append caller
	b = append(b, caller.Bytes()...)
	// append tIndex
	binary.BigEndian.PutUint32(tmp4, tIndex)
	b = append(b, tmp4...)
	// append amount
	m := common.LeftPadBytes(amount.Bytes(), 32)
	b = append(b, m...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// SignForRepair used to call AddRepair or SubRepair, when subRepair, label is "s"; when addRepair, label is "a"
func SignForRepair(sk string, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, label string) ([]byte, error) {
	ecdsaSk, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Fatal(err)
	}

	// (npIndex, _start, end, _size, nonce, tIndex, sprice, label)的哈希值
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

	fmt.Println("hash:", hash)

	// 私钥签名
	sig, err := crypto.Sign(hash, ecdsaSk)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// hash(uIndex, pIndex, nonce, _start, end, _size, sPrice)
func SignForOrder(
	uID uint64,
	pID uint64,
	nonce uint64,
	start uint64,
	end uint64,
	sz uint64,
	tIndex uint32,
	price *big.Int,
	accSk string,
) ([]byte, error) {
	// string to ecdsa
	skEcdsa, err := crypto.HexToECDSA(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(uIndex, pIndex, nonce, _start, end, _size, sPrice)
	buf := make([]byte, 8)
	buf4 := make([]byte, 4)
	d := sha3.NewLegacyKeccak256()
	binary.BigEndian.PutUint64(buf, uID)
	d.Write(buf)
	binary.BigEndian.PutUint64(buf, pID)
	d.Write(buf)
	binary.BigEndian.PutUint64(buf, nonce)
	d.Write(buf)
	binary.BigEndian.PutUint64(buf, start)
	d.Write(buf)
	binary.BigEndian.PutUint64(buf, end)
	d.Write(buf)
	binary.BigEndian.PutUint64(buf, sz)
	d.Write(buf)
	binary.BigEndian.PutUint32(buf4, tIndex)
	d.Write(buf4)
	d.Write(LeftPadBytes(price.Bytes(), 32))
	hash := d.Sum(nil)

	fmt.Printf("hash for add order:\n%x\n", hash)

	// sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(pIndex, tIndex, pay, lost)
func SignForProWithdraw(
	pIndex uint64,
	tIndex uint32,
	pay *big.Int,
	lost *big.Int,
	accSk string,
) ([]byte, error) {
	// string to ecdsa
	skEcdsa, err := crypto.HexToECDSA(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(pIndex, tIndex, pay, lost)
	var buf8 = make([]byte, 8)
	var buf4 = make([]byte, 4)
	d := sha3.NewLegacyKeccak256()
	binary.BigEndian.PutUint64(buf8, pIndex)
	d.Write(buf8)
	binary.BigEndian.PutUint32(buf4, tIndex)
	d.Write(buf4)
	d.Write(LeftPadBytes(pay.Bytes(), 32))
	d.Write(LeftPadBytes(lost.Bytes(), 32))

	hash := d.Sum(nil)

	fmt.Printf("hash for proWithdraw:\n%x\n", hash)

	// sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func LeftPadBytes(slice []byte, l int) []byte {
	if l <= len(slice) {
		return slice
	}

	padded := make([]byte, l)
	copy(padded[l-len(slice):], slice)

	return padded
}
