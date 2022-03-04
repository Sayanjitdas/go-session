package main

import (
	"fmt"
	"go-session/session"
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
		usr.SetPassword(r.FormValue("password"))
		err := usr.CreateUser()
		if err != nil {
			http.Redirect(w, r, "/create-user?error="+err.Error(), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if len(r.FormValue("error")) > 0 {
		tpl.ExecuteTemplate(w, "createuser.html", r.FormValue("error"))
		return
	}
	tpl.ExecuteTemplate(w, "createuser.html", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		usr, err := user.GetUserByEmail(r.FormValue("username"))
		if err != nil {
			http.Redirect(w, r, "/login?error="+err.Error(), http.StatusSeeOther)
			return
		}
		if !usr.CheckPassword(r.FormValue("password")) {
			http.Redirect(w, r, "/login?error=invalid password", http.StatusSeeOther)
			return
		}
		session.SetSession(&w, usr.GetUserId())
		http.Redirect(w, r, "/success", http.StatusSeeOther)
		return
	}
	if len(r.FormValue("error")) > 0 {
		tpl.ExecuteTemplate(w, "login.html", r.FormValue("error"))
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.DeleteSession(&w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Success(w http.ResponseWriter, r *http.Request) {
	if session.HasSession(r) {
		sid, _ := r.Cookie("session")
		uid, _ := session.GetSessionUser(sid.Value)
		usr, _ := user.GetSpecificUser(uid)
		tpl.ExecuteTemplate(w, "success.html", usr.Email)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/create-user", CreateUser)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/success", Success)
	http.HandleFunc("/logout", Logout)
}

func main() {
	fmt.Println("RUNNING SERVER ON PORT 8000...")
	http.ListenAndServe("0.0.0.0:8000", nil)
}
