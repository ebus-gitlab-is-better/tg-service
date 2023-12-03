package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	tele "gopkg.in/telebot.v3"
)

type User struct {
	ID     uint32 `gorm:"primaryKey"`
	ChatID int64
}

type UserRepo interface {
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, uint32) error
	GetByChatID(context.Context, int64) (*User, error)
	List(context.Context) ([]*User, error)
}

type UserUseCase struct {
	repo   UserRepo
	bot    *tele.Bot
	logger *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		logger: log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) error {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) error {
	return uc.repo.Update(ctx, user)
}

func (uc *UserUseCase) Delete(ctx context.Context, id uint32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUseCase) GetByChatID(ctx context.Context, chatId int64) (*User, error) {
	return uc.repo.GetByChatID(ctx, chatId)
}

func (uc *UserUseCase) SendAll(ctx context.Context, message string) error {
	users, err := uc.repo.List(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		chat, err := uc.bot.ChatByID(user.ChatID)
		if err != nil {
			continue
		}
		uc.bot.Send(chat, message)
	}
	return nil
}
