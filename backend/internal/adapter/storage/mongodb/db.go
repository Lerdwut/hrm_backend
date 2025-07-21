package mysql

import (
	"context"
	"fmt"
	"hr_management/internal/adapter/config"
	"hr_management/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	DB *mongo.Database
}

var tables = []interface{}{
	&domain.User{},
	&domain.Leave{},
}

func NewDatabase(config *config.DB) (*database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// สร้าง MongoDB URI จาก config
	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", config.Host, config.Port, config.Database)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// ทดสอบการเชื่อมต่อ
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Printf("Connected to MongoDB at %s:%s, database: %s\n", config.Host, config.Port, config.Database)
	return &database{DB: client.Database(config.Database)}, nil
}

func (db *database) Close() error {
	client := db.DB.Client()
	client.Disconnect(context.Background())
	fmt.Println("Database close!")
	return nil
}
