package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configFile []byte

type AppConfig struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("conf/application.yaml")

	if err != nil {
		log.Fatalf("init application.yaml failed, err is %v", err)
	}
}

func GetConfig() (a AppConfig, err error) {
	err = yaml.Unmarshal(configFile, &a)
	return a, err
}
