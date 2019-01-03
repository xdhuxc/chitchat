package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetLogin(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/login", Login)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/login", nil)
	mux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
	body := w.Body.String()
	if strings.Contains(body, "Sign In") == false {
		t.Errorf("Body does not contain Sing In")
	}
}
