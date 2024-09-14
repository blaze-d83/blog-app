package types

import	"time"

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Pass     string `gorm:"not null"`
}

type Post struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Title     string     `gorm:"not null"`
	Date      time.Time  `gorm:"not null"`
	AdminID   uint       `gorm:"not null"`
	Category  []Category `gorm:"many2many:post_categories"`
	Summary   string     `gorm:"type:text"`
	Content   Content    `gorm:"embedded"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type Content struct {
	Background string `gorm:"type:text"`
	Events     string `gorm:"type:text"`
	MainBody   string `gorm:"type:text"`
	Conclusion string `gorm:"type:text"`
	Sources    string `gorm:"type:text"`
}

type Category struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"unique;not null"`
	Posts []Post `gorm:"many2many:post_categories"`
}

