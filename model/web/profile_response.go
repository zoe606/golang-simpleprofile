package web

type ProfileResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}
