package validators

import (
	"errors"
	"strings"
)

func CategoriesValidator(categories []string) error {
	if len(categories) == 0 || len(categories) > 7 {
		return errors.New("invalid category! ")
	}

	TrueCategories := []string{"Tech Support", "General Discussion", "Tutorials", "Announcements", "Gaming", "Job Listings", "Hobbies & Interests"}
	for _, category := range categories {
		bol := false
		for _, truecat := range TrueCategories {
			if strings.EqualFold(category, truecat) {
				bol = true
			}
		}
		if !bol {
			return errors.New("invalid category! ")
		}

	}

	return nil
}

func TitleValidator(title string) error {
	if len(strings.TrimSpace(title)) > 40 {
		return errors.New("title is too long")
	} else if len(strings.TrimSpace(title)) < 4 {
		return errors.New("title is too short")
	}
	return nil
}

func ValidContent(content string) error {
	if len(strings.TrimSpace(content)) > 1000 {
		return errors.New("content is too long")
	} else if len(strings.TrimSpace(content)) < 1 {
		return errors.New("content is too short")
	}
	return nil
}
