package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

// Create a string representation of the signature.
func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}

// String2BigIntTuple converts a hexadecimal string into a tuple of big integers.
// The input string is expected to have a length of 128 characters (64 characters for x and 64 characters for y).
// It returns the tuple (x, y) as pointers to big.Int values.
func String2BigIntTuple(s string) (big.Int, big.Int) {
	// Decode the first 64 characters of the string into bytes.
	bx, _ := hex.DecodeString(s[:64])
	by, _ := hex.DecodeString(s[64:])

	var bix big.Int
	var biy big.Int

	// Set the value of bix bix to the big integer representation of bx & by.
	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)

	// Return the pointers to bix and biy.
	return bix, biy
}

func SignatureFromString(s string) *Signature {
	x, y := String2BigIntTuple(s)

	return &Signature{
		&x,
		&y,
	}
}

// Create a public key from a string.
func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := String2BigIntTuple(s)
	return &ecdsa.PublicKey{elliptic.P256(), &x, &y}
}

// Create a private key from a string.
func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	b, _ := hex.DecodeString(s)
	var bi big.Int
	_ = bi.SetBytes(b)
	return &ecdsa.PrivateKey{*publicKey, &bi}
}
