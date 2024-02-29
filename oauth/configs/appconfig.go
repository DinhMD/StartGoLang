package configs

import (
	"encoding/json"
	"io"
	"os"
)

var ENV *Server

type Configuration struct {
	Server Server
}

type Server struct {
	Host string
	Port int
}

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

	appconfig := Configuration{}
	err = json.Unmarshal(conf, &appconfig)
	if err != nil {
		panic(err)
	}
	ENV = &appconfig.Server
}
