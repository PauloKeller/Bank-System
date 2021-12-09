package interactors

import (
	"users_service/entities"
	"users_service/models"
	"users_service/repositories"
)

type CreateUserInteractorInterface interface {
	Create(model models.CreateUserModel) *models.UsersServiceError
}

type CreateUserInteractor struct {
	Repository repositories.UserRepositoryInterface
}

func (interactor *CreateUserInteractor) Create(model *models.CreateUserModel) *models.UsersServiceError {
	var serviceError *models.UsersServiceError

	entity := &entities.UserEntity{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Username:  model.Username,
		Email:     model.Email,
		Password:  model.Password,
	}

	_, err := interactor.Repository.Insert(entity)

	if err != nil {
		serviceError = &models.UsersServiceError{
			Code:    1,
			Err:     err,
			Message: "Cannot insert into database.",
		}
		return serviceError
	}

	return serviceError
}
