package routes

import (
	"e-commerce-1/domain"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App *fiber.App
	Middleware domain.Middleware
	FileHandler	domain.FileHandler
}

func NewRoutes(app *fiber.App, middleware *domain.Middleware, fileHandler *domain.FileHandler) *Routes {
	return &Routes{
		App:     app,
		Middleware: *middleware,
		FileHandler:    *fileHandler,
	}
}


func (r *Routes) SetupRoutes() {
    // Group API routes	
    api := r.App.Group("/api")

	files := api.Group("/files")

	// Route dengan middleware JWT
    protected := api.Group("/protected")

	// buat api login dengan Auth sebagai handlernya
	api.Post("/login", r.FileHandler.Upload)
	// nanti di handler auth akan memanggil generate token dari pkg middleware

	// salah, harusnya disini pengecekan token bukan auth
    protected.Use(r.Middleware.Auth())

	files.Post("/", r.FileHandler.Upload)
    files.Get("/:id", r.FileHandler.GetFile)
    files.Get("/", r.FileHandler.ListFiles)
    files.Delete("/:id", r.FileHandler.DeleteFile)

}