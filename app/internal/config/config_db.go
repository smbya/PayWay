package config

import "os"

type db struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
}

func createDb() *db {
	return &db{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
	}
}

func (d *db) GetConnectString() string {
	return "postgres://" + d.user + ":" + d.password + "@" + d.host + ":" + d.port + "/" + d.dbname + "?sslmode=disable"
}
