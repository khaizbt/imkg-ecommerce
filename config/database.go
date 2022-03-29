package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/olivere/elastic"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func GetElasticDB() *elastic.Client {
	esClient, err := elasticConnnect()
	if err != nil {
		log.Printf("Error Connecting to Elastic")
	}

	return esClient
}

func GetMongoDB() *mongo.Client {
	client, err := mongodbConnect()

	if err != nil {
		log.Printf("Error Connecting to MongoDB")
	}

	return client
}

func mongodbConnect() (mongoClient *mongo.Client, err error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGODB_HOST"), os.Getenv("MONGODB_PORT"))),

		options.Client().SetMaxConnIdleTime(5*time.Second))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func mysqlConnect() error {
	_ = godotenv.Load(".env")

	databaseInisialisation := "" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(databaseInisialisation), &gorm.Config{})

	if err != nil {
		return err
	}

	db = database

	return nil
}

func elasticConnnect() (esClient *elastic.Client, err error) {

	url := fmt.Sprintf("%s://%s:%s", os.Getenv("ELASTIC_SCHEME"), os.Getenv("ELASTIC_HOST"), os.Getenv("ELASTIC_PORT"))
	elasticURL := elastic.SetURL(url)
	sniff := elastic.SetSniff(false)
	client, err := elastic.NewClient(elasticURL, sniff)

	return client, err
}

func init() {

	err := mysqlConnect()

	if err != nil {
		log.Println(err)
	}
	log.Printf("Connected to the MYSQL %s:%s with database %s...", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	_, err = mongodbConnect()

	if err != nil {
		log.Println(err)
	}
	log.Printf("Connected to the MongoDB %s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))

}
