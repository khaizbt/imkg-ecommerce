package service

import (
	"errors"
	"runtime"

	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/helper"
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
		CreateUser(input entity.DataUserInput) (bool, error)
	}

	BodyEmail struct {
		Name string
		Body string
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

func (s *service) CreateUser(input entity.DataUserInput) (bool, error) {
	cekUser, _ := s.user_repository.FindByEmail(input.Email)

	if cekUser.ID != 0 {
		return false, errors.New("email has been registered")
	}

	cekUsername, _ := s.user_repository.FindByUsername(input.Username)

	if cekUsername.ID != 0 {
		return false, errors.New("username has been taken")
	}

	user := model.User{
		Name:       input.Name,
		Email:      input.Email,
		Username:   input.Username,
		IdUserType: 2,
		Status:     "Active",
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return false, err
	}

	user.Password = string(password)

	_, err = s.user_repository.CreateUser(user)

	if err != nil {
		return false, err
	}

	bodyEmail := BodyEmail{
		input.Name,
		"Terima Kasih sudah Registrasi di IMKG, Semoga betah",
	}
	runtime.GOMAXPROCS(2)

	//Send Email
	go helper.SendEmail(input.Email, "Registration Success", "html/email_verification.html", bodyEmail)

	return true, nil
}
