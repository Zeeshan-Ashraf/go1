package dao

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/models"
	"gorm.io/gorm"
)

//pupose of DAO layer is to interact with DB (DAO is final / nearest layer to DB)

//init() func automatically called (no need to call this func) whenever this package dao is imported & it is called only once
func Init() { //Automatically migrate your schema, to keep your schema up to date. NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	//course := models.Course{gorm.Model{}, "c1", 100}
	fmt.Printf("connections.DB %+v\n", connections.DB)
	connections.DB.AutoMigrate(&models.Course{})
	fmt.Printf("init func executed. connections.DB.AutoMigrate(&models.Course{}) is called")
}

func GetCourse(db *gorm.DB, id uint) []models.Course {
	var course []models.Course
	db.Find(&course)
	return course
}

func InsertCourse(db *gorm.DB, course *models.Course) *gorm.DB {
	return db.Create(&course)
}
