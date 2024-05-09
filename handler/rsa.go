package handler

import (
	"net/http"
)

// RsaCipherHandler is a struct that defines the handler for the RSA cipher
type RsaCipherHandler struct{}

func (r RsaCipherHandler) GenerateKeyHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Generate Key"))
}

func (r RsaCipherHandler) EncryptHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Encrypt"))
}

func (r RsaCipherHandler) DecryptHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Decrypt"))
}
