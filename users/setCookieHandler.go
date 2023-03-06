package forum

import (
	"fmt"
	"net/http"
)

type cookies struct {
	UuidUser string
}

var COOKIES cookies

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  USER.Uuid,
		Path:   "/",
		Domain: "",
		// Expires:    time.Time{},
		// RawExpires: "",
		MaxAge:   99999999999,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Raw:      "",
		Unparsed: []string{},
	}

	http.SetCookie(w, &cookie)
	fmt.Println("cookie set")
}
