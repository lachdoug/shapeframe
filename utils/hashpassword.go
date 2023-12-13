package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash a password to a 16 byte string
func HashPassword(password string) (h string) {
	m := md5.Sum([]byte(password))
	h = hex.EncodeToString(m[:])
	return
}
