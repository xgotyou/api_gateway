package dtos

import "time"

type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	BirthDate *time.Time `json:"birthDate,omitempty"`
	Role      Role       `json:"role"`
}

type Role string

const (
	Customer Role = "Customer"
	Admin    Role = "Admin"
	Manager  Role = "Manager"
)
