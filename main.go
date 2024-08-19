package main

import (
	"student/entity"
	"student/server"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.Database = map[string]entity.Student{}

	r := gin.Default()

	r.POST("/student", server.AddStudent)

	r.GET("/student", server.ShowStudents)

	r.DELETE("/student/:id", server.RemoveStudent)

	r.PUT("/student/:id", server.ShowStudenDues)

	r.Run(":8080")
}
