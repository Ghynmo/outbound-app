package firebase

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"e-commerce-1/pkg/firebase"
)

type userImageRepository struct {
	storage *firebase.Storage
}

func NewUserImageRepository(storage *firebase.Storage) *userImageRepository {
	return &userImageRepository{
		storage: storage,
	}
}

func (r *userImageRepository) UploadPhoto(ctx context.Context, file *multipart.FileHeader) (filename, url string, err error) {
	// Buka file
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	filename = fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))

	url, err = r.storage.UploadFile(ctx, filename, src)
	if err != nil {
		return "", "", err
	}

	return filename, url, nil
}

func (r *userImageRepository) GetPhotoURL(ctx context.Context, filename string) (string, error) {
	return r.storage.GetSignedURL(ctx, filename)
}

func (r *userImageRepository) UpdatePhotoURL(ctx context.Context, filename, data string) (string, error) {
	return r.storage.GetSignedURL(ctx, filename)
}

func (r *userImageRepository) DeletePhotoURL(ctx context.Context, filename string) error {
	return r.storage.DeleteFile(ctx, filename)
}
