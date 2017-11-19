package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

type SessionResponse struct {
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func generateRandomBytes(n int) ([]byte, error) {
	byteArray := make([]byte, n)
	_, err := rand.Read(byteArray)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func generateBase64Token(length int) (string, error) {
	bytes, err := generateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}

func renderError(w http.ResponseWriter, message string, status int) {
	res := ErrorResponse{
		Error: message,
	}

	bytes, _ := json.Marshal(res)
	w.WriteHeader(status)
	w.Write(bytes)
}
