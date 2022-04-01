package repository

import (
	"github.com/khaizbt/imkg-ecommerce/config"
	"github.com/olivere/elastic"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

//type (
//	repository struct {
//		db *gorm.DB
//	}
//
//	UserRepository interface {
//		//ListUser() ([]model.User, error)
//		FindByEmail(email string) (model.User, error)
//		FindById(userId int) (model.User, error)
//	}
//)

//func NewRepository() *repository {
//	return &repository{config.GetSqlDB()}
//}

type repository struct {
	db *gorm.DB
}

type elasticRepo struct {
	Client *elastic.Client
}

type mongoRepo struct {
	MongoClient *mongo.Client
}

func NewElasticRepo() *elasticRepo {
	return &elasticRepo{config.GetElasticDB()} //Elastic
}
func NewRepository() *repository {
	return &repository{config.GetDB()} //MySql
}

func NewMongoRepository() *mongoRepo {
	return &mongoRepo{config.GetMongoDB()}
}
