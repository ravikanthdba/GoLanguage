package models

type UserDB struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Id        int    `json:"id"`
}
