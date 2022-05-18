package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port         int               `json:"port"`
	DocumentRoot string            `json:"documentRoot"`
	Route        map[string]string `json:"route"`
}

func MakeConfig(confFile string) (*Config, error) {
	viper.SetConfigFile(confFile)
	viper.SetDefault("Port", 80)
	viper.SetDefault("MaxPackageSize", 100)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	//如果需要对config内容进行校验，在这里编写校验逻辑
	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
