package models

type Shan6Info struct {
	UserID     string `json:"user_id"`
	DeviceIP   string `json:"device_ip"`
	DeviceID   string `json:"device_id"`
	Name       string `json:"name,omitempty"`
	BankCode   string `json:"bank_code,omitempty"`
	Tel        string `json:"tel,omitempty"`
	DeviceType string `json:"device_type"`
	AliName    string `json:"ali_name,omitempty"`
}

type Shan6DepositeRequest struct {
	MerchantID string    `json:"merchant_id"`
	OrderID    string    `json:"order_id"`
	NotifyUrl  string    `json:"notify_url"`
	BillPrice  string    `json:"bill_price"`
	Extra      string    `json:"extra,omitempty"`
	Info       Shan6Info `json:"info"`
}

type Shan6DepositeResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	FlashID   string `json:"flashid"`
	QrCode    string `json:"qr_code"`
	Expire    string `json:"expire"`
	PayeeName string `json:"payee_name"`
}

type Shan6DepositeCallbackRequest struct {
	FlashID         string `json:"flashid"`
	Merchant        string `json:"merchant"`
	Status          int    `json:"status"`
	PayedMoney      string `json:"payed_money"`
	UpstreamOrder   string `json:"upstream_order"`
	RepeatPay       string `json:"repeat_pay"`
	MerchantOrderID string `json:"merchant_order_id"`
	Message         string `json:"message"`
	OrderFee        string `json:"order_fee"`
}
