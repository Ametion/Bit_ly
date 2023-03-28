package models

type ResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LinksResponse struct {
	OriginalLink string `json:"originalLink"`
	ShortedLink  string `json:"shortedLink"`
}

type AccountInfoResponse struct {
	Links []LinksResponse `json:"links"`
}
