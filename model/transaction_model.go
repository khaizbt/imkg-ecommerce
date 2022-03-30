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
)
