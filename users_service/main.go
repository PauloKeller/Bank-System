package main

import (
	"log"
	"net"
	"os"

	"users_service/handler"
	"users_service/interactors"
	"users_service/repositories"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	core_grpc_api "users_service/proto"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbdriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	serverPort := ":" + os.Getenv("SERVER_PORT")

	dbData := &repositories.PostgresConnectionData{
		DBName:   dbName,
		Port:     dbPort,
		Password: dbPassword,
		Username: dbUser,
		Driver:   dbdriver,
		Host:     dbHost,
	}

	repositories, err := repositories.NewRepositories(dbData)

	repositories.Automigrate()

	handler := &handler.Handler{
		CreateUserInteractor: interactors.NewCreateUserInteractor(repositories.User),
		GetAllUserInteractor: &interactors.GetAllUserInteractor{
			Repository: repositories.User,
		},
		GetUserByIDInteractor: &interactors.GetUserByIDInteractor{
			Repository: repositories.User,
		},
	}

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	core_grpc_api.RegisterUsersServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	log.Println("Starting server on port", serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
