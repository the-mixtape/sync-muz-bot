package main

import (
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"net/url"
	"sync-muz-bot/pkg/cache/local_cache"
	"sync-muz-bot/pkg/repository"
	"sync-muz-bot/pkg/repository/postgres"
	"sync-muz-bot/pkg/telegram"
	"sync-muz-bot/pkg/util"
	"time"
)

const (
	cacheDefaultExpiration = 5 * time.Minute
	cacheCleanupInterval   = 10 * time.Minute
)

func main() {
	botConfig := readConfig()
	bot := createTgBotApi(botConfig.Token, botConfig.Socks5Proxy, botConfig.Debug)

	botCache := local_cache.NewBotCache(cacheDefaultExpiration, cacheCleanupInterval)

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     botConfig.DBHost,
		Port:     botConfig.DBPort,
		Username: botConfig.DBUsername,
		Password: botConfig.DBPassword,
		DBName:   botConfig.DBName,
		SSLMode:  botConfig.DBSSLMode,
	})
	if err != nil {
		slog.Error(err.Error())
	}

	repos := repository.NewRepository(db)

	telegramBot := telegram.NewBot(bot, botCache, repos)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func readConfig() util.Config {
	configPath := flag.String("config", "./configs", "bot.env config path")
	flag.Parse()

	botConfig, err := util.LoadConfig(*configPath)
	if err != nil {
		slog.Error("cannot load config:" + err.Error())
	}

	return botConfig
}

func createTgBotApi(token string, socks5 string, debug bool) *tgbotapi.BotAPI {
	client := &http.Client{}
	if len(socks5) > 0 {
		tgProxyURL, err := url.Parse(socks5)
		if err != nil {
			slog.Error("Failed to parse proxy URL: " + err.Error())
		}
		tgTransport := &http.Transport{Proxy: http.ProxyURL(tgProxyURL)}
		client.Transport = tgTransport
	}

	bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, client)
	if err != nil {
		slog.Error(err.Error())
	}
	bot.Debug = debug
	return bot
}
