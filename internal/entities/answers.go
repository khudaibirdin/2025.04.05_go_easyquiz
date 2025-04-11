package entities

import "gorm.io/gorm"

// Модель ответа от пользователя на вопрос
type Answer struct {
	gorm.Model
	UserID           uint
	User             User `gorm:"ForeignKey:UserID"`
	QuizID           uint
	Quiz             Quiz `gorm:"ForeignKey:QuizID"`
	QuestionID       uint
	Question         Question `gorm:"ForeignKey:QuestionID"`
	AnswerVariantsID uint
	AnswerVariant    AnswerVariant `gorm:"ForeignKey:AnswerVariantsID"`
}
