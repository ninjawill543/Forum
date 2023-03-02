package forum

import (
	"net/http"
)

func Logout(r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("logOutButton") == "logout" {
			USER.BirthDate = ""
			USER.CreationDate = ""
			USER.Email = ""
			USER.Username = ""
			USER.Uuid = ""
		}
	}
}
