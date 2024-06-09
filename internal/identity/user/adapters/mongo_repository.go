package adapters

import (
	"common/mongo_utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"identity/user"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

type UserEntity struct {
	ID           string `bson:"_id"`
	Username     string
	Email        string
	PasswordHash string
	IsDeleted    bool
}

func (m *MongoUserRepository) AddUser(ctx context.Context, user *user.User) error {
	ue := &UserEntity{
		ID:           user.Id().String(),
		Username:     user.Username(),
		Email:        user.Email(),
		PasswordHash: user.Password(),
		IsDeleted:    user.IsDeleted(),
	}
	_, err := m.collection.InsertOne(ctx, ue)
	return err
}

func (m *MongoUserRepository) UpdateUser(ctx context.Context, user *user.User) error {
	filter, err := mongo_utils.CreateIdFilter(user.Id().String())
	if err != nil {
		return err
	}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "username", Value: user.Username()},
				{Key: "email", Value: user.Email()},
				{Key: "password_hash", Value: user.Password()},
				{Key: "is_deleted", Value: user.IsDeleted()},
			}}}
	_, err = m.collection.UpdateOne(ctx, filter, update)
	return err
}

func (m *MongoUserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	filter, err := mongo_utils.CreateIdFilter(id.String())
	if err != nil {
		return nil, err
	}
	var ue UserEntity
	err = m.collection.FindOne(ctx, filter).Decode(&ue)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	idParsed, err := uuid.Parse(ue.ID)
	if err != nil {
		return nil, err
	}
	return user.Assemble(idParsed, ue.Username, ue.Email, ue.PasswordHash, ue.IsDeleted), nil
}

func (m *MongoUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	filter, err := mongo_utils.CreateIdFilter(id.String())
	if err != nil {
		return err
	}
	_, err = m.collection.DeleteOne(ctx, filter)
	return err
}
