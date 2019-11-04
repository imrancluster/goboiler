package config

import (
	"time"

	"github.com/spf13/viper"
)

// RedisConf redis struct
type RedisConf struct {
	Host        string
	Password    string
	DB          int
	Prefix      string
	Port        int
	PinResetTTL time.Duration
}

var redis RedisConf

// Redis exportable function
func Redis() *RedisConf {
	return &redis
}

// LoadRedis loading redis connection
func LoadRedis() {
	redis = RedisConf{
		Host:        viper.GetString("redis.host"),
		Port:        viper.GetInt("redis.port"),
		Password:    viper.GetString("redis.password"),
		DB:          viper.GetInt("redis.db"),
		Prefix:      viper.GetString("redis.prefix"),
		PinResetTTL: viper.GetDuration("app.request_timeout") * time.Second,
	}
}
