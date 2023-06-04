package models

type Room struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Messages []Message
}
