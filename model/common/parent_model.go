package common

type ParentModel struct {
	ParentID   string `json:"parentId" gorm:"type:char(36);index;column:parent_id"`
	ParentName string `json:"parentName" gorm:"-"`
}

type ChildrenMode[T any] struct {
	Children []T `json:"children" gorm:"-"`
}
