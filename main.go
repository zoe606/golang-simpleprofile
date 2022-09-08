package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"simpleProfile/app"
	"simpleProfile/controllers"
	"simpleProfile/helpers"
	"simpleProfile/middleware"
	"simpleProfile/repositories"
	"simpleProfile/services"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitServer()
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}

func InitServer() *http.Server {
	userRepository := repositories.NewUserRepository()
	db := app.NewDB()
	validate := validator.New()
	userService := services.NewUserServiceImpl(userRepository, db, validate)
	userController := controllers.NewUserController(userService)
	router := app.NewRouter(userController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}
