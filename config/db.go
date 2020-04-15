package config

import (
	"fmt"
	"github.com/Favemus/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB = nil

func InitDB() {
	dbArgs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		Conf.DatabaseHost, Conf.DatabasePort, Conf.DatabaseUser, Conf.DatabaseName, Conf.DatabasePassword, Conf.DatabaseSSlActivated)
	db, err := gorm.Open(Conf.DatabaseDialect, dbArgs)
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	// Migrate the schema
	DB.AutoMigrate(&model.User{}, &model.Token{})

	// Add Foreign keys
	db.Model(&model.Token{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
