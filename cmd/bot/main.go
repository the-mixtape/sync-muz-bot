package main

import (
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"net/url"
	"sync-muz-bot/pkg/cache/local_cache"
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

	bot.Debug = true
	telegramBot := telegram.NewBot(bot, botCache)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func readConfig() util.Config {
	configPath := flag.String("config", "./configs", "bot.env config path")
	flag.Parse()

	botConfig, err := util.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return botConfig
}

func createTgBotApi(token string, socks5 string, debug bool) *tgbotapi.BotAPI {
	client := &http.Client{}
	if len(socks5) > 0 {
		tgProxyURL, err := url.Parse(socks5)
		if err != nil {
			log.Printf("Failed to parse proxy URL:%s\n", err)
		}
		tgTransport := &http.Transport{Proxy: http.ProxyURL(tgProxyURL)}
		client.Transport = tgTransport
	}

	bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = debug
	return bot
}
