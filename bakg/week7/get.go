package main

import (
	applicationError "bakg.six/lib/application-error"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson" // Binary JSON의 줄임말로, MongoDB에서 데이터를 직렬화(Serialize)하고 역직렬화(Deserialize)하는 데 사용되는 이진(binary) 형태의 데이터 포맷
)

func GetEmployees(c *fiber.Ctx) error {
	query := bson.D{{}} // key:value 형태의 슬라이스

	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return applicationError.New(500, err.Error(), "Something went wrong. Please try again later.")
	}

	employees := make([]Employee, 0)

	if err := cursor.All(c.Context(), &employees); err != nil {
		return applicationError.New(500, err.Error(), "Something went wrong. Please try again later.")
	}

	return c.JSON(employees)
}
