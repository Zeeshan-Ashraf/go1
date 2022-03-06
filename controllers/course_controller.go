package controllers

import (
	"github.com/Zeeshan-Ashraf/go1/models"
	"github.com/Zeeshan-Ashraf/go1/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func GetCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("error in converting id string %s to int64 %s", c.Param("id"), err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return //without return all the c.<xyz> will be sent to user
	}
	row, err := services.GetCourse(id)
	if err != nil {
		if row == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return //without return all the c.<xyz> will be sent to user
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, *row)
	return
}

func GetCourseByName(c *gin.Context) {
	name := c.Query("name") //also c.Query("name") is same as c.Request.URL.Query().Get("name")
	log.Printf("query name=%s", name)
	if name == "" {
		log.Printf("provide name query in URL ?name=name")
		c.AbortWithStatusJSON(http.StatusBadRequest, "provide name query in URL ?name=course_name")
		return //without return all the c.<xyz> will be sent to user
	}
	row, err := services.GetCourseByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return //without return all the c.<xyz> will be sent to user
	}
	c.JSON(200, row)
	return
}

func GetCourseToGenericMap(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error in parsing id from URL: "+c.Request.RequestURI+"  Error:"+err.Error())
		return //without return all the c.<xyz> will be sent to user
	}
	row, err := services.GetCourseToGenericMap(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return //without return all the c.<xyz> will be sent to user
	}
	c.JSON(200, *row)
	return
}

func GetRawSqlFromCourses(c *gin.Context) {
	row, err := services.GetRawSql()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return //without return all the c.<xyz> will be sent to user
	}
	if *row == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "no records found")
		return //without return all the c.<xyz> will be sent to user
	}
	c.JSON(200, *row)
	return
}

func GetPaginatedCourse(c *gin.Context) {
	page, err := strconv.ParseInt(c.Query("page"), 10, strconv.IntSize) //also c.Query("name") is same as c.Request.URL.Query().Get("name")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "err in parsing page in query parameter err: "+err.Error())
		return
	}
	log.Printf("page no=%d", page)

	offset, err := strconv.ParseInt(c.Query("offset"), 10, strconv.IntSize) //also c.Query("name") is same as c.Request.URL.Query().Get("name")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "err in parsing offset in query parameter err: "+err.Error())
		return
	}
	log.Printf("count in one page=%d", offset)

	row, err := services.GetPaginateCourse(int(page), int(offset))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return //without return all the c.<xyz> will be sent to user
	}
	if row == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "no records found")
		return //without return all the c.<xyz> will be sent to user
	}
	c.JSON(200, row)
	return
}
