package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaginatedUserList struct {
	Items  []User
	Total  int
	Limit  int
	Offset int
}
