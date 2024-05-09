package handler

import "net/http"

type AutokeyCipherHandler struct{}

func (a AutokeyCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encrypt"))
}

func (a AutokeyCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Decrypt"))
}
