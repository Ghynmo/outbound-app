package middleware

import (
	"e-commerce-1/domain"
	pkg_middleware "e-commerce-1/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type middleware struct {
    jwtService *pkg_middleware.JWTMiddleware
}

func NewMiddleware(jwt *pkg_middleware.JWTMiddleware) domain.Middleware {
    return &middleware{
        jwtService: jwt,
    }
}

// Penamaannya bukan auth tapi checking token
func (m *middleware) Auth() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized",
            })
        }
        return c.Next()
    }
}

func (m *middleware) Logger() fiber.Handler {
    return logger.New()
}

func (m *middleware) Recover() fiber.Handler {
    return recover.New()
}
