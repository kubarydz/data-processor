package user

import "time"

type User struct {
	Id      string
	Name    string
	Created time.Time
}

func NewUser(id, name string, tmpstmp time.Time) *User {
	return &User{Id: id, Name: name}
}
