package main

func init() {
	tplMap.Store("vgam", Template{
		Resp: `{"status":"ok"}`,
		CallbackURL: `{{ index .query.notify_url 0 }}` +
			`?time={{ index .query.time 0 }}` +
			`&sign={{ index .query.sign 0 }}`,
		// header 首字 要大寫 ex: X-Time, X-Sign
		CallbackHeader: `{` +
			`"Time": ["{{ index .header.Time 0 }}"]` +
			`,"Sign": ["{{ index .header.Sign 0 }}"]` +
			`}`,
		CallbackBody: `{` +
			`"flashid":"{{.body.flashid}}"` +
			`,"merchant":"{{ index .query.merchant 0 }}"` +
			`,"status":"{{ index .query.status 0 }}"` +
			`,"payed_money":"{{ index .query.payed_money 0 }}"` +
			`,"upstream_order":"{{ index .query.upstream_order 0 }}"` +
			`,"repeat_pay":"{{ index .query.repeat_pay 0 }}"` +
			`,"merchant_order_id":"{{ index .query.merchant_order_id 0 }}"` +
			`,"message":"{{.body.message}}"` +
			`,"order_fee":"{{ index .query.order_fee 0 }}"` +
			`}`,
	})
}
