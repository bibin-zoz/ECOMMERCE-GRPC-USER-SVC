package models

type UserLogin struct {
	Email    string
	Password string
}

type UserDetails struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}

type UserSignUp struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Password  string
}
type UserDetail struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Password  string
}
