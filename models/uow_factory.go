package models

import "gorm.io/gorm"

type IUoWFactory interface {
	GetUoW() IUoW
	GetDefaultUoW() IUoW
}

type UoWFactory struct {
	DB *gorm.DB
}

func NewUoWFactory(db *gorm.DB) *UoWFactory {
	return &UoWFactory{
		DB: db,
	}
}
func (f *UoWFactory) GetUoW() IUoW {
	return &UoW{
		DB: f.DB.Begin(),
	}
}
func (f *UoWFactory) GetDefaultUoW() IUoW {
	return &UoW{
		DB: f.DB,
	}
}
