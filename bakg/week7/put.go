package main

import (
	applicationError "bakg.six/lib/application-error"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateEmployee(c *fiber.Ctx) error {
	idParam := c.Params("id")
	employee := new(Employee)
	if err := c.BodyParser(employee); err != nil {
		return applicationError.New(400, err.Error(), "Invalid request body.")
	}

	query := bson.D{{Key: "_id", Value: idParam}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err := mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return applicationError.New(404, err.Error(), "Employee not found.")
		}

		return applicationError.New(500, err.Error(), "Something went wrong. Please try again later.")
	}

	employee.ID = idParam
	return c.Status(200).JSON(employee)
}
