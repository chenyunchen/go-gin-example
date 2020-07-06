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
