package validators

import (
	"errors"
	"fmt"
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

func Allowed(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		// redirect awla chi laeba
		return 0, err
	}
	fmt.Println(cookie)
	user_id, err := queries.Logged(cookie.Value)
	fmt.Println(user_id)
	if err != nil {
		// redirect awla chi laeba
		return 0, err
	}

	return user_id, nil
}

func TitleValidator(title string) error {
	return nil
}
