package forum

import (
	"fmt"
	t "forum/users"
	"net/http"
	"time"
)

type cookies struct {
	Value string
}

var COOKIES cookies

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.
	cookie := http.Cookie{
		Name:       "session",
		Value:      t.USER.Username,
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     99999999999,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   http.SameSiteLaxMode,
		Raw:        "",
		Unparsed:   []string{},
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	fmt.Println("cookie set!")
}
