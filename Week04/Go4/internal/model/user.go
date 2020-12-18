package model

// User type
type User struct {
	id     uint64
	name   string
	gender bool
	age    uint8
}

// NewUser : User ctor
func NewUser(id uint64, name string, gender bool, age uint8) *User {
	return &User{
		id:     id,
		name:   name,
		gender: gender,
		age:    age,
	}
}
