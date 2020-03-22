package main

import (
	"github.com/costap/windstats/internal/config"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	DBAdrr string `mapstructure:"dbaddr"`
	DBUser string `mapstructure:"dbuser"`
	DBName string `mapstructure:"dbname"`
	DBPass string `mapstructure:"dbpass"`
}

// ReadConfig ...
func ReadConfig() *Config {
	viper.SetDefault("dbaddr", "http://influxdb:8086")
	viper.SetDefault("dbname", "windstats")
	viper.SetDefault("dbuser", "windstats")
	viper.SetDefault("dbpass", "windstats")
	var C Config
	config.MustReadConfig("WINDSTATS", &C)
	C.DBAdrr = viper.GetString("dbaddr")
	C.DBPass = viper.GetString("dbpass")
	return &C
}
