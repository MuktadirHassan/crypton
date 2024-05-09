package handler

import "net/http"

type AffineCipherHandler struct{}

func (a AffineCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encrypt"))
}

func (a AffineCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Decrypt"))
}
