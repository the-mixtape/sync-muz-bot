package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initializeUpdatesChannel(0, 60)
	err := b.handleUpdates(updates)

	return err
}

func (b *Bot) initializeUpdatesChannel(offset int, timeout int) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(offset)
	u.Timeout = timeout

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message)
			if err != nil {
				return err
			}
			continue
		}

		err := b.handleMessage(update.Message)
		if err != nil {
			return err
		}
	}
	return nil
}
