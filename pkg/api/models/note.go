package models

/**
	Модели для раздела Блог
 */

import "github.com/jinzhu/gorm"

// Отдельная запись в блоге
type Note struct {
	gorm.Model
	Code string `gorm:"type:varchar(255);unique_index" json:"code"`
	Active bool `gorm:"type:boolean" json:"active"`
	Category string `gorm:"type:varchar(255);" json:"category"`
	Title string `gorm:"type:varchar(255);" json:"title"`
	Media string `gorm:"type:text;" json:"media"`
	PreviewText string `gorm:"type:text;" json:"preview_text"`
	DetailText string `gorm:"type:text;" json:"detail_text"`
	HTML bool `gorm:"type:boolean" json:"html"`
}

func (n *Note) GetAll(db *gorm.DB, page int) []Note {
	t := []Note{}
	offset, limit := PaginatorNum(page)
	db.Select("code, title, media, preview_text, created_at, category").Order("created_at desc").Offset(offset).Limit(limit).Find(&t, Note{
		Active: true,
	})
	return t
}

func (n *Note) GetByCode(code string, db *gorm.DB)  {
	db.First(n, Note{
		Code: code,
		Active: true,
	})
}