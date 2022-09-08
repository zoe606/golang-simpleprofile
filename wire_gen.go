// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"simpleProfile/app"
	"simpleProfile/controllers"
	"simpleProfile/middleware"
	"simpleProfile/repositories"
	"simpleProfile/services"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

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