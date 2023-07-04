package utils

type RegisterResponse struct {
	Status  string          `json:"status"`
	Payload RegisterPayload `json:"payload"`
}

type RegisterPayload struct {
	Id       string `json:"id"`
	Username string `json:username`
	Email    string `json:email`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  string       `json:"status"`
	Payload LoginPayload `json:"data"`
}

type LoginPayload struct {
	Id       string `json:"id"`
	Username string `json:username`
	Email    string `json:email`
}
