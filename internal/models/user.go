package models

import "github.com/google/uuid"

type UserRole int

const (
	AdminRole UserRole = iota
	BasicRole
)

var roles = [...]string{
	AdminRole: "admin",
	BasicRole: "user",
}

var rolesWordToTypeMap map[string]UserRole
var rolesTypeToWord map[UserRole]string

func init() {
	rolesTypeToWord = make(map[UserRole]string)
	rolesWordToTypeMap = make(map[string]UserRole)

	for i := 0; i < len(roles); i++ {
		value := roles[i]
		userRoleIndex := UserRole(i)
		rolesWordToTypeMap[value] = userRoleIndex
		rolesTypeToWord[userRoleIndex] = value
	}
}

func LookUpRoleString(role UserRole) string {
	roleStr, ok := rolesTypeToWord[role]

	if !ok {
		return ""
	}

	return roleStr
}

func LookUpRole(roleStr string) (UserRole, bool) {
	role, ok := rolesWordToTypeMap[roleStr]

	return role, ok
}

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Role     UserRole
}

type TokenizedUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"` // for the frontend
}

type CreateUser struct {
	Username string
	Password string
	Role     UserRole
}

type LoginUser struct {
	Username string
	Password string
}
