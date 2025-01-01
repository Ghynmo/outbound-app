package firebase

import (
	"e-commerce-1/domain"
	"e-commerce-1/domain/user"
	"e-commerce-1/pkg/firebase"
)

type firebaseRepositoryFactory struct {
	storage *firebase.Storage
}

func NewFirebaseRepositoryFactory(storage *firebase.Storage) domain.RepositoryFactory {
	return &firebaseRepositoryFactory{
		storage: storage,
	}
}

func (f *firebaseRepositoryFactory) NewUserImageRepository() user.UserImageRepository {
	return NewUserImageRepository(f.storage)
}
