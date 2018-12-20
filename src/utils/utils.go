package utils

import (
	"log"
)

type Configuration struct {
	Address      string `yaml:"Address"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
	Static       string `yaml:"Static"`
}

var conf Configuration
var logger *log.Logger

func StringInSlice(s string, slice []string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}
