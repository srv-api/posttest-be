package multiple

type MultipleBatchRequest struct {
	Questions []MultipleRequest `json:"questions" validate:"required,min=1,max=100,dive"`
}

type MultipleBatchResponse struct {
	SuccessCount int                         `json:"success_count"`
	FailedCount  int                         `json:"failed_count"`
	TotalCount   int                         `json:"total_count"`
	Results      []MultipleBatchItemResponse `json:"results"`
}

type MultipleBatchItemResponse struct {
	Index        int               `json:"index"`
	Success      bool              `json:"success"`
	QuestionID   string            `json:"question_id,omitempty"`
	QuestionText string            `json:"question_text,omitempty"`
	Data         *MultipleResponse `json:"data,omitempty"`
	Error        string            `json:"error,omitempty"`
}
