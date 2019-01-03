package routes

import (
	"fmt"
	"github.com/xdhuxc/chitchat/src/models"
	"github.com/xdhuxc/chitchat/src/utils"
	"net/http"
)

/**
Show the new thread form page
GET /threads/new
*/
func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := utils.Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		utils.GenerateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

/**
Create the user account
POST /signup
*/
func CreateThread(w http.ResponseWriter, r *http.Request) {
	session, err := utils.Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			utils.Danger(err, "Can not parse form.")
		}
		user, err := session.User()
		if err != nil {
			utils.Danger(err, "Can not get user from session.")
		}
		topic := r.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			utils.Danger(err, "Can not create thread.")
		}
		http.Redirect(w, r, "/", 302)
	}
}

/**
Show the details of the thread, including the posts and the form to write a post
*/
func ReadThread(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	uuid := values.Get("id")
	thread, err := models.FindThreadByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(w, r, "Can not read thread")
	} else {
		_, err := utils.Session(w, r)
		if err != nil {
			utils.GenerateHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

func PostThread(w http.ResponseWriter, r *http.Request) {
	session, err := utils.Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			utils.Danger(err, "Can not parse form.")
		}
		user, err := session.User()
		if err != nil {
			utils.Danger(err, "Can not get user from session.")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := models.FindThreadByUUID(uuid)
		if err != nil {
			utils.ErrorMessage(w, r, "Can not read thread.")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			utils.Danger(err, "Can not create post.")
		}
		url := fmt.Sprintf("/thread/read?id=%s", uuid)
		http.Redirect(w, r, url, 302)
	}
}
