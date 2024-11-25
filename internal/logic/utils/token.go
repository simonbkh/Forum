package utils

import (
	"crypto/rand"
	// "github.com/gofrs/uuid"
	"encoding/hex"
)

func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// token := uuid.NewV4().String()
	// return token, nil
	return hex.EncodeToString(bytes), nil
}
