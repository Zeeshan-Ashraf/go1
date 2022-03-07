package dao

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/models"
	"gorm.io/gorm"
)

func init() {
	//Automatically migrate your schema, to keep your schema up to date. NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	fmt.Printf("connections.DB %+v\n", connections.DB)
	connections.DB.AutoMigrate(&models.CourseDetails{}) //automatically creates table in DB(Schema) if doesn't exists but schema must be created manually
	fmt.Printf("init func executed. connections.DB.AutoMigrate(&models.CourseDetails{}) is called")
}

func GetCourseDetails(db *gorm.DB) ([]models.CourseDetails, error) {
	var course_detail []models.CourseDetails
	tx := db.Find(&course_detail) //since course is struct & it has gorm.Model defined inside struct hence no need to provide table name, tableName at DB will be struct name with camelCase of struct converted to snake_case at DB
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return course_detail, nil
}

func InsertCourseDetails(db *gorm.DB, course_d *models.CourseDetails) (*gorm.DB, error) {
	tx := db.Create(&course_d)
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return tx, nil
}
