package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var DestinationAccountCallbackRequest = Type("DestinationAccountCallbackRequest", func() {
	Description("Destination Account Callback Request")

	Attribute("ReferenceNumber", String, func() {
		Description("Unique posting reference for the transaction leg")
		MinLength(1)
		MaxLength(30)
		Example("40ca18c6765086089a1_1")
	})
	Attribute("MobileNumber", String, func() {
		Description("Recipient M-Pesa mobile number")
		MinLength(10)
		MaxLength(13)
		Example("2547xxxxxxxx")
	})
	Attribute("Amount", Float64, func() {
		Description("Transaction Amount")
		Minimum(0)
		Maximum(999999)
		Example(777)
	})
	Attribute("Narration", String, func() {
		Description("Transaction posting narration")
		MinLength(1)
		MaxLength(25)
		Example("Electricity Payment")
	})
	Attribute("TransactionID", String, func() {
		Description("Co-operative Bank's processed transaction number concatenated with M-Pesa transaction number")
		Example("116e68e0af0c38zY_NKE7HBQEIL")
	})
	Attribute("ResponseCode", String, func() {
		Description("Posting leg response code")
		Example("0")
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})

	Required("ReferenceNumber", "MobileNumber", "Amount", "Narration")
})
