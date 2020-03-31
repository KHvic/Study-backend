package dao

import (
	"github.com/KHvic/quiz-backend/models"
	"github.com/jinzhu/gorm"
)

// QuestionDAO ...
type QuestionDAO interface {
	GetByID(id int64) (*models.Question, error)
	GetBySubCat(sub string) ([]*models.Question, error)
	MGetBySubCat(sub string, count int) ([]*models.Question, error)
	GetByOffsetAndLimit(offset, limit int64) ([]*models.Question, error)
}

// QuestionDAOImpl ...
type QuestionDAOImpl struct{}

// NewQuestionDAO ...
func NewQuestionDAO() QuestionDAO {
	return &QuestionDAOImpl{}
}

// GetByID get question by id
func (*QuestionDAOImpl) GetByID(id int64) (*models.Question, error) {
	question := &models.Question{}
	err := db.Where("id = ?", id).First(question).Error
	if err != nil {
		return nil, err
	}
	return question, nil
}

// GetByOffsetAndLimit get questions by offset and limit
func (*QuestionDAOImpl) GetByOffsetAndLimit(offset, limit int64) ([]*models.Question, error) {
	var questions []*models.Question
	err := db.Offset(offset).Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// GetBySubCat get questions by sub category
func (*QuestionDAOImpl) GetBySubCat(sub string) ([]*models.Question, error) {
	var questions []*models.Question
	err := db.Where("sub_category = ?", sub).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// MGetBySubCat get a batch of random questions by sub category
func (*QuestionDAOImpl) MGetBySubCat(sub string, count int) ([]*models.Question, error) {
	var questions []*models.Question
	err := db.Limit(count).Where("sub_category = ?", sub).Order(gorm.Expr("rand()")).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
