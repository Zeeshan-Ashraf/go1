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
	connections.DB.AutoMigrate(&models.Course{}) //automatically creates table in DB(Schema) if doesn't exists but schema must be created manually
	fmt.Printf("init func executed. connections.DB.AutoMigrate(&models.Course{}) is called")
}

func GetCourses(db *gorm.DB) ([]models.Course, error) {
	var course []models.Course
	tx := db.Find(&course)
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return course, nil
}

func InsertCourse(db *gorm.DB, course *models.Course) (*gorm.DB, error) {
	tx := db.Create(&course)
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return tx, nil
}

func GetCourse(db *gorm.DB, id int64) (*models.Course, error) {
	var course models.Course
	tx := db.First(&course, id)
	/* different ways of filter data
	tx := db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	tx := db.First(&user, "10") //string can be passed to id even if id is int64
	// SELECT * FROM users WHERE id = 10;

	tx := db.Find(&users, []int{1,2,3})
	// SELECT * FROM users WHERE id IN (1,2,3);
	*/
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return &course, nil
}
