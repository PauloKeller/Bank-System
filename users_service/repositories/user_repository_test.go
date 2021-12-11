package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"users_service/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var repository *UserRepository
var mock sqlmock.Sqlmock

var db *sql.DB
var err error
var dialector gorm.Dialector
var gormDB *gorm.DB

func TestMain(m *testing.M) {
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer db.Close()

	dialector = postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "mysql",
	})

	gormDB, err = gorm.Open(dialector)

	// open gorm db
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository = NewUserRepository(gormDB)

	os.Exit(m.Run())
}

func TestGetAllProducts(t *testing.T) {
	var data []entities.UserEntity

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password"}).AddRow(uuid.NewString(), "paulo", "keller", "paulo", "test@gmail.com", "123456"))

	data, err = repository.GetAll()

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(data)

	fmt.Println(data)

	if len(data) < 1 {
		t.Errorf("Test failed. Should return at least one user entity.")
	}

	if data[0].Username != "paulo" {
		t.Errorf("Test failed. Wrong username")
	}

	if data[0].Email != "test@gmail.com" {
		t.Errorf("Test failed. Wrong email.")
	}

	mock.ExpectClose()

	data, err = repository.GetAll()

	if data != nil || err == nil {
		t.Errorf("Test failed. Request query fail.")
	}
}
