package services

import (
	"github.com/Zeeshan-Ashraf/go1/connections"
	"github.com/Zeeshan-Ashraf/go1/dao"
	"github.com/Zeeshan-Ashraf/go1/models"
)

func InsertCourse(course models.Course) (int64, error) {
	db_tx := dao.InsertCourse(connections.DB, &course)
	if db_tx.Error != nil {
		return 0, db_tx.Error
	}
	return db_tx.RowsAffected, nil
}
