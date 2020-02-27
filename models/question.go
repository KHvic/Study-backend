package models

// QuestionType represent type of question
type QuestionType int

const (
	// FillBlank ...
	FillBlank = 1
	// MCQ ...
	MCQ = 2
	// MultiSelect ...
	MultiSelect = 3
)

// Question ...
type Question struct {
	ID          int64
	Category    string
	SubCategory string
	Type        QuestionType
	Description string
	Question    string // e.g [["option a", "option b"], ["option c, option d"]]
	Answer      string // e.g [[1][0]]
	Image       *string
}

// TableName ...
func (*Question) TableName() string {
	return "question"
}
