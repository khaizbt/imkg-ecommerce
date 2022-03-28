package config

import (
	"context"

	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func mongodbConnect(host, port, db string) (mongoClient *mongo.Client, err error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)),
		options.Client().SetAppName("MONGODB_GO_STOCK"),
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

func elasticConnect(scheme, host, port string) (esClient *elastic.Client, err error) {
	url := fmt.Sprintf("%s://%s:%s", scheme, host, port)
	elasticURL := elastic.SetURL(url)
	sniff := elastic.SetSniff(false)
	client, err := elastic.NewClient(elasticURL, sniff)

	return client, err
}

// Initialize MySQL
func mysqlConnect(user, password, host, port, dbName string) (DB *sql.DB, err error) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName)

	return db, err
}

// func mySqlConnect(user, password, host, port, dbName string) (DB *mysql.DB)

var projectFolder *string

func init() {
	projectFolder = flag.String("folder", "./", "absolute path of project folder")
	flag.Parse()

	if *projectFolder == "" {
		// Get file path
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		a := fmt.Sprintf("%s", dir+"/")
		projectFolder = &a
	}

	// Load ENV Config
	gotenv.Load(*projectFolder + ".env")
	log.Printf("Initializing {project_name} Service...")

	// Should be initialize Mongo Connection here
	_, err := mongodbConnect(os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"))
	failOnError(err, "Error when dialing MongoDB")
	log.Printf("Connected to the MongoDB %s with database %s...", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_DB"))

	_, err = mongodbConnect(os.Getenv("MONGO_HOST_HILDA"), os.Getenv("MONGO_PORT_HILDA"), os.Getenv("MONGO_DB_HILDA"))
	failOnError(err, "Error when dialing MongoDB Main")
	log.Printf("Connected to the MongoDB %s with database %s...", os.Getenv("MONGO_HOST_HILDA"), os.Getenv("MONGO_DB_HILDA"))

	// Initialize MySQL Connection here
	mysqlClient, err := mysqlConnect(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	failOnError(err, "Error when connecting to MySQL")
	defer mysqlClient.Close()
	log.Printf("Connected to the MYSQL %s:%s with database %s...", os.Getenv("DB_HOST_MASTER"), os.Getenv("DB_PORT_MASTER"), os.Getenv("DB_NAME_MASTER"))

	// Initialize Elasticsearch Client here
	_, err = elasticConnect(os.Getenv("ELASTIC_SCHEME"), os.Getenv("ELASTIC_HOST"), os.Getenv("ELASTIC_PORT"))
	failOnError(err, "Error when connecting to Elasticsearch")
	log.Printf("Connected to the ElasticSearch %s:%s...", os.Getenv("ELASTIC_HOST"), os.Getenv("ELASTIC_PORT"))
}
