package models

import (
	"github.com/jinzhu/gorm"
	"sort"
)

/**
	Модели для раздела Резюме
 */

// Резюме
type Resume struct {
	gorm.Model
	PreviewText string `gorm:"type:text;" json:"preview_text"`
	Skills []ResumeSkill `json:"skills"`
	Experience []ResumeWork `json:"experience"`
	Education []ResumeEducation `json:"education"`
}

// Список навыков
type ResumeSkill struct {
	ID uint `gorm:"primary_key" binding:"-"`
	Name string `gorm:"type:varchar(80);" json:"name"`
	Level int `gorm:"type:int" json:"level"`
	Resume Resume `json:"-"`
	ResumeID uint `json:"-"`
}

// Список для "опыт работы"
type ResumeWork struct {
	ID uint `gorm:"primary_key" binding:"-"`
	Years string `gorm:"type:varchar(9);" json:"years"`
	Organization string `gorm:"type:varchar(255)" json:"work"`
	Position string `gorm:"type:varchar(255)" json:"position"`
	Location string `gorm:"type:varchar(255)" json:"location"`
	Description string `gorm:"type:text" json:"description"`
	Resume Resume `json:"-"`
	ResumeID uint `json:"-"`
}

// Список для "Обучение"
type ResumeEducation struct {
	ID uint `gorm:"primary_key" binding:"-"`
	Years string `gorm:"type:varchar(9);" json:"years"`
	Organization string `gorm:"type:varchar(255)" json:"work"`
	Position string `gorm:"type:varchar(255)" json:"position"`
	Location string `gorm:"type:varchar(255)" json:"location"`
	Description string `gorm:"type:text" json:"description"`
	Resume Resume `json:"-"`
	ResumeID uint `json:"-"`
}

func (r *Resume) Get(db *gorm.DB)  {
	resumeSkills := []ResumeSkill{}
	resumeWorks := []ResumeWork{}
	resumeEducation := []ResumeEducation{}
	db.Last(&r).Related(&resumeSkills).Related(&resumeWorks).Related(&resumeEducation)

	// Вомзожно, не лучшее решение, просто быстро и без запарки
	sort.Slice(resumeSkills, func(i, j int) bool {
		return resumeSkills[i].Level > resumeSkills[j].Level
	})

	r.Skills = resumeSkills
	r.Experience = resumeWorks
	r.Education = resumeEducation
}