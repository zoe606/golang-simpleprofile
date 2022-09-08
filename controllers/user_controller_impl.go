package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simpleProfile/helpers"
	"simpleProfile/model/web"
	"simpleProfile/services"
	"strconv"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (u UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	regisCreateReq := web.RegisterCreateRequest{}
	helpers.ReadFormRequestBody(r, &regisCreateReq)

	registerResponse := u.UserService.Register(r.Context(), regisCreateReq)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   registerResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (u UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	loginReq := web.LoginRequest{}
	helpers.ReadFormRequestBody(r, &loginReq)

	loginRes := u.UserService.Login(r.Context(), loginReq)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   loginRes,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (u UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profileUpdateReq := web.ProfileUpdateRequest{}
	helpers.ReadFormRequestBody(r, &profileUpdateReq)

	categoryId := params.ByName("profileId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	profileUpdateReq.Id = id

	profileRes := u.UserService.Update(r.Context(), profileUpdateReq)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   profileRes,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (u UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profileId := params.ByName("profileId")
	id, err := strconv.Atoi(profileId)
	helpers.PanicIfError(err)

	profileRes := u.UserService.FindById(r.Context(), id)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   profileRes,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (u UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	logoutReq := web.LogoutRequest{}
	helpers.ReadFormRequestBody(r, &logoutReq)

	logoutRes := u.UserService.Logout(logoutReq)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   logoutRes,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
