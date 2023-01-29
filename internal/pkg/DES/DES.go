package DES

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func Encrypt(plainText []byte, key []byte) (string, error) {
	key = PKCS5Padding(key, 8)[0:8]
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	key = PKCS5Padding(key, block.BlockSize())
	plainText = PKCS5Padding(plainText, block.BlockSize())
	// plainText must be a multiple of block size
	if len(plainText)%block.BlockSize() != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}

	cipherText := make([]byte, len(plainText))
	dst := cipherText
	for len(plainText) > 0 {
		block.Encrypt(dst, plainText[:block.BlockSize()])
		plainText = plainText[block.BlockSize():]
		dst = dst[block.BlockSize():]
	}
	cipherText = []byte(base64.StdEncoding.EncodeToString(cipherText))
	return string(cipherText), nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Decrypt(cipherText_ string, key []byte) (string, error) {
	key = PKCS5Padding(key, 8)[0:8]
	cipherText, err := base64.StdEncoding.DecodeString(cipherText_)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	key = PKCS5Padding(key, block.BlockSize())

	// cipherText must be a multiple of block size
	if len(cipherText)%block.BlockSize() != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	plainText := make([]byte, len(cipherText))
	dst := plainText
	for len(cipherText) > 0 {
		block.Decrypt(dst, cipherText[:block.BlockSize()])
		cipherText = cipherText[block.BlockSize():]
		dst = dst[block.BlockSize():]
	}
	plainText = PKCS5UnPadding(plainText)
	return string(plainText), nil
}

func main() {
	key := []byte("123456789")
	plainText := []byte("hello world")
	cipherText, err := Encrypt(plainText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("cipherText:", cipherText)
	decodeText, err := Decrypt(cipherText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodeText:", decodeText)
}
