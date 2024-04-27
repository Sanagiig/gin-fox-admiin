package internal

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type Author struct {
	gorm.Model
	Name     string
	Age      int
	Articles []Article `gorm:"ForeignKey:AuthorID"`
}

func (a *Author) TableName() string {
	return "author"
}

type Article struct {
	gorm.Model
	Name     string
	Des      string
	AuthorID uint
}

func (a *Article) TableName() string {
	return "article"
}

func CreateOne2Many(db *gorm.DB) {
	author := Author{
		Name:     "x1",
		Articles: make([]Article, 3),
	}
	for i := 0; i < 3; i++ {
		author.Articles = append(author.Articles, Article{Name: "book" + strconv.Itoa(i)})
	}
	err := db.Create(&author).Error
	if err != nil {
		panic(err)
	}
}

func GetOne2Many(db *gorm.DB) {
	author := Author{
		Model: gorm.Model{
			ID: 3,
		},
	}

	err := db.Where("id = ?", author.ID).Association("Articles").
		Find(&author)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("get :", db.RowsAffected)
	fmt.Printf("author is : %v\n", author)
}
