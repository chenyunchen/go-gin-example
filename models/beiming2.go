package models

type Beiming2DepositeRequest struct {
	MerchantCode  string `json:"merchant_code"`
	MerchantOrder string `json:"merchant_order"`
	NotifyUrl     string `json:"notify_url"`
	BillPrice     string `json:"bill_price"`
	UserId        string `json:"user_id"`
	UserName      string `json:"user_name,omitempty"`
	DeviceIP      string `json:"device_ip,omitempty"`
	DeviceID      string `json:"device_id,omitempty"`
	DeviceType    string `json:"device_type,omitempty"`
	Tel           string `json:"tel,omitempty"`
}

type Beiming2DepositeResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    Beiming2DepositeData `json:"data"`
}

type Beiming2DepositeData struct {
	MerchantCode  string `json:"merchant_code"`
	MerchantOrder string `json:"merchant_order"`
	FlashID       string `json:"flash_id"`
	PayUrl        string `json:"pay_url"`
	Status        int    `json:"status"`
	Message       string `json:"message"`
	BillPrice     string `json:"bill_price"`
	CardNum       string `json:"card_num"`
	AcceptName    string `json:"accept_name"`
	BankName      string `json:"bank_name"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Node          string `json:"node"`
	CreatedAt     string `json:"created_at"`
}

type Beiming2DepositeCallbackRequest struct {
	MerchantCode  string `json:"merchant_code"`
	MerchantOrder string `json:"merchant_order"`
	FlashId       string `json:"flash_id"`
	ReceiptPrice  string `json:"receipt_price"`
	PayedTime     string `json:"payed_time"`
	PayInfo       string `json:"pay_info"`
}
