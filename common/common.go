package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5(plaintext string) string {
	h := md5.New()
	io.WriteString(h, plaintext)
	return hex.EncodeToString(h.Sum(nil))
}
