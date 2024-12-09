package services

import (
	"errors"
	"net/http"

	"forum/internal/data/queries"

	//"forum/internal/data/utils"
	"forum/internal/logic/utils"
	"forum/internal/logic/validators"

	"github.com/gofrs/uuid"
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
	session_token := ""
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
	err = GenerateSessionToken(&session_token)
	if err != nil {
		return errors.New("couldn't generate token")
	}
	// fmt.Println(email)
	queries.InsertSession(email, session_token)
	utils.SetTokenCookie(w, session_token)

	return nil
}

// Generate UUID v4
func GenerateSessionToken(token *string) error {
	uuid, err := uuid.NewV4()
	*token = uuid.String()
	if err != nil {
		return err
	}
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.String() == "" {
		// deber
		return err
	}
	err = queries.Logout(cookie.String())
	utils.SetTokenCookie(w, "")
	if err != nil {
		// handli zeb
		return err
	}
	return nil
}
