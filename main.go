package main

import (
	"fmt"
	"go-session/user"
	"html/template"
	"net/http"
)

var tpl *template.Template

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		usr := user.New()
		usr.FirstName = r.FormValue("firstname")
		usr.LastName = r.FormValue("lastname")
		usr.Email = r.FormValue("email")
		usr.Password = r.FormValue("password")
		usr.CreateUser()
		fmt.Println(usr)
		fmt.Println(usr.GetUserId())
		http.Redirect(w, r, "/create-user", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "createuser.html", nil)

}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/create-user", CreateUser)
}

func main() {
	http.ListenAndServe("0.0.0.0:8000", nil)
}
