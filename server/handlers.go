package server

import (
	"fmt"
	"net/http"
	"student/entity"

	"github.com/gin-gonic/gin"
)

func AddStudent(ctx *gin.Context) {
	fmt.Println("Adding Student...")
	var cstudent entity.Student
	if err := ctx.BindJSON(&cstudent); err != nil {
		fmt.Println("Failed to add student")
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	entity.Database[cstudent.Id] = cstudent
	fmt.Println("Student added")
	ctx.IndentedJSON(http.StatusCreated, gin.H{"Student added": cstudent})
}

func ShowStudents(ctx *gin.Context) {
	var ListStudent []entity.Student
	for _, sstudent := range entity.Database {
		ListStudent = append(ListStudent, sstudent)
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"List of student": ListStudent})
}

func RemoveStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, found := entity.Database[id]; !found {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Invalid student information": id})
		return
	}
	ThisStudent := entity.Database[id]
	delete(entity.Database, id)
	fmt.Println("Student information deleted")
	ctx.IndentedJSON(http.StatusOK, gin.H{"Student informaion deleted": ThisStudent})
}

func ShowStudenDues(ctx *gin.Context) {
	iid := ctx.Param("id")
	if _, found := entity.Database[iid]; !found {
		fmt.Println("No such student")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Failed to retrive  student data": iid})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"Student Current Dues": entity.Database[iid]})
}
