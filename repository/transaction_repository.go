package repository

import (
	"context"
	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"log"
	"time"
)

type TransactionRepository interface {
	AddToCart(post model.Cart) error
}

func (r *mongoRepo) AddToCart(post model.Cart) error {
	cartId, err := helper.GenerateNumber(6)

	if err != nil {
		return err
	}

	cart := map[string]interface{}{
		"_id":        cartId,
		"item_name":  post.ItemName,
		"seller":     post.Seller,
		"qty":        post.Qty,
		"user_id":    post.UserID,
		"created_at": time.Time{},
		"updated_at": time.Time{},
	}

	_, err = r.MongoClient.Database("imkg_database").Collection("carts").InsertOne(context.Background(), cart)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
