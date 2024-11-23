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

// func Password_Validator(password string) error{
// 	// if len(password) > 20 {
// 	// 	return errors.New("long password")
// 	// }else if len(password) < 3 {
// 	// 	return errors.New("short password")
// 	// }

// 	// return nil
// }

func Email_Validator(email string) error {
	if len(email) > 20 {
		return errors.New("long password")
	}

	return nil
}
