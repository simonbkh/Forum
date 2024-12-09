package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"forum/internal/data/modles"
	"forum/internal/data/queries"
)

// / generate session token
func GenerateSessionToken() (string, error) {
	tokne := make([]byte, 32)
	_, err := rand.Read(tokne)
	if err != nil {
		return "", errors.New("creation sissiontoken")
	}

	tokn := base64.URLEncoding.EncodeToString(tokne)
	return tokn, nil
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
				fmt.Println(er)
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
func CheckUserSession(r *http.Request) error {
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
