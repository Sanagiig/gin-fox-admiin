package internal

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Name       string
	PersonInfo PersonInfo `gorm:"ForeignKey:PersonID;References:ID"`
}

func (Person) TableName() string {
	return "person"
}

type PersonInfo struct {
	ID       uint `gorm:"PrimaryKey"`
	PersonID uint
	Age      int
	Sex      int
}

func (*PersonInfo) TableName() string {
	return "person_info"
}

func (p *PersonInfo) BeforeCreate(tx *gorm.DB) error {
	p.ID = uint(time.Now().Unix())
	return nil
}

func CreateOne2One(db *gorm.DB) {
	pi := PersonInfo{
		Age: 1,
		Sex: 1,
	}
	p := Person{
		Name:       "xx1",
		PersonInfo: pi,
	}

	if err := db.Create(&p).Error; err != nil {
		panic(err)
	}
}

func GetOne2One(db *gorm.DB) {
	p := Person{
		Model: gorm.Model{ID: 2},
	}

	err := db.Preload("PersonInfo").Find(&p).Error
	//err := db.Model(&p).Association("PersonInfo").Find(&p.PersonInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("person: ", p)
}
