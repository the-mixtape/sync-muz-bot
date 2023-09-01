package util

import "github.com/spf13/viper"

type Config struct {
	Token       string `mapstructure:"TOKEN"`
	Socks5Proxy string `mapstructure:"SOCKS5_PROXY"`
	Debug       bool   `mapstructure:"DEBUG"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("bot")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
