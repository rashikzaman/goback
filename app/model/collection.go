package model

type Collection struct {
	Items      interface{} `json:"items"`
	PageNumber int         `json:"page_number"`
	PageSize   int         `json:"page_size"`
}
