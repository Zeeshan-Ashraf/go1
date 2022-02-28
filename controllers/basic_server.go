package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
simplest get request, returns string (not json) = "Hello There" to users
*/
func Hello(c *gin.Context) { //first letter upper case means method is public, c contains request details like url host body cookies headers etc, c is Context which can also validates and serializes JSON
	c.String(200, "Hello There")
}

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
		//panic(err)
	}
	fmt.Printf("url: %s\nresp status code:%s\n", url_wttr, resp.StatusCode)
	content, err := ioutil.ReadAll(resp.Body)                           //ioutil. ReadAll is a useful io utility function for reading all data from a io. Reader until EOF. It's often used to read data such as HTTP response body, files and other data sources which implement io. Reader interface
	fmt.Printf("response from url %s :\n%s", url_wttr, string(content)) //dont use Sprintf to print []byte / json string
	//c.String(200, string(content)) //works but it'll render html as text on browser so all tag labels will be visible
	//c.HTML(http.StatusOK, "Weather Report", string(content))
	c.Data(http.StatusOK, "text/html; charset=utf-8", content) //requires data in []byte form not string or json or object to sent and & contentType="text/html; charset=utf-8" to tell browser abt what type of []byte it is
}

/*
returns {"Name":"Zeeshan","Id":101} to user
*/
func SendJsonUsingGinH(c *gin.Context) {
	// H is a shortcut for map[string]interface{}
	//c.JSON is best way to convert any data type to JSON and send it to user
	c.JSON(200, gin.H{"name": "Zeeshan"}) // cant be used in c.Data as c.Data requires []byte but gin.H is map so one can send any data type in c.JSON
	// so this gives err c.Data(200, "application/json", gin.H{"name": "Zeeshan"}) as gin.H is not []byte
	//P.S: any object can be converted to JSON using json.Marshal(any_var_zee) & result will be ([]byte,,err) this []byte contains json
}

/*
returns {"Name":"Zeeshan","Id":[101,102],"Addresses":[{"City":"Kolkata","Pincode":700046},{"City":"Bangalore","Pincode":800002}]} to user
*/
func SendJsonUsingMarshal(c *gin.Context) {
	type Address struct { //note: fisrt letter must be upper case (address instead of Address) if Marshaling otherwise if pvt (unexported) then json.Masrshal wont be able to access that key & hence wont be able to convert to json
		City    string
		Pincode int
	}
	struct_data := struct {
		Name      string
		Id        []int
		Addresses []Address
	}{
		"Zeeshan",
		[]int{101, 102}, //note a comma in end is required if there is composite data structure
		[]Address{{"Kolkata", 700046}, {"Bangalore", 800002}}, //note a comma in end is required if there is composite data structure
	}

	json_in_byteSlice, err := json.Marshal(struct_data) //automatically convert to json []byte of any interface
	fmt.Printf("json_in_byteSlice: %s", json_in_byteSlice)
	if err != nil {
		log.Printf("error occured: %s", err.Error())
		c.AbortWithStatusJSON(500, err)
	}
	c.Data(200, "application/json", json_in_byteSlice) //not using c.JSON since c.JSON auto converts directly any object to json without marshal //c.Data requires []byte to send & contentType="application/json" to tell browser abt what type of []byte it is
}

/*
returns {"naam":"Zeeshan","id":[101,102],"Addresses":[{"city":"Kolkata","pincode":700046},{"city":"Bangalore","pincode":800002}]} to user
*/
func SendJsonWithCustomkeyName(c *gin.Context) {
	type Address struct { //note: fisrt letter must be upper case (address instead of Address) if Marshaling otherwise if pvt (unexported) then json.Masrshal wont be able to access that key & hence wont be able to convert to json
		City    string `json:"city"`
		Pincode int    `json:"pincode,omitempty"` //omitempty means remove this field if it is nil note: dont keep space after `pincode,`<no space>omitempty
	}
	struct_data := struct {
		Name      string    `json:"naam"`
		Id        []int     `json:"id"` //hide value of this field by ***
		Password  int       `json:"-"`  //remove this field while marshalling
		Addresses []Address //no json naming so it'll use variable name as it is i.e Addresses
	}{
		"Zeeshan",
		[]int{101, 102}, //note a comma in end is required if there is composite data structure
		12345678,
		[]Address{{"Kolkata", 700046}, {"Bangalore", 800002}}, //note a comma in end is required if there is composite data structure
	}
	//c.JSON is fast way to convert any data type to JSON and send it to user
	c.JSON(200, struct_data) // cant be used in c.Data as c.Data requires []byte but gin.H is map so one can send any data type in c.JSON. so this gives err c.Data(200, "application/json", gin.H{"name": "Zeeshan"}) as gin.H is not []byte
}

func SendSimpleJSON(c *gin.Context) {
	struct_data := struct {
		Name string
		Id   int
	}{
		"Zeeshan",
		101,
	}
	//c.JSON is best way to convert any data type to JSON and send it to user
	c.JSON(200, struct_data) // cant be used in c.Data as c.Data requires []byte but gin.H is map so one can send any data type in c.JSON
	// so this gives err c.Data(200, "application/json", gin.H{"name": "Zeeshan"}) as gin.H is not []byte
}

func SendPostReqWithData() {
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
