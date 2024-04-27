package main

import (
	"gin-one/testUtils/gorm/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	dsn := "root:7q8w9e@tcp(127.0.0.1:3306)/gin_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	sql := mysql.New(mysql.Config{
		DSN: dsn,
	})
	db, err := gorm.Open(sql, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		//&internal.User{},
		&internal.Student{},
		//&internal.StudentInfo{},
		&internal.Person{},
		&internal.PersonInfo{},

		&internal.Author{},
		&internal.Article{},

		&internal.SmallStudent{},
		&internal.Course{},
		&internal.StudentCourse{},
	)
	if err != nil {
		panic(err)
	}
	return db
}
