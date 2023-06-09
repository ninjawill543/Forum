package forum

import (
    "fmt"
    "net/http"
)

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
    // Retrieve the cookie from the request using its name (which in our case is
    // "exampleCookie"). If no matching cookie is found, this will return a
    // http.ErrNoCookie error. We check for this, and return a 400 Bad Request
    // response to the client.
    cookie, err := r.Cookie("session")
    if err != nil {
        fmt.Println(err)
        USER.Username = ""
        USER.Uuid = ""
        USER.Email = ""
        USER.CreationDate = ""
        USER.Admin = 0
        USER.BirthDate = ""
    } else {
        if cookie.Value != "" {
            COOKIES.UuidUser = cookie.Value
            LoginWithCookie(cookie.Value)
        } else {
            USER.Username = ""
            USER.Uuid = ""
            USER.Email = ""
            USER.CreationDate = ""
            USER.Admin = 0
            USER.BirthDate = ""
            USER.Uuid = ""    
        }
    }
}