package models

import "time"

type PolypayInfo struct {
	Telephone  string `binding:"required" json:"telephone"`
	PlayerID   string `binding:"required" json:"player_id"`
	Name       string `binding:"required" json:"name"`
	DeviceIP   string `binding:"required" json:"device_ip"`
	DeviceID   string `binding:"required" json:"device_id"`
	DeviceType string `binding:"required" json:"device_type"`
	PayAccount string `binding:"required" json:"pay_account"`
	Metadata   string `json:"metadata;omitempty"`
}

type PolypayDepositeRequest struct {
	MerchantID string      `binding:"required" json:"merchant_id"`
	OrderID    string      `binding:"required" json:"order_id"`
	NotifyURL  string      `binding:"required" json:"notify_url"`
	Payment    string      `binding:"required" json:"payment"`
	BillPrice  string      `binding:"required" json:"bill_price"`
	Info       PolypayInfo `binding:"required" json:"info"`
}

type PolypayDepositeResponse struct {
	Message string `json:"message"`
	OrderID string `json:"order_id,omitempty"`
	PayURL  string `json:"pay_url,omitempty"`
	Status  int    `json:"status"`
	TradeID string `json:"trade_id,omitempty"`
	Account string `json:"payee_account,omitempty"`
	Name    string `json:"payee_realname,omitempty"`
}

type PolypayParam struct {
	MerchantID   uint64    `gorm:"column:merchant_id;type:bigint;NOT NULL"`
	PlatformID   string    `gorm:"column:platform_id;type:varchar(20);NOT NULL"`
	Platform     string    `gorm:"column:platform;type:varchar(30);NOT NULL"`
	Host         string    `gorm:"column:host;type:varchar(100);NOT NULL"`
	PolypayID    string    `gorm:"column:polypay_id;type:varchar(50);NOT NULL"`
	PolypayToken string    `gorm:"column:polypay_token;type:varchar(100);NOT NULL"`
	Origin       string    `gorm:"column:origin;type:varchar(20);NOT NULL"`
	OriginToken  string    `gorm:"column:origin_token;type:varchar(100);NOT NULL"`
	IsEnable     bool      `gorm:"column:is_enable;type:boolean;NOT NULL;DEFAULT:true;"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;DEFAULT:now()"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;DEFAULT:NULL"`
}

type PolypayDepositeCallbackRequest struct {
	TradeID       string `json:"trade_id"` //platform order number
	OrderID       string `json:"order_id"` //tracking number
	Status        int    `json:"status"`
	TradeAmount   string `json:"trade_amount"`   //estimated cost
	ReceiptAmount string `json:"receipt_amount"` //actual cost
	UpstreamOrder string `json:"upstream_order"` //payment number
	RepeatPay     bool   `json:"repeat_pay"`
}
