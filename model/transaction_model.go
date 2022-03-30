package model

import "time"

type (
	Cart struct {
		ID        string    `json:"id" bson:"_id"`
		ProductID string    `json:"product_id" bson:"product_id"`
		Qty       int       `json:"qty" bson:"qty"`
		UserID    int       `json:"user_id" bson:"user_id"`
		CreatedAt time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	}

	Address struct {
		Name        string `json:"name" bson:"name"`
		KecamatanID int    `json:"kecamatan_id" bson:"kecamatan_id"`
		CityId      int    `json:"city_id" bson:"city_id"`
		ZipCode     int    `json:"zip_code" bson:"zip_code"`
	}

	TransactionDetail struct {
		ProductID string `json:"product_id" bson:"product_id"`
		Qty       int    `json:"qty" bson:"qty"`
	}

	Transaction struct {
		ID                string `json:"id" bson:"_id"`
		UserID            int    `json:"user_id" bson:"user_id"`
		TransactionDetail `json:"transaction_detail" bson:"transaction_detail"`
		Address           `json:"address" bson:"address"`
		CreatedAt         time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
	}
)
