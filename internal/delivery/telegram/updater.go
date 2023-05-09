package telegram

import (
	"github.com/KutsDenis/logpac"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Update(bot *tgbotapi.BotAPI, l *logpac.Logger) {
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60

	updates := bot.GetUpdatesChan(upd)

	for update := range updates {
		// Перехват колбэков
		if update.CallbackQuery != nil {
			// todo
		}

		// Игнорирование любых обновлений кроме сообщений
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		// Команды
		// todo
	}
}
