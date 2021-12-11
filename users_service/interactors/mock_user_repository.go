package interactors

import (
	"fmt"
	"users_service/entities"
)

type MockUserRepository struct{}

func (m *MockUserRepository) Insert(entity *entities.UserEntity) error {
	if entity.FirstName == "fail" {
		return fmt.Errorf("some error")
	}

	return nil
}

func (m *MockUserRepository) GetByID(string) (*entities.UserEntity, error) {
	return nil, nil
}

func (m *MockUserRepository) GetAll() ([]entities.UserEntity, error) {
	return nil, nil
}
