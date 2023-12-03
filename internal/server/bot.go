package server

import (
	"context"
	"fmt"
	"tgbot-service/internal/biz"
	"tgbot-service/pkg/telebot"

	tele "gopkg.in/telebot.v3"
)

func NewTelebot(uc *biz.UserUseCase, bot *tele.Bot) *telebot.Bot {
	//TODO выбор маршрута
	bot.Handle("/start", func(c tele.Context) error {
		uc.Create(context.TODO(), &biz.User{
			ChatID: c.Chat().ID,
		})
		fmt.Println("Вы подписались на уведомления")
		return nil
	})
	return telebot.NewBot(bot)
}
