package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha1"
    "io"
    "fmt"
    "strings"
    "encoding/base64"
)

func encrypt(json []byte) []byte {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(json))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], json)

	// convert to base64
	return ciphertext
}

func decrypt(cryptoText []byte) []byte {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(cryptoText) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := cryptoText[:aes.BlockSize]
	cryptoText = cryptoText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cryptoText, cryptoText)
	return cryptoText
}

func createPass(params, secret string) (pass string) {
    params = params + secret
    sum := fmt.Sprintf("%x", sha1.Sum([]byte(params)))
    return strings.TrimRight(base64.StdEncoding.EncodeToString([]byte(sum)), "=")
}