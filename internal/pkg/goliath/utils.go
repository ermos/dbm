package goliath

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
)

// paddingPKCS5 adds PKCS#5 padding to the given data to make its length a multiple of the block size.
// https://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS#5_and_PKCS#7
func paddingPKCS5(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, pad...)
}

// computeHMAC computes the HMAC-SHA256 tag for the given data and key.
func computeHMAC(data []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
