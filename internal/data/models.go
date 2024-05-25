package data

import (
	"github.com/jinzhu/gorm"
)

type Models struct {
	Courses CourseModel
	Modules ModuleModel
	Lessons LessonModel
}

func NewModels(db *gorm.DB) Models {
	return Models{
		Courses: CourseModel{DB: db},
		Modules: ModuleModel{DB: db},
		Lessons: LessonModel{DB: db},
	}
}
