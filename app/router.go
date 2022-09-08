package app

import (
	"github.com/julienschmidt/httprouter"
	"simpleProfile/controllers"
	"simpleProfile/exception"
)

func NewRouter(userController controllers.UserController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.POST("/logout", userController.Logout)
	router.GET("/profile/:profileId", userController.FindById)
	router.PUT("/profile/:profileId", userController.Update)
	router.PanicHandler = exception.ErrorHandler
	return router
}
