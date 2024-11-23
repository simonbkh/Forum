package validators

import "errors"

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

func EmailValidator(email string) error {
	if len(email) > 25 {
		return errors.New("email is too long (max: 25 characters)")
	} else if len(email) < 13 {
		return errors.New("email is too short (min: 13 characters)")
	}
	return nil
}
