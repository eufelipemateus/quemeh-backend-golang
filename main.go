package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	//"github.com/eufelipemateus/quemeh-backend-golang/database"
	config "github.com/eufelipemateus/quemeh-backend-golang/configs"
	"github.com/eufelipemateus/quemeh-backend-golang/routes"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Erro on load config.toml"))
	}
	// database.OpenConnection()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(config.GetServerPort())
}
