package model

type TokenAuthBody struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}
