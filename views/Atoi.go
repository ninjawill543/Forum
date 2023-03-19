package forum

import (
	"fmt"
	"strconv"
)

func Atoi(date string) int {
	marks, err := strconv.Atoi(date)
	if err != nil {
		fmt.Println(err)
	}
	return marks
}
