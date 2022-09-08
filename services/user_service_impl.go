package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"simpleProfile/exception"
	"simpleProfile/helpers"
	"simpleProfile/model/domain"
	"simpleProfile/model/web"
	"simpleProfile/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (u UserServiceImpl) Register(ctx context.Context, request web.RegisterCreateRequest) web.RegisterResponse {
	err := u.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := u.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	password := []byte(request.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		helpers.PanicIfError(err)
	}

	user := domain.Profile{
		Username:  request.Username,
		Password:  string(hashedPassword),
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Age:       request.Age,
		Phone:     request.Phone,
		Address:   request.Address,
		Email:     request.Email,
	}

	user = u.UserRepository.Register(ctx, tx, user)

	return helpers.ToRegisterResponse(user)
}

func (u UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := u.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := u.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	req := domain.Profile{
		Username: request.Username,
		Password: request.Password,
	}

	login, err := u.UserRepository.Login(ctx, tx, req)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	reqPwd := []byte(request.Password)
	userPwd := []byte(login.Password)
	err = bcrypt.CompareHashAndPassword(userPwd, reqPwd)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helpers.ToLoginResponse(login)
}

func (u UserServiceImpl) Update(ctx context.Context, request web.ProfileUpdateRequest) web.ProfileResponse {
	err := u.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := u.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	profile, err := u.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if len(request.Password) > 0 {
		password := []byte(request.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
		if err != nil {
			helpers.PanicIfError(err)
		}
		profile.Password = string(hashedPassword)
	}

	if len(request.Username) > 0 {
		profile.Username = request.Username
	}

	if len(request.Firstname) > 0 {
		profile.Firstname = request.Firstname
	}

	if len(request.Lastname) > 0 {
		profile.Lastname = request.Lastname
	}

	if request.Age > 0 {
		profile.Age = request.Age
	}

	if len(request.Phone) > 0 {
		profile.Phone = request.Phone
	}

	if len(request.Email) > 0 {
		profile.Email = request.Email
	}

	if len(request.Address) > 0 {
		profile.Address = request.Address
	}
	profile = u.UserRepository.Update(ctx, tx, profile)

	return helpers.ToProfileResponse(profile)
}

func (u UserServiceImpl) FindById(ctx context.Context, profileId int) web.ProfileResponse {
	tx, err := u.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	profile, err := u.UserRepository.FindById(ctx, tx, profileId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helpers.ToProfileResponse(profile)
}

func (u UserServiceImpl) Logout(request web.LogoutRequest) web.LogoutResponse {
	err := u.Validate.Struct(request)
	helpers.PanicIfError(err)

	return helpers.ToLogoutResponse()
}
