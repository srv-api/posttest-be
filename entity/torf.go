package entity

import (
	"time"
)

type TorfQuestion struct {
	ID              string   `gorm:"primaryKey;type:uuid"`
	UserID          string   `gorm:"type:uuid;not null;index"`
	DetailID        string   `gorm:"type:varchar(100);not null"`
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
