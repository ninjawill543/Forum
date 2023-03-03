package forum

import (
	"fmt"
	"net/http"
)

func Logout(r *http.Request) {
	fmt.Println("logout")
	USER.BirthDate = ""
	USER.CreationDate = ""
	USER.Email = ""
	USER.Username = ""
	USER.Uuid = ""
}
