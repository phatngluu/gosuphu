package models

import "gorm.io/plugin/soft_delete"

type Model struct {
	CreatedAt int64                 `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt int64                 `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt soft_delete.DeletedAt `json:"deletedAt"`
}
