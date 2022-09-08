//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simpleProfile/app"
	"simpleProfile/controllers"
	"simpleProfile/middleware"
	"simpleProfile/repositories"
	"simpleProfile/services"
)

func InitServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		repositories.NewUserRepository,
		services.NewUserServiceImpl,
		controllers.NewUserController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil

}
