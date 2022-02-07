package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	Address Address            `bson:"address"`
}

type Address struct {
	Address  string `bson:"address"`
	City     string `bson:"city"`
	Province string `bson:"province"`
}
