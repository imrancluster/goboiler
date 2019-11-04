package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	// importing for remote connection
	_ "github.com/spf13/viper/remote"
)

var muLock sync.Mutex

// Init : global init function that loads env, db and app settings
func Init() {
	viper.SetEnvPrefix("goboiler")
	viper.BindEnv("env")
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	consulURL := viper.GetString("consul_url")
	if consulURL == "" {
		log.Fatal("CONSUL_URL missing")
	}

	consulPath := viper.GetString("consul_path")
	if consulPath == "" {
		log.Fatal("CONSUL_PATH missing")
	}

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("yml")

	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatal(fmt.Sprintf("%s occured at %s, url: %s", err.Error(), consulPath, consulURL))
	}

	LoadDB()
	LoadRedis()
	LoadApp()
	LoadAws()
}
