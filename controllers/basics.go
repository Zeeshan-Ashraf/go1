package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
)

/*
simplest get request, returns string (not json) = "Hello There" to users
*/
func Hello(c *gin.Context) { //first letter upper case means method is public, c contains request details like url host body cookies headers etc, c is Context which can also validates and serializes JSON
	c.String(200, "Hello There")
}

func GetPathEnvVar(c *gin.Context) {
	path_value := os.Getenv("PATH")
	c.String(200, "PATH env variable value: %s", path_value)
}
