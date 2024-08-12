package types

import (
	"time"
)

// format JSON
type MetaData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
type FullResponse struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data"`
}

// User
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) (*User, error)
}
type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Nim       string    `json:"nim"`
	Semester  int       `json:"semester"`
	Major     string    `json:"major"`
	Faculty   string    `json:"faculty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type RegisterUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=8"`
	Name     string `json:"name" validate:"required"`
	Nim      string `json:"nim" validate:"required"`
	Semester int    `json:"semester" validate:"required"`
	Major    string `json:"major" validate:"required"`
	Faculty  string `json:"faculty" validate:"required"`
}
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=8"`
}

// Course
type CourseStore interface {
	GetCourses() ([]Course, error)
	GetCourseById(id int) (*Course, error)
}
type Course struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Credit    int       `json:"credit"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Student Plan Card (spc)
type SPCStore interface {
	CreateSPC(SPC) (*SPC, error)
	GetSPCByUserId(userId int) (*SPC, error)
}
type SPC struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	CourseId  int       `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type SPCPayload struct {
	UserId   int `json:"user_id" validate:"required"`
	CourseId int `json:"course_id" validate:"required"`
}
