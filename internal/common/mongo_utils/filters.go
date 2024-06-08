package mongo_utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateIdFilter(hexId string) (*bson.D, error) {
	oId, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, err
	}
	return &bson.D{{Key: "_id", Value: oId}}, nil
}
