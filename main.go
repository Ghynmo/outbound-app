package main

import (
	"e-commerce-1/config"
	"e-commerce-1/domain"
	"e-commerce-1/handler"
	"e-commerce-1/middleware"
	"e-commerce-1/pkg/firebase"
	pkg_middleware "e-commerce-1/pkg/middleware"
	"e-commerce-1/pkg/mysql"
	"e-commerce-1/repository"
	"e-commerce-1/routes"
	"e-commerce-1/service"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
    cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Inisialisasi MySQL
	db, err := mysql.NewConnection(&cfg.MySQL)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

    // Inisialisasi Firebase
	storage, err := firebase.NewStorage(&cfg.Firebase)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

    // Database table migration
    if err = db.AutoMigrate(
        &domain.File{},
    ); err != nil {log.Fatal("Failed to auto-migrate:", err) 
    }


    // Initialize Repositories
    fileStorageRepo := repository.NewFileFirebaseRepo(storage)
    fileDBRepo := repository.NewFileRepository(db)

    // Initialize Service with both repositories
    fileService := service.NewFileService(fileStorageRepo, fileDBRepo)

    // Initialize Handler
    fileHandler := handler.NewFileHandler(fileService)

    
    // Setup Fiber
    app := fiber.New()
    
	// Inisialisasi JWT Middleware
    jwtMiddleware := pkg_middleware.NewJWTMiddleware(&cfg.JWT)

    // Init middleware
    middleware := middleware.NewMiddleware(jwtMiddleware)
    
    // Init routes
    NewRoutes := routes.NewRoutes(app, &middleware, &fileHandler)
    NewRoutes.SetupRoutes()

    log.Fatal(app.Listen(":8080"))    
}