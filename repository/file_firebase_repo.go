package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"e-commerce-1/domain"
	"e-commerce-1/pkg/firebase"
)

type firebaseRepository struct {
    storage *firebase.Storage
}

func NewFileFirebaseRepo(storage *firebase.Storage) domain.FileFirebaseRepo {
    return &firebaseRepository{
        storage: storage,
    }
}

func (r *firebaseRepository) Upload(ctx context.Context, file *multipart.FileHeader) (filename string, url string, err error) {
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

func (r *firebaseRepository) GetURL(ctx context.Context, filename string) (string, error) {
    return r.storage.GetSignedURL(ctx, filename)
}

func (r *firebaseRepository) Delete(ctx context.Context, filename string) error {
    return r.storage.DeleteFile(ctx, filename)
}