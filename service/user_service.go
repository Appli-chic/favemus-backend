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
	err := config.DB.Select("id, email, name, created_at, updated_at").Where("users.id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserService) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *UserService) GetUserFromRefreshToken(refreshToken string) (model.User, error) {
	user := model.User{}
	err := config.DB.
		Joins("left join tokens on tokens.user_id = users.id").
		Where("tokens.token = ?", refreshToken).
		First(&user).Error
	return user, err
}

func (u *UserService) Save(user *model.User) error {
	config.DB.NewRecord(user)
	err := config.DB.Create(&user).Error
	return err
}
