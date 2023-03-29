package models

type CreateLinkRequest struct {
	OriginalLink string `json:"originalLink"`
}

type AccountInfoRequest struct {
	Login string `json:"login"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}
