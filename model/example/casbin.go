package example

import "gorm.io/gorm"

type CasbinPolice struct {
	gorm.Model
	Ptype string `gorm:"size:50;uniqueIndex:unique_index"`
	V0    string `gorm:"size:50;uniqueIndex:unique_index"`
	V1    string `gorm:"size:50;uniqueIndex:unique_index"`
	V2    string `gorm:"size:50;uniqueIndex:unique_index"`
	V3    string `gorm:"size:50;uniqueIndex:unique_index"`
	V4    string `gorm:"size:50;uniqueIndex:unique_index"`
	V5    string `gorm:"size:50;uniqueIndex:unique_index"`
}
