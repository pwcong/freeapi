package model

type Attachment struct {
	BaseModel
	Symbol   string `json:"symbol"`
	Filename string `json:"filename"`
	Year     string `json:"year"`
	Month    string `json:"month"`
	Date     string `json:"date"`
	ExtName  string `json:"extname"`
}
