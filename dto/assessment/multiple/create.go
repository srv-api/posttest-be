package multiple

type MultipleRequest struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	DetailID        string   `json:"detail_id"`
	CreatedBy       string   `json:"created_by"`
	QuestionType    string   `json:"question_type"`
	QuestionText    string   `json:"question_text"`
	AnswerOptions   []string `json:"answer_options"`
	Explanation     string   `json:"explanation"`
	PlaceholderText string   `json:"placeholder_text"`
	ImageBase64     string   `json:"image_base64"`   // Base64 encoded image (optional)
	ImageFilename   string   `json:"image_filename"` // Original filename (optional)
	ImageURL        string   `json:"image_url"`      // Will be set after processing
}

type MultipleResponse struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	DetailID        string   `json:"detail_id"`
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
	ID              string   `json:"id"`
	QuestionType    string   `json:"question_type"`
	QuestionText    string   `json:"question_text"`
	AnswerOptions   []string `json:"answer_options"`
	Explanation     string   `json:"explanation"`
	PlaceholderText string   `json:"placeholder_text"`
	ImageBase64     string   `json:"image_base64"`
	ImageFilename   string   `json:"image_filename"`
	ImageURL        string   `json:"image_url"`
}
