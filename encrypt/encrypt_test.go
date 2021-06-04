package encrypt

import (
	"crypto"
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {

	md5,_:=MD5("123")
	fmt.Println("MD5: ", md5)

	sha1,_:=SHA1("123")
	fmt.Println("SHA1: ", sha1)

	hmacSha1,_:=HmacSHA1("123", "encrypt")
	fmt.Println("HmacSHA1: ", hmacSha1)

	hmacSha256,_:=HmacSHA256("123", "encrypt")
	fmt.Println("HmacSHA256: ", hmacSha256)
}

func TestAES(t *testing.T) {
	type People struct {
		Name string
	}

	jsonByte, _ := json.Marshal(People{Name: "JJ"})
	secret := "asdfreqw34thv123"
	//iv := "dgy7reqw34thv123"

	a := NewAES()
	// encrypt
	cipher, err := a.Encrypt(jsonByte, secret)
	if err != nil {
		t.Log(err)
	}
	t.Log("AES Encrypt: ", string(cipher))

	// decrypt
	data, aErr := a.Decrypt(cipher, secret)
	if aErr != nil {
		t.Log(aErr)
	}
	t.Log("AES Decrypt: ", string(data))

	d := People{}
	json.Unmarshal(data, &d)
	t.Log("AES Decrypt: ", d)

}



func TestRSA(t *testing.T) {

	type People struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	r := NewRSA()
	//pubKey, priKey, genPemErr := r.GeneratePem(2048, "./pem")
	pubKey, priKey, genPemErr := r.GeneratePem(2048)
	if genPemErr != nil {
		panic(genPemErr)
	}
	fmt.Println("rsa Public Key: ", pubKey)
	fmt.Println("rsa Private Key: ", priKey)

	dataByte, _ := json.Marshal(People{Name: "Dwarf"})
	cipherByte := r.Encrypt(dataByte, []byte(pubKey))
	fmt.Println("rsa Encrypt: ", string(cipherByte))
	fmt.Println("rsa Decrypt: ", r.Decrypt(cipherByte, []byte(priKey)))

	sign, _ := r.Signature(dataByte, []byte(priKey), crypto.SHA256)
	fmt.Println("rsa signature: ", sign)
	fmt.Println("rsa verify signature: ", r.Verify(dataByte, sign, []byte(pubKey), crypto.SHA256) == nil)
}
