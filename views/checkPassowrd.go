package forum

import (
	"fmt"
	"regexp"
	"unicode"
)

func CheckPassword(testpassword string) bool {
	oui := 0
	oui2 := 0
	oui3 := 0

	var length = len([]rune(testpassword))
	for _, rune := range testpassword {
		if unicode.IsUpper(rune) {
			oui = 1
			break
		} else {
			oui = 0
		}
	}
	if oui == 0 {
		fmt.Println("no maj")
		return false
	}

	for _, rune2 := range testpassword {
		if unicode.IsLower(rune2) {
			oui2 = 1
			break
		} else {
			oui2 = 0
		}
	}
	if oui2 == 0 {
		fmt.Println("no min")
		return false
	}

	for _, char := range testpassword {
		if unicode.IsNumber(char) {
			oui3 = 1
			break
		} else {
			oui3 = 0
		}
	}
	if oui3 == 0 {
		fmt.Println("no numb")
		return false
	}

	for _, rune2 := range testpassword {
		if unicode.IsLower(rune2) {
			oui2 = 1
			break
		} else {
			oui2 = 0
		}
	}
	if oui2 == 0 {
		fmt.Println("no min")
		return false
	}

	is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(testpassword)

	if is_alphanumeric {
		fmt.Println("no special char")
		return false
	}

	if length < 8 {
		fmt.Println("2 short")
		return false
	}
	return true
}
