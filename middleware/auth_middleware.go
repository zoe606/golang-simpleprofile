package middleware

import (
	"net/http"
	"simpleProfile/helpers"
	"simpleProfile/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "INITOKENRAHASIA" == request.Header.Get("X-API-KEY") || request.URL.Path == "/login" || request.URL.Path == "/register" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized!",
		}
		helpers.WriteToResponseBody(writer, webResponse)
	}
}
