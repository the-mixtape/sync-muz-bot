package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) buildMsgMainMenuView(interlocutorName string) (string, tgbotapi.InlineKeyboardMarkup) {
	var msgText = fmt.Sprintf(
		`Привет, %s! Меня зовут <b>%s</b> 🤖

Моя цель - синхронизация музыки из разных сервисов 🎵
Вы можете привязать несколько сервисов и я буду присылать вам обновления ваших плейлистов, по расписанию 🕙
Настроить сервисы и время проверки ваших плейлистов можно ниже ⬇️⬇️⬇️

Сервисы с которыми я умею работать:
<i><b>• Vk Music</b></i>
<i><b>• Yandex Music</b></i>
`,
		interlocutorName,
		b.bot.Self.UserName,
	)

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔄 Синхронизировать", callbackStartSync),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎵 Vk Music", callbackVkMuzView),
			tgbotapi.NewInlineKeyboardButtonData("🎵 Yandex Music", callbackYandexMuzView),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⚙️ Настройки", callbackSettingsView),
		),
	)

	return msgText, numericKeyboard
}

func (b *Bot) buildMsgHelpView() (string, tgbotapi.InlineKeyboardMarkup) {
	var msgText = fmt.Sprintf(`Я <b>%s</b> 🤖

Сервисы с которыми я умею работать:
<i><b>• Vk Music</b></i>
<i><b>• Yandex Music</b></i>
`,
		b.bot.Self.UserName,
	)

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Главное меню", callbackMainMenuView),
		),
	)

	return msgText, numericKeyboard
}

func (b *Bot) buildMsgVkMuzView() (string, tgbotapi.InlineKeyboardMarkup) {
	return b.buildMsgWIP()
}

func (b *Bot) buildMsgYandexMuzView() (string, tgbotapi.InlineKeyboardMarkup) {
	return b.buildMsgWIP()
}

func (b *Bot) buildMsgWIP() (string, tgbotapi.InlineKeyboardMarkup) {
	var msgText = `Данный раздел в разработке 👷🏻👷🏿
Попробуйте позже 🕙
Спасибо 🖐
`

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Главное меню", callbackMainMenuView),
		),
	)

	return msgText, numericKeyboard
}
