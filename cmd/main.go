package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"githukudenis.com/db-test-pooling/internal/config"
	"githukudenis.com/db-test-pooling/internal/handler"
	"githukudenis.com/db-test-pooling/internal/routes"
	"githukudenis.com/db-test-pooling/internal/service"
	"githukudenis.com/db-test-pooling/internal/store"
)

func main() {
	// db := config.Connect()
	ctx := context.Background()
	pool, err := config.NewPostgresPool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	repo := store.NewPostgresUserRepository(pool)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	r := gin.Default()
	routes.RegisterUserRoutes(r, handler)
	r.Run(":8080")
}
