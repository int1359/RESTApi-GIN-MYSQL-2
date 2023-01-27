package routes

import (
	"rest/api/controller"
	"rest/api/dao"
	"rest/api/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	studentDao := dao.NewStudentDaoImpl()
	studentService := service.NewStudentServiceImpl(studentDao)
	studContoller := controller.NewStudentController(studentService)

	r := gin.Default()

	r.GET("Students", studContoller.GetStudents)
	r.GET("Student", studContoller.GetStudentsWithCourses)
	r.GET("Student/:id", studContoller.GetStudentById)
	r.POST("Student", studContoller.CreateStudent)

	return r
}
