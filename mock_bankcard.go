package main

func init() {
	tplMap.Store("bankcard", Template{
		Resp: `{` +
			`"message":"{{ index .query.message 0 }}"` +
			`,"flashid":"{{ index .query.flashid 0 }}"` +
			`,"status":"{{ index .query.status 0 }}"` +
			`}`,
		CallbackURL: `{{ index .query.notify_url 0 }}` +
			`?time={{ index .query.time 0 }}` +
			`&sign={{ index .query.sign 0 }}` +
			`&origin={{ index .query.origin 0 }}`,
		// header 首字 要大寫 ex: X-Time, X-Sign
		CallbackHeader: `{` +
			`"time": ["{{ index .header.Time 0 }}"]` +
			`,"sign": ["{{ index .header.Sign 0 }}"]` +
			`}`,
		CallbackBody: `{` +
			`"flashid":"{{.body.flashid}}"` +
			`,"merchant":"{{.body.merchant}}"` +
			`,"status":"{{.body.status}}"` +
			`,"payed_money":"{{.body.payed_money}}"` +
			`,"merchant_order_id":"{{.body.merchant_order_id}}"` +
			`,"payed_time":"{{.body.payed_time}}"` +
			`,"message":"{{.body.message}}"` +
			`,"order_fee":"{{.body.order_fee}}"` +
			`}`,
	})
}
