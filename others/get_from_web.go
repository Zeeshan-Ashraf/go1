package others

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetSimpleMsgFromWeb() string {
	url_wttr := "https://wttr.in/Kolkata"
	resp, err := http.Get(url_wttr)
	defer resp.Body.Close() //will be executed in last when func will be exiting (if multiple defer then first defer will be executed in last)
	if err != nil {
		fmt.Sprintf("error in http.Get err: %s\n", err.Error())
		fmt.Sprintf("response code: %d\n", resp.StatusCode)
		//panic(err)
	}
	fmt.Printf("url: %s\nresp status code:%s\n", url_wttr, resp.StatusCode)
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response from url %s :\n%s", url_wttr, string(content)) //dont use Sprintf to print []byte or json string
	return string(content)
}
