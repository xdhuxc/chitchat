package routes

import (
	"github.com/xdhuxc/chitchat/src/models"
	"github.com/xdhuxc/chitchat/src/utils"
	"net/http"
)

/**
Shows the error message page
GET /error?message=
*/
func Error(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	_, err := utils.Session(w, r)
	if err != nil {
		utils.GenerateHTML(w, values.Get("message"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(w, values.Get("message"), "layout", "private.navbar", "error")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err != nil {
		utils.ErrorMessage(w, r, "Can not get threads.")
	} else {
		_, err := utils.Session(w, r)
		if err != nil {
			utils.GenerateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}
