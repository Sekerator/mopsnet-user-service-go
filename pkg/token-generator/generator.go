package token_generator

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAuthToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}
