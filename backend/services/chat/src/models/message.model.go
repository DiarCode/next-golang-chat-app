package models

type Message struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	SenderId int64  `json:"senderId"`
	Content  string `json:"content"`
	RoomId   int64  `json:"roomId"`
}

// string room_id = 1;
// string user_id = 2;
// string content = 3;
