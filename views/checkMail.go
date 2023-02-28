package forum

import (
	"fmt"
	"net/mail"
)

func CheckMail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
