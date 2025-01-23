package validators

import (
	"errors"
	"forum/internal/data/queries"
	"net/http"
)

func CategoriesValidator(categories []string) error {
	istruecat := false
	if len(categories) == 0 {
		return errors.New("maximum categories to choose is 3")
	} else if len(categories) > 3 {
		return errors.New("maximum categories to choose is 3")
	}
	TrueCategories := []string{"Tech Support", "General Discussion", "Tutorials", "Announcements","Gaming","Job Listings","Hobbies & Interests"}
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
	cookie, err := r.Cookie("SessionToken")
	if err != nil {
		// redirect awla chi laeba
		return 0, err
	}
	user_id, err := queries.Hh(cookie.Value)
	if err != nil {
		// redirect awla chi laeba
		return 0, err
	}

	return user_id, nil
}

func TitleValidator(title string) error {

	if len(title) > 40 {
		return errors.New("title is too long")
	} else if len(title) < 4 {
		return errors.New("title is too short")
	}
	return nil
}
