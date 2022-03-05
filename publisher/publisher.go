package main

import (
	"log"

	"github.com/ahmadaidil/go-rabbitmq-best-practice/pkg/handler"
	rbmq "github.com/ahmadaidil/go-rabbitmq-best-practice/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// app server
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())

	// rbmq connection
	conn := rbmq.NewConnection(
		"my-publisher",              //name
		"my-exchange",               //exchange
		[]string{"user", "product"}, //queues
	)
	if err := conn.Connect(); err != nil {
		panic(err)
	}
	if err := conn.BindQueue(); err != nil {
		panic(err)
	}

	// app route and handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("app publisher is up!")
	})
	app.Post("/user", handler.PublishUserMessage(conn))
	app.Post("/product", handler.PublishProductMessage(conn))

	// app listen
	if err := app.Listen(":8081"); err != nil {
		log.Panicf("Failed to run app server: %s", err.Error())
	}
}
