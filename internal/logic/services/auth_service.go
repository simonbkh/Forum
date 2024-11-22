package services

import (
	"net/http"

	"forum/internal/logic/utils"
	"forum/internal/logic/validators"
)

// Authentication logic

func Register_Service(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	validators := []func(string) error{
		validators.Username_Validator,
		validators.Email_Validator,
		validators.Password_Validator,
	}

	inputs := []string{username, email, password}

	for i, validate := range validators {
		if err := validate(inputs[i]); utils.IsErrors(err) {
			return err
		}
	}

	return nil
}

func Login_Service(w http.ResponseWriter, r *http.Request) error {

	email := r.FormValue("email")
	password := r.FormValue("password")

	validators := []func(string) error{
		validators.Email_Validator,
		validators.Password_Validator,
	}

	inputs := []string{email, password}

	for i, validate := range validators {
		if err := validate(inputs[i]); utils.IsErrors(err) {
			return err
		}
	}
	

	return nil
}
