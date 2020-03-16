package models

import (
	"database/sql/driver"
	"encoding/json"
)

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

// Options represent options type
type Options [][]string

// Value indicate how database store Options type
func (o Options) Value() (driver.Value, error) {
	bytes, _ := json.Marshal(o)
	return string(bytes), nil
}

// Scan indicate how application read Options type
func (o *Options) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, o)
}

// Answers represent answer type
type Answers [][]int

// Value indicate how database store Answers type
func (a Answers) Value() (driver.Value, error) {
	bytes, _ := json.Marshal(a)
	return string(bytes), nil
}

// Scan indicate how application read Answers type
func (a *Answers) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, a)
}

// Question ...
type Question struct {
	ID          int64        `json:"id"`
	Category    string       `json:"category"`
	SubCategory string       `json:"subcategory"`
	Type        QuestionType `json:"type"`
	Description string       `json:"description"`
	Options     Options      `json:"options"` // e.g [["option a", "option b"], ["option c, option d"]]
	Answers     Answers      `json:"answers"` // e.g [[1][0]]
	Image       *string      `json:"image,omitempty"`
}

// TableName ...
func (*Question) TableName() string {
	return "question"
}
