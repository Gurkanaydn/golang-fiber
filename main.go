package main

import (
	"github.com/gofiber/fiber/v2"
)

type employee struct {
	Lastame   string `json:"lastname" xml:"lastname" form:"lastname" query:"lastname"`
	FirstName string `json:"firstname" xml:"firstname" form:"firstname" query:"firstname"`
	Age       int    `json:"age" xml:"age" form:"age" query:"age"`
	Ismarried bool   `json:"ismarried" xml:"ismarried" form:"ismarried" query:"age"`
}

var employees = []employee{
	{FirstName: "Arthur", Lastame: "Oliver", Age: 19, Ismarried: false},
	{FirstName: "Jack", Lastame: "Ethan", Age: 19, Ismarried: false},
	{FirstName: "Harry", Lastame: "Alexander", Age: 20, Ismarried: false},
	{FirstName: "Charlie", Lastame: "Daniel", Age: 18, Ismarried: false},
}

func main() {
	app := fiber.New()
	app.Get("/employees", func(c *fiber.Ctx) error {
		return c.JSON(employees)
	})

	app.Post("/employees", func(c *fiber.Ctx) error {
		e := new(employee)
		if err := c.BodyParser(e); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		employees = append(employees, *e)
		return c.JSON("Succes")

	})

	app.Listen(":4400")
}
