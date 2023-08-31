package main

import (
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
	"net/url"
	"sync-muz-bot/pkg/telegram"
	"sync-muz-bot/pkg/util"
)

func main() {
	botConfig := readConfig()
	bot := createTgBotApi(botConfig.Token, botConfig.Socks5Proxy)

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
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

func createTgBotApi(token string, socks5 string) *tgbotapi.BotAPI {
	client := &http.Client{}
	if len(socks5) > 0 {
		tgProxyURL, err := url.Parse(socks5)
		if err != nil {
			log.Printf("Failed to parse proxy URL:%s\n", err)
		}
		tgDialer, err := proxy.FromURL(tgProxyURL, proxy.Direct)
		if err != nil {
			log.Printf("Failed to obtain proxy dialer: %s\n", err)
		}
		tgTransport := &http.Transport{
			Dial: tgDialer.Dial,
		}
		client.Transport = tgTransport
	}

	bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Panic(err)
	}
	return bot
}
