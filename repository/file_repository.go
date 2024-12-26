package repository

import (
	"context"
	"e-commerce-1/domain"

	"gorm.io/gorm"
)

type fileRepository struct {
    db *gorm.DB
}

func NewFileRepository(db *gorm.DB) domain.FileRepository {
    return &fileRepository{
        db: db,
    }
}

func (r *fileRepository) Create(ctx context.Context, fileReq *domain.FileUploadRequest, filename string, url string) (*domain.File, error) {
    // Mengisi field yg akan dikirim ke database
    file := domain.File{
        Name: filename,
        URL: url,
        Size: fileReq.File.Size,
    }

    // Query
    result := r.db.WithContext(ctx).Model(&file).Create(&file)
    if result.Error != nil {
        return nil, result.Error
    }

    return &file, nil
}

func (r *fileRepository) GetByID(ctx context.Context, id string) (*domain.File, error) {
    var fileModel domain.File
    if err := r.db.WithContext(ctx).First(&fileModel, "id = ?", id).Error; err != nil {
        return nil, err
    }

    return &domain.File{
        ID:        fileModel.ID,
        Name:      fileModel.Name,
        URL:       fileModel.URL,
        Size:      fileModel.Size,
        CreatedAt: fileModel.CreatedAt,
    }, nil
}

func (r *fileRepository) List(ctx context.Context, limit, offset int) (*[]domain.File, error) {
    var fileModels []domain.File
    if err := r.db.WithContext(ctx).
        Limit(limit).
        Offset(offset).
        Order("created_at DESC").
        Find(&fileModels).Error; err != nil {
        return nil, err
    }

    files := make([]domain.File, len(fileModels))
    for i, fmodel := range fileModels {
        files[i] = domain.File{
            ID:        fmodel.ID,
            Name:      fmodel.Name,
            URL:       fmodel.URL,
            Size:      fmodel.Size,
            CreatedAt: fmodel.CreatedAt,
        }
    }

    return &files, nil
}

func (r *fileRepository) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&domain.File{}, "id = ?", id).Error
}