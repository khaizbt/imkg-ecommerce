package model

import "time"

type (
	Cart struct {
		ID          string    `json:"id" bson:"_id"`
		ProductID   string    `json:"product_id" bson:"product_id"`
		ProductName string    `json:"product_name,omitempty" `
		Qty         int       `json:"qty" bson:"qty"`
		UserID      int       `json:"user_id" bson:"user_id"`
		CreatedAt   time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	}

	Address struct {
		Name         string `json:"name" bson:"name"`
		Address      string `json:"address" bson:"address"`
		ProvinceId   int    `json:"province_id" bson:"province_id"`
		CityId       int    `json:"city_id" bson:"city_id"`
		CityName     string `json:"city_name,omitempty"`
		ProvinceName string `json:"province_name,omitempty"`
		ZipCode      int    `json:"zip_code" bson:"zip_code"`
	}

	TransactionDetail struct {
		ProductID   string `json:"product_id" bson:"product_id"`
		Qty         int    `json:"qty" bson:"qty"`
		ProductName string `json:"product_name,omitempty" `
	}

	Transaction struct {
		ID                string              `json:"id" bson:"_id"`
		UserID            int                 `json:"user_id" bson:"user_id"`
		PaymentType       string              `json:"payment_type" bson:"payment_type"`
		TransactionDetail []TransactionDetail `json:"transaction_detail" bson:"transaction_detail"`
		Address           Address             `json:"address" bson:"address"`
		Courier           Courier             `json:"courier" bson:"courier"`
		CreatedAt         time.Time           `json:"created_at" bson:"created_at"`
		UpdatedAt         time.Time           `json:"updated_at" bson:"updated_at"`
	}

	Courier struct {
		CourierName string `json:"courier_name" bson:"courier_name"`
		Waybill     string `json:"waybill" bson:"waybill"`
		CourierType string `json:"courier_type" bson:"courier_type"`
	}
)
