package models

import "github.com/google/uuid"

type UserRole int

const (
	AdminRole UserRole = iota
	BasicRole
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Role     UserRole
}

type CreateUser struct {
	Username string
	Password string
	Role     UserRole
}
