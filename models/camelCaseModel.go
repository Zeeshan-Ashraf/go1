package models

import "gorm.io/gorm"

type CourseDetails struct { //will create table with name course_details in DB (gorm translates camelCase to snake_case )
	gorm.Model
	CourseName  string `json:"name"`
	CoursePrice int
	ISBN        int
}
