package handler

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/MuktadirHassan/crypton/pkg/rsa"
)

// RsaCipherHandler is a struct that defines the handler for the RSA cipher
type RsaCipherHandler struct{}

func (r RsaCipherHandler) GenerateKeyHandler(w http.ResponseWriter, req *http.Request) {
	n, public, private := rsa.GenerateKeyPair(2048)
	var response = struct {
		N       string `json:"n"`
		Public  string `json:"public"`
		Private string `json:"private"`
	}{
		N:       n.String(),
		Public:  public.String(),
		Private: private.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r RsaCipherHandler) EncryptHandler(w http.ResponseWriter, req *http.Request) {
	var request struct {
		N       string `json:"n"`
		Public  string `json:"public"`
		Message string `json:"message"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	n := new(big.Int)
	n.SetString(request.N, 10)

	public := new(big.Int)
	public.SetString(request.Public, 10)

	ciphertext := rsa.Enctrypt(n, public, request.Message)

	var response = struct {
		Ciphertext string `json:"ciphertext"`
	}{
		Ciphertext: ciphertext,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (r RsaCipherHandler) DecryptHandler(w http.ResponseWriter, req *http.Request) {
	var request struct {
		N          string `json:"n"`
		Private    string `json:"private"`
		Ciphertext string `json:"ciphertext"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	n := new(big.Int)
	n.SetString(request.N, 10)

	private := new(big.Int)
	private.SetString(request.Private, 10)

	ciphertext := rsa.Decrypt(n, private, request.Ciphertext)

	var response = struct {
		Message string `json:"plaintext"`
	}{
		Message: ciphertext,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
