package goliath

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

// EncryptData encrypts the given data using AES in CBC mode with a key derived from the master password.
// It returns a byte array that contains the encrypted data, IV, salt, and HMAC tag.
func EncryptData(data []byte, masterPassword string) (s string, err error) {
	// Generate a random IV (Initialization Vector)
	// https://en.wikipedia.org/wiki/Initialization_vector
	iv := make([]byte, aes.BlockSize)
	if _, err = rand.Read(iv); err != nil {
		return
	}

	// Generate a random salt for key derivation (prevent dictionary attacks)
	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return
	}

	// Derive a key from the master password using PBKDF2 and the salt
	key := pbkdf2.Key([]byte(masterPassword), salt, iterationPBKDF2, 32, hashAlg)

	// Encrypt the data using AES in CBC mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	paddedData := paddingPKCS5(data, blockSize)
	cipherText := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedData)

	// Compute the HMAC-SHA256 authentication tag
	hmacKey := sha256.Sum256(key)
	hmacTag := computeHMAC(cipherText, hmacKey[:])

	// Concatenate the encrypted data, IV, salt, and HMAC tag into a single byte array
	result := make([]byte, len(cipherText)+aes.BlockSize+16+sha256.Size)
	copy(result, cipherText)
	copy(result[len(cipherText):], iv)
	copy(result[len(cipherText)+aes.BlockSize:], salt)
	copy(result[len(cipherText)+aes.BlockSize+16:], hmacTag)

	return base64.StdEncoding.EncodeToString(result), nil
}
