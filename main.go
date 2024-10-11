package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var students = []Student{
	{Id: 1, Name: "Zeyno", Class: "12-A", Teacher: "Zekk"},
	{Id: 2, Name: "Ali", Class: "12-B", Teacher: "Apo"},
	{Id: 3, Name: "Ceren", Class: "12-A", Teacher: "Zekk"},
}

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

func ListStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var stu Student
	if err := c.BindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Student can't be created..."})
		return
	} else {
		students = append(students, stu)
		c.JSON(http.StatusCreated, students)
	}
}

func DeleteStudent(c *gin.Context) {
	sId := c.Query("stuID")
	id, err := strconv.Atoi(sId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "ID can't be found..."})
	}
	students = append(students[:id], students[id+1:]...)
	c.JSON(http.StatusOK, students)
}

func main() {
	router := gin.Default()
	router.GET("/students", ListStudents)
	router.POST("/students/add", CreateStudent)
	router.DELETE("/students/delete", DeleteStudent)
	router.Run("localhost:1801")
}
