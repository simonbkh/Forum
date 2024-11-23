package validators

import "errors"

func Username_Validator(username string) error{
	if len(username) > 10 {
		return errors.New("long username")
	}else if len(username) < 3 {
		return errors.New("short username")
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

func Email_Validator(email string) error{
	if len(email) > 20 {
		return errors.New("long password")
	}

	return nil
}