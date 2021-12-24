package db

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	*mongo.Collection
}

