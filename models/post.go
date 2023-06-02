package models

// Post model

type Post struct {
	BaseModel
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content"`
}
