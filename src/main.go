package main

import (
	"context"
	"log"
	"net"

	"etnasys.com/clientService/protogen/go/user"
	"etnasys.com/clientService/src/adapters/repository"
	"etnasys.com/clientService/src/core/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	loggerOptions := options.
		Logger().
		SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetLoggerOptions(loggerOptions)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Print(`Database connected`)
	dbName := "clientService"
	collectionName := "users"
	repo := repository.NewUserRepository(client, dbName, collectionName)

	// Setup gRPC server
	grpcServer := grpc.NewServer()
	userService := usecases.NewUserService(repo)
	user.RegisterUserServiceServer(grpcServer, userService)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Printf("Server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
