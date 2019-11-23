package main

import (
	"bytes"
//	"crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"

	"crypto/aes"
)

const byteNum = 128 / 8

type uint128 [byteNum]byte

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
	for i := 0; i < 2; i++ {//, _ := range end {
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
		plaintext = incrementByte(plaintext)
		fmt.Printf("%-10s: %s\n", "Plaintext", encodeHex(plaintext))
		cipher.Encrypt(ciphertext[:], plaintext[:])
		fmt.Printf("%-10s: %s\n", "Ciphertext", encodeHex(ciphertext))
		cipher.Decrypt(deciphertext[:], ciphertext[:])
		fmt.Printf("%-10s: %s\n", "Deciphered", encodeHex(deciphertext))
	}
}

func incrementByte(item uint128) uint128 {
	higher, lower := uint128Touint642(item)

	lower++
	if lower == 0 {
		higher++
	}

	return uint642Touint128(higher, lower)
}

func uint128Touint642(item uint128) (uint64, uint64) {
	lower := binary.BigEndian.Uint64(item[byteNum / 2:])
	higher := binary.BigEndian.Uint64(item[:byteNum / 2])
	return higher, lower
}


func uint642Touint128(higher, lower uint64) uint128 {
	var item uint128
	binary.BigEndian.PutUint64(item[byteNum / 2:], lower)
	binary.BigEndian.PutUint64(item[:byteNum / 2], higher)
	return item
}
