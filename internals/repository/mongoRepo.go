package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/User_model"
)

// MongoRepo struct definition
type MongoRepo struct {
	db  *mongo.Database
	ctx context.Context
}

func (r *MongoRepo) RegisterUser(user models.User) (int64, error) {
	collection := r.db.Collection("users")
	result, err := collection.InsertOne(r.ctx, user)
	if err != nil {
		return 0, err
	}
	insertedID := result.InsertedID.(int64)
	return insertedID, nil
}
