package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MuktadirHassan/crypton/pkg/affine"
)

type AffineCipherHandler struct{}

func (a AffineCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Plaintext string `json:"plaintext"`
		Key1      int    `json:"key1"`
		Key2      int    `json:"key2"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ciphertext := affine.Encrypt(request.Plaintext, request.Key1, request.Key2)

	var response = struct {
		Ciphertext string `json:"ciphertext"`
	}{
		Ciphertext: ciphertext,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a AffineCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Ciphertext string `json:"ciphertext"`
		Key1       int    `json:"key1"`
		Key2       int    `json:"key2"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plaintext := affine.Decrypt(request.Ciphertext, request.Key1, request.Key2)

	var response = struct {
		Plaintext string `json:"plaintext"`
	}{
		Plaintext: plaintext,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
