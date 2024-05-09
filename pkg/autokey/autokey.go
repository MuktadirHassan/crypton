package autokey

func Encrypt(plaintext string, key int) string {
	var result string
	key = key % 26

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			plainChar := int(char) - 'a'
			encryptedChar := (plainChar + key) % 26
			result += string(encryptedChar + 'a')
			key = plainChar
		} else if char >= 'A' && char <= 'Z' {
			encryptedChar := (int(char) - 'A' + key) % 26
			result += string(encryptedChar + 'A')
			key = int(char) - 'A'
		} else {
			result += string(char)
		}
	}

	return result

}

func Decrypt(ciphertext string, key int) string {
	var result string
	key = key % 26

	for _, char := range ciphertext {
		if char >= 'a' && char <= 'z' {
			decryptedChar := (int(char) - 'a' - key + 26) % 26
			result += string(decryptedChar + 'a')
			key = decryptedChar
		} else if char >= 'A' && char <= 'Z' {
			decryptedChar := (int(char) - 'A' - key + 26) % 26
			result += string(decryptedChar + 'A')
			key = decryptedChar
		} else {
			result += string(char)
		}
	}

	return result
}
