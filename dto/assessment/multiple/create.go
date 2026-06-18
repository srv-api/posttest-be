package multiple

import "mime/multipart"

type MultipleRequest struct {
	ID              string                `form:"id" json:"id"`
	UserID          string                `form:"user_id" json:"user_id"`
	Link            string                `form:"link" json:"link"`
	Status          string                `form:"status" json:"status"`
	CreatedBy       string                `form:"created_by" json:"created_by"`
	QuestionType    string                `form:"question_type" json:"question_type"`
	QuestionText    string                `form:"question_text" json:"question_text"`
	AnswerOptions   []string              `form:"answer_options" json:"answer_options"`
	Explanation     string                `form:"explanation" json:"explanation"`
	PlaceholderText string                `form:"placeholder_text" json:"placeholder_text"`
	ImageURL        string                `form:"image_url" json:"image_url"`
	Image           *multipart.FileHeader `form:"image" json:"-"`
	ImageIndex      int                   `form:"image_index" json:"image_index"` // -1 = no image, 0,1,2 = file index
}

type MultipleResponse struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	Status          string   `form:"status" json:"status"`
	Link            string   `json:"link"`
	CreatedBy       string   `json:"created_by"`
	QuestionType    string   `json:"question_type"`
	QuestionText    string   `json:"question_text"`
	AnswerOptions   []string `json:"answer_options"`
	Explanation     string   `json:"explanation"`
	PlaceholderText string   `json:"placeholder_text"`
	ImageURL        string   `json:"image_url"`
	CreatedAt       string   `json:"created_at,omitempty"`
	UpdatedAt       string   `json:"updated_at,omitempty"`
}

type MultipleUpdateRequest struct {
	ID              string                `form:"id" json:"id"`
	Link            string                `json:"link"`
	QuestionType    string                `form:"question_type" json:"question_type"`
	QuestionText    string                `form:"question_text" json:"question_text"`
	AnswerOptions   []string              `form:"answer_options" json:"answer_options"`
	Explanation     string                `form:"explanation" json:"explanation"`
	PlaceholderText string                `form:"placeholder_text" json:"placeholder_text"`
	ImageURL        string                `form:"image_url" json:"image_url"`
	Image           *multipart.FileHeader `form:"image" json:"-"`
}
