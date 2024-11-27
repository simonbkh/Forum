package validators

import (
	"errors"
	// "fmt"
	"regexp"
)

func User_Validator(username, email, password string) error {
	if username != "" {
		err := UsernameValidator(username)
		if err != nil {
			return err
		}
	}
	err := PasswordValidator(password)
	if err != nil {
		return err
	}
	err = EmailValidator(email)
	if err != nil {
		return err
	}
	return nil
}

func UsernameValidator(username string) error {
	re, err := regexp.Compile(`^[a-zA-Z][a-zA-Z0-9_]{2,15}$`)
	if err != nil {
		return errors.New("username is not corect ---")
	}
	match := re.MatchString(username)
	if !match {
		return errors.New("username is not corect ---")
	}
	if len(username) > 16 {
		return errors.New("username is too long (max: 10 characters)")
	} else if len(username) < 3 {
		return errors.New("username is too short (min: 3 characters)")
	}
	return nil
}

func PasswordValidator(password string) error {
	if len(password) > 15 {
		return errors.New("password is too long (max: 15 characters)")
	} else if len(password) < 5 {
		return errors.New("password is too short (min: 5 characters)")
	}
	return nil
}

func EmailValidator(email string) error {
	re, err := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return errors.New("the format  email not corect ")
	}
	match := re.MatchString(email)
	if !match {
		return errors.New("the email not corect !!")
	}
	if len(email) > 25 {
		return errors.New("email is too long (max: 25 characters)")
	} else if len(email) < 13 {
		return errors.New("email is too short (min: 13 characters)")
	}
	return nil
}
