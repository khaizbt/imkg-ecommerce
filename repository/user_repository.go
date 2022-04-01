package repository

import (
	"github.com/khaizbt/imkg-ecommerce/model"
)

type UserRepository interface {
	FindByEmail(email string) (model.User, error)
	FindById(userId int) (model.User, error)
	UserFeature(feature model.UserTypeFeature) (model.UserTypeFeature, error)
	CreateUser(user model.User) (model.User, error)
	FindByUsername(Username string) (model.User, error)

	//Master
	GetBrandIdByName(name string) (model.ProductBrand, error)
	GetBrandById(brandId int) (model.ProductBrand, error)
	GetCategoryIdByName(name string) (model.ProductCategory, error)
	GetCategoryById(categoryId int) (model.ProductCategory, error)

	GetProvinceList() ([]model.Province, error)
	GetCityList() ([]model.City, error)
	GetProvinceById(provinceId int) (model.Province, error)
	GetCityById(cityId int) (model.City, error)
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}

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

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByUsername(Username string) (model.User, error) {
	var user model.User

	err := r.db.Where("username = ?", Username).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
