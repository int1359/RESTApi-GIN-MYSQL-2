package service

import (
	"rest/api/dao"
	model "rest/api/models"
)

type StudentService interface {
	GetAllStudent() ([]*model.Student, error)
	GetStudentsWithCourses() ([]*model.Student, error)
	GetStudentById(string) ([]*model.Student, error)
	CreateStudent(*model.Student) error
}

type StudentServiceImpl struct {
	studdao dao.StudentDao
}

func NewStudentServiceImpl(dao dao.StudentDao) StudentService {
	return &StudentServiceImpl{studdao: dao}
}

func (studsvc *StudentServiceImpl) GetAllStudent() ([]*model.Student, error) {

	studentlist, err := studsvc.studdao.GetAllStudents()
	if err != nil {
		return nil, err
	}
	return studentlist, nil
}
func (studsvc *StudentServiceImpl) GetStudentsWithCourses() ([]*model.Student, error) {

	studentlist, err := studsvc.studdao.GetStudentsWithCourses()
	if err != nil {
		return nil, err
	}
	return studentlist, nil
}
func (studsvc *StudentServiceImpl) GetStudentById(id string) ([]*model.Student, error) {

	studentlist, err := studsvc.studdao.GetStudentById(id)
	if err != nil {
		return nil, err
	}
	return studentlist, nil
}
func (studsvc *StudentServiceImpl) CreateStudent(student *model.Student) error {
	err := studsvc.studdao.CreateStudent(student)
	if err != nil {
		return err

	}
	return nil

}
