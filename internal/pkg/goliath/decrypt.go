package goliath

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/pbkdf2"
)

// DecryptData decrypts the given data using AES in CBC mode with a key derived from the master password.
func DecryptData(content string, masterPassword string) (b []byte, err error) {
	data, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return
	}

	// Split the byte array into the encrypted data, IV, salt, and HMAC tag
	if len(data) < aes.BlockSize+16+sha256.Size {
		return nil, errors.New("invalid data")
	}

	cipherText := data[:len(data)-aes.BlockSize-16-sha256.Size]
	iv := data[len(cipherText) : len(cipherText)+aes.BlockSize]
	salt := data[len(cipherText)+aes.BlockSize : len(cipherText)+aes.BlockSize+16]
	hmacTag := data[len(data)-sha256.Size:]

	// Derive a key from the master password using PBKDF2 and the salt
	key := pbkdf2.Key([]byte(masterPassword), salt, iterationPBKDF2, 32, hashAlg)

	// Compute the HMAC-SHA256 authentication tag
	hmacKey := sha256.Sum256(key)
	expectedHMAC := computeHMAC(cipherText, hmacKey[:])
	if !hmac.Equal(hmacTag, expectedHMAC) {
		return nil, errors.New("invalid HMAC")
	}

	// Decrypt the data using AES in CBC mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	// Remove padding
	padLen := int(plainText[len(plainText)-1])
	if padLen > aes.BlockSize || padLen > len(plainText) {
		return nil, errors.New("invalid padding")
	}

	b = plainText[:len(plainText)-padLen]
	return
}
