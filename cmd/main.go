package main

import (
	"github.com/KutsDenis/anti-smoking/internal/config"
	"github.com/KutsDenis/anti-smoking/internal/delivery/telegram"
	"github.com/KutsDenis/logpac"
)

func main() {
	l := logpac.GetLogger()

	config.GetConfig(l)
	telegram.StartBot(l)
	bot := *telegram.Bot

	telegram.Update(&bot, l)
}
