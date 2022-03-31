package entity

import "github.com/khaizbt/imkg-ecommerce/model"

type (
	CartInput struct {
		UserID    int
		ProductID string `json:"product_id"`
		Qty       int    `json:"qty"`
		CartId    string `json:"cart_id"`
	}

	Checkout struct {
		CourierID      int              `json:"courier_id"`
		UserID         int              `json:"user_id"`
		PaymentType    string           `json:"payment_type"`
		CheckoutDetail []CheckoutDetail `json:"checkout_detail"`
		Address        model.Address    `json:"address" bson:"address"`
	}

	CheckoutDetail struct {
		CartId string `json:"cart_id"`
	}

	// Address struct {
	// 	Name        string `json:"name" bson:"name"`
	// 	KecamatanID int    `json:"kecamatan_id" bson:"kecamatan_id"`
	// 	CityId      int    `json:"city_id" bson:"city_id"`
	// 	ZipCode     int    `json:"zip_code" bson:"zip_code"`
	// }

	TransactionStatus struct {
		TransactionId string        `json:"transaction_id" bson:"transaction_id"`
		Status        string        `json:"status" bson:"status"`
		Courier       model.Courier `json:"courier"`
	}
)
