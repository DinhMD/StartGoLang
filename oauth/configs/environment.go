package configs

import (
	"io"
	"os"

	yaml "gopkg.in/yaml.v3"

	"gorm.io/gorm"
)

type Configuration struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type Server struct {
	Port      int    `json:"port"`
	Host      string `json:"host"`
	SecretKey string `json:"secret_key"`
}

var DB *gorm.DB

var ENV *Server

func InitializeConfig() {
	confFile, err := os.Open("application_config.yaml")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	conf, err := io.ReadAll(confFile)
	if err != nil {
		panic(err)
	}
	appconfig := Configuration{}
	err = yaml.Unmarshal(conf, &appconfig)
	if err != nil {
		panic(err)
	}
	db, err := Connect(appconfig.Database)
	if err != nil {
		panic(err)
	}
	DB = db
	ENV = &appconfig.Server
}
