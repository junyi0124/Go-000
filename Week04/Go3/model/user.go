package model

//User type
type User struct {
	id     uint64
	name   string
	gender bool
	age    int
}

func NewUser(name string, gender bool, age int) *User {
	return &User{
		id:     0,
		name:   name,
		gender: gender,
		age:    age,
	}
}

func (user *User) Id() uint64 {
	return user.id
}

func (user *User) Name() string {
	return user.name
}

func (user *User) Gender() bool {
	return user.gender
}

func (user *User) Age() int {
	return user.age
}

func (user *User) SetId(id uint64) {
	user.id = id
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) SetGender(gender bool) {
	user.gender = gender
}

func (user *User) SetAge(age int) {
	user.age = age
}
