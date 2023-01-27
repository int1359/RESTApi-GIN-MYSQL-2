package controller

import (
	"fmt"
	"net/http"
	model "rest/api/models"
	"rest/api/service"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentservice service.StudentService
}

func NewStudentController(studsvc service.StudentService) *StudentController {
	return &StudentController{studentservice: studsvc}
}

func (studcont *StudentController) GetStudents(c *gin.Context) {

	student, err := studcont.studentservice.GetAllStudent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
func (studcont *StudentController) GetStudentsWithCourses(c *gin.Context) {

	student, err := studcont.studentservice.GetStudentsWithCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
func (studcont *StudentController) GetStudentById(c *gin.Context) {
	id := c.Param("id")

	student, err := studcont.studentservice.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
func (studcont *StudentController) CreateStudent(c *gin.Context) {
	var student model.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(student)
	err := studcont.studentservice.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
