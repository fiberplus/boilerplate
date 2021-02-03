package models

// User
type User struct {
	ID int `gorm:"primaryKey"`
	Name string
}
