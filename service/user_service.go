package service

import (
	"github.com/Favemus/config"
	"github.com/Favemus/model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserById(id uint64) (*model.User, error) {
	user := model.User{}
	err := config.DB.Select("id, email, created_at, updated_at").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserService) Save(user *model.User) error {
	config.DB.NewRecord(user)
	err := config.DB.Create(&user).Error
	return err
}
