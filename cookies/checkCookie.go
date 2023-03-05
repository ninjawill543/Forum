package forum

import (
	"fmt"
	"net/http"
)

func CheckCookie(r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("cookie doesnt exist")
	} else {
		COOKIES.Value = cookie.Value
	}
}
