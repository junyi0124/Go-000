package services

import (
	"net/http"
	"regexp"

	"github.com/junyi0124/Go-000/Week02/dao"
)

type StudentController struct {
	studentIDPattern *regexp.Regexp
	studentDao *dao.StudentDao
}

func newStudentController() (s *StudentController, err error) {
	*s.studentDao, _ = dao.newStudnentDao()
	return &StudentController{
		studentIDPattern: regexp.MustCompile(`^/students/(\d+)/?`),
	}, nil
}



func (uc StudentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte("hello from User controller!!!"))
}

