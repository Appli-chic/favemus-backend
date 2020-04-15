package service

import (
	"github.com/Favemus/config"
	"github.com/Favemus/model"
)

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (t *TokenService) Save(token model.Token) (model.Token, error) {
	config.DB.NewRecord(token)
	err := config.DB.Create(&token).Error
	return token, err
}

func (t *TokenService) GetTokenByUserId(userId interface{}) (model.Token, error) {
	token := model.Token{}
	err := config.DB.Where("user_id = ?", userId).First(&token).Error
	return token, err
}
