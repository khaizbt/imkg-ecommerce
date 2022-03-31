package repository

import (
	"context"
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
	Checkout(post model.Transaction) error
	ListTransaction(userId int) ([]model.Transaction, error)
	ListSalesTransaction() ([]model.Transaction, error)
	UpdateTransaction(transactionId, key string, data interface{}) error
	DeleteBulkCart(carts []string) error
}

//NOTES I'm sorry for not Implement DRY Concept in Here
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

func (r *mongoRepo) Checkout(post model.Transaction) error {
	checkoutId, err := helper.GenerateNumber(6)

	if err != nil {
		return err
	}

	checkout := map[string]interface{}{
		"_id":                checkoutId,
		"user_id":            post.UserID,
		"transaction_detail": post.TransactionDetail,
		"address":            post.Address,
		"payment_type":       post.PaymentType,
		"status":             "Dibayar",
		"courier":            nil,
		"created_at":         time.Time{},
		"updated_at":         time.Time{},
	}

	_, err = r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_TRANSACTION")).InsertOne(context.Background(), checkout)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (r *mongoRepo) ListTransaction(userId int) ([]model.Transaction, error) {
	// collection := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART"))

	filter := bson.M{"user_id": userId}

	// mongoData, err := collection.Find(ctx, filter)

	mongoData, _ := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_TRANSACTION")).Find(context.TODO(), filter)

	var results []model.Transaction

	if err := mongoData.All(context.TODO(), &results); err != nil {
		return results, err
	}

	return results, nil
}

func (r *mongoRepo) ListSalesTransaction() ([]model.Transaction, error) {
	var listTransaction []model.Transaction

	filter := bson.D{}

	mongoData, _ := r.MongoClient.
		Database(os.Getenv("MONGODB_DATABASE")).
		Collection(os.Getenv("MONGODB_COLLECTION_TRANSACTION")).
		Find(context.Background(), filter)

	// listTransaction := make([]model.Transaction, 0)
	// for mongoData.Next(context.Background()) {
	// 	t := model.Transaction{}
	// 	err := mongoData.Decode(&t)

	// 	if err != nil {
	// 		return listTransaction, err
	// 	}

	// 	listTransaction = append(listTransaction, t)
	// }
	if err := mongoData.All(context.TODO(), &listTransaction); err != nil {
		return listTransaction, err
	}

	return listTransaction, nil

}

func (r *mongoRepo) UpdateTransaction(transactionId, key string, data interface{}) error {
	filter := bson.D{{"_id", transactionId}}

	updater := bson.D{{"$set",
		bson.D{
			{key, data},
		},
	}}

	_, err := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_TRANSACTION")).UpdateOne(context.TODO(), filter, updater)

	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepo) DeleteBulkCart(carts []string) error {
	filter := bson.M{"_id": bson.M{"$in": carts}}

	_, err := r.MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION_CART")).DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil

}
