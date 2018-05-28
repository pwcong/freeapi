package model

type Area struct {
	BaseModel
	Name     string `json:"ip"`
	ParentID uint   `json:"action"`
	Depth    uint   `json:"depth"`
}
