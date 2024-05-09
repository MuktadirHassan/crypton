package additive

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
	if key < 0 {
		key = 26 + key
	}
	if key > 26 {
		key = key % 26
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
