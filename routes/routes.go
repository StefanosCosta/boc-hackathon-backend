package routes

import (
	"OpenBankingAPI/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/businessCustomers", controllers.GetBusinessCustomers)
	app.Post("/api/requestAccess", controllers.RequestAccess)
	app.Post("/api/provideAccess", controllers.ProvideAccess)
	app.Get("/api/accessRequests", controllers.GetAuditorsAccessRequests)
	app.Get("/api/balances/:id", controllers.GetBalances)

}
