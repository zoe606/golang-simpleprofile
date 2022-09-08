package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"simpleProfile/helpers"
	"simpleProfile/middleware"
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
