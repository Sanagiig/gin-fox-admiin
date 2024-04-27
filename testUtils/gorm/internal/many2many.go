package internal

import "gorm.io/gorm"

type SmallStudent struct {
	gorm.Model
	UUID    uint `gorm:"index"`
	Name    string
	Age     int
	Courses []Course `gorm:"many2many:student_course;"`
}

type Course struct {
	gorm.Model
	ClsName      string
	Des          string
	SmallStudent []SmallStudent `gorm:"many2many:student_course"`
}

type StudentCourse struct {
	gorm.Model
	StudentID    uint         `gorm:"column:student_id"`
	CourseID     uint         `gorm:"column:course_id"`
	SmallStudent SmallStudent `gorm:"foreignKey:StudentID"`
	Course       Course       `gorm:"foreignKey:CourseID"`
}

func (s *StudentCourse) TableName() string {
	return "student_course"
}
