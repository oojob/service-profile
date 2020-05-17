package db

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database base type struct
type Database struct {
	mongo *mongo.Database
	redis *redis.Client
}

// NewMongoDB generates a new mongodb database connection
func NewMongoDB(config *Config) (*mongo.Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.DatabaseURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	oojob := client.Database("test")

	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to MongoDB database")
	}

	// Check the connection
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	return oojob, err
}

// NewRedisDB generates a new redis connection
func NewRedisDB(config *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: viper.GetString("redisuri"),
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to redis database")
	}

	return client, nil
}

// New new mongodb database instanse
func New(config *Config) (*Database, error) {
	oojobMongo, err := NewMongoDB(config)
	if err != nil {
		return nil, errors.Wrap(err, "database initialization error")
	}

	oojobRedis, err := NewRedisDB(config)
	if err != nil {
		return nil, errors.Wrap(err, "database initialization error")
	}

	return &Database{
		mongo: oojobMongo,
		redis: oojobRedis,
	}, nil
}

// Close :- close all database active connection
func (db *Database) Close() error {
	err := db.redis.Close()
	err = db.mongo.Client().Disconnect(context.Background())

	return err
}
