package handlers

import (
	"html/template"
	"net/http"
	"inventory-app/database"
	"inventory-app/models"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("assets/login.html")
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user models.User
	database.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
    http.Redirect(w, r, "/login?error=Invalid+username+or+password", http.StatusSeeOther)
    return
	}

	session, _ := store.Get(r, "session")
	session.Values["user_id"] = user.ID
	session.Save(r, w)
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)

	if err := r.ParseForm(); err != nil {
    http.Error(w, "Invalid form submission", http.StatusBadRequest)
    return
	}

	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}
	session.Values["user_id"] = user.ID
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Could not save session", http.StatusInternalServerError)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "user_id")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
