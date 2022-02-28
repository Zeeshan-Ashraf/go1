package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
