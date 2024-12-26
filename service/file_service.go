package service

import (
	"context"
	"e-commerce-1/domain"
)

type fileService struct {
    // Daftar kontrak/list dari repository yg akan digunakan di service
    storageRepo domain.FileFirebaseRepo  // Firebase Storage
    dbRepo      domain.FileRepository // MySQL
}

// gw butuh repo, dari situ gw bisa buat jadi struct service
// struct service berisi hal yg gw butuh untuk menjalankan fuction gw
func NewFileService(storageRepo domain.FileFirebaseRepo, dbRepo domain.FileRepository) domain.FileService {
    // Disini parameter repo dimasukkan kedalam struct fileService agar
    // terhubung langsung dengan function fileService
    return &fileService{
        storageRepo: storageRepo,
        dbRepo:      dbRepo,
    }
}

func (s *fileService) UploadFile(ctx context.Context, fileReq *domain.FileUploadRequest) (*domain.File, error) {
    // Upload ke Firebase
    filename, url, err := s.storageRepo.Upload(ctx, &fileReq.File)
    if err != nil {
        return nil, err
    }

    // Simpan metadata ke MySQL
    result, err := s.dbRepo.Create(ctx, fileReq, filename, url)
    if err != nil {
        // Jika gagal menyimpan ke database, hapus file dari storage
        _ = s.storageRepo.Delete(ctx, filename)
        return nil, err
    }

    return result, nil
}

func (s *fileService) GetFile(ctx context.Context, id string) (*domain.File, error) {
    return s.dbRepo.GetByID(ctx, id)
}

func (s *fileService) ListFiles(ctx context.Context, page, limit int) (*[]domain.File, error) {
    offset := (page - 1) * limit
    return s.dbRepo.List(ctx, limit, offset)
}

func (s *fileService) DeleteFile(ctx context.Context, id string) error {
    // Hapus dari storage
    if err := s.storageRepo.Delete(ctx, id); err != nil {
        return err
    }

    // Hapus dari database
    return s.dbRepo.Delete(ctx, id)
}