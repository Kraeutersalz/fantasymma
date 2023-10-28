package controller

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/kreaeutersalz/fantasymma/managers"
	"github.com/kreaeutersalz/fantasymma/models"
)

func accountInit(app *fiber.App) {
	app.Get("/account/:id", GetAccount)
	app.Post("/account", CreateAccount)
}

func GetAccount(c *fiber.Ctx) {
	// Later use Authenticaten Token to get the account
	fmt.Println(c.Params("id"))
	err, account := managers.GetAccount(c.Params("id"))
	if err != nil {
		fmt.Println(err)
		c.Status(404).SendString("Account not found")
		return
	}
	c.JSON(account)
}

func CreateAccount(c *fiber.Ctx) {
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		c.Status(400).SendString("Bad Request")
		return
	}
	err, account := managers.CreateAccount(account)
	if err != nil {
		c.Status(400).SendString("Bad Request")
		return
	}
	c.JSON(account)
}
