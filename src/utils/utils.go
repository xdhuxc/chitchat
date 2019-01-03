package utils

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xdhuxc/chitchat/src/models"
	"html/template"
	"net/http"
	"strings"
)

func P(x ...interface{}) {
	fmt.Print(x)
}

func StringInSlice(s string, slice []string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}

/**
to redirect to the error message page
*/
func ErrorMessage(w http.ResponseWriter, r *http.Request, message string) {
	url := []string{"/error?message=", message}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

/**
Checks if the user is logged in and has a session
*/
func Session(w http.ResponseWriter, r *http.Request) (models.Session, error) {
	var s models.Session
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return s, err
	}
	s = models.Session{UUID: cookie.Value}
	if ok, _ := s.Check(); !ok {
		return s, errors.New("Invalid session. ")
	}
	return s, nil
}

/**
Parse HTML templates
*/
func ParseTemplateFiles(filenames ...string) *template.Template {
	var files []string
	var t *template.Template
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return t
}

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func Info(args ...interface{}) {
	logrus.Infoln(args...)
}

func Danger(args ...interface{}) {
	logrus.Error(args...)
}

func Warning(args ...interface{}) {
	logrus.Warning(args...)
}

/**
the version information
*/
func Version() string {
	return "0.0.1"
}
