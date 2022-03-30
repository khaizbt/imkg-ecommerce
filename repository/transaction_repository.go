package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionRepository interface {
	AddToCart(post model.Cart) error
	UpdateCart(cartId string, qty int) error
	ListCart(userId int) ([]model.Cart, error)
	DetailCart(cartId string) (model.Cart, error)
	DeleteCart(cartId string) error
}

func (r *mongoRepo) AddToCart(post model.Cart) error {
	cartId, err := helper.GenerateNumber(6)

	if err != nil {
		return err
	}

	cart := map[string]interface{}{
		"_id":        cartId,
		"product_id": post.ProductID,
		"qty":        post.Qty,
		"user_id":    post.UserID,
		"created_at": time.Time{},
		"updated_at": time.Time{},
	}

	_, err = r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART")).InsertOne(context.Background(), cart)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (r *mongoRepo) ListCart(userId int) ([]model.Cart, error) {
	// collection := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer r.MongoClient.Disconnect(ctx)

	filter := bson.M{"user_id": userId}

	// mongoData, err := collection.Find(ctx, filter)

	mongoData, _ := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART")).Find(context.Background(), filter)

	var results []model.Cart

	if err := mongoData.All(ctx, &results); err != nil {
		return results, err
	}

	fmt.Println(results)

	return results, nil
}

func (r *mongoRepo) UpdateCart(cartId string, qty int) error {
	filter := bson.D{{"_id", cartId}}

	updater := bson.D{{"$set",
		bson.D{
			{"qty", qty},
		},
	}}

	_, err := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART")).UpdateOne(context.TODO(), filter, updater)

	if err != nil {
		return err
	}

	return nil

}

func (r *mongoRepo) DeleteCart(cartId string) error {
	filter := bson.D{primitive.E{"_id", cartId}}

	_, err := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART")).DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepo) DetailCart(cartId string) (model.Cart, error) {
	var result model.Cart
	filter := bson.D{primitive.E{"_id", cartId}}

	err := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART")).FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}
