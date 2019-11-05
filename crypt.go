package main

import (
	"crypto/des"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha1"
    "io"
    "fmt"
    "strings"
    "encoding/base64"
	"bytes"
)

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(
		origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func getBlock(secret []byte) (cipher.Block) {
	for len(secret) < 24 {
		secret = append(secret, secret...)
	}
	secret = secret[:24]

	block, err := des.NewTripleDESCipher(secret)
	if err != nil {
		panic(err)
	}

	return block
}

func encrypt(text []byte, secret []byte) []byte {
	block := getBlock(secret)
	text = zeroPadding(text, block.BlockSize())

	cipherText := make([]byte, block.BlockSize()+len(text))
	iv := cipherText[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[block.BlockSize():], text)

	return cipherText
}

func decrypt(cryptedText []byte, secret []byte) []byte {
	block := getBlock(secret)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(cryptedText) < block.BlockSize() {
		panic("ciphertext too short")
	}
	iv := cryptedText[:block.BlockSize()]
	cryptedText = cryptedText[block.BlockSize():]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cryptedText, cryptedText)
	return zeroUnPadding(cryptedText)
}

func createPass(params string, secret []byte, length int) (pass string) {
    params = params + string(secret)
    sum := fmt.Sprintf("%x", sha1.Sum([]byte(params)))
	pass = strings.TrimRight(base64.StdEncoding.EncodeToString([]byte(sum)), "=")
	if (len(pass) < length) {
		pass = pass + createPass(pass, secret, length)
	} else {
		pass = pass[:length]
	}
	return pass
}