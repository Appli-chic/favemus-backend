package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string
	Hash  string
}
