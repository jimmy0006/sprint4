package router

import (
	"context"
	"fmt"
	redisConnection "sprint4/redisConnection"
	tokenizer "sprint4/token"

	"github.com/gofiber/fiber/v2"
)

type JoinRequest struct {
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

var ctx = context.Background()

// using Oauth
func Login(c *fiber.Ctx) error {
	token, err := tokenizer.CreateToken("testEmail@naver.com")
	if err != nil {
		return nil
	}
	c.Cookie(&fiber.Cookie{
		Name:  "Auth",
		Value: token,
	})
	return c.SendString("HI!")
}

func Join(c *fiber.Ctx) error {
	request := new(JoinRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	var dbConnector = new(redisConnection.DBconnector)
	fmt.Println(request.Email)
	fmt.Println(request.Name)
	dbConnector.SetHash(ctx, request.Email, request.Name, 12)
	return c.SendString("Joined!")
}

func TokenToString(c *fiber.Ctx) error {
	fmt.Println("step1")
	value, err := tokenizer.ExtractTokenEmail(c.Cookies("Auth"))
	if err != nil {
		fmt.Println("error at Verify token!")
		return nil
	}
	return c.SendString(value.Email)
}
