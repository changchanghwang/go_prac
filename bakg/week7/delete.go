package main

import (
	applicationError "bakg.six/lib/application-error"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteEmployee(c *fiber.Ctx) error {
	employeeID := c.Params("id")

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), query)
	if err != nil {
		return applicationError.New(500, err.Error(), "Something went wrong. Please try again later.")
	}

	if result.DeletedCount < 1 {
		return applicationError.New(
			404,
			"Employee not found.",
			"Employee not found.",
		)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Employee deleted successfully"})
}
