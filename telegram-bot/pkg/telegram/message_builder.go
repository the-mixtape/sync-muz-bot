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

func (b *Bot) buildMsgVkMuzView(vkId *int64, vkName *string) (string, tgbotapi.InlineKeyboardMarkup) {
	var numericKeyboard = tgbotapi.InlineKeyboardMarkup{}

	var msgText = ""
	if vkId == nil {
		msgText = "На данный момент страница <b>ВК</b> для синхронизации не указана"

		newButton := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("🔗 Привязать страницу ВК", callbackEditVkId),
		}
		numericKeyboard.InlineKeyboard = append(numericKeyboard.InlineKeyboard, newButton)

	} else {
		msgText = fmt.Sprintf(`Привязан аккаунт <b>ВК</b>:
Имя - %s
ID - %d
`, *vkName, vkId)

		newButtons := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("🔗 Изменить привязку ВК", callbackEditVkId),
			tgbotapi.NewInlineKeyboardButtonData("❌ Удалить привязку ВК", callbackDeleteVkId),
		}
		numericKeyboard.InlineKeyboard = append(numericKeyboard.InlineKeyboard, newButtons)
	}

	mainMenuButton := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("🔙 Главное меню", callbackMainMenuView),
	}
	numericKeyboard.InlineKeyboard = append(numericKeyboard.InlineKeyboard, mainMenuButton)

	return msgText, numericKeyboard
}

func (b *Bot) buildMsgEditVkId() string {
	return `В следующем сообщении отправьте ссылку на профиль <b>ВК</b>, который вы хотите привязать.
Важный момент, чтобы синхронизация музыки работала, доступ к ней должен быть открыт.`
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
