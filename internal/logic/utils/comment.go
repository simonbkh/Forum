package utils

func Check_Comment(comment string) {
	count := 0
	for _, v := range comment {
		if v == '\n' {
			continue
		}
		count++
	}

	if comment == "" || count > 1000 {
		return
	}
}
