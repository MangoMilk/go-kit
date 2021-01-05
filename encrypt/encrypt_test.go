package encrypt

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {

	fmt.Println("MD5: ", MD5("123"))

	fmt.Println("SHA1: ", SHA1("123"))

	fmt.Println("HmacSHA1: ", HmacSHA1("123", "encrypt"))

	fmt.Println("HmacSHA256: ", HmacSHA256("123", "encrypt"))
}
