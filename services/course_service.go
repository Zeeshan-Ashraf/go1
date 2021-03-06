package services

import (
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/dao"
	"github.com/Zeeshan-Ashraf/go1/models"
	"gorm.io/gorm"
)

func InsertCourse(course *models.Course) (int64, error) {
	dao_tx, err := dao.InsertCourse(connections.DB, course)
	if err != nil {
		return 0, err
	}
	return dao_tx.RowsAffected, nil
}

func InsertCourseIfNotExist(course *models.Course) (*gorm.DB, error) {
	dao_tx, err := dao.InsertCourseIfNotExist(connections.DB, course)
	if err != nil {
		return nil, err
	}
	return dao_tx, nil
}

func GetAllCourses() ([]models.Course, error) {
	dao_tx, err := dao.GetCourses(connections.DB)
	if err != nil {
		return nil, err
	}
	return dao_tx, nil
}

func GetCourse(id int64) (*models.Course, error) {
	dao_tx, err := dao.GetCourse(connections.DB, id)
	if err != nil {
		return nil, err
	}
	return dao_tx, nil
}

func GetCourseByName(name string) (models.Course, error) {
	dao_tx, err := dao.GetCourseByName(name)
	if err != nil {
		return dao_tx, err
	}
	return dao_tx, nil
}

func GetCourseToGenericMap(id int64) (*map[string]interface{}, error) {
	return dao.GetCourseToGenericMap(id)

}

func GetRawSql() (*map[string]interface{}, error) {
	return dao.RawSql()

}

func GetPaginateCourse(page int, count_peer_pg int) ([]models.Course, error) {
	return dao.GetPaginatedResult(page, count_peer_pg)

}
