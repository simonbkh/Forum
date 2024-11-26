package validators

import (
	"errors"
	"net/http"
	"forum/internal/data/queries"
)

func CategoriesValidator(categories []string) error {
	istruecat := false
	TrueCategories := []string{"general", "games", "sports", "fashion", "travel", "food", "health"}
	for _, category := range categories {
		for _, truecat := range TrueCategories {
			if truecat == category {
				istruecat = true
			}
		}
	}
	if !istruecat {
		return errors.New("invalid category! ")
	}
	return nil
}
func Allowed(r *http.Request, ) error {
	cookie, err := r.Cookie("token")
	if err != nil || !queries.Logged(cookie.Value) {
		// redirect awla chi laeba
		return errors.New("you are not logged in! ")
	}

	return nil
}

func TitleValidator(title string) error {

	return nil
}
