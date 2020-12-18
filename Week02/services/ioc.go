package services

import (
	"net/http"
)

func RegisterController() {
	uc, _ := newStudentController()

	http.Handle("/students", *uc)
	http.Handle("/students/", *uc)

}
