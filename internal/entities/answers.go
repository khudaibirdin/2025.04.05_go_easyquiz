package entities

import "gorm.io/gorm"

type Answers struct {
	gorm.Model
	UserID     uint
	QuizID     uint
	QuestionID uint
	Answer     int
	Result     bool
}
