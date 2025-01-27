package utils

import (
	"fmt"
	"os"
	"strconv"
)

func Convstr(s string) int {
	fmt.Println(s)
	nb, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return nb
}
