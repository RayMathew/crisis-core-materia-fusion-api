package crisiscoremateriafusion

type User struct {
	Name   string
	Exists bool
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}
