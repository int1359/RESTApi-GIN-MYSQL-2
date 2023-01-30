package model

type Student struct {
	StudentID int32    `json:"studentid"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Dept      string   `json:"dept"`
	DOB       string   `json:"dob"`
	PhoneNo   int32    `json:"phoneno"`
	Courses   []Course `json:"courses"`
}
type Course struct {
	CourseID   int32   `json:"courseid"`
	CourseName string  `json:"coursename"`
	CourseFee  float32 `json:"coursefee"`
}
type Enrolment struct {
	EnrolmentID string `json:"enrolmentid"`
	StudentID   int32  `json:"studentid"`
	CourseID    int32  `json:"courseid"`
}
type StudentMap struct {
	StudentID int32  `json:"studentid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Dept      string `json:"dept"`
	DOB       string `json:"dob"`
	PhoneNo   int32  `json:"phoneno"`
}
