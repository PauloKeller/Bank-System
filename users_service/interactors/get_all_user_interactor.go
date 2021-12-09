package interactors

import (
	"users_service/entities"
	"users_service/repositories"
)

type GetAllUserInteractorInterface interface {
	GetAll() ([]entities.UserEntity, error)
}

type GetAllUserInteractor struct {
	Repository repositories.UserRepositoryInterface
}

func (interactor *GetAllUserInteractor) GetAll() ([]entities.UserEntity, error) {
	data, err := interactor.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}
