package encrypt

import (
	"encoding/json"
	//"github.com/MangoMilk/go-kit/id"
	"testing"
)

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
