package validators

import (
	"errors"
	"net/http"

	"forum/internal/data/queries"
)

func CategoriesValidator(categories []string) error {
	istruecat := false
	TrueCategories := []string{"general", "games", "sports", "fashion", "travel", "food", "health", "all"}
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
	user_id, err := queries.Logged(cookie.Value)
	if err != nil {
		// redirect awla chi laeba
		return 0, err
	}

	return user_id, nil
}
func TitleValidator(title,content string) error{
	if len(title) > 15 {
		return errors.New("title is too long")
	}else if len(title) < 4 {
		return errors.New("title is too short")
	}
	if len(content) > 25 {
		return errors.New("content is too long")
	}else if len(content) < 4 {
		return errors.New("content is too short")
	}
	return nil
}
