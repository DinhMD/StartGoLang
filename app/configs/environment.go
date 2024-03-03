package configs

import (
	"io"
	"os"

	"starter_go/app/configs/models"

	yaml "gopkg.in/yaml.v3"

	"gorm.io/gorm"
)

var DB *gorm.DB

var ENV *models.Server

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
	appconfig := models.Configuration{}
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
