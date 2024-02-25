package models

type Configuration struct {
	Server   Server `json:"server"`
	Database DB     `json:"database"`
}

type DB struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type Server struct {
	Port int
	Host string
}
