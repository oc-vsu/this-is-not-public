package main

import (
	"log"

	"opitz-consulting.com/examples/unittesting/mocking/internal/app/mocking"
	"opitz-consulting.com/examples/unittesting/mocking/internal/app/mocking/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	c, err := mongo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	db := c.Database("database")
	mongoRepo := repository.New(db.Collection("collection"))

	// create new logic struct with mongodb repository
	logic := mocking.NewLogic(mongoRepo)

	err = logic.DoSomething()
	if err != nil {
		log.Fatal(err)
	}
}
