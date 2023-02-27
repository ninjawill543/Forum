package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", Handler_index)

	//url of our funcs
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	//listening on port 8080
	http.ListenAndServe(":8080", nil)
}

func Handler_index(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/index.html"))

	if r.Method == "POST" {
		username := r.FormValue("input_username")
		password := r.FormValue("input_password")

		fmt.Print(username, "\n")
		fmt.Print(password, "\n")

	}

	// db, err := sql.Open("sqlite3", ":forum:")

	// if err != nil {
	// 	log.Fatal(err)
	// } else {

	// }

	// defer db.Close()

	tmpl1.Execute(w, "babygirl")
}
