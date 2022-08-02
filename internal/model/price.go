package model

// Price struct company
type Price struct {
	CompanyID string `json:"CompanyID,omitempty"`
	Name      string `json:"Name,omitempty" `
	Ask       uint32 `json:"Ask,omitempty"`
	Bid       uint32 `json:"Bid,omitempty"`
	Time      string `json:"Time,omitempty"`
}
