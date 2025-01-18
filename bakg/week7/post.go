package main

import (
	applicationError "bakg.six/lib/application-error"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEmployee(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")
	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return applicationError.New(400, err.Error(), "Invalid request body.")
	}

	employee.ID = primitive.NewObjectID().Hex()
	insertionResult, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return applicationError.New(500, err.Error(), "Something went wrong. Please try again later.")
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &Employee{}
	createdRecord.Decode(createdEmployee)
	return c.Status(201).JSON(createdEmployee)
}
