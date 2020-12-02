package main

import (
	"errors"
	pkg_errs "github.com/pkg/errors"
	"fmt"
	"net/http"

	"github.com/junyi0124/Go-000/Week02/services"
)

func main() {
	err := errors.New("this is error")
	fmt.Printf("%T %v \n", err, err)
	fmt.Println("hello world")

	// Wrap 会带有调用栈信息
	err1 := pkg_errs.Wrap(err, "this is wrapped!")
	pkg_err := pkg_errs.WithMessage(err1, "this is with")
	fmt.Printf("print err, type: %T value:%+v", pkg_err, pkg_err)

	services.RegisterController()
	http.ListenAndServe(":3000", nil)
}
