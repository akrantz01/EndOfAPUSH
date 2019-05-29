package database

import "github.com/jinzhu/gorm"

// Stores the user's information in a PostgreSQL table
type User struct {
	gorm.Model
	Name     string
	Username string
	Password string
	Inbox    []Message `gorm:"foreignkey:ToID"`
	Outbox   []Message `gorm:"foreignkey:FromID"`
}

// Stores a message in a PostgreSQL table
type Message struct {
	gorm.Model
	Subject   string
	Message   string
	Algorithm uint
	ToID      uint
	To        User
	FromID    uint
	From      User
}

// Stores a user authentication token in a PostgreSQL table
type Token struct {
	gorm.Model
	SigningKey string
	UserId     uint
	User       User
}
