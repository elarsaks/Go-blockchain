package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PrivateKey
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, publicKey := ecdsa.GenerateKeyPair(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.publicKey.PublicKey
	return w
}

func (w *Wallet) privateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) privateKeyStr() *ecdsa.PublicKey {
	return fmt.Printf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) publicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) publicKeyStr() *ecdsa.PublicKey {
	return fmt.Printf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
