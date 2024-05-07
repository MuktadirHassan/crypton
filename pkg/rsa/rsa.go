package rsa

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Enctrypt() {
	// do something

}

func Decrypt() {
	// do something
}

/*
 * Generate RSA key pair
 e - public key (1 < e < φ(n) and gcd(e, φ(n)) = 1)
 d - private key (d = e^-1 mod φ(n))
 n - e * d
*/

func generateRSAKeyPair(bits int) (e, d, n *big.Int, err error) {
	// do something
	p, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Println(p)

	return nil, nil, nil, nil
}

// find out extended gcd
func extendedGCD(a, eulerN *big.Int) (x, y *big.Int) {
	// do something
	return nil, nil
}
