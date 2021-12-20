package interactors

import (
	"testing"
	"users_service/dtos"
)

func TestFailToCreate(t *testing.T) {
	interactor := &CreateUserInteractor{
		repository: &MockUserRepository{},
	}

	err := interactor.Create(&dtos.CreateUserDto{})

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}

func TestCreate(t *testing.T) {
	interactor := &CreateUserInteractor{
		repository: &MockUserRepository{},
	}

	err := interactor.Create(&dtos.CreateUserDto{
		FirstName: "Paulo",
		LastName:  "Keller",
		Username:  "paulo",
		Email:     "paulo@keller.com",
		Password:  "123456",
	})

	if err != nil {
		t.Errorf("Test failed. Should create.")
	}
}

func TestFailToCreateWhenInsert(t *testing.T) {
	interactor := &CreateUserInteractor{
		repository: &MockUserRepository{},
	}

	err := interactor.Create(&dtos.CreateUserDto{
		FirstName: "fail",
		LastName:  "Keller",
		Username:  "paulo",
		Email:     "paulo@keller.com",
		Password:  "123456",
	})

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}
