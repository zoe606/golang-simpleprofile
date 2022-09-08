package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"simpleProfile/helpers"
	"simpleProfile/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.Response{
			Code:   http.StatusBadRequest,
			Status: "bad request",
			Data:   exception.Error(),
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.Response{
			Code:   http.StatusNotFound,
			Status: "not found",
			Data:   exception.Error,
		}
		helpers.WriteToResponseBody(w, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.Response{
		Code:   http.StatusInternalServerError,
		Status: "internal server error",
		Data:   err,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
