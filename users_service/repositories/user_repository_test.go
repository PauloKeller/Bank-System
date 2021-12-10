package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

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

	repository = &UserRepository{
		db: gormDB,
	}
}

func TestGetAllProducts(t *testing.T) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.NewString()))

	data, err := repository.GetAll()

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(data)

	if len(data) < 1 {
		t.Errorf("Test failed. Wrong Quantity.")
	}
}
