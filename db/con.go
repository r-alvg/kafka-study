package db

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"sync"
)

type MongoDatabase struct {
	con *mongo.Database
}

var (
	DB       *mongo.Database
	instance *MongoDatabase
	once     sync.Once
)

func Start() {
	DB = connect()
}

func connect() *mongo.Database {
	once.Do(func() {
		ctx := context.Background()
		clientOpt := options.Client().ApplyURI("mongodb+srv://0uno:qy0nSuKoGijRUNSQ@twitter-ej3ca.mongodb.net/test?retryWrites=true&w=majority")
		client, err := mongo.NewClient(clientOpt)

		if err != nil {
			logrus.Error(err)
		}
		if err := client.Connect(ctx); err != nil {
			logrus.Fatal("error on connect")
		}
		if err := client.Ping(context.TODO(), nil); err != nil {
			logrus.Fatal("error on ping cause: ", err)
		}

		instance = &MongoDatabase{client.Database("kafka_db")}
	})
	return instance.con
}

func NewRepository(name string) Repository {
	return Repository{DB.Collection(name)}
}
