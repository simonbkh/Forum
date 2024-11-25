package validators

import (
	"errors"
	"regexp"
)

func Username_Validator(username string) error {
	reg, err := regexp.Compile(`\W+`)
	if err != nil {
		return err
	}
	guud := reg.MatchString(username)

	if !guud {
		if len(username) > 10 {
			return errors.New("long username")
		} else if len(username) < 4 {
			return errors.New("short username")
		}
	} else {
		return errors.New("invalid username")
	}
	return nil
}

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
	if len(username) > 10 {
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

// func Password_Validator(password string) error{
// 	// if len(password) > 20 {
// 	// 	return errors.New("long password")
// 	// }else if len(password) < 3 {
// 	// 	return errors.New("short password")
// 	// }

// 	// return nil
// }

func EmailValidator(email string) error{
	if len(email) > 20 {
		return errors.New("long password")
	}
	return nil
}
