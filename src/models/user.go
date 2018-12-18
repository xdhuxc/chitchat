package models

import "time"

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID int
	UUID string
	Email string
	UserId string
	CreatedAt time.Time
	UpdatedAt time.Time
}


