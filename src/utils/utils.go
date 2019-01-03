package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

type Configuration struct {
	Address      string `yaml:"Address"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
	Static       string `yaml:"Static"`
}

func p(x ...interface{}) {
	fmt.Print(x)
}

func LoadConfig(path string) (Configuration, error) {
	var conf Configuration
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}
	if err = yaml.Unmarshal(data, &conf); err != nil {
		return conf, err
	}

	return conf, nil
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

func session(w http.ResponseWriter, r *http.Request) (Session, error) {

}
