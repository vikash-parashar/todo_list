package models

type User struct {
	Username string
	Password []byte
	Todos    []Todo
}
