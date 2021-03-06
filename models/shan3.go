package models

type Shan3DepositeResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	TradeID string `json:"trade_id"`
	OrderID string `json:"order_id"`
	PayURL  string `json:"pay_url"`
}

type Shan3DepositeCallbackRequest struct {
	Tid           string `json:"tid"`
	Status        int    `json:"status"`
	TradeAmount   string `json:"trade_amount"`
	ReceiptAmount string `json:"receipt_amount"`
	TradeID       string `json:"trade_id"`
	OrderID       string `json:"order_id"`
	UserData      string `json:"userdata"`
	OrderFee      string `json:"order_fee"`
}
