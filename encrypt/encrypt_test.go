package encrypt

import (
	"fmt"
	"testing"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestEncrypt(t *testing.T) {
	fmt.Println("MD5: ", MD5("123"))

	fmt.Println("SHA1: ", SHA1("123"))

	fmt.Println("HmacSHA1: ", HmacSHA1("123", "encrypt"))

	fmt.Println("HmacSHA256: ", HmacSHA256("123", "encrypt"))

	base64Text := Base64Encoded("123")
	fmt.Println("Base64Encoded: ", base64Text)

	txt, err := Base64Decoded(base64Text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Base64Decoded: ", txt)
}
