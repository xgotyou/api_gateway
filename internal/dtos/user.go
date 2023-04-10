package dtos

import "time"

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
	Role      Role      `json:"role"`
}

type Role int8

const (
	Customer Role = iota
	Admin
	Manager
)
