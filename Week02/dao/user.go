package dao

import (
	"github.com/junyi0124/Go-000/Week02/models"
	//"fmt"
	"github.com/pkg/errors"
	"database/sql"
	// 
	_ "github.com/go-sql-driver/mysql"
)

type StudentDao struct {
}

func newStudnentDao() (d *StudentDao, err error) {
	return &StudentDao{}, nil
}

func GetStudent(id int) (models.Student, error) {
	// mysql connection string
	dsn := "www:2345678()@tcp(172.31.75.15:3306)/go-demo?charset=utf8mb4&parseTime=True&loc=Local"
	// open connection
	db, openErr := sql.Open("mysql", dsn)
	student := models.Student{}
	if openErr != nil {
		newErr := errors.Wrap(openErr, "this is wrapped!")
		return student, newErr
	}

	queryErr := db.QueryRow("SELECT id,`name`,gender FROM student", 1).Scan(&student.ID, &student.Name, &student.Gender)
	if queryErr != nil {
		newErr := errors.Wrap(queryErr, "this is wrapped!")
		return student, newErr
	}
	
	return student, nil
}

// func AddUser(u User) (User, error) {
// 	u.ID = nextID
// 	nextID++
// 	users = append(users, &u)
// 	return u, nil
// }
