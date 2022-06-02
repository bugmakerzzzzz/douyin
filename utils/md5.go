package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var salt = "zhsypd_douyin"

func MD5_SALT(str string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}