package main

import (
	"e-commerce-1/config"
	"e-commerce-1/domain/user"
	"e-commerce-1/handler"
	"e-commerce-1/middleware"
	"e-commerce-1/pkg/firebase"
	pkg_middleware "e-commerce-1/pkg/middleware"
	"e-commerce-1/pkg/mysql"
	"e-commerce-1/repository"
	repo_firebase "e-commerce-1/repository/firebase"
	"e-commerce-1/routes"
	"e-commerce-1/service"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {

	// =================================[ CONFIGURATION ]===============================

	// Mengambil configurasi dari .Env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// =================================[ INITIALIZE DB ]===============================

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

	// =================================[ ORM MIGRATION ]===============================

	// Database table migration
	if err = db.AutoMigrate(
		&user.User{},
	); err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	// =================================[ REPOSITORIES ]================================

	// Create repository factories
	firebaseFactory := repo_firebase.NewFirebaseRepositoryFactory(storage)

	// Initialize Repositories
	userImageRepository := firebaseFactory.NewUserImageRepository()
	userRepository := repository.NewUserRepository(db)

	// ===================================[ SERVICES ]==================================

	// Initialize Service with both repositories
	userService := service.NewUserService(userImageRepository, userRepository)

	// ===================================[ HANDLER ]===================================

	// Initialize Handler
	userHandler := handler.NewUserHandler(userService)

	// =====================================[ APP ]=====================================

	// Setup Fiber
	app := fiber.New()

	// =================================[ MIDDLEWARE ]==================================

	// Inisialisasi JWT Middleware
	jwtMiddleware := pkg_middleware.NewJWTMiddleware(&cfg.JWT)

	// Init middleware
	middleware := middleware.NewMiddleware(jwtMiddleware)

	// ==================================[ ROUTES ]====================================

	// Init routes
	NewRoutes := routes.NewRoutes(app, &middleware, &userHandler)
	NewRoutes.SetupRoutes()

	// ==================================[ SERVER ]====================================

	log.Fatal(app.Listen(":8080"))
}
