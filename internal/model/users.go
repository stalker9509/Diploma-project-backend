package model

type Users struct {
	ID          string `json:"ID" `
	FirstName   string `json:"FirstName" binding:"required" `
	LastName    string `json:"LastName" binding:"required" `
	Patronymic  string `json:"Patronymic" `
	Email       string `json:"Email" binding:"required" `
	Password    string `json:"Password" binding:"required" `
	IsProfessor bool   `json: "IsProfessor" binding:"required"`
	Enrollment  string `json: "Enrollment" binding:"required"`
	Graduation  string `json: "Graduation" binding:"required"`
	GroupId     string `json:"GroupID" binding:"required"`
	JobTitle    string `json:"JobTitle" binding:"required"`
	ImageAvatar string `json:"ImageAvatar" `
}
