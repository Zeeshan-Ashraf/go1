package connections

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:password@tcp(127.0.0.1:3306)/go1?charset=utf8mb4&parseTime=True&loc=Local" //<user>:<password>@<protocol>(ip:port)/<databaseName also called schemaName>?charset=utf8mb4&parseTime=True&loc=Local

func ConnectToDB() (*gorm.DB, error) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	fmt.Printf("DB conected DNS: %s\n", DNS)
	if err != nil {
		fmt.Printf("error in connecting DB: %s\n", err.Error())
		return nil, err
	}
	return DB, nil
}
