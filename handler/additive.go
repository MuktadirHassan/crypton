package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MuktadirHassan/crypton/pkg/additive"
)

type AdditiveCipherHandler struct{}

func (a AdditiveCipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Key       string `json:"key"`
		Plaintext string `json:"plaintext"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		// todo: send json err
		http.Error(w,
			"Invalid request body",
			http.StatusBadRequest)
		return
	}
	// convert key to int
	key, err := strconv.Atoi(reqBody.Key)
	if err != nil {
		http.Error(w,
			"Invalid key",
			http.StatusBadRequest)
		return
	}
	encryptedText := additive.Encrypt(reqBody.Plaintext, key)

	var resBody struct {
		EncryptedText string `json:"ciphertext"`
	}

	resBody.EncryptedText = encryptedText

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a AdditiveCipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Key        string `json:"key"`
		Ciphertext string `json:"ciphertext"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w,
			"Invalid request body",
			http.StatusBadRequest)
		return
	}

	key, err := strconv.Atoi(reqBody.Key)
	if err != nil {
		http.Error(w,
			"Invalid key",
			http.StatusBadRequest)
		return
	}
	text := reqBody.Ciphertext

	decryptedText := additive.Decrypt(text, key)

	var resBody struct {
		DecryptedText string `json:"plaintext"`
	}

	resBody.DecryptedText = decryptedText

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
