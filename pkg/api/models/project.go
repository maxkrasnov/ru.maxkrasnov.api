package models

import "github.com/jinzhu/gorm"

/**
	Модели для раздела Мои работы
 */

type Project struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);" json:"title"`
	Active bool `gorm:"type:boolean" json:"active"`
	Media string `gorm:"type:text;" json:"media"`
	Category string `gorm:"type:varchar(255);" json:"category"`
	Link string `gorm:"type:varchar(255);" json:"link"`
	Description string `gorm:"type:text;" json:"description"`
}

func (p *Project) GetAll(db *gorm.DB, page int) []Project {
	t := []Project{}
	offset, limit := PaginatorNum(page)
	db.Offset(offset).Limit(limit).Find(&t, Project{
		Active: true,
	})
	return t
}
