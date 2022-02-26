package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Hello(c *gin.Context) { //first letter upper case means method is public, c contains request details like url host body cookies headers etc, c is Context which can also validates and serializes JSON
	c.String(200, "Hello There")
}

func Get_weather_by_location(c *gin.Context) {
	loc := c.Param("loc")
	url_wttr := "https://wttr.in/" + loc
	resp, err := http.Get(url_wttr)
	defer resp.Body.Close()
	if err != nil {
		fmt.Sprintf("error in http.Get err: %s\n", err.Error())
		fmt.Sprintf("response code: %d\n", resp.StatusCode)
		//panic(err)
	}
	fmt.Printf("url: %s\nresp status code:%s\n", url_wttr, resp.StatusCode)
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response from url %s :\n%s", url_wttr, string(content)) //dont use Sprintf to print []byte / json string
	//c.String(200, string(content)) //works but it'll render html as text on browser so all tag labels will be visible
	//c.HTML(http.StatusOK, "Weather Report", string(content))
	c.Data(http.StatusOK, "text/html; charset=utf-8", content) //requires data in []byte not string to sent
}
