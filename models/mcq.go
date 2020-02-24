package models

type Question struct {
	ID           int64
	Question     string
	Category     string
	SubCategory  string
	QuestionType string
	Answer       string
}
