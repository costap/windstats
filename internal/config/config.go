package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func MustReadConfig(envPrevix string, c interface{}) {
	viper.SetDefault("configPath", ".")

	viper.SetEnvPrefix(envPrevix)
	viper.AutomaticEnv()

	if ecp := viper.GetString("CONFIG_PATH"); ecp != "" {
		viper.Set("configPath", ecp)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(viper.GetString("configPath"))
	if err := viper.ReadInConfig(); err != nil {
		if cfnferr, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("%v\n", cfnferr.Error())
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	err := viper.Unmarshal(c)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}
