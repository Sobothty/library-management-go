package Model

import "time"

type Book struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title"`
	IsBorrow string    `gorm:"default:false" json:"isBorrow"`
	UpdateAt time.Time `json:"updateAt"`
}
