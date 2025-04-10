package entities

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	QuizID        uint
	Quiz          Quiz
	Number        int
	Text          string
	RightAnswerID uint
}

type AnswerVariants struct {
	gorm.Model
	QuestionID uint
	Question   Question
	Text       string
}

type Quiz struct {
	gorm.Model
	UserID         uint
	User           User
	Theme          string
	TimeOutMinutes int
}
