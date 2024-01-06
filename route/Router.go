package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fiber-postgre/controller"
)


func Start(){

	app := fiber.New()

	api := app.Group("/api")
	usertype := api.Group("/usertype")
	users := api.Group("/users")

	usertype.Post("/",controller.AddUsertype)

	users.Post("/register",controller.Register)
	users.Post("/login",controller.Login)
	users.Post("/logout",controller.Logout)


	users.Get("/",controller.Authorize(controller.GetUsers,1))

	//users.Delete("/",controller.DeleteUser)

	app.Listen(":8182")

}