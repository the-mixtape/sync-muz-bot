package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart = "start"
	//commandSettings  = "settings"
	//commandVkBinds   = "vk_binds"
	//commandStartSync = "start_sync"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(query.Message.Chat.ID, query.Data)
	_, err := b.bot.Send(msg)
	return err
}
