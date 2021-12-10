package interactors

import (
	"users_service/entities"
	"users_service/models"
	"users_service/repositories"
	"users_service/utils"
)

type CreateUserInteractorInterface interface {
	Create(model *models.CreateUserModel) *utils.Error
}

type CreateUserInteractor struct {
	Repository repositories.UserRepositoryInterface
}

func (interactor *CreateUserInteractor) Create(model *models.CreateUserModel) *utils.Error {
	var serviceError *utils.Error

	isValid, msg := model.IsValid()

	if !isValid {
		serviceError = &utils.Error{
			Code:    utils.InvalidDataErrorCode,
			Err:     nil,
			Message: msg,
		}
		return serviceError
	}

	entity := &entities.UserEntity{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Username:  model.Username,
		Email:     model.Email,
		Password:  model.Password,
	}

	_, err := interactor.Repository.Insert(entity)

	if err != nil {
		serviceError = &utils.Error{
			Code:    utils.UnknownErrorCode,
			Err:     err,
			Message: "Cannot insert into database.",
		}
		return serviceError
	}

	return serviceError
}
