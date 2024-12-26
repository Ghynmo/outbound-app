package domain

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	Auth() fiber.Handler
	Logger() fiber.Handler
	Recover() fiber.Handler
}