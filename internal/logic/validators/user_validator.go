package validators

import (
	"errors"
	"strings"

	"forum/internal/logic/utils"
)

func User_Validator(username, email, password string) error {
	if username != "" && email != "" && password != "" {
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

	} else {
		return errors.New("username or email or password is empty")
	}

	return nil
}

func validUser(username string) bool { //// hadi bax nxofo username wax fih ri4 mn [a-z] o [A-Z] o [0-9] o [- o _ o .]
	for _, v := range username {
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') && (v < '0' || v > '9') && (v != '.' && v != '-' && v != '_') {
			return false
		}
	}
	return true
}

func UsernameValidator(username string) error {
	if !validUser(username) {
		return errors.New("username should only contain lowercase letters, uppercase letters, numbers, '-', '_', and '.' symbols")
	}
	if len(username) > 20 {
		return errors.New("username is too long (max: 10 characters)")
	} else if len(username) < 3 {
		return errors.New("username is too short (min: 3 characters)")
	}
	return nil
}

func PasswordValidator(password string) error {
	if len(password) > 20 {
		return errors.New("password is too long (max: 20 characters)")
	} else if len(password) < 6 {
		return errors.New("password is too short (min: 5 characters)")
	}
	return nil
}

func EmailValidator(email string) error {
	if len(email) > 100 {
		return errors.New("email is too long (max: 100 characters)")
	}
	var str string
	cont := 0
	for i := len(email) - 1; i >= 0; i-- {
		if email[i] == '@' {
			str = email[i:]
			cont++
		}
	}
	if cont != 1 {
		return errors.New("email should contain only one '@' symbol")
	}
	if strings.Contains(str, "@") && strings.Contains(str, ".") && (strings.Index(str, "@") < strings.Index(str, ".")) {
		if strings.Index(str, "@")+1 == strings.Index(str, ".") {
			return errors.New("email should not contain consecutive '@' or '.' symbols")
		}
		if str[len(str)-1] == '.' {
			return errors.New("email should not end with '.' symbol")
		}
		return nil
	} else {
		return errors.New("email is not valid")
	}
}

func Login_Validat(email, password string) error {
	err := EmailValidator(email)
	if utils.IsErrors(err) {
		return err
	}
	err = PasswordValidator(password)
	if utils.IsErrors(err) {
		return err
	}
	return nil
}
