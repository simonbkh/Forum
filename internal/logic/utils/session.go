package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func GenerateSessionToken() (string, error) {
	tokne := make([]byte, 32)
	_, err := rand.Read(tokne)
	if err != nil {
		return nil, errors.New("creation sissiontoken")
	}

	tokne = base64.URLEncoding.EncodeToString(tokne)
	return tokne, nil
}
