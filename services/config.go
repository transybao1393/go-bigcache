package services

import "github.com/spf13/viper"

type Config struct {
	Shards      string `mapstructure:"Shards"`
	LifeWindow  string `mapstructure:"LifeWindow"`
	CleanWindow string `mapstructure:"CleanWindow"`
	Verbose     string `mapstructure:"Verbose"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("staging")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
