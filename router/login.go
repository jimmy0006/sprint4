package router

import (
	"fmt"
	tokenFactory "sprint4/token"

	"github.com/gofiber/fiber/v2"
)

// using Oauth
func Login(c *fiber.Ctx) error {
	token, err := tokenFactory.CreateToken("testEmail@naver.com")
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
	return c.SendString("Join!")
}

func TokenToString(c *fiber.Ctx) error {
	fmt.Println("step1")
	value, err := tokenFactory.VerifyToken(c.Cookies("Auth"))
	if err != nil {
		fmt.Println("error at Verify token!")
		return nil
	}
	return c.SendString(value.Raw)
}
