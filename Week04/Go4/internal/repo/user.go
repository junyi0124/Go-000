package repo

import (
	"github.com/junyi0124/Go4/internal/model"
)

var (
	users []*model.User
)


type UserRepo struct{}

// Users array
func Users() []*model.User,error {
	return users,nil
}

//
func AddUser(user model.User) model.User, error {
	return nil,nil
}
