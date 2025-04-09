package models

type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Image     string
	CreatedAt string
	RoleId    uint
}
