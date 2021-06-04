package encode

import (
	"encoding/base64"
	"errors"
)

func Base64Encoded(msg string) string {
	headerStrByte := []byte(msg)
	return base64.StdEncoding.EncodeToString(headerStrByte)
}

func Base64Decoded(msg string) (string, error) {
	msgByte, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", errors.New("base64 decode fail")
	}

	return string(msgByte), nil
}