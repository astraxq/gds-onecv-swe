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