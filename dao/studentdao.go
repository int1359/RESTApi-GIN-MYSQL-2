package dao

import (
	"fmt"
	"log"
	"rest/api/config"
	model "rest/api/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type StudentDao interface {
	GetAllStudents() ([]*model.StudentResponse, error)
	GetStudentsWithCourses() ([]*model.StudentResponse, error)
	CreateStudent(*model.Student) error
	GetStudentById(string) ([]*model.StudentResponse, error)
}

type StudentDaoImpl struct {
}

func NewStudentDaoImpl() *StudentDaoImpl {
	return &StudentDaoImpl{}
}

func (studdao *StudentDaoImpl) GetAllStudents() ([]*model.StudentResponse, error) {

	var studentlist []*model.StudentResponse
	rows, err := config.DB.Query("SELECT students.studentid,name,email,dept,dob,phoneno,COALESCE(courses.courseid,-1),COALESCE(coursename,''),COALESCE(coursefee,-1) FROM students left join enrolments on students.studentid = enrolments.studentid  left join courses on enrolments.courseid = courses.courseid")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &stud.Courses.CourseID, &stud.Courses.CourseName, &stud.Courses.CourseFee)
		if err != nil {
			return nil, err
		} else {
			studentlist = append(studentlist, &stud)

		}
	}

	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentsWithCourses() ([]*model.StudentResponse, error) {

	var studentlist []*model.StudentResponse
	rows, err := config.DB.Query("SELECT students.studentid,name,email,dept,dob,phoneno,courses.courseid,coursename,coursefee FROM students  join enrolments on students.studentid = enrolments.studentid   join courses on enrolments.courseid = courses.courseid")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &stud.Courses.CourseID, &stud.Courses.CourseName, &stud.Courses.CourseFee)
		if err != nil {
			return nil, err
		} else {
			studentlist = append(studentlist, &stud)

		}
	}

	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentById(id string) ([]*model.StudentResponse, error) {

	var studentlist []*model.StudentResponse = []*model.StudentResponse{}
	rows, err := config.DB.Query("SELECT students.studentid,name,email,dept,dob,phoneno,COALESCE(courses.courseid,-1),COALESCE(coursename,''),COALESCE(coursefee,-1) FROM students  left join enrolments on students.studentid = enrolments.studentid left  join courses on enrolments.courseid = courses.courseid where students.studentid =?", id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &stud.Courses.CourseID, &stud.Courses.CourseName, &stud.Courses.CourseFee)
		if err != nil {
			return nil, err
		} else {
			studentlist = append(studentlist, &stud)

		}
	}

	return studentlist, nil
}
func (studdao *StudentDaoImpl) CreateStudent(student *model.Student) (err error) {
	dateString := student.DOB
	date, errs := time.Parse("2006-01-02", dateString)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	courses := student.Courses

	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec("INSERT INTO students(studentid,name,email,dept,dob,phoneno)VALUES(?,?,?,?,?,?)", student.StudentID, student.Name, student.Email, student.Dept, date, student.PhoneNo)
	if err != nil {

		return err
	}

	for _, course := range courses {
		enlormentId := uuid.New()
		fmt.Print(enlormentId)
		if err != nil {

			return err
		}

		_, err = tx.Exec("INSERT INTO courses(courseid,coursename,coursefee)VALUES(?,?,?)", course.CourseID, course.CourseName, course.CourseFee)
		if err != nil {

			return err
		}
		_, err = tx.Exec("INSERT INTO enrolments(enrolmentid,studentid,courseid)VALUES(?,?,?)", enlormentId, student.StudentID, course.CourseID)
		if err != nil {

			return err
		}
	}

	return nil

}
