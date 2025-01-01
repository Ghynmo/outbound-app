package repository

import (
	"context"
	"e-commerce-1/domain/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, data *user.User) (*user.User, error) {
	var user user.User
	// Query
	result := r.db.WithContext(ctx).Model(&user).Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	var user user.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) (*[]user.User, error) {
	var user []user.User
	if err := r.db.WithContext(ctx).
		// Limit(limit).
		// Offset(offset).
		Order("created_at DESC").
		Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (bool, error) {
	var user user.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepository) Update(ctx context.Context, id string, data *user.UpdateRequest) (*user.User, error) {
	var user user.User
	result := r.db.WithContext(ctx).Model(&user).Update(id, &data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	var user user.User
	return r.db.WithContext(ctx).Delete(&user, "id = ?", id).Error
}
