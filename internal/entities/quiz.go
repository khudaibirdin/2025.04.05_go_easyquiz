package entities

import (
	"gorm.io/gorm"
)

// Модель вопроса в Квизе
// Имеет текст вопроса, последовательный номер
type Question struct {
	gorm.Model
	QuizID uint
	Quiz   Quiz `gorm:"ForeignKey:QuizID"`
	Number int
	Text   string
}

// Вариант ответа для Квиза
// Связан с вопросом
// Имеет текст ответа
// Имеет статус "правильный/не правильный"
type AnswerVariant struct {
	gorm.Model
	QuestionID uint
	Question   Question `gorm:"ForeignKey:QuestionID"`
	Text       string
	IsRight    bool
}

// Модель Квиза
// Связана с пользователем, создавшим квиз
type Quiz struct {
	gorm.Model
	UserID         uint
	User           User `gorm:"ForeignKey:UserID"`
	Theme          string
	TimeOutMinutes int // время в минутах на выполнение квиза
}
