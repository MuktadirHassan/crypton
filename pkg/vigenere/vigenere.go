package vigenere

func Encrypt(plaintext, keyword string) string {
	var result string
	keyIndex := 0

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			plainChar := int(char - 'a')
			keyChar := int(keyword[keyIndex] - 'a')
			encryptedChar := (plainChar + keyChar) % 26
			result += string(encryptedChar + 'a')
			keyIndex = (keyIndex + 1) % len(keyword)
		} else if char >= 'A' && char <= 'Z' {
			plainChar := int(char - 'A')
			keyChar := int(keyword[keyIndex] - 'a')
			encryptedChar := (plainChar + keyChar) % 26
			result += string(encryptedChar + 'A')
			keyIndex = (keyIndex + 1) % len(keyword)
		} else {
			result += string(char)
		}
	}

	return result
}

func Decrypt(ciphertext, keyword string) string {
	var result string
	keyIndex := 0

	for _, char := range ciphertext {
		if char >= 'a' && char <= 'z' {
			cipherChar := int(char - 'a')
			keyChar := int(keyword[keyIndex] - 'a')
			decryptedChar := (cipherChar - keyChar + 26) % 26
			result += string(decryptedChar + 'a')
			keyIndex = (keyIndex + 1) % len(keyword)
		} else if char >= 'A' && char <= 'Z' {
			cipherChar := int(char - 'A')
			keyChar := int(keyword[keyIndex] - 'a')
			decryptedChar := (cipherChar - keyChar + 26) % 26
			result += string(decryptedChar + 'A')
			keyIndex = (keyIndex + 1) % len(keyword)
		} else {
			result += string(char)
		}
	}

	return result
}
