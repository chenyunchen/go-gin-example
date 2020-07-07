package models

type Beiming1DepositeResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    Beiming1DepositeData `json:"data"`
}

type Beiming1DepositeData struct {
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
	Remark        string `json:"remark"`
}

type Beiming1DepositeCallbackRequest struct {
	MerchantCode  string `json:"merchant_code"`
	MerchantOrder string `json:"merchant_order"`
	FlashID       string `json:"flash_id"`
	ReceiptPrice  string `json:"receipt_price"`
	PayedTime     string `json:"payed_time"`
	OrderType     int    `json:"order_type"`
	OrderMode     int    `json:"order_mode"`
	UserLevel     int    `json:"user_level"`
	TargetAccount string `json:"target_account"`
	TargetName    string `json:"target_name"`
	TradeTime     string `json:"trade_time"`
	RepeatPay     string `json:"repeat_pay"`
}
