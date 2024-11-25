package validators

import "errors"

func CategoriesValidator(categories []string) error {
	istruecat := false
	TrueCategories := []string{"general", "games", "sports", "fashion", "travel", "food", "health"}
	for _,category := range categories {
		for _, truecat :=range TrueCategories {
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

func TitleValidator(title string) error{
	



	return nil
}