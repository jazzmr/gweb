package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var conf AppConfig

type AppConfig struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
}

func init() {
	var err error
	configFile, err := ioutil.ReadFile("conf/application.yaml")

	if err != nil {
		log.Fatalf("init application.yaml failed, err is %v", err)
	}

	err = yaml.Unmarshal(configFile, &conf)
}

func GetConfig() *AppConfig {
	return &conf
}

func GetContextPath() string {
	return conf.Server.ContextPath
}
