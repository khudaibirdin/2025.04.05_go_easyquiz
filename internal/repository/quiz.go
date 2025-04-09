package repository

import (
	"app/internal/entities"
	"errors"

	"gorm.io/gorm"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{
		db: db,
	}
}

func (r *QuizRepository) CreateQuiz(quiz entities.Quiz) (uint, error) {
	result := r.db.Create(&quiz)
	return quiz.ID, result.Error
}

func (r *QuizRepository) CreateQuestions(questions []entities.Question) ([]uint, error) {
	result := r.db.Create(questions)
	if result.Error != nil {
		return nil, result.Error
	}
	var ids []uint
	for _, question := range questions {
		ids = append(ids, question.ID)
	}
	return ids, nil
}

func (r *QuizRepository) GetAllQuestions(quizID uint) (*[]entities.Question, error) {
	var questions []entities.Question
	result := r.db.Where("quiz_id = ?", quizID).Find(questions)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &questions, result.Error
}

func (r *QuizRepository) GetQuestion(quizID, questionID uint) (*entities.Question, error) {
	question := &entities.Question{}
	result := r.db.First(question, questionID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return question, result.Error
}

func (r *QuizRepository) GetQuestionByNumber(quizID uint, number int) (*entities.Question, error) {
	question := &entities.Question{}
	result := r.db.Where("quiz_id = ? AND number = ?", quizID, number).First(question)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return question, result.Error
}
