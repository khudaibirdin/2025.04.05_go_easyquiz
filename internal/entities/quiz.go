package entities

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuizID  uint
	Number  int
	Text    string
	Answers []string
	Right   int
}

type Quiz struct {
	gorm.Model
	User  User
	Theme string
}
