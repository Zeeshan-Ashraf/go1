package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
receives html file from website and convert it to []byte and prints in log
*/
func GetReqFromWeb() {
	url_wttr := "https://wttr.in/Kolkata"
	resp, err := http.Get(url_wttr)
	defer resp.Body.Close()
	if err != nil {
		fmt.Sprintf("error in http.Get err: %s\n", err.Error())
		fmt.Sprintf("response code: %d\n", resp.StatusCode)
		//panic(err)
	}
	fmt.Printf("url: %s\nresp status code:%s\n", url_wttr, resp.StatusCode)
	content, err := ioutil.ReadAll(resp.Body)                           //ioutil. ReadAll is a useful io utility function for reading all data from a io. Reader until EOF. It's often used to read data such as HTTP response body, files and other data sources which implement io. Reader interface
	fmt.Printf("response from url %s :\n%s", url_wttr, string(content)) //dont use Sprintf to print []byte / json string
	//c.String(200, string(content)) //works but it'll render html as text on browser so all tag labels will be visible
	//c.HTML(http.StatusOK, "Weather Report", string(content))
}

/*
send post request to website https://reqbin.com/sample/post/json with json data & website returns success msg in json
*/
func SendPostReqToWebWithData() {
	myUrl := "https://reqbin.com/sample/post/json"
	requestBody := strings.NewReader(`
		{
			"sample":"json",
			"sample_no":101
		}
	`) // strings.NewReader returns *Reader * http.POST requires io.Reader P.S: io.Reader interface represents the read end of a stream of data

	response, _ := http.Post(myUrl, "application/json", requestBody)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body) //content is now []byte
	fmt.Printf("response : %s", string(content))
}

/*
get call from web with timeout and send the result back to user
*/
func GetReqWithTimeoutFromWeb(c *gin.Context) {
	tim, err := strconv.ParseInt((c.Param("tim")), 10, 64)
	if err != nil {
		log.Printf("error in converting time string to int64 %s", err.Error())
		c.AbortWithStatusJSON(500, err)
	}
	url_50mb_json := "http://postman-echo.com/bytes/50/mb?type=json" //this contains json which takes about 3 sec+ to fetch
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Duration(tim) * 1000 * time.Millisecond,
	}
	resp, err := client.Get(url_50mb_json)
	defer resp.Body.Close()
	if err != nil {
		fmt.Sprintf("error in http.Get err: %s\n", err.Error())
		fmt.Sprintf("response code: %d\n", resp.StatusCode)
		c.AbortWithStatusJSON(500, err)
		//panic(err)
	}
	fmt.Printf("url: %s\nresp status code:%s\n", url_50mb_json, resp.StatusCode)
	content, err := ioutil.ReadAll(resp.Body) //ioutil. ReadAll is a useful io utility function for reading all data from a io. Reader until EOF. It's often used to read data such as HTTP response body, files and other data sources which implement io. Reader interface
	c.Data(http.StatusOK, "application/json", content)
}
