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

/**
1. Choose two very large prime numbers p and q
   - how big should p and q be?
	- p and q should be of similar bit-length (e.g., 2048 bits)
	- n = p * q, so n will be of 2 * bit-length (e.g., 4096 bits)
2. Calculate phi(n) = (p-1)(q-1)
	- why phi(n) is important?
	- because any number x raised to the power of phi(n) and then mod n will be 1 (x^phi(n) mod n = 1)
	- it will allow use to derive the inverse of e
3. Choose an integer e such that 1 < e < phi(n) and gcd(e, phi(n)) = 1
	- why gcd(e, phi(n)) = 1?:w
	- because e and phi(n) should be coprime
	- e is the public exponent
	- e is usually 3, 5, 17 or 65537
4. Calculate d such that d * e mod phi(n) = 1
	- d is the private exponent

*/

// GeneratePrimeNumber generates a prime number of bits length
func GeneratePrimeNumber(bits int) (n *big.Int) {
	n, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		fmt.Println(err)
		return
	}

	return n
}
