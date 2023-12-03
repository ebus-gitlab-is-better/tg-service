package telebot

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot *tele.Bot
}

func NewBot(bot *tele.Bot) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start(ctx context.Context) error {
	b.bot.Start()
	return nil
}

func (b *Bot) Stop(ctx context.Context) error {
	b.bot.Stop()
	return nil
}
