package entities

import "gorm.io/gorm"

// Результат выполнения квиза для User
type Result struct {
	gorm.Model
	UserID                  uint
	User                    User `gorm:"ForeignKey:UserID"`
	QuizID                  uint
	Quiz                    Quiz `gorm:"ForeignKey:QuizID"`
	QuestionsAmount         int  // количество вопросов в квизе
	QuestionsAnsweredAmount int  // количество правильных ответов
	Percent                 int  // процент праивильных ответов
}
