package repository

import (
	"fmt"
	"github.com/khaizbt/imkg-ecommerce/model"
)

type UserRepository interface {
	//ListUser() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(userId int) (model.User, error)
}

//func (r *repository) ListUser() ([]model.User, error) {
//	results := []model.User{}
//
//	users, err := r.db.Query("select * from users")
//
//	if err != nil {
//		return results, err
//	}
//
//	defer users.Close()
//
//	for users.Next() {
//		var each model.User
//
//		err = users.Scan(&each.ID, &each.Name, &each.Email, each.Phone)
//
//		if err != nil {
//			return results, err
//		}
//
//		results = append(results, each)
//	}
//
//	return results, nil
//}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	fmt.Println("check12")
	err := r.db.Where("email = ?", email).First(&user).Error
	fmt.Println("check2")
	if err != nil {
		return user, err
	}
	fmt.Println("check3")
	return user, nil
}

func (r *repository) FindById(userId int) (model.User, error) {
	var user model.User

	err := r.db.Where("id = ?", userId).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
