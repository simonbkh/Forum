package validators

import (
	"errors"
	"forum/internal/data/queries"
	"net/http"
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
func Allowed(r *http.Request) (error, int) {
	cookie, err := r.Cookie("token")
	user_id := queries.Logged(cookie.Value)
	if err != nil || user_id == 0 {
		// redirect awla chi laeba
		return errors.New("you are not logged in! ") , 0
	}

	return nil, user_id
}

func TitleValidator(title string) error {

	return nil
}
