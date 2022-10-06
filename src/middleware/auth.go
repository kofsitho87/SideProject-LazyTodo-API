package middleware

import (
	"gofiber-todo/utils/jwt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	log.Println(token)

	if token == "" {
		return fiber.ErrUnauthorized
	}

	// Spliting the header
	// chunks := strings.Split(token, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	// if len(chunks) < 2 {
	// 	return fiber.ErrUnauthorized
	// }

	// Verify the token which is in the chunks
	user, err := jwt.Verify(token)

	log.Println(err)

	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("USER", user.ID)

	return c.Next()
}
