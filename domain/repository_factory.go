package domain

import "e-commerce-1/domain/user"

// Contract Interface
type RepositoryFactory interface {
	NewUserImageRepository() user.UserImageRepository
}
