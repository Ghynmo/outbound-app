package user

type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRequest struct {
	Fullname   string `json:"fullname"`
	Password   string `json:"password"`
	ProfileImg string `json:"profile_img"`
}
