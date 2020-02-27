package dao

import (
	"github.com/KHvic/study-backend/models"
	"github.com/jinzhu/gorm"
)

// QuestionDAO ...
type QuestionDAO interface {
	GetByID(id int64) (*models.Question, error)
	GetBySubCat(sub string) ([]*models.Question, error)
	GetBySubCatRandK(sub string, k int) ([]*models.Question, error)
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

// GetBySubCat get questions by sub category
func (*QuestionDAOImpl) GetBySubCat(sub string) ([]*models.Question, error) {
	var questions []*models.Question
	err := db.Where("sub_category = ?", sub).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// GetBySubCatRandK get K random questions by sub category
func (*QuestionDAOImpl) GetBySubCatRandK(sub string, k int) ([]*models.Question, error) {
	var questions []*models.Question
	err := db.Limit(k).Where("sub_category = ?", sub).Order(gorm.Expr("rand()")).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
