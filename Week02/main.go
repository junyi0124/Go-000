package main

import (
	"fmt"
	"net/http"
	"services"
	"models"
)


func main() { 
	fmt.Println("hello world")
	services.RegisterController()
	http.ListenAndServe(":3000", nil)
}