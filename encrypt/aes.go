// AES: Advanced Encrypt Standard
//
// Desc: base block, 128 bit per block.
//
// Three Factor:
// 1.secret size(bit): 128、192、256
// 2.padding: NoPadding:
//	PKCS5Padding(default)
//	PKCS7Padding
//	ISO10126Padding
// 3.work mode:
//	ECB(Electronic Codebook Book, default)
//	CBC(Cipher Block Chaining)
//	CTR(Counter)
//	CFB(Cipher FeedBack)
//	OFB(Output FeedBack)

package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type Mode uint8

const (
	// block encrypt
	CBC Mode = 1 + iota
	ECB
	CTR

	// stream encrypt
	CFB
	OFB
)

var (
	ErrWorkModeNotSupport = errors.New("work mode not support ")
)

type aseHelper struct {
	mode Mode
}

func NewAES(mode Mode) *aseHelper {
	return &aseHelper{
		mode: mode,
	}
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

	switch a.mode {
	case ECB:
		blockSize := block.BlockSize()
		plaintext = a.PKCS7Padding(plaintext, blockSize)
		ciphertext = make([]byte, len(plaintext))
		// block encrypt
		for start, end := 0, blockSize; start < len(plaintext); start, end = start+blockSize, end+blockSize {
			block.Encrypt(ciphertext[start:end], plaintext[start:end])
		}
		break
	case CBC:
		blockSize := block.BlockSize()
		plaintext = a.PKCS7Padding(plaintext, blockSize)
		ciphertext = make([]byte, len(plaintext))
		iv := []byte(secret)[:blockSize]
		encrypt := cipher.NewCBCEncrypter(block, iv)
		encrypt.CryptBlocks(ciphertext, plaintext)
		break
	case CTR:
		break
	case CFB:
		break
	case OFB:
		break
	default:
		err = ErrWorkModeNotSupport
		break
	}

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

	switch a.mode {
	case ECB:
		blockSize := block.BlockSize()
		plaintext = make([]byte, len(ciphertext))
		for start, end := 0, blockSize; start < len(ciphertext); start, end = start+blockSize, end+blockSize {
			block.Decrypt(plaintext[start:end], ciphertext[start:end])
		}
		plaintext = a.PKCS7UnPadding(plaintext)
		break
	case CBC:
		blockSize := block.BlockSize()
		plaintext = make([]byte, len(ciphertext))
		vi := []byte(secret)[:blockSize]
		encrypt := cipher.NewCBCDecrypter(block, vi)
		encrypt.CryptBlocks(plaintext, ciphertext)

		plaintext = a.PKCS7UnPadding(plaintext)
		break
	case CTR:
		break
	case CFB:
		break
	case OFB:
		break
	default:
		err = ErrWorkModeNotSupport
		break
	}

	return
}

// PKCS5Padding（是 PKCS7Padding 的子集，块大小固定为8字节）
// 原理：如果明文块少于 blockSize 个字节（按16字节分组即128bit，一般按秘钥字节数位分组步长），在明文块末尾补足相应数量的字符，且每个字节的值等于缺少的字符数。
// @param []byte ciphertext
// @param int blockSize
//
// @return []byte
func (a *aseHelper) PKCS5Padding(ciphertext []byte) []byte {
	return a.PKCS7Padding(ciphertext, 8)
}

// PKCS5UnPadding
// @param []byte origData
//
// @return []byte
func (a *aseHelper) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// PKCS7Padding
// 原理：如果明文块少于 blockSize 个字节（按16字节分组即128bit，一般按秘钥字节数位分组步长），在明文块末尾补足相应数量的字符，且每个字节的值等于缺少的字符数。
// @param []byte ciphertext
// @param int blockSize
//
// @return []byte
func (a *aseHelper) PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7UnPadding
func (a *aseHelper) PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// ISO10126填充
// 原理：如果明文块少于 blockSize 个字节（按16字节分组即128bit，一般按秘钥字节数位分组步长），在明文块末尾补足相应数量的字节，最后一个字符值等于缺少的字符数，其他字符填充随机数。
func (a *aseHelper) ISO10126Padding() {

}

// ISO10126解填充
func (a *aseHelper) ISO10126UnPadding() {

}
