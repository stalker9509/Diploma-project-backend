package model

type Users struct {
	ID          int    `json:"ID" `
	FirstName   string `json:"FirstName" binding:"required" `
	LastName    string `json:"LastName" binding:"required" `
	Patronymic  string `json:"Patronymic" `
	Email       string `json:"Email" binding:"required" `
	Password    string `json:"Password" binding:"required" `
	ImageAvatar string `json:"ImageAvatar" `
}
