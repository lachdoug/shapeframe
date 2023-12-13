package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func CipherDecrypt(cipherKey string, in string) (out string) {
	var err error
	var cipherBlock cipher.Block
	cipherText, _ := hex.DecodeString(in)
	if cipherBlock, err = aes.NewCipher([]byte(cipherKey)); err != nil {
		panic(err)
	}
	plainText := make([]byte, len(cipherText))
	cipherBlock.Decrypt(plainText, cipherText)
	out = string(plainText[:])
	return
}
