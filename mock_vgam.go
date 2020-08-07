package main

func init() {
	tplMap.Store("vgam", Template{
		Resp: `{"data":{` +
			`"orderUID":"{{ index .query.orderUID 0 }}"` +
			`,"expectedAmount":"{{ index .query.expectedAmount 0 }}"` +
			`,"mappingAmount":"{{ index .query.mappingAmount 0 }}"` +
			`,"expiredAt":"{{ index .query.expiredAt 0 }}"` +
			`,"qrcodeURL":"{{ index .query.qrcodeURL 0 }}"` +
			`,"qrcodeqrcodeTrackingURL":"{{ index .query.qrcodeqrcodeTrackingURL 0 }}"` +
			`}}`,
		CallbackURL: `{{ index .query.notify_url 0 }}` +
			`?time={{ index .query.time 0 }}` +
			`&sign={{ index .query.sign 0 }}`,
		// header 首字 要大寫 ex: X-Time, X-Sign
		CallbackHeader: `{` +
			`"X-Utctime": ["{{ index .header.Time 0 }}"]` +
			`,"X-sign": ["{{ index .header.Sign 0 }}"]` +
			`}`,
		CallbackBody: `{` +
			`"id":"{{.body.id}}"` +
			`,"orderUID":"{{.body.orderUID}}"` +
			`,"newOrderID":"{{.body.newOrderID}}"` +
			`,"merchantID":"{{.body.merchantID}}"` +
			`,"merchantName":"{{.body.merchantName}}"` +
			`,"businessID":"{{.body.businessID}}"` +
			`,"merchantOrderNo":"{{.body.merchantOrderNo}}"` +
			`,"agentID":"{{.body.agentID}}"` +
			`,"customerServiceID":"{{.body.customerServiceID}}"` +
			`,"customerServiceAccountID":"{{.body.customerServiceAccountID}}"` +
			`,"strategy":"{{.body.strategy}}"` +
			`,"deviceType":"{{.body.deviceType}}"` +
			`,"paymentMode":"{{.body.paymentMode}}"` +
			`,"paymentType":"{{.body.paymentType}}"` +
			`,"accountRec":"{{.body.accountRec}}"` +
			`,"accountUIDRec":"{{.body.accountUIDRec}}"` +
			`,"accountPay":"{{.body.accountPay}}"` +
			`,"accountUIDPay":"{{.body.accountUIDPay}}"` +
			`,"accountCityRec":"{{.body.accountCityRec}}"` +
			`,"accountCityPay":"{{.body.accountCityPay}}"` +
			`,"bankNameRec":"{{.body.bankNameRec}}"` +
			`,"bankNamePay":"{{.body.bankNamePay}}"` +
			`,"bankCodeRec":"{{.body.bankCodeRec}}"` +
			`,"bankCodePay":"{{.body.bankCodePay}}"` +
			`,"userNameRec":"{{.body.userNameRec}}"` +
			`,"userNamePay":"{{.body.userNamePay}}"` +
			`,"remark":"{{.body.remark}}"` +
			`,"qrcodeType":"{{.body.qrcodeType}}"` +
			`,"qrcodeRealID":"{{.body.qrcodeRealID}}"` +
			`,"transID":"{{.body.transID}}"` +
			`,"playerUID":"{{.body.playerUID}}"` +
			`,"currencyCode":"{{.body.currencyCode}}"` +
			`,"expectedAmount":"{{.body.expectedAmount}}"` +
			`,"mappingAmount":"{{.body.mappingAmount}}"` +
			`,"actualAmount":"{{.body.query.actualAmount}}"` +
			`,"fee":"{{.body.fee}}"` +
			`,"feeType":"{{.body.feeType}}"` +
			`,"actualFee":"{{.body.actualFee}}"` +
			`,"balance":"{{.body.balance}}"` +
			`,"paidAt":"{{.body.paidAt}}"` +
			`,"manualTrigerAt":"{{.body.manualTrigerAt}}"` +
			`,"status":"{{.body.status}}"` +
			`,"errorCode":"{{.body.errorCode}}"` +
			`,"errorMsg":"{{.body.errorMsg}}"` +
			`,"imageURL":"{{.body.imageURL}}"` +
			`,"isRetriable":"{{.body.isRetriable}}"` +
			`,"callbackBsURL":"{{.body.callbackBsURL}}"` +
			`,"isCallbackBsSuccess":"{{.body.isCallbackBsSuccess}}"` +
			`,"callbackBsFailedCount":"{{.body.callbackBsFailedCount}}"` +
			`,"lastCallbackBsAt":"{{.body.lastCallbackBsAt}}"` +
			`,"isCallbackAgentSuccess":"{{.body.isCallbackAgentSuccess}}"` +
			`,"callbackAgentFailedCount":"{{.body.callbackAgentFailedCount}}"` +
			`,"lastCallbackAgentAt":"{{.body.lastCallbackAgentAt}}"` +
			`,"qrcodeSeenTimes":"{{.body.qrcodeSeenTimes}}"` +
			`,"lastQrcodeSeenIP":"{{.body.lastQrcodeSeenIP}}"` +
			`,"lastQrcodeSeenCountry":"{{.body.lastQrcodeSeenCountry}}"` +
			`,"lastQrcodeSeenCity":"{{.body.lastQrcodeSeenCity}}"` +
			`,"lastQrcodeSeenAt":"{{.body.lastQrcodeSeenAt}}"` +
			`,"clientIP":"{{.body.clientIP}}"` +
			`,"clientCountry":"{{.body.clientCountry}}"` +
			`,"clientCity":"{{.body.clientCity}}"` +
			`,"createdAt":"{{.body.createdAt}}"` +
			`,"expiredAt":"{{.body.expiredAt}}"` +
			`,"updatedAt":"{{.body.updatedAt}}"` +
			`,"updaterID":"{{.body.updaterID}}"` +
			`,"updaterName":"{{.body.updaterName}}"` +
			`,"orderCompletedAt":"{{.body.orderCompletedAt}}"` +
			`,"note":"{{.body.note}}"` +
			`,"noteImageURL":"{{.body.noteImageURL}}"` +
			`,"childOrders":"{{.body.childOrders}}"` +
			`}`,
	})
}
