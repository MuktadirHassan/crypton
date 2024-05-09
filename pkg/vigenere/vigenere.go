package vigenere

func Encrypt(plaintext string, key string) string {
	var result string
	keyIndex := 0
	for _, c := range plaintext {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a'+rune(key[keyIndex]-'a'))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A'+rune(key[keyIndex]-'A'))%26 + 'A')
		} else {
			result += string(c)
		}
		keyIndex++
		if keyIndex == len(key) {
			keyIndex = 0
		}
	}

	return result
}

func Decrypt(ciphertext string, key string) string {
	var result string
	keyIndex := 0
	for _, c := range ciphertext {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a'-rune(key[keyIndex]-'a')+26)%26 + 'a') // +26 to handle negative values
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A'-rune(key[keyIndex]-'A')+26)%26 + 'A')
		} else {
			result += string(c)
		}
		keyIndex++
		if keyIndex == len(key) {
			keyIndex = 0
		}
	}
	return result
}
