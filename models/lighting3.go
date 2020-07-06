package models

type Lighting3DepositeResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	TradeID string `json:"trade_id"`
	OrderID string `json:"order_id"`
	PayURL  string `json:"pay_url"`
}
