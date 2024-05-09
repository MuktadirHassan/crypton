package handler

import "net/http"

type MultiplicativeCipherHandler struct{}

func (m MultiplicativeCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encrypt"))
}

func (m MultiplicativeCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Decrypt"))
}
