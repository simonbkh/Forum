package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"forum/internal/data/queries"
	"forum/internal/logic/utils"
	"forum/internal/logic/validators"
)

// Authentication logic
var Text = "tt"

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
	// tier 4 cookie
	var sessionToken string
	var expryTime time.Time
	token, errr := r.Cookie("SessionToken")
	sessionToken, err = utils.GenerateSessionToken()
	if err != nil {
		return err
	}
	expryTime = time.Now().Add(time.Hour * 24)
	if errr != nil || token.Value == "" {
		exit, _ := queries.IssesionidAvailable("", email)
		if exit {
			er := queries.Removesesionid("", email)
			if er != nil {
				fmt.Println(er)
				return er
			}
		}
		er := queries.Insersessions(sessionToken, email, expryTime)
		if er != nil {
			return er
		}
	} else {
		exit, _ := queries.IssesionidAvailable(token.Value, email)
		if exit {
			er := queries.UpdiateSesiontoken(sessionToken, email, expryTime)
			if er != nil {
				return er
			}
		} else {
			er := queries.Insersessions(sessionToken, email, expryTime)
			if er != nil {
				return er
			}
		}
	}
	Text = ""
	http.SetCookie(w, &http.Cookie{
		Name:    "SessionToken",
		Value:   sessionToken,
		Expires: expryTime,
	})

	return nil
}

// //service logout
func Logout_Service(w http.ResponseWriter, r *http.Request) error {
	token, er := r.Cookie("SessionToken")
	if er != nil {
		return er
	}
	err := queries.Removesesionid(token.Value, "")
	if err != nil {
		return err
	}
	return nil
}
