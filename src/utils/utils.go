package utils

import (
	"fmt"
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

func p(s ...interface{}) {
	fmt.Println(s)
}

func LoadConfig() {

}
