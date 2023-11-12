package config

import (
	"gopkg.in/yaml.v2"
)

const (
	CONFIG = `
local:
  log:
    dir: /Users/apple/workspace/lgo-go/log/
    filename: my.log
    logName: myapp
    level: debug
`
)

type PhaseConfig struct {
	Local AppConf `yaml:"local"`
	Dev   AppConf `yaml:"dev"`
	Real  AppConf `yaml:"real"`
}

type LogConfig struct {
	Dir      string `yaml:"dir"`
	Filename string `yaml:"filename"`
	LogName  string `yaml:"logName"`
	Level    string `yaml:"level"`
}

type AppConf struct {
	DB  DBConfig  `yaml:"db"`
	Log LogConfig `yaml:"log"`
}

type DBConfig struct {
	Server   string `yaml:"server"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

var conf *AppConf

func LoadConfig() error {
	config := PhaseConfig{}

	err := yaml.Unmarshal([]byte(CONFIG), &config)
	if err != nil {
		return err
	}

	// load configuration
	var appConf AppConf
	appConf = config.Local
	conf = &appConf
	return nil
}

func Config() *AppConf {
	return conf
}
