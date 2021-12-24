package card

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveCardPgt(r *mongo.Collection ,cardPayload Payload) (string, error) {
	result, err := r.InsertOne(context.Background(), cardPayload)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.Hex(), nil
}
