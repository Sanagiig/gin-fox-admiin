package common

type IndexModel struct {
	Index string `json:"index" gorm:"index;type:int;"`
}
