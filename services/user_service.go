package services

import (
	"context"
	"simpleProfile/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.RegisterCreateRequest) web.RegisterResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	Update(ctx context.Context, request web.ProfileUpdateRequest) web.ProfileResponse
	FindById(ctx context.Context, profileId int) web.ProfileResponse
	Logout(request web.LogoutRequest) web.LogoutResponse
}
