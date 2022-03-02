package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*simplest way to send object in JSON form*/
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

//declearing struct as global so that it can be used at multiple places
type Address struct { //note: fisrt letter must be upper case (address instead of Address) if Marshaling otherwise if pvt (unexported) then json.Masrshal wont be able to access that key & hence wont be able to convert to json
	City    string
	Pincode int
}

type Adhaar struct {
	Name      string
	Id        []int
	Addresses []Address
}

/*
returns {"Name":"Zeeshan","Id":[101,102],"Addresses":[{"City":"Kolkata","Pincode":700046},{"City":"Bangalore","Pincode":800002}]} to user
*/
func SendJsonUsingMarshal(c *gin.Context) {

	pincode1 := 700046
	struct_data := Adhaar{
		"Zeeshan",
		[]int{101, 102}, //note a comma in end is required if there is composite data structure
		[]Address{{"Kolkata", pincode1}, {"Bangalore", 800002}}, //note a comma in end is required if there is composite data structure
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
	//P.S Address is declared globally as well by local gets preference
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

//===============UNMARSHAL=======================================================
//create fake json in []byte format to mimic as if data is recieved from web
var Json_data_from_web_correct []byte = []byte(`
{
  "Name": "Zeeshan",
  "Id": [101,102],
  "Addresses": [
    {
      "City": "Kolkata",
      "Pincode": 700046
    },
    {
      "City": "Bangalore",
      "Pincode": 800002
    }]
}
`)

var Json_data_from_web_wrong []byte = []byte(`
{
  "Name": "Zeeshan",
  "Id": [101,102],
  "Addresses": [
    {
      "City": "Kolkata",
      "Pincode": 700046
    },
    {
      "City": "Bangalore",
      "Pincode": 800002
    }
}
`)

var Json_data_from_web_correct_but_not_Adhaar []byte = []byte(`
{
  "Name": "Zeeshan",
  "Id": "Ashraf",
  "Addresses": [
    {
      "City": "Kolkata",
      "Pincode": 700046
    },
    {
      "City": "Bangalore",
      "Pincode": 800002
    }]
}
`)

/*
validate json from web to check if json received is in json format //here we are not checking if json is in inline with a obj (that will be done by json.Unmarshal or c.ShouldBindWith)
*/
func ValidateJsonFromWeb(data []byte) error {
	checkValidity := json.Valid(data)
	if checkValidity == false {
		fmt.Printf("Data is not a valid json\ndata:%s", data)
		return fmt.Errorf("Data is not a valid json\nData:%s", data)
	}
	return nil
}

func UnmarshalJson() {
	/*positive case: works fine*/
	var adhaar Adhaar //create a variable without any data
	err := json.Unmarshal(Json_data_from_web_correct, &adhaar)
	if err != nil {
		fmt.Printf("\nerror in unmarshal\nerr:%s\n", err.Error())
	} else {
		fmt.Printf("adhaar: %s\n", adhaar)
	}

	/*negative case: data itself is not in json form*/
	var adhaar2 Adhaar //create a variable without any data
	err2 := json.Unmarshal(Json_data_from_web_wrong, &adhaar2)
	if err2 != nil {
		fmt.Printf("\nerror in unmarshal\nerr:%s\n", err2.Error())
	} else {
		fmt.Printf("adhaar: %s\n", adhaar2)
	}

	/*negative case: data is in json form but not according to Adhaar struct type*/
	var adhaar3 Adhaar //create a variable without any data
	err3 := json.Unmarshal(Json_data_from_web_correct_but_not_Adhaar, &adhaar3)
	if err3 != nil {
		fmt.Printf("\nerror in unmarshal\nerr:%s\n", err3.Error())
	} else {
		fmt.Printf("adhaar: %s\n", adhaar3)
	}
}

func UnmarshalJsonUsingShouldBind(c *gin.Context) {
	var adhar Adhaar
	if err := c.ShouldBindJSON(&adhar); err != nil { //c.ShouldBindJSON(&adhar) validates if it is json & unmarshal it to adhar
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, adhar)
}

func UnmarshalAnyJsonToMap(c *gin.Context) {
	var json_data_from_web map[string]interface{}                 //json field is always string hence map[string]
	if err := c.ShouldBindJSON(&json_data_from_web); err != nil { //c.ShouldBindJSON(&adhar) validates if it is json & unmarshal it to adhar
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, json_data_from_web)
}
