package rsa

import (
	"crypto/rand"
	"math/big"
)

// https://en.wikipedia.org/wiki/Fermat_number (Fermat number)

func Enctrypt() {

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
	- d is the modular multiplicative inverse of e modulo phi(n)
5. Public key is (n, e)
6. Private key is (n, d)
7. Encrypt message m using public key: c = m^e mod n
8. Decrypt message c using private key: m = c^d mod n

*/

// GeneratePrimeNumber generates a prime number with given bits
func generatePrime(bits int) *big.Int {
	// generate a random number of given bits
	randomNumber, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return randomNumber
}

// Genrates a public and private key pair with given bits
// bits: bit-length of the key
// returns n (modulus), e (public exponent), d (private exponent)
func GenerateKeyPair(bits int) (n *big.Int, public *big.Int, private *big.Int) {
	// generate two prime numbers each of half the bit-length
	p := generatePrime(bits / 2)
	q := generatePrime(bits / 2)
	n = new(big.Int).Mul(p, q)
	// calculate phi_n
	phi_n := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	// choose e
	eList := []int{3, 5, 17, 65537}
	e := big.NewInt(int64(eList[3]))
	for i := 0; i < len(eList); i++ {
		if new(big.Int).GCD(nil, nil, e, phi_n).Cmp(big.NewInt(1)) == 0 {
			break
		}
		e = big.NewInt(int64(eList[i]))
	}

	// calculate d
	d := new(big.Int).ModInverse(e, phi_n)

	return n, e, d
}
