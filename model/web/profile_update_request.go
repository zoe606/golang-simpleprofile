package web

type ProfileUpdateRequest struct {
	Id        int    `validate:"required"`
	Username  string `validate:"max=50" json:"username"`
	Password  string `validate:"max=255" json:"password"`
	Firstname string `validate:"max=50" json:"firstname"`
	Lastname  string `validate:"max=50" json:"lastname"`
	Age       int    `json:"age"`
	Phone     string `validate:"max=20" json:"phone"`
	Address   string `validate:"max=255" json:"address"`
	Email     string `validate:"max=50" json:"email"`
}
