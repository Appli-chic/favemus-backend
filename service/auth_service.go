package service

import (
	"github.com/Favemus/config"
	"github.com/Favemus/model"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) GetUserByEmailAndPassword(email string, password string) (*model.User, error) {
	user := model.User{}
	err := config.DB.Select("id, email").Where("email = ? AND hash = ?", email, password).First(&user).Error
	return &user, err
}
