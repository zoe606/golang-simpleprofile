package web

type RegisterCreateRequest struct {
	Username  string `validate:"required,max=50,min=3" json:"username"`
	Password  string `validate:"required,max=255,min=8" json:"password"`
	Firstname string `validate:"required,max=50,min=1" json:"firstname"`
	Lastname  string `validate:"max=50" json:"lastname"`
	Age       int    `json:"age"`
	Phone     string `validate:"max=20" json:"phone"`
	Address   string `validate:"max=255" json:"address"`
	Email     string `validate:"max=50" json:"email"`
}
