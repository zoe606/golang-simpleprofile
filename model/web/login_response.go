package web

type LoginResponse struct {
	Id    int         `json:"id"`
	Token interface{} `json:"token"`
}
