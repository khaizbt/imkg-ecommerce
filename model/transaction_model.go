package model

import "time"

type (
	Cart struct {
		ID        string    `json:"id" bson:"id"`
		ItemName  string    `json:"item_name" bson:"item_name"`
		Seller    string    `json:"seller" bson:"seller"`
		Qty       int       `json:"qty" bson:"qty"`
		UserID    int       `json:"user_id" bson:"user_id"`
		CreatedAt time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	}
)
