package repository

import "github.com/khaizbt/imkg-ecommerce/model"

func (r *repository) GetProvinceList() ([]model.Province, error) {
	var province []model.Province
	err := r.db.Find(&province).Error

	if err != nil {
		return province, err
	}

	return province, nil
}

func (r *repository) GetCityList() ([]model.City, error) {
	var city []model.City
	err := r.db.Find(&city).Error

	if err != nil {
		return city, err
	}

	return city, nil
}

func (r *repository) GetProvinceById(provinceId int) (model.Province, error) {
	var province model.Province

	err := r.db.Where("id_province = ?", provinceId).First(&province).Error

	if err != nil {
		return province, err
	}

	return province, nil
}

func (r *repository) GetCityById(cityId int) (model.City, error) {
	var city model.City

	err := r.db.Where("id_city = ?", cityId).First(&city).Error

	if err != nil {
		return city, err
	}

	return city, nil
}
