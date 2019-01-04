package models

type Configuration struct {
	Server Server `yaml:"Server"`
}

type Server struct {
	Address      string `yaml:"Address"`
	ReadTimeout  int64  `yaml:"ReadTimeout"`
	WriteTimeout int64  `yaml:"WriteTimeout"`
	Static       string `yaml:"Static"`
}
