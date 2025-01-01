package service

import (
	"context"
	"e-commerce-1/domain/user"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	// Daftar kontrak/list dari repository yg akan digunakan di service
	imageRepo user.UserImageRepository // Image database
	dbRepo    user.UserRepository      // MySQL
}

// gw butuh repo, dari situ gw bisa buat jadi struct service
// struct service berisi hal yg gw butuh untuk menjalankan fuction gw
func NewUserService(imageRepo user.UserImageRepository, dbRepo user.UserRepository) user.UserService {
	// Disini parameter repo dimasukkan kedalam struct UserService agar
	// terhubung langsung dengan function UserService
	return &UserService{
		imageRepo: imageRepo,
		dbRepo:    dbRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, data *user.RegisterRequest) (*user.User, error) {

	// Check if email already exists
	existingUser, err := s.dbRepo.FindByEmail(ctx, data.Email)
	if err == nil && existingUser {
		return nil, errors.New("email already registered")
	}

	// Hash password di service layer
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Buat user baru
	newUser := &user.User{
		Email:     data.Email,
		Password:  string(hashedPassword), // Password yang sudah di-hash
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := s.dbRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	result, err := s.dbRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) GetUsers(ctx context.Context) (*[]user.User, error) {
	result, err := s.dbRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (bool, error) {
	result, err := s.dbRepo.FindByEmail(ctx, email)
	if err != nil {
		return true, err
	}

	return result, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, data *user.UpdateRequest) (*user.User, error) {
	result, err := s.dbRepo.Update(ctx, id, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	err := s.dbRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
