package common

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CommonModel struct {
	ID        uint           `gorm:"primary_key;unique" json:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`                    // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`                    // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`               // 删除时间
}

type StatusModel struct {
	Status string `json:"status" form:"status" gorm:"type:char(20);default:active;comment:数据状态 active , disabled"`
}

type UUIDModel struct {
	ID string `gorm:"primary_key;unique;type:char(36);" json:"id"` // 主键ID
}

func (u *UUIDModel) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		UUID, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		u.ID = UUID.String()
	}
	return nil
}

type DateModel struct {
	CreatedAt time.Time      `json:"createdAt"`      // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type UuidDateModel struct {
	UUIDModel
	DateModel
}
