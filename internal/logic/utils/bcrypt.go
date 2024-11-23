// Password encryption

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func ComparePassAndHashedPass(password, hashedpassword string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedpassword))
	if err != nil {
		return err
	}
	return  nil
}
