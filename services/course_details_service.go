package services

import (
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/dao"
	"github.com/Zeeshan-Ashraf/go1/models"
)

func InsertCourseDetails(course_d *models.CourseDetails) (int64, error) {
	dao_tx, err := dao.InsertCourseDetails(connections.DB, course_d)
	if err != nil {
		return 0, err
	}
	return dao_tx.RowsAffected, nil
}

func GetAllCourseDetails() ([]models.CourseDetails, error) {
	dao_tx, err := dao.GetCourseDetails(connections.DB)
	if err != nil {
		return nil, err
	}
	return dao_tx, nil
}
