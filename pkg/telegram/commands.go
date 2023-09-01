package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync-muz-bot/pkg/models"
)

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msgText, keyboard := b.buildMsgMainMenuView(message.Chat.FirstName)

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "html"

	user := models.User{
		Id:       message.Chat.ID,
		Username: message.Chat.UserName,
	}

	// repository test
	//_, err := b.repository.CreateUser(user)
	//if err != nil {
	//	user2, _ := b.repository.GetUser(message.Chat.ID)
	//	print(user2.Username)
	//}

	_, err = b.bot.Send(msg)
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
