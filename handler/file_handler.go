package handler

import (
	"e-commerce-1/domain"

	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
    fileService domain.FileService
}

func NewFileHandler(fileService domain.FileService) domain.FileHandler {
    return &FileHandler{
        fileService: fileService,
    }
}

func (h *FileHandler) Upload(c *fiber.Ctx) error {
    var fileReq domain.FileUploadRequest
    // Parsing JSON request body ke dalam struct File
    if err := c.BodyParser(&fileReq); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request payload",
        })
    }

    // Parsing formFile ke dalam struct File
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Error getting file from request",
        })
    }
    
    fileReq.File = *file

    result, err := h.fileService.UploadFile(c.Context(), &fileReq)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    fileResp := domain.FileUploadResponse{
        ID: result.ID,
        Name: result.Name,
        URL: result.URL,
    }

    return c.JSON(fileResp)
}

func (h *FileHandler) GetFile(c *fiber.Ctx) error {
    return c.JSON("")

}

func (h *FileHandler) ListFiles(c *fiber.Ctx) error {
    return c.JSON("")

}

func (h *FileHandler) DeleteFile(c *fiber.Ctx) error {
    return c.JSON("")

}