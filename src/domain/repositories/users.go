package repositories

import (
	"context"
	ds "go-fiber-template/src/domain/datasources"
	"go-fiber-template/src/domain/entities"
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	Context    context.Context
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
	return &usersRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users"),
	}
}

func (repo *usersRepository) InsertNewUser(data *entities.UserDataFormat) error {
	if _, err := repo.Collection.InsertOne(repo.Context, data); err != nil {
		fiberlog.Errorf("Users -> InsertNewUser: %s \n", err)
		return err
	}
	return nil
}

func (repo *usersRepository) FindAll() (*[]entities.UserDataFormat, error) {
	cursor, err := repo.Collection.Find(repo.Context, bson.M{})
	if err != nil {
		fiberlog.Errorf("Users -> FindAll: %s \n", err)
		return nil, err
	}
	defer cursor.Close(repo.Context)

	var userData []entities.UserDataFormat
	for cursor.Next(repo.Context) {
		var user entities.UserDataFormat
		if err := cursor.Decode(&user); err != nil {
			continue
		}
		userData = append(userData, user)
	}

	return &userData, nil
}

func (repo *usersRepository) UpdateUser(userID string, data *entities.NewUserBody) error {
	if _, err := repo.Collection.UpdateOne(repo.Context, bson.M{"user_id": userID}, bson.M{"$set": data}); err != nil {
		fiberlog.Errorf("Users -> UpdateUser: %s \n", err)
		return err
	}
	return nil
}

func (repo *usersRepository) DeleteUser(userID string) error {
	if _, err := repo.Collection.DeleteOne(repo.Context, bson.M{"user_id": userID}); err != nil {
		fiberlog.Errorf("Users -> DeleteUser: %s \n", err)
		return err
	}
	return nil
}

func (repo *usersRepository) GetUser(userID string) (*entities.UserDataFormat, error) {
	var user entities.UserDataFormat
	if err := repo.Collection.FindOne(repo.Context, bson.M{"user_id": userID}).Decode(&user); err != nil {
		fiberlog.Errorf("Users -> GetUser: %s \n", err)
		return nil, err
	}
	return &user, nil
}
