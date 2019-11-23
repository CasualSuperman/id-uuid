package main

import (
	"bytes"
	"crypto/aes"
	"math/rand"
	"testing"
)

func BenchmarkAesByteArrEncrypt(b *testing.B) {
	//rand.Seed(42)

	var key uint128
	for i, _ := range key {
		key[i] = uint8(rand.Intn(0xFF))
	}

	var end uint128 = uint642Touint128(0, uint64(b.N))

	cipher, err := aes.NewCipher(key[:])
	if err != nil {
		b.Error(err)
	}

	var plaintext uint128
	var ciphertext uint128
	b.ResetTimer()
	for bytes.Compare(plaintext[:], end[:]) < 0 {
		plaintext = incrementByte(plaintext)
		cipher.Encrypt(ciphertext[:], plaintext[:])
	}
}

func BenchmarkAesByteArrDecrypt(b *testing.B) {
	//rand.Seed(42)

	var key uint128
	for i, _ := range key {
		key[i] = uint8(rand.Intn(0xFF))
	}

	var end uint128 = uint642Touint128(0, uint64(b.N))

	cipher, err := aes.NewCipher(key[:])
	if err != nil {
		b.Error(err)
	}

	var ciphertext uint128
	var deciphertext uint128
	b.ResetTimer()
	for bytes.Compare(ciphertext[:], end[:]) < 0 {
		ciphertext = incrementByte(ciphertext)
		cipher.Decrypt(deciphertext[:], ciphertext[:])
	}
}
