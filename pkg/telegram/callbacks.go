package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCallbackMainMenu(query *tgbotapi.CallbackQuery) error {
	text, keyboard := b.buildMsgMainMenuView(query.Message.Chat.FirstName)
	editMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, text, keyboard)
	editMsg.ParseMode = "html"

	_, err := b.bot.Send(editMsg)
	return err
}

func (b *Bot) handleCallbackVkMuzView(query *tgbotapi.CallbackQuery) error {
	text, keyboard := b.buildMsgVkMuzView()
	editMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, text, keyboard)
	editMsg.ParseMode = "html"

	_, err := b.bot.Send(editMsg)
	return err
}

func (b *Bot) handleCallbackYandexMuzView(query *tgbotapi.CallbackQuery) error {
	text, keyboard := b.buildMsgYandexMuzView()
	editMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, text, keyboard)
	editMsg.ParseMode = "html"

	_, err := b.bot.Send(editMsg)
	return err
}

func (b *Bot) handleCallbackUnknown(query *tgbotapi.CallbackQuery) error {
	var msgText = fmt.Sprintf("Unknown callback: <b><i>%s</i></b>", query.Data)
	msg := tgbotapi.NewMessage(query.Message.Chat.ID, msgText)
	msg.ParseMode = "html"

	_, err := b.bot.Send(msg)
	return err
}
