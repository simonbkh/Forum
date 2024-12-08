package services

import (
	"errors"
	"net/http"
	"time"

	models "forum/internal/data/database/modles"
	"forum/internal/data/queries"
	"forum/internal/logic/utils"
	"forum/internal/logic/validators"
)

// service register
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
	err = queries.InserUser(username, email, hashedpass)
	if err != nil {
		return err
	}

	return nil
}

// service login
func Login_Service(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// tier 2 logic
	err := validators.Login_Validat(email, password)
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
	sessionToken, expryTime, err := utils.ManageSessionToken(email, r)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "SessionToken",
		Value:   sessionToken,
		Expires: expryTime,
	})
	models.UserStatus = true
	return nil
}

// service logout
func Logout_Service(w http.ResponseWriter, r *http.Request) error {
	token, er := r.Cookie("SessionToken")
	if er != nil {
		return er
	}
	err := queries.Removesesionid(token.Value, "")
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "SessionToken",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	return nil
}
