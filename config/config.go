package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Databases []DatabaseSetting
	Server    ServerSetting
	App       ApplicationSetting
}

type DatabaseSetting struct {
	DbName      string
	Url         string
	Collections *map[string]string
}
type ServerSetting struct {
	Services []ServiceSetting
}

type ServiceSetting struct {
	Name string
	Port string
}

type ApplicationSetting struct {
	Name string
	Jwt  JwtSetting
}

type JwtSetting struct {
	Key string
	Exp int
}

func ReadConfiguration() Configuration {

	//Get ENVIRONMENT variable
	environment := os.Getenv("ENVIRONMENT_MY_GO")

	if environment == "" {
		environment = "development"
	}
	fmt.Println("ENVIRONMENT:", environment)

	var config Configuration
	configpath := "/config."

	// Check OS for setting work path
	if os.Getenv("OS") == "Windows_NT" {
		configpath = "\\config."
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error Getwd: %v", err)
	}
	yamlFile, err := ioutil.ReadFile(cwd + configpath + environment + ".yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err.Error())
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("Error unmarshaling YAML data: %v", err)
	}

	return config
}
