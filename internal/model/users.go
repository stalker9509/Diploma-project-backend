package model

type Users struct {
	ID          int    `json:"ID" `
	FirstName   string `json:"FirstName" binding:"required" `
	LastName    string `json:"LastName" binding:"required" `
	Patronymic  string `json:"Patronymic" `
	Email       string `json:"Email" binding:"required" `
	Password    string `json:"Password" binding:"required" `
	IsProfessor bool   `json: "isProfessor" binding:"required"`
	Enrollment  string `json: "enrollment" binding:"required"`
	Graduation  string `json: "graduation" binding:"required"`
	GroupId     string `json:"GroupID" binding:"required"`
	JobTitle    string `json:"JobTitle" binding:"required"`
	ImageAvatar string `json:"ImageAvatar" `
}
