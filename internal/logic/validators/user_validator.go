package validators

import (
	"errors"
	"regexp"
)

// function validator emal & password & username of user
func User_Validator(username, email, password string) error {
	err := UsernameValidator(username)
	if err != nil {
		return err
	}

	err = PasswordValidator(password)
	if err != nil {
		return err
	}
	err = EmailValidator(email)
	if err != nil {
		return err
	}
	return nil
}

// check username is corect or not
func UsernameValidator(username string) error {
	re, err := regexp.Compile(`^[a-zA-Z][a-zA-Z0-9_]{2,14}$`)
	if err != nil {
		return err
	}
	match := re.MatchString(username)
	if !match {
		return errors.New("username is not corect ")
	}
	return nil
}

// check password is corect or not
func PasswordValidator(password string) error {
	if len(password) > 21 {
		return errors.New("password is too long (max: 20 characters)")
	} else if len(password) < 5 {
		return errors.New("password is too short (min: 5 characters)")
	}
	return nil
}

// check email is corect or not
func EmailValidator(email string) error {
	re, err := regexp.Compile(`^[a-zA-Z0-9._]{3,}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return err
	}
	match := re.MatchString(email)
	if !match {
		return errors.New("email is not corect ")
	}
	return nil
}

func Login_Validat(email, password string) error {
	err := EmailValidator(email)
	if err != nil {
		return err
	}
	err = PasswordValidator(password)
	if err != nil {
		return err
	}
	return nil
}
