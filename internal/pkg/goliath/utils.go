package goliath

import (
	"bytes"
	"crypto/aes"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
)

// addPaddingPKCS5 adds PKCS#5 padding to the given data to make its length a multiple of the block size.
// https://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS#5_and_PKCS#7
func addPaddingPKCS5(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, pad...)
}

// removePaddingPKCS5 remove PKCS#5 padding.
func removePaddingPKCS5(data []byte) ([]byte, error) {
	padLen := int(data[len(data)-1])
	if padLen > aes.BlockSize || padLen > len(data) {
		return nil, errors.New("invalid padding")
	}

	return data[:len(data)-padLen], nil
}

// computeHMAC computes the HMAC-SHA256 tag for the given data and key.
func computeHMAC(data []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
