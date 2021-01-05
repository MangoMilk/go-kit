package encrypt

import (
	"encoding/json"
	"fmt"
	"testing"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestRSA(t *testing.T) {

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
}
