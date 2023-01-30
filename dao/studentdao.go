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
	GetAllStudents() ([]*model.Student, error)
	GetStudentsWithCourses() ([]*model.Student, error)
	CreateStudent(*model.Student) error
	GetStudentById(string) ([]*model.Student, error)
}

type StudentDaoImpl struct {
}

func NewStudentDaoImpl() *StudentDaoImpl {
	return &StudentDaoImpl{}
}

func (studdao *StudentDaoImpl) GetAllStudents() ([]*model.Student, error) {

	var studentlist []*model.Student
	rows, err := config.DB.Query("SELECT students.studentid,name,email,dept,dob,phoneno,COALESCE(courses.courseid,-1),COALESCE(coursename,''),COALESCE(coursefee,-1) FROM students left join enrolments on students.studentid = enrolments.studentid  left join courses on enrolments.courseid = courses.courseid")

	if err != nil {
		log.Print(err)
	}

	coursemap := make(map[int32][]model.Course)
	studmap := make(map[int32]model.StudentMap)

	for rows.Next() {
		var stud model.StudentMap
		var studCourse model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &studCourse.CourseID, &studCourse.CourseName, &studCourse.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if studCourse.CourseID != -1 {
				coursemap[stud.StudentID] = append(coursemap[stud.StudentID], studCourse)
			}

			studmap[stud.StudentID] = stud

		}
	}
	for k, v := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: v.StudentID, Name: v.Name, Email: v.Email, Dept: v.Dept, DOB: v.DOB, PhoneNo: v.PhoneNo, Courses: coursemap[k],
		})
	}
	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentsWithCourses() ([]*model.Student, error) {

	var studentlist []*model.Student
	rows, err := config.DB.Query("SELECT students.studentid,name,email,dept,dob,phoneno,courses.courseid,coursename,coursefee FROM students  join enrolments on students.studentid = enrolments.studentid   join courses on enrolments.courseid = courses.courseid")

	if err != nil {
		log.Print(err)
	}

	coursemap := make(map[int32][]model.Course)
	studmap := make(map[int32]model.StudentMap)

	for rows.Next() {
		var stud model.StudentMap
		var studCourse model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &studCourse.CourseID, &studCourse.CourseName, &studCourse.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if studCourse.CourseID != -1 {
				coursemap[stud.StudentID] = append(coursemap[stud.StudentID], studCourse)
			}

			studmap[stud.StudentID] = stud

		}
	}
	for k, v := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: v.StudentID, Name: v.Name, Email: v.Email, Dept: v.Dept, DOB: v.DOB, PhoneNo: v.PhoneNo, Courses: coursemap[k],
		})
	}
	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentById(id string) ([]*model.Student, error) {

	var studentlist []*model.Student = []*model.Student{}
	rows, err := config.DB.Query("SELECT students.studentid,name,email,dept,dob,phoneno,COALESCE(courses.courseid,-1),COALESCE(coursename,''),COALESCE(coursefee,-1) FROM students  left join enrolments on students.studentid = enrolments.studentid left  join courses on enrolments.courseid = courses.courseid where students.studentid =?", id)

	if err != nil {
		log.Print(err)
	}

	coursemap := make(map[int32][]model.Course)
	studmap := make(map[int32]model.StudentMap)

	for rows.Next() {
		var stud model.StudentMap
		var studCourse model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &studCourse.CourseID, &studCourse.CourseName, &studCourse.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if studCourse.CourseID != -1 {
				coursemap[stud.StudentID] = append(coursemap[stud.StudentID], studCourse)
			}

			studmap[stud.StudentID] = stud

		}
	}
	for k, v := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: v.StudentID, Name: v.Name, Email: v.Email, Dept: v.Dept, DOB: v.DOB, PhoneNo: v.PhoneNo, Courses: coursemap[k],
		})
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
	defer func() error {
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit()
		return nil
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
