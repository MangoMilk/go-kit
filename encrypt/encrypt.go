package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func MD5(msg string) (string,error) {
	h := md5.New()
	if _, err := h.Write([]byte(msg));err != nil {
		return "",err
	}

	return hex.EncodeToString(h.Sum(nil)),nil
}

func SHA1(msg string) (string,error) {
	s := sha1.New()
	if _, err := s.Write([]byte(msg));err != nil {
		return "",err
	}

	return hex.EncodeToString(s.Sum(nil)),nil
}

func HmacSHA1(msg string, key string) (string,error) {
	h := hmac.New(sha1.New, []byte(key))
	if _, err := h.Write([]byte(msg));err != nil {
		return "",err
	}

	return hex.EncodeToString(h.Sum(nil)),nil
}

func HmacSHA256(msg string, key string) (string,error) {
	h := hmac.New(sha256.New, []byte(key))
	if _, err := h.Write([]byte(msg));err != nil {
		return "",err
	}

	return hex.EncodeToString(h.Sum(nil)),nil
}
