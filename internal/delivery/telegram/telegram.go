package telegram

import (
	"github.com/KutsDenis/anti-smoking/internal/config"
	"github.com/KutsDenis/logpac"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

var onceInit sync.Once
var Bot *tgbotapi.BotAPI

// StartBot запускает бота Telegram с заданным токеном и настройками отладки
func StartBot(l *logpac.Logger) {
	onceInit.Do(func() {
		var err error

		Bot, err = tgbotapi.NewBotAPI(config.Get.Bot.Token)
		if err != nil {
			l.Panicf("%s", err)
		}

		Bot.Debug = config.Get.Bot.Debug
	})
}
