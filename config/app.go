package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Version : Goboiler Version
var Version = "1.0.0"

// Env : Goboiler default server env
var Env = "prod"

// Application : application envs type
type Application struct {
	Env            string
	Port           int
	Sentry         string
	Version        string
	ReadTimeOut    time.Duration
	RequestTimeOut time.Duration
	AppKey         string
	UserToken      string
}

var app Application

// App : exportable function for app env
func App() *Application {
	return &app
}

// LoadApp : setup the application env to run the application
func LoadApp() {
	muLock.Lock()
	defer muLock.Unlock()

	serverEnv := Env
	if e := viper.GetString("app.env"); e != "" {
		serverEnv = e
	}

	version := Version
	if v := viper.GetString("app.version"); v != "" {
		version = v
	}

	app = Application{
		Env:            serverEnv,
		Port:           viper.GetInt("app.port"),
		Sentry:         viper.GetString("app.sentry"),
		Version:        version,
		ReadTimeOut:    viper.GetDuration("app.read_timeout") * time.Second,
		RequestTimeOut: viper.GetDuration("app.request_timeout") * time.Second,
		AppKey:         viper.GetString("app.app_key"),
		UserToken:      viper.GetString("app.user_token"),
	}

	// setup logs
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
