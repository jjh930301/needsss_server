package utils

import (
	"encoding/base64"
)

func EncryptBase64(pw []byte) string {
	encode := base64.StdEncoding.EncodeToString([]byte(pw))
	return string(encode)
}

func DecryptBase64(key string) string {
	decode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return ""
	}
	return string(decode)
}
