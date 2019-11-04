package config

import "github.com/spf13/viper"

// Database : all necessary credential for database
type Database struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

var db Database

// DB : exportable database connector
func DB() *Database {
	return &db
}

// LoadDB : setup the db connection to connect db to the application
func LoadDB() {
	db = Database{
		Host:         viper.GetString("db.host"),
		Port:         viper.GetInt("db.port"),
		Username:     viper.GetString("db.username"),
		Password:     viper.GetString("db.password"),
		DatabaseName: viper.GetString("db.name"),
	}
}
