# Mock Server
Mock Server for Polypay, Bankcard, Shan3, Shan6, Beiming1, Beiming2, Exchange

## API

* [Polypay 聚合支付](#header_1)
	* [Request](#header_1_1)
	    * [Example](#header_1_1_2)
* [Bankcard 銀行卡](#header_2)
	* [Request](#header_2_1)
	    * [Example](#header_2_1_2)
* [Shan3 閃3支付](#header_3)
	* [Request](#header_3_1)
	    * [Example](#header_3_1_2)
* [Shan6 閃6支付](#header_4)
	* [Request](#header_4_1)
	    * [Example](#header_4_1_2)
* [Beiming1 北冥1支付](#header_5)
	* [Request](#header_5_1)
	    * [Example](#header_5_1_2)
* [Beiming2 北冥2支付](#header_6)
	* [Request](#header_6_1)
	    * [Example](#header_6_1_2)
* [Exchange 支付寶兌換網關](#header_7)
	* [Request](#header_7_1)
	    * [Example](#header_7_1_2)

<a name="header_1"></a>
### Polypay 聚合支付

<a name="header_1_1"></a>
#### Request

```
GET /deposite/polypay
```

|Request Query-String|Description|
|:-:|:-:|
|status|交易狀態|
|message|訊息描述|
|trade_id|平台訂單號|
|order_id|商戶訂單號|
|pay_url|付款QRCode連結|
|account|用戶帳號|
|name|用戶名稱|
|trade_amount|訂單金額|
|receipt_amount|實際付款金額|
|upstream_order|交易流水號|
|repeat_pay|是否重複支付|
|notify_url|接收回調地址|
|sign|加密簽章|
|time|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_1_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/deposite/polypay?
status=10000&
message=success&
trade_id=2018080900473244E406&
order_id=A001-2304010&
pay_url=https://t.cn/pay&
account=13838383838&
name=alex&
trade_amount=0.01&
receipt_amount=0.01&
upstream_order=201800125481245844575112541&
repeat_pay=false&
notify_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
time=1492257443&
retry=10
```

<a name="header_2"></a>
### Bankcard 銀行卡

<a name="header_2_1"></a>
#### Request

```
GET /withdraw/bankcard
```

|Request Query-String|Description|
|:-:|:-:|
|status|交易狀態|
|message|訊息描述|
|flash_id|平台訂單號|
|merchant|商戶ID|
|payed_money|訂單金額|
|merchant_order_id|商戶訂單號|
|payed_time|訂單時間|
|notify_url|接收回調地址|
|sign|加密簽章|
|time|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_2_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/withdraw/bankcard?
status=10000&
message=请求成功&
flash_id=fla1234&
merchant=A0001&
payed_money=1000&
merchant_order_id=ord1234&
payed_time=2020-01-02 15:15:03&
notify_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
time=1492257443&
retry=10
```

<a name="header_3"></a>
### Shan3 閃3支付

<a name="header_3_1"></a>
#### Request

```
GET /deposite/shan3
```

|Request Query-String|Description|
|:-:|:-:|
|status|交易狀態|
|message|訊息描述|
|trade_id|平台訂單號|
|order_id|商戶訂單號|
|pay_url|付款QRCode連結|
|tid|平台訂單號|
|trade_amount|訂單金額|
|receipt_amount|實際付款金額|
|user_data|用戶資料|
|notify_url|接收回調地址|
|sign|加密簽章|
|time|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_3_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/deposite/shan3?
status=10000&
message=success&
trade_id=20171104193517B84121&
order_id=e2dca4fcfc0e4e139b31c203a666b6cc&
pay_url=http://127.0.0.1/api/pay.jsp?id=2%26c=e920dc57680ce4bf2b89654440adfee5&
tid=546564&
trade_amount=0.01&
receipt_amount=0.01&
user_data=abcd&
notify_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
time=1492257443&
retry=10
```

<a name="header_4"></a>
### Shan6 閃6支付

<a name="header_4_1"></a>
#### Request

```
GET /deposite/shan6
```

|Request Query-String|Description|
|:-:|:-:|
|status|交易狀態|
|message|訊息描述|
|flash_id|平台訂單號|
|qr_code|付款QRCode連結|
|payee_name|用戶名稱|
|expire|到期時間|
|merchant|商戶ID|
|payed_money|實際付款金額|
|upstream_order|交易流水號|
|repeat_pay|是否重複支付|
|merchant_order_id|商戶訂單號|
|notify_url|接收回調地址|
|sign|加密簽章|
|time|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_4_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/deposite/shan6?
status=10000&
message=订单生成成功&
flash_id=DA20170929042702499412xxxxxxxx&
qr_code=https://qr.alipay.com/xxxxxxxxxx&
payee_name=alex&
expire=2000-01-02 03:04:05&
merchant=xxxx&
payed_money=9.78&
upstream_order=201800125481245844575112541&
repeat_pay=no&
merchant_order_id=368571021&
notify_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
time=1492257443&
retry=10
```

<a name="header_5"></a>
### Beiming1 北冥1支付

<a name="header_5_1"></a>
#### Request

```
GET /deposite/beiming1
```

|Request Query-String|Description|
|:-:|:-:|
|status|交易狀態|
|message|訊息描述|
|merchant_code|商戶ID|
|merchant_order|商戶訂單號|
|flash_id|平台訂單號|
|pay_url|付款QRCode連結|
|bill_price|訂單金額|
|card_num|收款卡號|
|accept_name|收款人姓名|
|bank_name|銀行名稱|
|province|州、省|
|city|城市|
|country|國家|
|node|銀行分行|
|created_at|創建時間|
|remark|備註|
|receipt_price|實際付款金額|
|payed_time|訂單時間|
|order_type|訂單來源|
|order_mode|訂單模式|
|user_level|用戶等級|
|target_account|用戶帳號|
|target_name|用戶名稱|
|trade_time|交易時間|
|repeat_pay|是否重複支付|
|notify_url|接收回調地址|
|sign|加密簽章|
|time|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_5_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/deposite/beiming1?
status=1&
message=创建成功&
merchant_code=ONION&
merchant_order=test_order3&
flash_id=DA20170929042702499412xxxxxxxx&
pay_url=https://receipt.trqsxt.com/receipt-html/wait/&
bill_price=10&
card_num=9560580597580076&
accept_name=alex&
bank_name=支付宝支付&
province=&
city=taipei&
country=taiwan&
node=taipei branch&
created_at=2020-01-02 15:15:03&
remark=remark&
receipt_price=10&
payed_time=2020-05-04 15:54:13&
order_type=2&
order_mode=0&
user_level=2&
target_account=9560580597580076&
target_name=alex&
trade_time=2020-01-02 15:15:03&
repeat_pay&
notify_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
time=1492257443&
retry=10
```

<a name="header_6"></a>
### Beiming2 北冥2支付

<a name="header_6_1"></a>
#### Request

```
GET /deposite/beiming2
```

|Request Query-String|Description|
|:-:|:-:|
|status|交易狀態|
|message|訊息描述|
|merchant_code|商戶ID|
|merchant_order|商戶訂單號|
|flash_id|平台訂單號|
|pay_url|付款QRCode連結|
|bill_price|訂單金額|
|card_num|收款卡號|
|accept_name|收款人姓名|
|bank_name|銀行名稱|
|province|州、省|
|city|城市|
|country|國家|
|node|銀行分行|
|created_at|創建時間|
|receipt_price|實際付款金額|
|payed_time|訂單時間|
|target_account|用戶帳號|
|target_name|用戶名稱|
|notify_url|接收回調地址|
|sign|加密簽章|
|time|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_6_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/deposite/beiming2?
status=1&
message=订单创建成功&
merchant_code=ONION&
merchant_order=test_order3&
flash_id=123456&
pay_url=https://receipt.trqsxt.com/receipt-html/wait/&
bill_price=10&
card_num=9560580597580076&
accept_name=alex&
bank_name=支付宝支付&
province=&
city=taipei&
country=taiwan&
node=taipei branch&
created_at=2020-01-02 15:15:03&
receipt_price=10&
payed_time=2020-05-04 15:54:13&
target_account=9560580597580076&
target_name=alex&
notify_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
time=1492257443&
retry=10
```

<a name="header_7"></a>
### Exchange 支付寶兌換網關

<a name="header_7_1"></a>
#### Request

```
GET /withdraw/exchange
```

|Request Query-String|Description|
|:-:|:-:|
|body|請求回覆內容|
|code|交易狀態|
|message|訊息描述|
|order_id|平台訂單號|
|trade_No|商戶訂單號|
|service_charge|手續費|
|order_amout|訂單金額|
|pay_account|用戶帳號|
|pay_time|訂單時間|
|callback_url|接收回調地址|
|sign|加密簽章|
|t|Unix時間戳|
|retry|回調地址重試次數|

<a name="header_7_1_2"></a>
Example:

```
https://mockserver-gbhzzlku2a-de.a.run.app/withdraw/exchange?
body=Success&
code=10000&
message=Success&
order_id=ord1234&
trade_No=20170113110070001502510003065349&
service_charge=0&
order_amout=100&
pay_account=lusu@qq.com&
pay_time=2020-01-02 15:15:03&
callback_url=http://localhost:8000/&
sign=1cf245d08e3a2fac8fc93ab964542a1&
t=1492257443&
retry=10
```
