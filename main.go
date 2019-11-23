package main

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"

	"crypto/aes"
)

var one = big.NewInt(1)

type uint128 [128 / 8]byte

func encodeHex(data uint128) string {
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data[:])
	return string(dst)
}

func main() {
	rand.Seed(42)

	var key uint128
	for i, _ := range key {
		key[i] = uint8(rand.Intn(0xFF))
	}
	fmt.Printf("%-10s : %s\n", "Key", encodeHex(key))

	var end uint128
	for i := 0; i < 1; i++ {//, _ := range end {
		end[len(end) - i - 1] = 0xFF
	}

	cipher, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}
	var plaintext uint128
	var ciphertext uint128
	var deciphertext uint128
	for bytes.Compare(plaintext[:], end[:]) < 0 {
		plaintext = increment(plaintext)
		fmt.Printf("%-10s: %s\n", "Plaintext", encodeHex(plaintext))
		cipher.Encrypt(ciphertext[:], plaintext[:])
		fmt.Printf("%-10s: %s\n", "Ciphertext", encodeHex(ciphertext))
		cipher.Decrypt(deciphertext[:], ciphertext[:])
		fmt.Printf("%-10s: %s\n", "Deciphered", encodeHex(deciphertext))
	}
}

func increment(item uint128) uint128 {
	i := big.NewInt(0)
	//fmt.Println(i)
	i.SetBytes(item[:])
	//fmt.Println(i)
	i = i.Add(i, one)
	//fmt.Println(i)
	return bigToUint128(i);
}

func bigToUint128 (i *big.Int) uint128 {
	var dest uint128
	slice := i.Bytes()
	l := len(slice)
	copy(dest[len(dest) - l:], slice)
	return dest
}

type FeistelCipher struct {
	key uint128
	tweak uint128
}

func (f FeistelCipher) BlockSize() int {
	return 128 / 8
}

func (f FeistelCipher) Encrypt(dst, src []byte) {

}

func (f FeistelCipher) Decrypt(dst, src []byte) {

}

func NewCipher() cipher.Block {
	return &FeistelCipher{}
}
