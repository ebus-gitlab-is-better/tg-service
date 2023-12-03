package service

import "tgbot-service/internal/biz"

type BotService struct {
	uc *biz.UserUseCase
}

func NewBotService(uc *biz.UserUseCase) *BotService {
	return &BotService{uc: uc}
}
