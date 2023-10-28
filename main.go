package main

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/kreaeutersalz/fantasymma/controller"
	"github.com/kreaeutersalz/fantasymma/database"
)

func main() {
	//init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	//defer close database
	defer database.CloseMongoDB()

	// generate app
	app := fiber.New()
	controller.Init(app)

	//get port from ENV
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func initApp() error {
	//setup ENV
	err := loadENV()
	if err != nil {
		return err
	}

	//setup Database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
