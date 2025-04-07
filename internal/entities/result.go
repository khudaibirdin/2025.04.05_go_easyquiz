package entities

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	UserID                  uint
	QuizID                  uint
	QuestionsAmount         int
	QuestionsAnsweredAmount int
	Percent                 int
}
