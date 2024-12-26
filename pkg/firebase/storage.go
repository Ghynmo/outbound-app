package firebase

import (
	"context"
	"fmt"
	"io"
	"time"

	"e-commerce-1/config"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Storage struct {
    bucket *storage.BucketHandle
}

func NewStorage(cfg *config.FirebaseConfig) (*Storage, error) {
    ctx := context.Background()

    config := &firebase.Config{
        ProjectID:     cfg.ProjectID,
        StorageBucket: cfg.BucketName,
    }

    opt := option.WithCredentialsFile(cfg.CredentialFile)
    app, err := firebase.NewApp(ctx, config, opt)
    if err != nil {
        return nil, fmt.Errorf("error initializing app: %v", err)
    }

    client, err := app.Storage(ctx)
    if err != nil {
        return nil, fmt.Errorf("error creating storage client: %v", err)
    }

    bucket, err := client.DefaultBucket()
    if err != nil {
        return nil, fmt.Errorf("error getting default bucket: %v", err)
    }

    return &Storage{
        bucket: bucket,
    }, nil
}

func (s *Storage) UploadFile(ctx context.Context, filename string, file io.Reader) (string, error) {
    obj := s.bucket.Object(filename)
    writer := obj.NewWriter(ctx)

    if _, err := io.Copy(writer, file); err != nil {
        return "", fmt.Errorf("error copying file to storage: %v", err)
    }

    if err := writer.Close(); err != nil {
        return "", fmt.Errorf("error closing writer: %v", err)
    }

    opts := &storage.SignedURLOptions{
        Scheme:  storage.SigningSchemeV4,
        Method:  "GET",
        Expires: time.Now().Add(time.Hour * 24 * 7),
    }
    
    // Generate signed URL
    return s.bucket.SignedURL(filename, opts)
}

func (s *Storage) GetSignedURL(ctx context.Context, filename string) (string, error) {
    opts := &storage.SignedURLOptions{
        Scheme:  storage.SigningSchemeV4,
        Method:  "GET",
        Expires: time.Now().Add(time.Hour * 24 * 7),
    }

    return s.bucket.SignedURL(filename, opts)
}

func (s *Storage) DeleteFile(ctx context.Context, filename string) error {
    return s.bucket.Object(filename).Delete(ctx)
}