package main

import (
	"errors"
	"fmt"

	"github.com/junyi0124/Go3/code"
	"github.com/junyi0124/Go3/repository"
	"github.com/junyi0124/Go3/model"
)

func main() {

	users, err := repository.QueryUser()
	if err != nil {
		if errors.Is(err, code.SqlQuery) {
			
		}
	}
	if(users!=nil && users[0]!=nil) {
		fmt.Println(users[0].Id(), users[0].Name(), users[0].Gender(), users[0].Age() )
	}

	user := model.NewUser("zhangxinyun", true, 6)
	count, err := repository.AddUser(user)
	fmt.Println(count, err)
}
