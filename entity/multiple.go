package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan StringArray: unsupported type %T", value)
	}
	return json.Unmarshal(bytes, s)
}

func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return json.Marshal(s)
}

type MultipleQuestion struct {
	ID              string      `gorm:"primary_key;type:varchar(39)" json:"id"`
	UserID          string      `gorm:"type:varchar(36);index" json:"user_id"`
	Status          string      `gorm:"type:varchar(50)" json:"status"`
	Link            string      `gorm:"type:varchar(255)" json:"link"`
	CreatedBy       string      `gorm:"type:varchar(100);not null" json:"created_by"`
	QuestionType    string      `gorm:"type:varchar(50)" json:"question_type"`
	QuestionText    string      `gorm:"type:text;not null" json:"question_text"`
	AnswerOptions   StringArray `gorm:"type:jsonb" json:"answer_options"`
	Explanation     string      `gorm:"type:text" json:"explanation"`
	PlaceholderText string      `gorm:"type:text" json:"placeholder_text"`
	ImageURL        string      `gorm:"type:varchar(500)" json:"image_url"`
	CreatedAt       time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

func (MultipleQuestion) TableName() string {
	return "multiple_questions"
}
