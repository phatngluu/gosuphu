package models

import "gorm.io/gorm"

type IUoW interface {
	GetDB() *gorm.DB
	Rollback()
	Commit() error
}
type UoW struct {
	DB *gorm.DB
}

func (u *UoW) GetDB() *gorm.DB {
	return u.DB
}
func (u *UoW) Commit() error {
	return u.DB.Commit().Error
}

func (u *UoW) Rollback() {
	u.DB.Rollback()
}

func NewUoW(db *gorm.DB) IUoW {
	return &UoW{
		DB: db,
	}
}
