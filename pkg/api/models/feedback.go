package models

import "time"

/**
	Модели для раздела Написать мне
 */

// Обратная связь
type Feedback struct {
	ID uint `binding:"-"`
	Name string `gorm:"varchar(150)" json:"name" binding:"required"`
	Email string `gorm:"varchar(150)" json:"email" binding:"required"`
	Question string `gorm:"varchar(255)" json:"question" binding:"required"`
	CreatedAt time.Time `binding:"-"`
}
