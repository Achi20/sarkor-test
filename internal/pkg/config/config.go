package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Conf struct {
	BaseUrl string `yaml:"base_url"`
	IP      string `yaml:"ip"`
	Port    string `yaml:"port"`
	JWTKey  string `yaml:"jwt_key"`
}

func GetConf() (cfg *Conf) {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Fatalf("yamlFile: %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return
}
