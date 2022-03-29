package model

import "time"

type (
	Product struct {
		ID          string    `bson:"id" json:"id"`
		UserID      int       `bson:"user_id" json:"user_id"`
		Slug        string    `bson:"slug" json:"slug" binding:"required"`
		CategoryId  int       `json:"category_id" bson:"category_id" binding:"required"`
		Title       string    `json:"title" bson:"title" binding:"required"`
		SKU         string    `json:"sku" bson:"sku" binding:"required"`
		Description string    `json:"description" bson:"description" binding:"required"`
		Price       string    `json:"price" binding:"required"`
		InStock     bool      `json:"in_stock" bson:"in_stock" binding:"required"`
		CreatedAt   time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
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
)
