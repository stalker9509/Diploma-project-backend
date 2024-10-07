package model

type Users struct {
	ID          string `json:"id" `
	FirstName   string `json:"firstName" binding:"required" `
	LastName    string `json:"lastName" binding:"required" `
	Patronymic  string `json:"patronymic" `
	Email       string `json:"email" binding:"required" `
	Password    string `json:"password" binding:"required" `
	IsProfessor bool   `json: "isProfessor" binding:"required"`
	GroupId     string `json:"groupID" binding:"required"`
	JobTitle    string `json:"jobTitle" binding:"required"`
	ImageAvatar string `json:"imageAvatar" `
}
