package forum

import (
	"net/http"
)

func Logout(r *http.Request) {
	USER.BirthDate = ""
	USER.CreationDate = ""
	USER.Email = ""
	USER.Username = ""
	USER.Uuid = ""
	USER.Admin = 0
}
