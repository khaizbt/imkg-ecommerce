package repository

import (
	"database/sql"
	"golang-template/config"
	"golang-template/entity"
)

type (
	repository struct {
		db *sql.DB
	}

	UserRepository interface {
		ListUser() ([]entity.User, error)
	}
)

func NewRepository() *repository {
	return &repository{config.GetSqlDB()}
}
