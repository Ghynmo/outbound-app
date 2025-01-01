package routes

import (
	"e-commerce-1/domain"
	"e-commerce-1/domain/user"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App         *fiber.App
	Middleware  domain.Middleware
	UserHandler user.UserHandler
}

func NewRoutes(app *fiber.App, middleware *domain.Middleware, userHandler *user.UserHandler) *Routes {
	return &Routes{
		App:         app,
		Middleware:  *middleware,
		UserHandler: *userHandler,
	}
}

func (r *Routes) SetupRoutes() {
	// Group API routes
	api := r.App.Group("/api")

	user := api.Group("/user")

	// Route dengan middleware JWT
	protected := api.Group("/protected")

	// buat api login dengan Auth sebagai handlernya
	api.Post("/login", r.UserHandler.Login)
	// nanti di handler auth akan memanggil generate token dari pkg middleware

	// salah, harusnya disini pengecekan token bukan auth
	protected.Use(r.Middleware.Auth())

	user.Post("/", r.UserHandler.Register)
	user.Get("/:id", r.UserHandler.GetUser)
	user.Get("/", r.UserHandler.GetUsers)
	user.Put("/:id", r.UserHandler.UpdateUser)
	user.Delete("/:id", r.UserHandler.DeleteUser)

}
