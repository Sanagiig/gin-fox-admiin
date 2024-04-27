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

type LinkModel struct {
	ID        string    `gorm:"primary_key;unique;type:char(36);" json:"id"` // 主键ID
	CreatedAt time.Time `json:"createdAt"`                                   // 创建时间
}

type UUIDModel struct {
	ID        string         `gorm:"primary_key;unique;type:char(36);" json:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`                                   // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`                                   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                              // 删除时间
}

func (u *UUIDModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		UUID, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		u.ID = UUID.String()
	}
	return
}
