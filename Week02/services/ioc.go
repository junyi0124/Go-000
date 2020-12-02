package services

import (
	"net/http"
)

func RegisterController() {
	uc := newStudentController()

	http.Handle("/students", *uc)
	http.Handle("/students/", *uc)

}
