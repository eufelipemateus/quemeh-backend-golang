package routes

import (
	"github.com/eufelipemateus/quemeh-backend-golang/controllers"
	"github.com/gofiber/fiber/v2"
	//"github.com/eufelipemateus/quemeh-backend-golang/middlewares"
)

//Setup configura las rutas de la api
func Setup(app *fiber.App) {

	app.Post("/whois", controllers.Whois)
	//app.Get("/profession", controllers.ListProfessions)

}
