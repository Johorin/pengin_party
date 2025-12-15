package config

import (
	"fmt"
	"os"
)

type db struct {
	host     string
	database string
	user     string
	password string
	port     string
}

var DB db

func LoadDB() {
	DB = db{
		host:     os.Getenv("DB_HOST"),
		database: os.Getenv("DB_DATABASE"),
		user:     os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		port:     os.Getenv("DB_PORT"),
	}
}

func (d *db) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", d.user, d.password, d.host, d.port, d.database)
}