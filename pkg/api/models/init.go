package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Инициализация коннекта к БД
func InitDB() *gorm.DB  {
	var err error
	db, err = gorm.Open("postgres", "host=postgres port=5444 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(1)
	}

	db.AutoMigrate(
		&Note{},
		&Resume{},
		&ResumeSkill{},
		&ResumeWork{},
		&ResumeEducation{},
		&Feedback{},
		&Project{},
	)

	return db
}

// функция для выдачи offset и limit для запросов
func PaginatorNum(page int) (int, int)  {
	size := 10
	if page == 0 {
		return 0, 0
	}
	if page == 1 {
		return 0, size
	}
	page--
	return page * 10, size
}