package helpers

import (
	"simpleProfile/model/domain"
	"simpleProfile/model/web"
)

func ToRegisterResponse(user domain.Profile) web.RegisterResponse {
	return web.RegisterResponse{
		Message: " Berhasil Registrasi dengan user name " + user.Username,
	}
}

func ToLogoutResponse() web.LogoutResponse {
	return web.LogoutResponse{
		Message: " Berhasil Logout!",
	}
}

func ToLoginResponse(user domain.Profile) web.LoginResponse {
	return web.LoginResponse{
		Id:    user.Id,
		Token: "INITOKENRAHASIA",
	}
}

func ToProfileResponse(user domain.Profile) web.ProfileResponse {
	return web.ProfileResponse{
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Phone:     user.Phone,
		Address:   user.Address,
		Email:     user.Email,
	}
}
