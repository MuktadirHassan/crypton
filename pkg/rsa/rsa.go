package rsa

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Enctrypt(plaintext string, e, n *big.Int) (ciphertext *big.Int, err error) {
	// do something
	return nil, nil
}

func Decrypt() {
	// do something
}

func generateRSAKeyPair(bits int) (e, d, n *big.Int, err error) {
	// do something
	p, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Println(p)

	return nil, nil, nil, nil
}
