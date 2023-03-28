package models

type ResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
