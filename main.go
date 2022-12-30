package main

import (
	"context"
	grpc "sprint4/grpc"
	"sprint4/redisConnection"
	"sprint4/router"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

type User struct {
	Email string `json:"email"`
}

type Join struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func tokening(c *fiber.Ctx) error {
	return c.SendString("HI!")
}

func main() {
	app := fiber.New(fiber.Config{AppName: "Test App v1.0.1"})
	var dbConnector = new(redisConnection.DBconnector)
	go grpc.GRPC()
	dbConnector.Setting()

	app.Get("/test", router.Login)
	app.Post("/join", router.Join)
	app.Get("/token", router.TokenToString)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.Send(c.Body())
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return err
		}
		value, err := dbConnector.GetHash(ctx, user.Email)
		if err != nil {
			c.SendString("failed!")
		}
		return c.SendString(value)
	})

	app.Post("/join", func(c *fiber.Ctx) error {
		join := new(Join)

		if err := c.BodyParser(join); err != nil {
			return err
		}
		dbConnector.SetHash(ctx, join.Email, join.Password)
		return c.SendString("Success!")
	})

	app.Listen(":3000")
}

// type Person struct {
// 	Name     string `json:"name" xml:"name" form:"name"`
// 	Password string `json:"password" xml:"password" form:"password"`
// }

// type PersonQ struct {
// 	Name     string   `query:"name"`
// 	Pass     string   `query:"pass"`
// 	Products []string `query:"products"`
// }

// type Cookie struct {
// 	Name        string    `json:"name"`
// 	Value       string    `json:"value"`
// 	Path        string    `json:"path"`
// 	Domain      string    `json:"domain"`
// 	MaxAge      int       `json:"max_age"`
// 	Expires     time.Time `json:"expires"`
// 	Secure      bool      `json:"secure"`
// 	HTTPOnly    bool      `json:"http_only"`
// 	SameSite    string    `json:"same_site"`
// 	SessionOnly bool      `json:"session_only"`
// }

// app.Get("/api1/:value", func(c *fiber.Ctx) error {
// 	return c.SendString("value: " + c.Params("value"))
// 	// => Get request with value: hello world
// })

// // GET http://localhost:3000/john
// app.Get("/api2/:name?", func(c *fiber.Ctx) error {
// 	if c.Params("name") != "" {
// 		return c.SendString("Hello " + c.Params("name"))
// 		// => Hello john
// 	}
// 	return c.SendString("Where is john?")
// })

// // GET http://localhost:3000/api/user/john
// app.Get("/api3/*", func(c *fiber.Ctx) error {
// 	return c.SendString("API path: " + c.Params("*"))
// 	// => API path: user/john
// })

// app.Get("/cookie", func(c *fiber.Ctx) error {
// 	cookie := new(fiber.Cookie)
// 	cookie.Name = "jin"
// 	cookie.Value = "youngmin"
// 	cookie.Expires = time.Now().Add(24 * time.Hour)
// 	c.Cookie(cookie)
// 	return c.SendString("Cookie!")
// })

// app.Get("/api4", func(c *fiber.Ctx) error {
// 	return c.SendString(c.Query("name"))
// 	// => API path: user/john
// })

// app.Get("/api5", func(c *fiber.Ctx) error {
// 	p := new(Person)

// 	if err := c.QueryParser(p); err != nil {
// 		return err
// 	}
// 	c.SendStatus(201)
// 	return c.SendString(p.Name) // john
// })
