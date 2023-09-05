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
	chatId := query.Message.Chat.ID

	user, err := b.repository.GetUser(chatId)
	if err != nil {
		return err
	}

	var userVkName *string = nil
	if user.VkId != nil {
		userInfo, err := b.vkApi.GetUserIdAndName(chatId)
		if err == nil {
			userVkName = &userInfo.UserName
		}
	}

	text, keyboard := b.buildMsgVkMuzView(user.VkId, userVkName)
	editMsg := tgbotapi.NewEditMessageTextAndMarkup(chatId, query.Message.MessageID, text, keyboard)
	editMsg.ParseMode = "html"

	_, err = b.bot.Send(editMsg)
	return err
}

func (b *Bot) handleCallbackEditVkId(query *tgbotapi.CallbackQuery) error {
	chatId := query.Message.Chat.ID

	text := b.buildMsgEditVkId()
	editMsg := tgbotapi.NewEditMessageText(chatId, query.Message.MessageID, text)
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
