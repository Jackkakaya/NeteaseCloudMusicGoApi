package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5Hex(message string) string {
	md := md5.New()
	md.Write([]byte(message))
	return hex.EncodeToString(md.Sum(nil))
}