package forum

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	cookie, err := r.Cookie("session")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			fmt.Println(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			fmt.Println(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body.
	// w.Write([]byte(cookie.Value))
	fmt.Println(cookie.Value)
}
