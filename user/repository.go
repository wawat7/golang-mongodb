package user

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-mongodb/app"
	"golang-mongodb/exception"
)

type Repository interface {
	Create(user User, address Address) User
	FindAll() (users []User)
	FindById(Id string) (user User)
	Update(user User) User
	Delete(Id string)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepository(database *mongo.Database) *repository {
	return &repository{Collection: database.Collection("users")}
}

func (repo *repository) Create(user User, address Address) User {
	ctx, cancel := app.NewMongoContext()
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, bson.M{
		"name":  user.Name,
		"email": user.Email,
		"address": bson.M{
			"address":  address.Address,
			"city":     address.City,
			"province": address.Province,
		},
	})
	exception.PanicIfNeeded(err)

	return user
}

func (repo *repository) FindAll() (users []User) {
	ctx, cancel := app.NewMongoContext()
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		fmt.Println(document["address"])
		user := User{
			Id:    document["_id"].(primitive.ObjectID),
			Name:  document["name"].(string),
			Email: document["email"].(string),
		}

		address := document["address"].(bson.M)
		user.Address = Address{
			Address:  address["address"].(string),
			City:     address["city"].(string),
			Province: address["province"].(string),
		}

		users = append(users, user)

	}

	return
}

func (repo *repository) FindById(Id string) (user User) {
	ctx, cancel := app.NewMongoContext()
	defer cancel()

	result := User{}
	objID, err := primitive.ObjectIDFromHex(Id)
	exception.PanicIfNeeded(err)
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	//Perform FindOne operation & validate against the error.
	err = repo.Collection.FindOne(ctx, filter).Decode(&result)
	exception.PanicIfNeeded(err)
	fmt.Println(result)
	//Return result without any error.
	return result

}

func (repo *repository) Update(user User) User {
	ctx, cancel := app.NewMongoContext()
	defer cancel()
	filter := bson.D{primitive.E{Key: "_id", Value: user.Id}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "name", Value: user.Name},
		primitive.E{Key: "email", Value: user.Email},
	}}}

	//Perform UpdateOne operation & validate against the error.
	_, err := repo.Collection.UpdateOne(ctx, filter, updater)
	exception.PanicIfNeeded(err)
	//Return success without any error.
	return user
}

func (repo *repository) Delete(Id string) {
	ctx, cancel := app.NewMongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(Id)
	exception.PanicIfNeeded(err)
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	_, err = repo.Collection.DeleteOne(ctx, filter)
	exception.PanicIfNeeded(err)

	return
}
