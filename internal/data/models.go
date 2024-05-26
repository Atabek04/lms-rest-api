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

func (m ModuleModel) GetWithLessons(id uint) (*Module, error) {
	var module Module
	if err := m.DB.Preload("Lessons").First(&module, id).Error; err != nil {
		return nil, err
	}
	return &module, nil
}

func (m CourseModel) GetWithModulesAndLessons(id uint) (*Course, error) {
	var course Course
	if err := m.DB.Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Lessons")
	}).First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
