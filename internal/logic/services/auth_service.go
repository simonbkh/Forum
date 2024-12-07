package services

import (
	"errors"
	"net/http"

	"forum/internal/data/queries"
	"forum/internal/logic/utils"
	"forum/internal/logic/validators"
	"forum/internal/presentation/templates"
)

// Authentication logic

func Register_Service(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	// tier 2 logic
	err := validators.User_Validator(username, email, password)
	if err != nil {
		return err
	}
	// tier 3 data
	if queries.IsUserExist(username, email) {
		return errors.New("user already exists")
	}
	hashedpass, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	queries.InserUser(username, email, hashedpass)

	return nil
}

func Login_Service(w http.ResponseWriter, r *http.Request) (string, error) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := validators.Login_Validat(email, password)
	if err != nil {
		return "", err
	}

	if !queries.Checkemail(email) {
		return "", errors.New("wrong email")
	}
	// tier 3 data

	HashPassword, err := queries.GetHashedPass(email)
	if err != nil {
		return "", err
	}

	if !utils.ComparePassAndHashedPass(HashPassword, password) {
		return "", errors.New("wrong password")
	}
	var tocken string
	// Set token in cookie
	str := queries.CheckeToken(email)

	if str != "" {
		tocken = str
		tocken, err = utils.GenerateToken(16)
		if err != nil {
			return "", err
		}
		err = queries.Insert_OR_remove_token(tocken, email)
		if err != nil {
			return "", err
		}
	}
	if str == "" {
		tocken, err = utils.GenerateToken(16)
		if err != nil {
			return "", err
		}
		err := queries.Insert_OR_remove_token(tocken, email)
		if err != nil {
			return "", err
		}
	}
	return tocken, nil
}

func Log_out_Service(w http.ResponseWriter, r *http.Request) error {
	// var cookie *http.Cookie
	if !validators.Check_cokes(r) {
		return errors.New("unauthorized")
	}
	cookie, err := r.Cookie("token")
	err = queries.Insert_OR_remove_token(cookie.Value, "")
	if err != nil {
		return err
	}
	return nil
}

func Method(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ErrorTemlate.Execute(w, "Method Not Allowed.....")
		return false
	}
	return true
}
