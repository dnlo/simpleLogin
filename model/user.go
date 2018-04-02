package model

type User struct {
	ID int `storm:"id,increment"`
	Name string `storm:"unique"`
	Password string
}
