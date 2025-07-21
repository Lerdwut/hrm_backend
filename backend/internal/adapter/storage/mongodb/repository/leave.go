package repository

import (
	"context"
	"hr_management/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoLeaveRepo struct {
	collection *mongo.Collection
}

func NewMongoLeaveRepo(db *mongo.Database) *MongoLeaveRepo {
	return &MongoLeaveRepo{collection: db.Collection("leaves")}
}

func (r *MongoLeaveRepo) Create(l *domain.Leave) error {
	// Set default values
	if l.CreatedAt.IsZero() {
		l.CreatedAt = time.Now()
	}
	if l.Status == "" {
		l.Status = domain.Pending
	}

	_, err := r.collection.InsertOne(context.Background(), l)
	return err
}

func (r *MongoLeaveRepo) GetAll() ([]domain.Leave, error) {
	var leaves []domain.Leave
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var leave domain.Leave
		if err := cursor.Decode(&leave); err != nil {
			return nil, err
		}
		leaves = append(leaves, leave)
	}
	return leaves, nil
}

func (r *MongoLeaveRepo) GetByID(id uint) (*domain.Leave, error) {
	var leave domain.Leave
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&leave)
	if err != nil {
		return nil, err
	}
	return &leave, nil
}

func (r *MongoLeaveRepo) Update(id uint, status domain.LeaveStatus) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": bson.M{"status": status}})
	return err
}
