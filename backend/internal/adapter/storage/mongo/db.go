package mongo

import (
	"context"
	"hr_management/internal/core/domain"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func ConnectMongoDB() (*Database, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, domain.ErrMongoURI
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		return nil, domain.ErrDatabaseName
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")

	return &Database{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (db *Database) GetCollection(collectionName string) *mongo.Collection {
	return db.Database.Collection(collectionName)
}

func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}
