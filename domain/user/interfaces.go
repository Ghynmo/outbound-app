package user

import (
	"context"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

// Contract Interface

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserService interface {
	CreateUser(ctx context.Context, data *RegisterRequest) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context) (*[]User, error)
	FindUserByEmail(ctx context.Context, email string) (bool, error)
	UpdateUser(ctx context.Context, id string, data *UpdateRequest) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type UserRepository interface {
	Create(ctx context.Context, data *User) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) (*[]User, error)
	FindByEmail(ctx context.Context, email string) (bool, error)
	Update(ctx context.Context, id string, data *UpdateRequest) (*User, error)
	Delete(ctx context.Context, id string) error
}

type UserImageRepository interface {
	UploadPhoto(ctx context.Context, file *multipart.FileHeader) (filename, url string, err error)
	GetPhotoURL(ctx context.Context, filename string) (string, error)
	UpdatePhotoURL(ctx context.Context, filename, data string) (string, error)
	DeletePhotoURL(ctx context.Context, filename string) error
}
