package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"password-go/utils"
)

type ConfigYaml struct {
	Gin struct {
		Mode string `yaml:"mode"`
	} `yaml:"gin"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"db"`
	} `yaml:"database"`
}

var fileUtil = utils.FileUtil{}

func Config() ConfigYaml {
	var config ConfigYaml
	configData := fileUtil.ReadFile("config.yaml")
	err := yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}
	return config
}
