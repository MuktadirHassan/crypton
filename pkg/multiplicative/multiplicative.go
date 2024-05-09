package multiplicative

func Encrypt(plaintext string, key int) string {
	var result string
	for _, c := range plaintext {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a')*rune(key)%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A')*rune(key)%26 + 'A')
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

	// Find the multiplicative inverse of key
	var inverse int
	for i := 0; i < 26; i++ {
		if (key*i)%26 == 1 {
			inverse = i
			break
		}
	}

	if inverse == 0 {
		return "Key is not coprime with 26"
	}

	for _, c := range ciphertext {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a')*rune(inverse)%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A')*rune(inverse)%26 + 'A')
		} else {
			result += string(c)
		}

	}
	return result
}
