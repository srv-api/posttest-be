package entity

import (
	"time"
)

type MultipleQuestion struct {
	ID              string   `gorm:"primary_key;type:varchar(39)" json:"id"`
	UserID          string   `gorm:"type:uuid;not null;index"`
	CreatedBy       string   `gorm:"type:varchar(100);not null"`
	QuestionType    string   `form:"question_type" json:"question_type"`
	QuestionText    string   `gorm:"type:text;not null"`
	AnswerOptions   []string `gorm:"type:jsonb"`
	Explanation     string   `gorm:"type:text"`
	PlaceholderText string   `gorm:"type:text"`
	ImageURL        string   `gorm:"type:varchar(500)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
