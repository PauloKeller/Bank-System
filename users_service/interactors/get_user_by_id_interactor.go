package interactors

import (
	"users_service/entities"
	"users_service/repositories"
)

type GetUserByIDInteractorInterface interface {
	GetByID(ID string) (*entities.UserEntity, error)
}

type GetUserByIDInteractor struct {
	Repository repositories.UserRepositoryInterface
}

func (interactor *GetUserByIDInteractor) GetByID(ID string) (*entities.UserEntity, error) {
	data, err := interactor.Repository.GetByID(ID)

	if err != nil {
		return nil, err
	}

	return data, nil
}
