package model

type Area struct {
	BaseModel
	Name     string `json:"name"`
	ParentID uint   `json:"parent_id"`
	Depth    uint   `json:"depth"`
}
