package forum

import (
	"fmt"
	"net/http"
)

type emailStorage struct {
	email string
}

var EMAILSTORAGE emailStorage

func EmailStorage(r *http.Request) {
	email := r.FormValue("input_mail")
	EMAILSTORAGE.email = email
	fmt.Println(EMAILSTORAGE.email)
}
