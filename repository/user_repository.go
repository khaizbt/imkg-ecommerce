package repository

import (
	"fmt"
	"github.com/khaizbt/imkg-ecommerce/model"
)

type UserRepository interface {
	FindByEmail(email string) (model.User, error)
	FindById(userId int) (model.User, error)
	UserFeature(feature model.UserTypeFeature) (model.UserTypeFeature, error)

	GetBrandIdByName(name string) (model.ProductBrand, error)
	GetCategoryIdByName(name string) (model.ProductCategory, error)
}

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

func (r *repository) UserFeature(feature model.UserTypeFeature) (model.UserTypeFeature, error) {
	err := r.db.Where("id_user_type = ?", feature.IDUserType).Where("id_feature = ?", feature.IDFeature).Find(&feature).Error

	if err != nil {
		return feature, err
	}

	return feature, nil
}
