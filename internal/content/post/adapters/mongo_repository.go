package adapters

import (
	"common/mongo_utils"
	"content/post"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "posts"
const timeFormat = time.RFC3339

type MongoPostRepository struct {
	collection *mongo.Collection
}

type PostEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Text      string             `bson:"text"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

func (m *MongoPostRepository) AddPost(ctx context.Context, post *post.Post) error {
	pe := PostEntity{
		ID:        primitive.NewObjectID(),
		Text:      post.Text(),
		CreatedAt: post.CreatedAt().Format(timeFormat),
		UpdatedAt: post.UpdatedAt().Format(timeFormat),
	}
	_, err := m.collection.InsertOne(ctx, pe)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoPostRepository) UpdatePost(ctx context.Context, post *post.Post) error {
	filter, err := mongo_utils.CreateIdFilter(post.Id())
	if err != nil {
		return err
	}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "text", Value: post.Text()},
		{Key: "updatedAt", Value: time.Now().UTC()},
	}}}
	_, err = m.collection.UpdateOne(ctx, filter, update)
	return err
}

func (m *MongoPostRepository) GetPostById(ctx context.Context, id string) (*post.Post, error) {
	filter, err := mongo_utils.CreateIdFilter(id)
	if err != nil {
		return nil, err
	}
	var pe PostEntity
	err = m.collection.FindOne(ctx, filter).Decode(&pe)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	createdAt, err := time.Parse(timeFormat, pe.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(timeFormat, pe.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return post.New(pe.ID.Hex(), pe.Text, createdAt, updatedAt), nil
}

func (m *MongoPostRepository) DeletePostById(ctx context.Context, id string) error {
	filter, err := mongo_utils.CreateIdFilter(id)
	if err != nil {
		return err
	}
	_, err = m.collection.DeleteOne(ctx, filter)
	return err
}

func NewMongoPostRepository(database *mongo.Database) (*MongoPostRepository, error) {
	c := database.Collection(collectionName)
	return &MongoPostRepository{c}, nil
}
