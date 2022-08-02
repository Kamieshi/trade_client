// Package model models from application
package model

// Position  model
type Position struct {
	ID               string `json:"id" readonly:"true"`
	UserID           string `json:"user_id"`
	Price            *Price `json:"price,omitempty" readonly:"true"`
	IsOpened         bool   `json:"is_opened" readonly:"true"`
	Profit           int64  `json:"profit" readonly:"true"`
	MaxProfit        int64  `json:"max_current_cost"`
	MinProfit        int64  `json:"min_current_cost"`
	CountBuyPosition uint32 `json:"count_buy_position"`
	IsSales          bool   `json:"is_sales"` // true/false : sale/buy
	IsFixes          bool   `json:"is_fixes"` // user limit or not
}
