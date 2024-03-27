package handler

import "net/http"

type AdditiveCipherHandler struct{}

func (a AdditiveCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encrypt"))
}

func (a AdditiveCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Decrypt"))
}
