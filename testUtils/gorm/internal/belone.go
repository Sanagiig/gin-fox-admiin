package internal

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Student struct {
	ID          string `gorm:"PrimaryKey;Size:256;"`
	Name        string
	No          string
	Age         uint
	InfoID      string
	StudentInfo StudentInfo `gorm:"ForeignKey:InfoID;"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) error {
	s.ID = strconv.FormatInt(time.Now().Unix(), 10)
	return nil
}

type StudentInfo struct {
	ID        string `gorm:"PrimaryKey;Size:256;"`
	Img       string
	Education string
	Protocol  string
	Address   string
}

func (si *StudentInfo) BeforeCreate(tx *gorm.DB) error {
	si.ID = strconv.FormatInt(time.Now().Unix(), 10)
	return nil
}

func CreateBelone(db *gorm.DB) {
	si := StudentInfo{
		Img:       "11.jpg",
		Education: "xxx",
	}
	s := Student{
		Name:        "x1",
		No:          "000001",
		Age:         11,
		StudentInfo: si,
	}

	err := db.Create(&s).Error
	if err != nil {
		panic(err)
	}
}

func GetBelone(db *gorm.DB) {
	s := Student{}

	err := db.Model(&s).Where("id like ?", "%1%").
		Preload("StudentInfo").First(&s).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
