package repository

import (
	"database/sql"

	"fmt"
	"strings"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/junyi0124/Go3/code"
	"github.com/junyi0124/Go3/model"
)

const (
	userName = "root"
	password = "23456789"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "mccdb"

	insertSQL    = "INSERT INTO user (`name`, `gender`, `age`) VALUES (?, ?, ?)"
	selectSQL    = "SELECT `id`,`name`,`gender`,`age` FROM `user`"
	selectOneSQL = "SELECT `name`,`gender`,`age` FROM `user` where `id`=?"
)

var (
	user  *model.User
	users []*model.User
)

// InitDB : init conn
func InitDB() *sql.DB {
	//Golang数据连接："用户名:密码@tcp(IP:端口号)/数据库名?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := sql.Open("mysql", path)
	if err != nil {
		//如果打开数据库错误，直接panic
		panic(err)
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)

	//验证连接
	if err := db.Ping(); err != nil {
		panic(err)
	}
	//将数据库连接的指针引用返回
	return db
}

// QueryUser : Query All User
func QueryUser() ([]*model.User, error) {
	db := InitDB()
	defer db.Close()

	rows, err := db.Query(selectSQL)
	if err != nil {
		fmt.Println("1st error:", err)
		return nil, errors.Wrapf(code.SqlQuery, fmt.Sprintf("sql: %s error: %v", selectSQL, err))
	}

	for rows.Next() {
		var id uint64
		var name string
		var gender bool
		var age int

		err := rows.Scan(&id, &name, &gender, &age)
		if err != nil {
			fmt.Println("2st error:", err)
			return nil, errors.Wrapf(code.ParseError, fmt.Sprintf("sql: %s error: %v", selectSQL, err))
		}
		user := model.NewUser(name, gender, age)
		user.SetId(id)
		users = append(users, user)
	}
	rows.Close()
	return users, nil
}

// Query user by id
func QueryUserById(id uint64) (*model.User, error) {
	db := InitDB()
	defer db.Close()

	state, err := db.Prepare(selectOneSQL)
	if err != nil {
		return nil, errors.Wrapf(code.SqlPrepare, fmt.Sprintf("sql: %s error: %v", insertSQL, err))
	}

	var name string
	var gender bool
	var age int

	err = state.QueryRow(id).Scan(&name, &gender, &age)
	if err != nil {
		fmt.Println("2st error:", err)
		return nil, errors.Wrapf(code.ParseError, fmt.Sprintf("sql: %s error: %v", selectSQL, err))
	}
	user := model.NewUser(name, gender, age)
	user.SetId(id)
	return user, nil
}

// AddUser to user table
func AddUser(user *model.User) (int64, error) {
	db := InitDB()
	defer db.Close()

	state, err := db.Prepare(insertSQL)
	if err != nil {
		return 0, errors.Wrapf(code.SqlPrepare, fmt.Sprintf("sql: %s error: %v", insertSQL, err))
	}

	ret, err := state.Exec(user.Name(), user.Gender(), user.Age())
	if err != nil {
		return 0, errors.Wrapf(code.SqlExec, fmt.Sprintf("sql: %s error: %v", insertSQL, err))
	}
	return ret.RowsAffected()
}
