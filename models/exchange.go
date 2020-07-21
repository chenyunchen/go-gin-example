package models

type ExchangeWithdrawRequest struct {
	OrderID     string `json:"order_id"`
	Type        int    `json:"type"`
	Amount      string `json:"amount"`
	FullName    string `json:"fullname"`
	Account     string `json:"account"`
	CallbackURL string `json:"callback_url"`
	DeviceType  string `json:"device_type"`
	DeviceID    string `json:"device_id"`
	DeviceIP    string `json:"device_ip"`
	PlayerID    string `json:"player_id"`
}

type ExchangeWithdrawCallbackRequest struct {
	OrderID       string `json:"order_id"`
	TradeNo       string `json:"trade_no"`
	ServiceCharge string `json:"service_charge"`
	OrderAmout    string `json:"order_amout"`
	PayAccount    string `json:"pay_account"`
	Code          int    `json:"code"`
	PayTime       string `json:"pay_time"`
	Message       string `json:"msg"`
	OrderFee      string `json:"order_fee"`
}
