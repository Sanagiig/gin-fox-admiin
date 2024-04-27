package internal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
	// 使用 CompanyRefer 作为外键
}

type Company struct {
	ID   int
	Name string
}
