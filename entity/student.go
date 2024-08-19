package entity

var Database map[string]Student

type Student struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CurrentDues int    `json:"dues"`
}
