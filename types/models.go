package types

import "time"

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Pass     string `gorm:"not null"`
}

type Post struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Title       string     `gorm:"not null"`
	Citation    string     `gorm:"not null"`
	Date        time.Time  `gorm:"not null"`
	AdminID     uint       `gorm:"not null"`
	Category    []Category `gorm:"many2many:post_categories"`
	Summary     string     `gorm:"type:text"`
	Content     string     `gorm:"type:text"`
	Sources     []Source   `gorm:"foreignKey:PostID"`
	PhotoIcon   string     `gorm:"type:text"`
	BannerImage string     `gorm:"type:text"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

type Source struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	PostID uint   `gorm:"not null"`
	Text   string `gorm:"not null"`
	URL    string `gorm:"type:text"`
	Order  uint   `gorm:"not null"`
}

type AdminPostListView struct {
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

type UserPostListView struct {
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Citation string    `json:"citation"`
	Summary  string    `json:"summary"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique;not null"`
}
