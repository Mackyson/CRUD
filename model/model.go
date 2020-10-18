package model

type User struct {
	Name     string `json:"name" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
