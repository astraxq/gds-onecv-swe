package controller

// create enum for status
type Status int

const (
	ACTIVE Status = iota
	INACTIVE
	SUSPENDED
)

// create enum for role
type Role int

const (
	ADMIN Role = iota
	TEACHER 
	STUDENT
)

type User struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role Role `json:"role"`
	Status Status `json:"status"`
	NotificationAllowed bool `json:"notification_allowed"`
}

// Create a tag struct between teacher-student
type UserTag struct {
	ID uint64 `json:"id"`
	TeacherID uint64 `json:"teacher_id"`
	StudentID uint64 `json:"student_id"`
}