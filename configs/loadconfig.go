package configs

import (
	"encoding/json"
	"io"
	"os"

	"starter_go/configs/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

var ENV *models.Server

func InitializeConfig() {
	confFile, err := os.Open("appconfig.json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	conf, err := io.ReadAll(confFile)
	if err != nil {
		panic(err)
	}
	appconfig := models.Configuration{}
	err = json.Unmarshal(conf, &appconfig)
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
