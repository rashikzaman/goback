package model

type Collection struct {
	Items      interface{} `json:"items"`
	PageNumber int64       `json:"page_number"`
	PageSize   int64       `json:"page_size"`
}
