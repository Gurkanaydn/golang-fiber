package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type employee struct {
	Id        string `json:"id" xml:"id" form:"id" query:"id"`
	FirstName string `json:"firstname" xml:"firstname" form:"firstname" query:"firstname"`
	Lastame   string `json:"lastname" xml:"lastname" form:"lastname" query:"lastname"`
	Age       int    `json:"age" xml:"age" form:"age" query:"age"`
	Ismarried bool   `json:"ismarried" xml:"ismarried" form:"ismarried" query:"age"`
}

var employees = []employee{
	{Id: "1", FirstName: "Arthur", Lastame: "Oliver", Age: 19, Ismarried: false},
	{Id: "2", FirstName: "Jack", Lastame: "Ethan", Age: 19, Ismarried: false},
	{Id: "3", FirstName: "Harry", Lastame: "Alexander", Age: 20, Ismarried: false},
	{Id: "4", FirstName: "Charlie", Lastame: "Daniel", Age: 18, Ismarried: false},
}

func main() {
	app := fiber.New()
	app.Get("/employees", func(c *fiber.Ctx) error {
		return c.JSON(employees)
	})
	app.Get("/employees/:id", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("id"))
		for i, e := range employees {
			if e.Id == c.Params("id") {
				return c.JSON(&employees[i])
			}
		}
		return c.JSON(
			fiber.Map{
				"message": "There is no employee with have such id",
			})
	})

	app.Post("/employees", func(c *fiber.Ctx) error {
		e := new(employee)
		if err := c.BodyParser(e); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		employees = append(employees, *e)
		return c.JSON(
			fiber.Map{
				"message": "Succes",
			})

	})

	app.Listen(":4400")
}
