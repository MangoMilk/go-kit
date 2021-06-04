package encrypt

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"strings"
)

type optionParams interface{}

type rsaHelper struct{}

func NewRSA() *rsaHelper {
	return &rsaHelper{}
}

const (
	generatePemPath = 0
)

var generatePemParams = []int{
	generatePemPath,
}

// GenerateRSAPem
//
// @param int bits
// @param string filePath
//
// @return publicKeyStr, privateKeyStr string, err error
func (r *rsaHelper) GeneratePem(bits int, opts ...optionParams) (publicKeyStr, privateKeyStr string, err error) {
	var isWriteIntoFile = false
	var filePath string

	if len(opts) > len(generatePemParams) {
		panic("Too many params.")
	}

	for k, v := range opts {
		if k == generatePemPath && v.(string) != "" {
			isWriteIntoFile = true
			filePath = v.(string)
			filePath = strings.TrimRight(filePath, "/")
			// create pem dir
			if isExist(filePath) != true {
				mkdirErr := os.Mkdir(filePath, 0755)
				if mkdirErr != nil {
					err = mkdirErr
					return
				}
			}
		}
	}

	// generate private key
	privateKey, genPriKeyErr := rsa.GenerateKey(rand.Reader, bits)
	if genPriKeyErr != nil {
		return
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// write into file
	if isWriteIntoFile {
		privateKeyFile, createFileErr := os.Create(filePath + "/private_key.pem")
		if privateKeyFile != nil {
			defer privateKeyFile.Close()
		}
		if createFileErr != nil {
			panic(createFileErr)
		}
		pemEncodeErr := pem.Encode(privateKeyFile, block)
		if pemEncodeErr != nil {
			return
		}
	}

	// write into buffer
	bufferPrivate := new(bytes.Buffer)
	pemEncodeErr := pem.Encode(bufferPrivate, block)
	if pemEncodeErr != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()

	// generate public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}

	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	// write into file
	if isWriteIntoFile {
		publicKeyFile, createFileErr := os.Create(filePath + "/public_key.pem")
		if publicKeyFile != nil {
			defer publicKeyFile.Close()
		}

		if createFileErr != nil {
			panic(createFileErr)
		}
		pemEncodeErr := pem.Encode(publicKeyFile, block)
		if pemEncodeErr != nil {
			return
		}
	}

	// write into buffer
	bufferPublic := new(bytes.Buffer)
	pubPemEncodeErr := pem.Encode(bufferPublic, block)
	if pubPemEncodeErr != nil {
		return
	}
	publicKeyStr = bufferPublic.String()

	return
}

// Encrypt
//
// @param []byte plaintext
// @param []byte pubKeyPem
//
// @return []byte ciphertext
func (r *rsaHelper) Encrypt(plaintext, pubKeyPem []byte) (ciphertext []byte) {

	//pem decode
	block, _ := pem.Decode(pubKeyPem)

	//x509 decode
	publicKeyInterface, x509DecodeErr := x509.ParsePKIXPublicKey(block.Bytes)
	if x509DecodeErr != nil {
		panic(x509DecodeErr)
	}

	pubKey := publicKeyInterface.(*rsa.PublicKey)

	ciphertext, encryptErr := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plaintext)
	if encryptErr != nil {
		panic(encryptErr)
	}
	//避免数据过长报错，故分段加密
	//partLen := pubKey.N.BitLen() / 8 - 11
	//chunks := split(data, partLen)
	//buffer := bytes.NewBufferString("")
	//for _, chunk := range chunks {
	//	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, chunk)
	//	if err != nil {
	//		return nil, err
	//	}
	//	buffer.Write(bytes)
	//}
	//return buffer.Bytes(), nil

	//keySize, srcSize := publicKey.(*rsa.PublicKey).Size(), len(src)
	//pub :=  publicKey.(*rsa.PublicKey)
	////logs.Debug("密钥长度：", keySize, "\t明文长度：\t", srcSize)
	////单次加密的长度需要减掉padding的长度，PKCS1为11
	//offSet, once := 0, keySize-11
	//buffer := bytes.Buffer{}
	//for offSet < srcSize {
	//	endIndex := offSet + once
	//	if endIndex > srcSize {
	//		endIndex = srcSize
	//	}
	//	// 加密一部分
	//	bytesOnce, err := rsa.EncryptPKCS1v15(myrand.Reader, pub, src[offSet:endIndex])
	//	if err != nil {
	//		return nil, err
	//	}
	//	buffer.Write(bytesOnce)
	//	offSet = endIndex
	//}
	//bytesEncrypt = buffer.Bytes()
	return
}

// Decrypt
// @param []byte ciphertext
// @param []byte priKeyPem
//
// @return []byte plaintext
func (r *rsaHelper) Decrypt(ciphertext, priKeyPem []byte) (plaintext []byte) {

	// pem decode
	block, _ := pem.Decode(priKeyPem)

	// x509 decode
	//privateKey, x509DecodeErr := x509.ParsePKCS8PrivateKey(block.Bytes)
	privateKey, x509DecodeErr := x509.ParsePKCS1PrivateKey(block.Bytes)
	if x509DecodeErr != nil {
		panic(x509DecodeErr)
	}

	plaintext, decryptErr := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if decryptErr != nil {
		panic(decryptErr)
	}
	//partLen := priKey.N.BitLen() / 8
	//chunks := split(data, partLen)
	//buffer := bytes.NewBufferString("")
	//for _, chunk := range chunks {
	//	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, chunk)
	//	if err != nil {
	//		return nil, err
	//	}
	//	buffer.Write(decrypted)
	//}
	//
	//return buffer.Bytes(), nil

	//keySize, srcSize := private.Size(), len(src)
	////logs.Debug("密钥长度：", keySize, "\t密文长度：\t", srcSize)
	//var offSet = 0
	//var buffer = bytes.Buffer{}
	//for offSet < srcSize {
	//	endIndex := offSet + keySize
	//	if endIndex > srcSize {
	//		endIndex = srcSize
	//	}
	//	bytesOnce, err := rsa.DecryptPKCS1v15(myrand.Reader, private, src[offSet:endIndex])
	//	if err != nil {
	//		return nil, err
	//	}
	//	buffer.Write(bytesOnce)
	//	offSet = endIndex
	//}
	//bytesDecrypt = buffer.Bytes()
	return
}

// Signature
//
// @param []byte plaintext
// @param []byte priKeyPem
// @param crypto.Hash hashAlg
//
// @return string sign, error err
func (r *rsaHelper) Signature(plaintext, priKeyPem []byte, hashAlg crypto.Hash) (sign string, err error) {
	sign = ""
	err = nil

	block, _ := pem.Decode(priKeyPem)

	//privateKey,x509DecodeErr :=x509.ParsePKCS8PrivateKey(block.Bytes)
	privateKey, x509DecodeErr := x509.ParsePKCS1PrivateKey(block.Bytes)
	if x509DecodeErr != nil {
		err = x509DecodeErr
		return
	}

	hashInstance := hashAlg.New()
	_, hashWriteErr := hashInstance.Write(plaintext)
	if hashWriteErr != nil {
		err = hashWriteErr
		return
	}
	hashed := hashInstance.Sum(nil)

	signByte, signErr := rsa.SignPKCS1v15(rand.Reader, privateKey, hashAlg, hashed)
	if signErr != nil {
		err = signErr
		return
	}

	sign = base64.StdEncoding.EncodeToString(signByte)

	return
}

// Verify
//
// @param []byte plaintext
// @param string sign
// @param []byte pubKeyPem
// @param crypto.Hash hashAlg
//
// @return error err
func (r *rsaHelper) Verify(plaintext []byte, sign string, pubKeyPem []byte, hashAlg crypto.Hash) (err error) {

	signBytes, base64DecodedErr := base64.StdEncoding.DecodeString(sign)
	if base64DecodedErr != nil {
		err = base64DecodedErr
		return
	}

	block, _ := pem.Decode(pubKeyPem)

	publicKey, x509DecodeErr := x509.ParsePKIXPublicKey(block.Bytes)
	if x509DecodeErr != nil {
		err = x509DecodeErr
		return
	}

	hashInstance := hashAlg.New()
	_, hashWriteErr := hashInstance.Write(plaintext)
	if hashWriteErr != nil {
		err = hashWriteErr
		return
	}
	hashed := hashInstance.Sum(nil)

	return rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), hashAlg, hashed, signBytes)
}

func isExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsExist(err) {
			return true
		}

		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	return true
}
