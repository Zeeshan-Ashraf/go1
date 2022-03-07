package controllers

import (
	"github.com/Zeeshan-Ashraf/go1/models"
	"github.com/Zeeshan-Ashraf/go1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCourseDetails(c *gin.Context) {
	var course_d models.CourseDetails
	if err := c.ShouldBindJSON(&course_d); err != nil { //c.ShouldBindJSON(&adhar) validates if it is json & unmarshal it to course
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	rows, err := services.InsertCourseDetails(&course_d)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, gin.H{"id": course_d.ID, "total_rows_inserted": rows, "status": "success"})
}

func GetAllCourseDetails(c *gin.Context) {
	rows, err := services.GetAllCourseDetails()
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, rows)
}
