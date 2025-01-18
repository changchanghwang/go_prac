package main

import (
	"context"
	"log"
	"time"

	"bakg.six/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id, omitempty"` // omitempty: 필드가 비어있으면 필드를 생략하는 옵션
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Connect to MongoDB
func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //30초 timeout context 생성 timeout 설정하지 않을 시 무한대기를 할 수 있음.
	// 함수 종료 시 컨텍스트 취소. -> 연결후에는 더이상 이 컨텍스트가 필요하지 않음.
	// 취소되지 않은 컨텍스트는 GC에 걸리지 않을 수 있기 때문에 cancel해주는게 좋음.
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	// MongoDB 연결 확인
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     client.Database(dbName),
	}
	return nil
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	app.Use(logger.New())

	app.Get("/employee", GetEmployees)
	app.Post("/employee", CreateEmployee)
	app.Put("/employee/:id", UpdateEmployee)
	app.Delete("/employee/:id", DeleteEmployee)

	app.Listen(":3000")
}
