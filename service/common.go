package service

import "github.com/jinzhu/gorm"

func NewHandle(db *gorm.DB) *Handlers {
	return &Handlers{
		db: db,
	}
}

type Handlers struct {
	db *gorm.DB
}
