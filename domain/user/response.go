package user

type RegisterResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Email string `json:"email"`
	} `json:"data"`
}
