// AES: Advanced Encrypt Standard
//
// Desc: base block, 128 bit per block.
//
// Three Factor:
// 1.secret size: 128、192、256
// 2.padding: NoPadding、PKCS5Padding(default)、ISO10126Padding
// 3.work mode: ECB(Electronic Codebook Book, default)、CBC(Cipher Block Chaining)、CTR(Counter)、CFB(Cipher FeedBack)、OFB(Output FeedBack)

package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type Mode uint8

const (
	CBC Mode = 1 + iota
	ECB
	CFB
	OFB
	CTR

	errMsgPrefix            = "[REA Error] "
	errMsgNotSupportEncrypt = errMsgPrefix + "not support this block mode"
)

type aseHelper struct{}

func NewAES() *aseHelper {
	return &aseHelper{}
}

// Encrypt
//
// @param []byte plaintext
// @param string secret
//
// @return []byte ciphertext, error err
func (a *aseHelper) Encrypt(plaintext []byte, secret string) (ciphertext []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(secret)); err != nil {
		return
	}
	blockSize := block.BlockSize()

	plaintext = PKCS5Padding(plaintext, blockSize)

	ciphertext = make([]byte, len(plaintext))
	iv := []byte(secret)[:blockSize]
	encrypt := cipher.NewCBCEncrypter(block, iv)
	encrypt.CryptBlocks(ciphertext, plaintext)

	return
}

// Decrypt
//
// @param []byte ciphertext
// @param string secret
//
// @return []byte plaintext, error err
func (a *aseHelper) Decrypt(ciphertext []byte, secret string) (plaintext []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(secret)); err != nil {
		return
	}
	blockSize := block.BlockSize()

	plaintext = make([]byte, len(ciphertext))
	vi := []byte(secret)[:blockSize]
	encrypt := cipher.NewCBCDecrypter(block, vi)
	encrypt.CryptBlocks(plaintext, ciphertext)

	plaintext = PKCS5Unpadding(plaintext)

	return
}

// PKCS5Padding
// 原理：如果明文块少于 blockSize 个字节（按16字节分组即128bit，一般按秘钥字节数位分组步长），在明文块末尾补足相应数量的字符，且每个字节的值等于缺少的字符数。
// @param []byte ciphertext
// @param int blockSize
//
// @return []byte
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Unpadding
//
// @param []byte origData
//
// @return []byte
func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// ISO10126填充
// 原理：如果明文块少于 blockSize 个字节（按16字节分组即128bit，一般按秘钥字节数位分组步长），在明文块末尾补足相应数量的字节，最后一个字符值等于缺少的字符数，其他字符填充随机数。
func ISO10126Padding() {

}

// ISO10126解填充
func ISO10126Unpadding() {

}
