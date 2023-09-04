package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotToken string `mapstructure:"BOT_TOKEN"`
	BotDebug bool   `mapstructure:"BOT_DEBUG"`

	CacheDefaultExpiration int `mapstructure:"CACHE_DEFAULT_EXPIRATION"`
	CacheCleanupInterval   int `mapstructure:"CACHE_CLEANUP_INTERVAL"`

	Socks5Proxy string `mapstructure:"SOCKS5_PROXY"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLM"`

	VkApiToken string `mapstructure:"VKAPI_TOKEN"`
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
