package dto

type Search struct {
	ID     uint64 `json:"id"`
	Column string `json:"column"`
	Action string `json:"action"`
	Query  string `json:"query"`
}
