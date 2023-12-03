package telebot

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot *tele.Bot
}

func (b *Bot) Start(ctx context.Context) {
	b.bot.Start()
}

func (b *Bot) Stop(ctx context.Context) {
	b.bot.Stop()
}
