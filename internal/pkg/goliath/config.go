package goliath

import "crypto/sha256"

var (
	hashAlg         = sha256.New
	iterationPBKDF2 = 1200000
)
