package model

type Group struct {
	ID        string `json:"ID" `
	Name      string `json:"Name" binding:"required" `
	Curator   string `json:"Curator" binding:"required" `
	Timetable string `json:"Timetable" `
}
