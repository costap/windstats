package webapp

import (
	"github.com/costap/windstats/internal/config"
	"github.com/ory/viper"
)

// Config ...
type Config struct {
	ListeningPort int    `mapstructure:"listeningport"`
	DataRootPath  string `mapstructure:"datarootpath"`
}

// ReadConfig ...
func ReadConfig() *Config {
	viper.SetDefault("datarootpath", "cmd/windstatsweb/kodata")
	var C Config
	config.MustReadConfig("WEBAPP", &C)
	return &C
}
