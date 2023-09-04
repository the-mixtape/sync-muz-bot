package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandHelp  = "help"
	commandSync  = "sync"
)

const (
	callbackStartSync     = "start_sync"
	callbackMainMenuView  = "main_menu"
	callbackVkMuzView     = "vk_muz_view"
	callbackYandexMuzView = "yandex_muz_view"
	callbackSettingsView  = "settings_view"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandSync:
		return b.handleSyncCommand(message)
	case commandStart:
		return b.handleStartCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	callback := tgbotapi.NewCallback(query.ID, query.Data)
	_, err := b.bot.Request(callback)
	if err != nil {
		return err
	}

	switch query.Data {
	case callbackMainMenuView:
		return b.handleCallbackMainMenu(query)
	case callbackVkMuzView:
		return b.handleCallbackVkMuzView(query)
	case callbackYandexMuzView:
		return b.handleCallbackYandexMuzView(query)
	default:
		return b.handleCallbackUnknown(query)
	}
}
