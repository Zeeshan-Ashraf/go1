package dao

import (
	"fmt"
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/models"
	"gorm.io/gorm"
	"log"
)

//pupose of DAO layer is to interact with DB (DAO is final / nearest layer to DB)

//init() func automatically called (no need to call this func) whenever this package dao is imported & it is called only once
func AutoMigrateCourse() { //Automatically migrate your schema, to keep your schema up to date. NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	//course := models.Course{gorm.Model{}, "c1", 100}
	fmt.Printf("connections.DB %+v\n", connections.DB)
	connections.DB.AutoMigrate(&models.Course{}) //automatically creates table in DB(Schema) if doesn't exists but schema must be created manually
	fmt.Printf("init func executed. connections.DB.AutoMigrate(&models.Course{}) is called")
}

func GetCourses(db *gorm.DB) ([]models.Course, error) {
	var course []models.Course
	tx := db.Find(&course) //since course is struct & it has gorm.Model defined inside struct hence no need to provide table name, tableName at DB will be struct name with camelCase of struct converted to snake_case at DB
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

func InsertCourseIfNotExist(db *gorm.DB, course *models.Course) (*gorm.DB, error) {
	tx := db.FirstOrCreate(&course, models.Course{Name: course.Name}) //checks if name exists or not if exists then don't insert tx.RowsAffected doesn't return 0 if rows already exists it always returns -1
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	log.Printf("%+v %+v\n", tx.RowsAffected, *tx.Row())
	return tx,nil
	/*
	//if used library: "github.com/jinzhu/gorm" instead of "gorm.io/gorm" then tx.RowsAffected return 1 if inserted else 0 "github.com/jinzhu/gorm" is old version of "gorm.io/gorm"
	if tx.Error != nil {
		return nil, tx.Error
	} else if tx.RowsAffected == 1 {
		log.Printf("New course created")
		return tx, nil
	} else {
		log.Printf("course with same name already exists")
		return tx, nil
	}
	return tx, nil
	*/
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

func GetCourseByName(name string) (models.Course, error) {
	var course models.Course
	// Get first matched record
	tx := connections.DB.Where("name = ?", name).First(&course) // SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return course, tx.Error
	}
	return course, nil
}

func GetCourseToGenericMap(id int64) (*map[string]interface{}, error) { //when db col are not known
	var course map[string]interface{}
	tx := connections.DB.Table("courses").Where("id = ?", id).Take(&course) //since course is map & it doesn't have gorm.Model defined inside struct hence need to provide table name
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return &course, nil
}

func RawSql() (*map[string]interface{}, error) {
	var result map[string]interface{}
	tx := connections.DB.Raw("SELECT id, name, price FROM courses WHERE name = ?", "ECE").Scan(&result)
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	log.Printf("executed raw Sql: %s,%s", "SELECT id, name, price FROM courses WHERE name = ?", "CSE")
	return &result, nil

}

func GetPaginatedResult(page int, count_per_pg int) ([]models.Course, error) {
	var course []models.Course
	tx := connections.DB.Limit(page * count_per_pg).Offset((page - 1) * count_per_pg).Find(&course) // SELECT * FROM users OFFSET 5 LIMIT 10; //Limit specify the max number of records to retrieve & Offset specify the number of records to skip before starting to return the records
	if tx.Error != nil {
		fmt.Printf("Error in db.Find, err: %+v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return course, nil

}

/*other gorm queries*/
/*

// Get all matched records
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';


//Raw sql

var result Result
db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

// Raw SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)


*/
