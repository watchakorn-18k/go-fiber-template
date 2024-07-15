package repositories

import (
	"context"
	"fmt"
	"os"
	"time"

	ds "go-fiber-template/src/domain/datasources"
	"go-fiber-template/src/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	Collection *mongo.Collection
}

type IUsersRepository interface {
	InsertNewUser(data *entities.UserDataFormat) error
	FindAll() (*[]entities.UserDataFormat, error)
	UpdateUser(userID string, data *entities.NewUserBody) error
	DeleteUser(userID string) error
	GetUser(userID string) (*entities.UserDataFormat, error)
}

func NewUsersRepository(db *ds.MongoDB) IUsersRepository {
	collection := db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users")
	return &usersRepository{
		Collection: collection,
	}
}

func (repo *usersRepository) InsertNewUser(data *entities.UserDataFormat) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := repo.Collection.InsertOne(ctx, data); err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}
	return nil
}

func (repo *usersRepository) FindAll() (*[]entities.UserDataFormat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding users: %v", err)
	}
	defer cursor.Close(ctx)

	var userData []entities.UserDataFormat
	for cursor.Next(ctx) {
		var user entities.UserDataFormat
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("error decoding user: %v", err)
		}
		userData = append(userData, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &userData, nil
}

func (repo *usersRepository) UpdateUser(userID string, data *entities.NewUserBody) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": data}

	if _, err := repo.Collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

func (repo *usersRepository) DeleteUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	if _, err := repo.Collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}

func (repo *usersRepository) GetUser(userID string) (*entities.UserDataFormat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.UserDataFormat
	filter := bson.M{"user_id": userID}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return &user, nil
}
