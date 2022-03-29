package entity

type (
	ProductInput struct {
		UserID      int    `json:"user_id"`
		Title       string `json:"title"`
		CategoryID  int    `json:"category_id"`
		SKU         string `json:"sku" bson:"sku"`
		Description string `json:"description" bson:"description"`
		Price       string `json:"price"`
		InStock     bool   `json:"in_stock" bson:"in_stock"`
		BrandID     int    `json:"brand_id" bson:"brand_id"`
	}

	SearchInput struct {
		SKU      string `json:"sku"`
		Category string `json:"category"`
		Brand    string `json:"brand"`
		Title    string `json:"title"`
	}
	IdUserInput struct {
		ID int `uri:"id" binding:"required"` //Ambil id dari URL

	}
)
