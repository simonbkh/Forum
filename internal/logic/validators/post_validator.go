package validators

import (
	"errors"
	"fmt"
	"forum/internal/data/queries"
	"net/http"
	"strings"
)

func CategoriesValidator(categories []string) error {
	istruecat := false
	fmt.Println(categories)
	if len(categories) == 0 || len(categories) > 7 {
		return errors.New("invalid category! ")
	}

	TrueCategories := []string{"Tech Support", "General Discussion", "Tutorials", "Announcements", "Gaming", "Job Listings", "Hobbies & Interests"}
	for _, category := range categories {
		bool := false
		for _, truecat := range TrueCategories {
			if strings.EqualFold(category, truecat) {
				bool = true
				istruecat = true
			}
		}
		if !bool {
			return errors.New("invalid category! ")
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
