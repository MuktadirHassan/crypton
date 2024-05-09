package handler

import "net/http"

// VigenereCipherHandler is a struct that defines the Vigenere cipher handler.
type VigenereCipherHandler struct{}

func (v VigenereCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encrypt"))
}

func (v VigenereCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Decrypt"))
}
