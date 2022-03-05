package handler

import (
	"encoding/json"

	rbmq "github.com/ahmadaidil/go-rabbitmq-best-practice/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PublishUserMessage(rc *rbmq.Connection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		body, _ := json.Marshal(user)

		m := rbmq.Message{
			Queue:       "user",
			Body:        body,
			ContentType: "text/json",
		}

		if err := rc.Publish(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"messsage": "user message published",
		})
	}
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func PublishProductMessage(rc *rbmq.Connection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		product := new(Product)
		if err := c.BodyParser(product); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		body, _ := json.Marshal(product)

		m := rbmq.Message{
			Queue:       "product",
			Body:        body,
			ContentType: "application/json",
		}

		if err := rc.Publish(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"messsage": "product message published",
		})
	}
}
