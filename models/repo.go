package models

import (
	"gorm.io/gorm"
)

type IRepo interface {
	SetUoW(uow IUoW)
}

type Repo struct {
	*gorm.DB
}

func NewRepo(uow IUoW) *Repo {
	return &Repo{
		DB: uow.GetDB(),
	}
}

func (r *Repo) SetUoW(uow IUoW) {
	r.DB = uow.GetDB()
}
