package model

import "time"

type (
	Product struct {
		ID           string `bson:"_id" json:"id"`
		UserID       int    `bson:"user_id" json:"user_id"`
		BrandName    string `json:"brand_name,omitempty"`
		CategoryName string `json:"category_name,omitempty"`
		BrandID      int    `json:"brand_id"`
		CategoryID   int    `json:"category_id"`
		Slug         string `bson:"slug" json:"slug" binding:"required"`
		// CategoryId  int       `json:"category_id" bson:"category_id" binding:"required"`
		Title       string    `json:"title" bson:"title" binding:"required"`
		SKU         string    `json:"sku" bson:"sku" binding:"required"`
		Description string    `json:"description" bson:"description" binding:"required"`
		Price       string    `json:"price" binding:"required"`
		InStock     bool      `json:"in_stock" bson:"in_stock" binding:"required"`
		Stock       int       `json:"stock" bson:"stock"`
		CreatedAt   time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
		Image       []Image   `json:"image"`
	}

	ProductBrand struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	ProductCategory struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	Province struct {
		IDProvince   int    `json:"id_province"`
		ProvinceName string `json:"province_name"`
	}

	City struct {
		IdCity   int    `json:"id_city"`
		CityName string `json:"city_name"`
	}

	Image struct {
		FileName  string `json:"file_name" bson:"file_name"`
		IsPrimary bool   `json:"is_primary" bson:"is_primary"`
	}
)
