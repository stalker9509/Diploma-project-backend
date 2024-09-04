package model

type Classes struct {
	ID           int    `json:"ID" `
	Title        string `json:"Title" binding:"required" `
	Subtitle     string `json:"Subtitle" `
	CreatorId    string `json:"CreatorId" binding:"required" `
	Description  string `json:"Description" `
	HrefTo       string `json:"HrefTo" `
	HrefLable    string `json:"HrefLable" `
	BgImageDark  string `json:"BgImageDark" `
	BgImageLight string `json:"BgImageLight" `
}
