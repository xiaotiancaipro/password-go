package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"password-go/utils"
)

type ConfigYaml struct {
	SecretKey string `yaml:"secret_key"`
	Gin       struct {
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

var log = utils.LoggerUtil{}.Logger()
var fileUtil = utils.FileUtil{}

func Config() (config ConfigYaml) {
	configData := fileUtil.ReadFile("config.yaml")
	err := yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Error(fmt.Sprintf("Error parsing config file: %v", err))
		os.Exit(1)
	}
	return
}
