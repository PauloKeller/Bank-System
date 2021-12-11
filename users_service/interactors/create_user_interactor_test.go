package interactors

import (
	"os"
	"testing"
	"users_service/dtos"
)

var interactor CreateUserInteractorInterface
var err error

func TestMain(m *testing.M) {
	interactor = &CreateUserInteractor{
		repository: &MockUserRepository{},
	}

	os.Exit(m.Run())
}

func TestFailToCreate(t *testing.T) {
	err := interactor.Create(&dtos.CreateUserDto{})

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}

func TestCreate(t *testing.T) {
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
