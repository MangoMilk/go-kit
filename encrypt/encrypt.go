package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func MD5(msg string) string {
	h := md5.New()
	signByte := []byte(msg)
	_, err := h.Write(signByte)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(msg string) string {
	s := sha1.New()
	signByte := []byte(msg)
	_, err := s.Write(signByte)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(s.Sum(nil))
}

func HmacSHA1(msg string, key string) string {

	var hSecretByte = []byte(key) // hmac

	h := hmac.New(sha1.New, hSecretByte)
	signByte := []byte(msg)
	_, err := h.Write(signByte)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func HmacSHA256(msg string, key string) string {
	var hSecretByte = []byte(key)

	h := hmac.New(sha256.New, hSecretByte)
	signByte := []byte(msg)
	_, err := h.Write(signByte)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}
