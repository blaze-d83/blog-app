package types

import "time"

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Pass     string `gorm:"not null"`
}


type AdminPostListView struct {
	Title     string    `json:"title"`
	Date      time.Time   `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}
