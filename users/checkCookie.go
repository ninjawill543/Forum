package forum

import (
	"fmt"
	"net/http"
)

func CheckCookie(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("cookie doesnt exist")
	} else {
		if COOKIES.UuidUser != "" {
			COOKIES.UuidUser = cookie.Value
			return true
		}
	}
	return false
}
