package app

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-mongodb/exception"
	"strconv"
	"time"
)

const MONGO_URI = "mongodb://mongo:mongo@localhost:27017"
const MONGO_DATABASE = "belajar"
const MONGO_POOL_MIN = "10"
const MONGO_POOL_MAX = "100"
const MONGO_MAX_IDLE_TIME_SECOND = "60"

func NewMongoDatabase() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(MONGO_POOL_MIN)
	exception.PanicIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(MONGO_POOL_MAX)
	exception.PanicIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(MONGO_MAX_IDLE_TIME_SECOND)
	exception.PanicIfNeeded(err)

	option := options.Client().
		ApplyURI(MONGO_URI).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(MONGO_DATABASE)
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
