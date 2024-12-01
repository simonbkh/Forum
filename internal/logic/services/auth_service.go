package services

import (
	"errors"
	"net/http"

	"forum/internal/data/queries"
	"forum/internal/logic/utils"
	"forum/internal/logic/validators"
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
	/// coockes
	var sessionToke string
	// var err error

	_, errr := r.Cookie("SessionToken")
	if errr != nil {
		sessionToke, err = utils.GenerateSessionToken()
		if err != nil {
			return err
		}
		queries.InserSisionToken(sessionToke)

	} else {
		if queries.IssissiontokenExit(email) {
			sessionToke, err = utils.GenerateSessionToken()
			if err != nil {
				return err
			}
			if err := queries.UpdiateSesiontoken(sessionToke, email); err != nil {
				return err
			}

		} else {
			sessionToke, err = utils.GenerateSessionToken()
			if err != nil {
				return err
			}
			queries.InserSisionToken(sessionToke)
		}
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "SessionToken",
		Value: sessionToke,
	})

	return nil
}

// //service logout
func Logout_Service(w http.ResponseWriter, r *http.Request) error {
	token, er := r.Cookie("SessionToken")
	if er != nil {
		return er
	}
	err := queries.Removesesionid(token.Value)
	if err != nil {
		return err
	}
	return nil
}
