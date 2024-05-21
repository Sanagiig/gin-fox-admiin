package common

type IndexModel struct {
	Index int `json:"index" gorm:"index;type:int;default 0"`
}
