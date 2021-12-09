package repositories

import (
	"fmt"

	"users_service/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnectionData struct {
	DBName   string
	Driver   string
	Username string
	Port     string
	Host     string
	Password string
}

type Repositories struct {
	db   *gorm.DB
	User UserRepositoryInterface
}

func NewRepositories(connectionData *PostgresConnectionData) (*Repositories, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", connectionData.Host, connectionData.Port, connectionData.Username, connectionData.DBName, connectionData.Password)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entities.UserEntity{})
}
