package entities

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	UserID           uint
	User             User
	QuizID           uint
	Quiz             Quiz
	QuestionID       uint
	Question         Question
	AnswerVariantsID uint
	AnswerVariants   AnswerVariants
	Result           bool
}
