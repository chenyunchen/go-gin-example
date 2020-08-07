package main

func init() {
	tplMap.Store("vgam", Template{
		Resp: `{` +
			`"message":"{{ index .query.message 0 }}"` +
			`,"order_id":"{{ index .query.order_id 0 }}"` +
			`,"pay_url":"{{ index .query.pay_url 0 }}"` +
			`,"status":"{{ index .query.status 0 }}"` +
			`,"trade_id":"{{ index .query.trade_id 0 }}"` +
			`,"payee_account":"{{ index .query.payee_account 0 }}"` +
			`,"payee_realname":"{{ index .query.payee_realname 0 }}"` +
			`}`,
		CallbackURL: `{{ index .query.notify_url 0 }}` +
			`?time={{ index .query.time 0 }}` +
			`&sign={{ index .query.sign 0 }}` +
			`&origin={{ index .query.origin 0 }}` +
			`&origin_sign={{ index .query.origin_sign 0 }}`,
		// header 首字 要大寫 ex: X-Time, X-Sign
		CallbackHeader: `{` +
			`"time": ["{{ index .header.Time 0 }}"]` +
			`,"sign": ["{{ index .header.Sign 0 }}"]` +
			`}`,
		CallbackBody: `{` +
			`"trade_id":"{{.body.trade_id}}"` +
			`,"order_id":"{{.body.order_id}}"` +
			`,"status":"{{.body.status}}"` +
			`,"trade_amount":"{{.body.trade_amount}}"` +
			`,"receipt_amount":"{{.body.receipt_amount}}"` +
			`,"upstream_order":"{{.body.upstream_order}}"` +
			`,"repeat_pay":"{{.body.repeat_pay}}"` +
			`,"order_fee":"{{.body.order_fee}}"` +
			`}`,
	})
}
