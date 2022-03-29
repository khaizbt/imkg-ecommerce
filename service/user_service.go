package service

import (
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	service struct {
		user_repository repository.UserRepository
	}

	UserService interface {
		Login(input entity.LoginInput) (model.User, error)
		GetUserById(userId int) (model.User, error)
		CheckFeature(idUserType int, idFeature int) (model.UserTypeFeature, error)
	}
)

func NewUserService(userRepository repository.UserRepository) *service {
	return &service{userRepository}
}

func (s *service) Login(input entity.LoginInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.user_repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(id int) (model.User, error) {
	user, err := s.user_repository.FindById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckFeature(idUserType int, idFeature int) (model.UserTypeFeature, error) {
	feature := model.UserTypeFeature{
		IDUserType: idUserType,
		IDFeature:  idFeature,
	}

	feature, err := s.user_repository.UserFeature(feature)

	if err != nil {
		return feature, err
	}

	return feature, nil
}
