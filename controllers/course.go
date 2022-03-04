package controllers

import (
	"github.com/Zeeshan-Ashraf/go1/models"
	"github.com/Zeeshan-Ashraf/go1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil { //c.ShouldBindJSON(&adhar) validates if it is json & unmarshal it to course
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	rows, err := services.InsertCourse(&course)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, gin.H{"id": course.ID, "total_rows_inserted": rows, "status": "success"})
}

func GetAllCourse(c *gin.Context) {
	rows, err := services.GetAllCourses()
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, rows)
}
