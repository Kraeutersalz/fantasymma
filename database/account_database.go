package database

import (
	"context"

	"github.com/kreaeutersalz/fantasymma/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAccount(id string) (error, *models.Account) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err, nil
	}
	collection := GetCollection("accounts")
	var account models.Account
	err = collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&account)
	if err != nil {
		return err, nil
	}
	return nil, &account
}

func CreateAccount(account *models.Account) (error, *models.Account) {
	collection := GetCollection("accounts")
	id, err := collection.InsertOne(context.Background(), account)
	if err != nil {
		return err, nil
	}
	account.Id = id.InsertedID.(primitive.ObjectID).Hex()
	return nil, account
}
