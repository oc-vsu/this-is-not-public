package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository is the interface, which defines all our method signatures that can be mocked.
type Repository interface {
	Insert(x interface{}) error
	FindOne(id string) (interface{}, error)
	Find() ([]interface{}, error)
}

// MongoDB is connected to an MongoDB collection.
// It includes all methods, defined by the Repository interface.
type MongoDB struct {
	col *mongo.Collection
}

func New(col *mongo.Collection) *MongoDB {
	return &MongoDB{
		col: col,
	}
}

func (m *MongoDB) Insert(x interface{}) error {
	_, err := m.col.InsertOne(context.Background(), x)

	return err
}

func (m *MongoDB) FindOne(id string) (interface{}, error) {
	res := m.col.FindOne(context.Background(), bson.M{"id": id})

	var model interface{}
	if err := res.Decode(&model); err != nil {
		return nil, err
	}
	return model, nil
}

func (m *MongoDB) Find() ([]interface{}, error) {
	cur, err := m.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var models []interface{}
	if err = cur.All(context.Background(), &models); err != nil {
		return nil, err
	}

	return models, nil
}
