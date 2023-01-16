package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const keyString = "testtesttesttest"

func Encrypt(plainTextByte []byte) []byte {
	// Key
	key := []byte(keyString)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	cipherTextByte := make([]byte, aes.BlockSize+len(plainTextByte))

	// Slice of first 16 bytes
	iv := cipherTextByte[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(cipherTextByte[aes.BlockSize:], plainTextByte)

	return cipherTextByte
}

func Decrypt(cipherTextByte []byte) []byte {
	// Key
	key := []byte(keyString)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(cipherTextByte) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := cipherTextByte[:aes.BlockSize]

	// Remove the IV from the ciphertext
	cipherTextByte = cipherTextByte[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(cipherTextByte, cipherTextByte)

	return cipherTextByte
}
