package model

type Group struct {
	ID        string `json:"id" `
	Name      string `json:"name" binding:"required" `
	Curator   string `json:"curator" binding:"required" `
	Timetable string `json:"timetable" `
}
