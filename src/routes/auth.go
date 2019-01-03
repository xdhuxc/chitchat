package routes

import (
	"github.com/xdhuxc/chitchat/src/models"
	"github.com/xdhuxc/chitchat/src/utils"
	"net/http"
)

/**
Show the login page
GET /login
*/
func Login(w http.ResponseWriter, r *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

/**
Show the sign up page
GET /signup
*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	utils.GenerateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

/**
Create the user account
POST /signup
*/
func SignUpAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		utils.Danger(err, "Parse form error.")
	}
	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		utils.Danger(err, "Create user error.")
	}
	http.Redirect(w, r, "/login", 302)
}

/**
Authenticate user by the given email and password
POST /authenticate
*/
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.FindUserByEmail(r.PostFormValue("email"))
	if err != nil {
		utils.Danger(err, "Can not find user.")
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			utils.Danger(err, "Can not create session.")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

/**
Logs the user out
GET /logout
*/
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.Warning(err, "Failed to get cookie.")
		session := models.Session{UUID: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
