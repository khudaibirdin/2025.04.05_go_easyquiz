package entities

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	UserID                  uint
	User                    User
	QuizID                  uint
	Quiz                    Quiz
	QuestionsAmount         int
	QuestionsAnsweredAmount int
	Percent                 int
}
