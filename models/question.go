package models

// Question ...
type Question struct {
	ID           int64
	Category     string
	SubCategory  string
	QuestionType string
	Description  string
	Question     string // TODO: custom type, e.g [["option a", "option b"], ["option"]]
	Answer       string // TODO: custom type, e.g [[1,2][1]]
	Image        *string
}

// TableName ...
func (*Question) TableName() string {
	return "question"
}
