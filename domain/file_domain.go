package domain

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Database Migration
type File struct {
    ID        string    `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name"`
    URL       string    `json:"url"`
    Size      int64     `json:"size"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `gorm:"default:null" json:"deleted_at"`
}

// File Upload
type FileUploadRequest struct {
    File      multipart.FileHeader `json:"file"`
}
type FileUploadResponse struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    URL       string    `json:"url"`
}

type FileHandler interface {
	Upload(c *fiber.Ctx) error
	GetFile(c *fiber.Ctx) error
	ListFiles(c *fiber.Ctx) error
	DeleteFile(c *fiber.Ctx) error
}

type FileService interface {
    UploadFile(ctx context.Context, fileReq *FileUploadRequest) (*File, error)
	GetFile(ctx context.Context, id string) (*File, error)
	ListFiles(ctx context.Context, page, limit int) (*[]File, error)
	DeleteFile(ctx context.Context, id string) error
}

type FileRepository interface {
	Create(ctx context.Context, filereq *FileUploadRequest, filename string, url string) (*File, error)
    GetByID(ctx context.Context, id string) (*File, error)
    List(ctx context.Context, limit, offset int) (*[]File, error)
    Delete(ctx context.Context, id string) error
}

type FileFirebaseRepo interface {
	Upload(ctx context.Context, file *multipart.FileHeader) (string, string, error)
	GetURL(ctx context.Context, filename string) (string, error)
	Delete(ctx context.Context, filename string) error
}