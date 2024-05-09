package affine

// Encrypt encrypts the plaintext using the affine cipher. Key1 is the multiplicative key and key2 is the additive key.
func Encrypt(plaintext string, key1, key2 int) string {
	var result string
	for _, c := range plaintext {
		if c >= 'a' && c <= 'z' {
			result += string((rune(key1)*(c-'a')+rune(key2))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((rune(key1)*(c-'A')+rune(key2))%26 + 'A')
		} else {
			result += string(c)
		}
	}

	return result
}

func Decrypt(ciphertext string, key1 int, key2 int) string {
	var result string
	inv := modInverse(key1, 26)
	for _, c := range ciphertext {
		if c >= 'a' && c <= 'z' {
			result += string((inv*int(c-'a'-rune(key2)+26))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((inv*int(c-'A'-rune(key2)+26))%26 + 'A')
		} else {
			result += string(c)
		}

	}
	return result
}

func modInverse(a, m int) int {
	a = a % m
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return 1
}
