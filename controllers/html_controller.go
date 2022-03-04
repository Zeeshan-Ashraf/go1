package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/*
receives html file from website and convert it to []byte and send it back as html
*/
func GetWeatherByLocation(c *gin.Context) {
	loc := c.Param("loc")
	url_wttr := "https://wttr.in/" + loc
	resp, err := http.Get(url_wttr)
	defer resp.Body.Close()
	if err != nil {
		fmt.Sprintf("error in http.Get err: %s\n", err.Error())
		fmt.Sprintf("response code: %d\n", resp.StatusCode)
		c.AbortWithStatusJSON(500, err.Error())
		return //without return all the c.<xyz> will be sent to user
	}
	fmt.Printf("url: %s\nresp status code:%s\n", url_wttr, resp.StatusCode)
	content, err := ioutil.ReadAll(resp.Body)                           //ioutil. ReadAll is a useful io utility function for reading all data from a io. Reader until EOF. It's often used to read data such as HTTP response body, files and other data sources which implement io. Reader interface
	fmt.Printf("response from url %s :\n%s", url_wttr, string(content)) //dont use Sprintf to print []byte / json string
	//c.String(200, string(content)) //works but it'll render html as text on browser so all tag labels will be visible
	//c.HTML(http.StatusOK, "Weather Report", string(content))
	c.Data(http.StatusOK, "text/html; charset=utf-8", content) //requires data in []byte form not string or json or object to sent and & contentType="text/html; charset=utf-8" to tell browser abt what type of []byte it is
}

/*
send local html files to user
*/
func SendLocalHtmlFile(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
