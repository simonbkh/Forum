package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

/// generate session token
func GenerateSessionToken() (string, error) {
	tokne := make([]byte, 32)
	_, err := rand.Read(tokne)
	if err != nil {
		return "", errors.New("creation sissiontoken")
	}

	tokn := base64.URLEncoding.EncodeToString(tokne)
	return tokn, nil
}
