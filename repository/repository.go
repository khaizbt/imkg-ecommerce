package repository

import (
	"github.com/khaizbt/imkg-ecommerce/config"
	"gorm.io/gorm"
)

//type (
//	repository struct {
//		db *gorm.DB
//	}
//
//	UserRepository interface {
//		//ListUser() ([]model.User, error)
//		FindByEmail(email string) (model.User, error)
//		FindById(userId int) (model.User, error)
//	}
//)

//func NewRepository() *repository {
//	return &repository{config.GetSqlDB()}
//}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{config.GetDB()}
}
