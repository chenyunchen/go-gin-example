package models

type BankcardInfo struct {
	Telephone  string `json:"tel"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	DeviceIP   string `json:"device_ip"`
	DeviceID   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	BankCard   string `json:"bank_card"`
	BankCode   string `json:"bank_code"`
}

type BankcardWithdrawRequest struct {
	MerchantID string       `json:"merchant_id"`
	OrderID    string       `json:"order_id"`
	NotifyURL  string       `json:"notify_url"`
	BillPrice  string       `json:"bill_price"`
	Info       BankcardInfo `json:"info"`
	Extra      string       `json:"extra"`
}

type BankcardWithdrawResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
	FlashID string `json:"flashid"`
}

type BankCardBankInfo struct {
	ID        string `json:"ID"`
	CreatedAt string `json:"CreatedAt"`
	Name      string `json:"Name"`
	BankCode  string `json:"BankCode"`
}

type BankcardWithdrawCallbackRequest struct {
	FlashID         string `json:"flashid"`
	Merchant        string `json:"merchant"`
	Status          int    `json:"status"`
	PayedMoney      string `json:"payed_money"`
	MerchantOrderID string `json:"merchant_order_id"`
	PayedTime       string `json:"payed_time"`
	Message         string `json:"message,omitempty"`
	OrderFee        string `json:"order_fee"`
}
