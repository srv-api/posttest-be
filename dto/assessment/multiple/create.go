package multiple

import (
	"mime/multipart"
)

type MultipleRequest struct {
	ID              string                `form:"id" json:"id"`
	UserID          string                `form:"user_id" json:"user_id"`
	DetailID        string                `form:"detail_id" json:"detail_id"`
	CreatedBy       string                `form:"created_by" json:"created_by"`
	QuestionType    string                `form:"question_type" json:"question_type"`
	QuestionText    string                `form:"question_text" json:"question_text"`
	AnswerOptions   []string              `form:"answer_options" json:"answer_options"`
	Explanation     string                `form:"explanation" json:"explanation"`
	PlaceholderText string                `form:"placeholder_text" json:"placeholder_text"`
	ImageURL        string                `form:"image_url"`
	Image           *multipart.FileHeader `form:"image"`
}

type MultipleResponse struct {
	ID              string   `form:"id" json:"id"`
	UserID          string   `form:"user_id" json:"user_id"`
	DetailID        string   `form:"detail_id" json:"detail_id"`
	CreatedBy       string   `form:"created_by" json:"created_by"`
	QuestionType    string   `form:"question_type" json:"question_type"`
	QuestionText    string   `form:"question_text" json:"question_text"`
	AnswerOptions   []string `form:"answer_options" json:"answer_options"`
	Explanation     string   `form:"explanation" json:"explanation"`
	PlaceholderText string   `form:"placeholder_text" json:"placeholder_text"`
	ImageURL        string   `form:"image_url"`
}
