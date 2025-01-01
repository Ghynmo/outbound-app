package handler

import (
	"e-commerce-1/domain/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(UserService user.UserService) user.UserHandler {
	return &UserHandler{
		userService: UserService,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var data user.RegisterRequest
	// Parsing JSON request body ke dalam struct User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validasi input
	if data.Email == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "All fields are required",
		})
	}

	if data.Password != data.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Passwords do not match",
		})
	}

	result, err := h.userService.CreateUser(c.Context(), &data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Buat response
	response := user.RegisterResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data: struct {
			Email string `json:"email"`
		}{
			Email: result.Email,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	return c.JSON("")

}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	return c.JSON("")

}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	return c.JSON("")

}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	return c.JSON("")

}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	return c.JSON("")

}
