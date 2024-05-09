package multiplicative

func Encrypt(plaintext string, key int) string {
	var result string
	for _, c := range plaintext {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a'+rune(key))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A'+rune(key))%26 + 'A')
		} else {
			result += string(c)
		}
	}

	return result
}

func Decrypt(ciphertext string, key int) string {
	var result string
	if key == 0 {
		return ciphertext
	}
	if key < 0 {
		key = 26 + key%26
	}

	if gcd := gcd(key, 26); gcd != 1 {
		return "Invalid . Key must be coprime with 26. Coprimes: 1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25"
	}

	for _, c := range ciphertext {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a'-rune(key)+26)%26 + 'a') // +26 to handle negative values
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A'-rune(key)+26)%26 + 'A')
		} else {
			result += string(c)
		}
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
