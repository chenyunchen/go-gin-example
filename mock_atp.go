package main

func init() {
	tplMap.Store("vgam", Template{
		Resp: `{"data":{` +
			`,"url":"{{ index .query.url 0 }}"` +
			`}}`,
		CallbackURL: `{{ index .query.notify_url 0 }}` +
			`?time={{ index .query.time 0 }}` +
			`&sign={{ index .query.sign 0 }}`,
		// header 首字 要大寫 ex: X-Time, X-Sign
		CallbackHeader: `{` +
			`"utctime": ["{{ index .header.Time 0 }}"]` +
			`,"sign": ["{{ index .header.Sign 0 }}"]` +
			`}`,
		CallbackBody: `{` +
			`"merchantOrderNo":"{{.body.merchantOrderNo}}"` +
			`,"orderUID":"{{.body.orderUID}}"` +
			`,"callbackMerchantURL":"{{.body.callbackMerchantURL}}"` +
			`,"playerID":"{{.body.playerID}}"` +
			`,"expectedAmount":"{{.body.expectedAmount}}"` +
			`,"actualAmount":"{{.body.actualAmount}}"` +
			`,"fee":"{{.body.fee}}"` +
			`,"feeType":"{{.body.feeType}}"` +
			`,"actualFee":"{{.body.actualFee}}"` +
			`,"status":"{{.body.status}}"` +
			`,"createdAt":"{{.body.createdAt}}"` +
			`,"updatedAt":"{{.body.updatedAt}}"` +
			`,"errorCode":"{{.body.errorCode}}"` +
			`}`,
	})
}
