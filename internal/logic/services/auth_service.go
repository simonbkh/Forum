package services

import (
	"errors"
	"forum/internal/data/queries"
	"forum/internal/logic/utils"
	"forum/internal/logic/validators"
	"net/http"
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
	hashedpass, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	// tier 3 data
	if queries.IsUserExist(username, email) {
		return errors.New("invalid credentiels")
	}
	queries.InserUser(username, email, hashedpass)

	return nil
}

func Login_Service(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")
	/// cooxkes
	sessionToke := r.cookes("SessionToken")
	if sessionToke == nil {
		fmt.Println("okkkkkk")
		return
	}
	sessionToken = utils.GenerateSessionToken()
	if queries.IsUserExist(sessionToken, email) {
		return errors.New("invalid sission token ")
	}
	queries.InserSisionToken(sessionToke)
	http.SetCookie(w, &http.Cookie{
		Name:  "SessionToken",
		Value: sessionToken,
	})

	// tier 2 logic
	err := validators.User_Validator("", email, password)
	if err != nil {
		return err
	}

	if !queries.Checkemail(email) {
		return errors.New("wrong email")
	}
	// tier 3 data

	HashPassword, err := queries.GetHashedPass(email)
	if err != nil {
		return err
	}

	if !utils.ComparePassAndHashedPass(HashPassword, password) {
		return errors.New("wrong password")
	}

	return nil
}
