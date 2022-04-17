package entity

import "github.com/khaizbt/imkg-ecommerce/model"

type (
	ProductInput struct {
		ID          string        `json:"product_id"`
		UserID      int           `json:"user_id"`
		Title       string        `json:"title"`
		Slug        string        `json:"slug"`
		CategoryID  int           `json:"category_id"`
		SKU         string        `json:"sku" bson:"sku"`
		Description string        `json:"description" bson:"description"`
		Price       string        `json:"price"`
		InStock     bool          `json:"in_stock" bson:"in_stock"`
		BrandID     int           `json:"brand_id" bson:"brand_id"`
		Stock       int           `json:"stock" bson:"stock"`
		Image       []model.Image `json:"image"`
	}

	SearchInput struct {
		SKU      string `json:"sku"`
		Category string `json:"category"`
		Brand    string `json:"brand"`
		Title    string `json:"title"`
	}
	IdUserInput struct {
		ID        int    `uri:"id"` //Ambil id dari URL
		ProductId string `uri:"product_id"`
	}
)
