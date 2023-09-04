package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msgText, keyboard := b.buildMsgMainMenuView(message.Chat.FirstName)

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "html"

	b.repository.CreateUserIfNotExists(message.Chat.ID, message.Chat.UserName)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msgText, keyboard := b.buildMsgHelpView()

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "html"

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleSyncCommand(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда")
	_, err := b.bot.Send(msg)
	return err
}
