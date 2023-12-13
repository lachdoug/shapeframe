package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func CipherEncrypt(cipherKey string, in string) (out string) {
	var err error
	var cipherBlock cipher.Block
	if cipherBlock, err = aes.NewCipher([]byte(cipherKey)); err != nil {
		panic(err)
	}
	cipherText := make([]byte, aes.BlockSize+len(in))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(cipherBlock, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(in))
	out = base64.RawStdEncoding.EncodeToString(cipherText)
	return
}
