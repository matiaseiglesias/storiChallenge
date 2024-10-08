package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIGPATH = "config.yml"

type Config struct {
	Server          Server
	DataBase        Database
	SmtpServer      Smtpserver
	EmailTemplate   EmailTemplate
	TransactionFile TransactionFile
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Schema string `yaml:"schema"`
	UserID string `yaml:"userID"`
	Pass   string `yaml:"pass"`
}

type Smtpserver struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

type EmailTemplate struct {
	Path string `yaml:"path"`
}

type TransactionFile struct {
	Directory string `yaml:"directory"`
}

func LoadConfig() *Config {
	// Read the file
	data, err := os.ReadFile(CONFIGPATH)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	// Create a struct to hold the YAML data
	var config Config

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	// Print the data
	fmt.Println(config)
	return &config
}
