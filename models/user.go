package models

// User model

type User struct {
	BaseModel
	Name  string `json:"name"`
	Phone string `json:"phone" gorm:"size:11;not null"`
}
