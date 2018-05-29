package model

type Page struct {
	PageNo      int         `json:"page_no"`
	PageSize    int         `json:"page_size"`
	CurrentSize int         `json:"current_size"`
	TotalSize   int         `json:"total_size"`
	Data        interface{} `json:"data"`
}
