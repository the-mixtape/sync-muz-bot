package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync-muz-bot/pkg/cache"
	"sync-muz-bot/pkg/repository"
	"sync-muz-bot/pkg/vk_api"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	botCache   cache.BotCache
	repository *repository.Repository
	vkApi      *vk_api.VkApi
}

func NewBot(bot *tgbotapi.BotAPI, c cache.BotCache, repos *repository.Repository, vkApi *vk_api.VkApi) *Bot {
	return &Bot{
		bot:        bot,
		botCache:   c,
		repository: repos,
		vkApi:      vkApi,
	}
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
		if update.Message != nil {
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
		} else if update.CallbackQuery != nil {
			err := b.handleCallbackQuery(update.CallbackQuery)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
