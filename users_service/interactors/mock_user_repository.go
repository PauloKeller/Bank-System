package interactors

import (
	"fmt"
	"time"
	"users_service/entities"

	"github.com/google/uuid"
)

type MockUserRepository struct {
	ShouldFailGetAll bool
}

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
	if m.ShouldFailGetAll {
		return nil, fmt.Errorf("some error")
	}

	id, _ := uuid.NewUUID()

	var data = []entities.UserEntity{
		{
			FirstName: "Paulo",
			LastName:  "Keller",
			ID:        id,
			Username:  "paulokeller",
			Email:     "paulo@keller.com",
			Password:  "123456",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return data, nil
}
