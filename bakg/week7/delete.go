package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteEmployee(c *fiber.Ctx) error {
	employeeID := c.Params("id")

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), query)
	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Employee deleted successfully"})
}
