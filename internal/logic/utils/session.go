package utils

import (
	"net/http"
	"time"

	"forum/internal/data/modles"
	"forum/internal/data/queries"

	"github.com/google/uuid"
)

// / generate session token
func GenerateSessionToken() (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

// manage session token in database is it available or not
func ManageSessionToken(email string, r *http.Request) (string, time.Time, error) {
	expryTime := time.Now().Add(time.Hour * 24)
	sessionToken, err := GenerateSessionToken()
	if err != nil {
		return "", time.Time{}, err
	}
	token, errr := r.Cookie("SessionToken")
	if errr != nil || token.Value == "" {

		exit, _ := queries.IssesionidAvailable("", email) // katchof wax kayn aluser
		if exit {
			er := queries.Removesesionid("", email)
			if er != nil {
				return "", time.Time{}, er
			}
		}
		er := queries.Insersessions(sessionToken, email, expryTime)
		if er != nil {
			return "", time.Time{}, er
		}
	} else {
		exit, _ := queries.IssesionidAvailable(token.Value, email)
		if exit {
			er := queries.UpdiateSesiontoken(sessionToken, email, expryTime)
			if er != nil {
				return "", time.Time{}, er
			}
		} else {
			er := queries.Insersessions(sessionToken, email, expryTime)
			if er != nil {
				return "", time.Time{}, er
			}
		}
	}

	return sessionToken, expryTime, nil
}

func CheckUserSession(r *http.Request) error { ////hna kayna xi zyada dyal error
	modles.UserStatus = false
	te, err := r.Cookie("SessionToken")
	if !modles.UserStatus {
		if err != nil || te.Value == "" {
			modles.UserStatus = false
		} else {
			bol, expiry := queries.IssesionidAvailable(te.Value, "")
			if bol && expiry.After(time.Now()) {
				modles.UserStatus = true
			} else {
				modles.UserStatus = false
			}
		}
	}
	return nil
}
