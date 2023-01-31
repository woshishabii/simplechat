package models

type Token struct {
	ID   string `json:"ID" binding:"required"`
	UUID string `json:"UUID" binding:"required,len=36"`
}
