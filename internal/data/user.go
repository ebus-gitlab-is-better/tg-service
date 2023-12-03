package data

import (
	"context"
	"tgbot-service/internal/biz"
)

type User struct {
	ID     uint32 `gorm:"primaryKey"`
	ChatID int64  `gorm:"unique"`
}

func (m User) TableName() string {
	return "users_tg"
}

func (m User) modelToResponse() *biz.User {
	return &biz.User{
		ID:     m.ID,
		ChatID: m.ChatID,
	}
}

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

// Create implements biz.UserRepo.
func (r *userRepo) Create(ctx context.Context, user *biz.User) error {
	var userDB User
	userDB.ChatID = user.ChatID
	if err := r.data.db.Create(&userDB).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements biz.UserRepo.
func (r *userRepo) Delete(ctx context.Context, id uint32) error {
	return r.data.db.Delete(&User{}, id).Error
}

// GetByChatID implements biz.UserRepo.
func (r *userRepo) GetByChatID(ctx context.Context, chadId int64) (*biz.User, error) {
	var userDB User
	if err := r.data.db.Where(&User{ChatID: chadId}).Find(&userDB).Error; err != nil {
		return nil, err
	}

	return userDB.modelToResponse(), nil
}

// List implements biz.UserRepo.
func (r *userRepo) List(context.Context) ([]*biz.User, error) {
	var usersDB []User
	localDB := r.data.db.Model(&User{})
	if err := localDB.Find(&usersDB).Error; err != nil {
		return nil, err
	}
	users := make([]*biz.User, 0)
	for _, u := range usersDB {
		users = append(users, u.modelToResponse())
	}
	return users, nil
}

// Update implements biz.UserRepo.
func (r *userRepo) Update(ctx context.Context, user *biz.User) error {
	var userDB User
	userDB.ChatID = user.ChatID
	userDB.ID = user.ID
	if err := r.data.db.Save(&userDB).Error; err != nil {
		return err
	}
	return nil
}
