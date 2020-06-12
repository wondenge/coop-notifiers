package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var SourceAccountCallbackRequest = Type("SourceAccountCallbackRequest", func() {
	Description("Source Account Callback Request")

	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
		Example("54321987654321")
	})
	Attribute("Amount", UInt64, func() {
		Description("Transaction Amount")
		Minimum(0.01)
		Maximum(999999.99)
		Example(777)
	})
	Attribute("TransactionCurrency", String, func() {
		Description("Posting account currency in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("KES")
	})
	Attribute("Narration", String, func() {
		Description("Posting account transaction narration")
		MinLength(1)
		MaxLength(25)
		Example("Supplier Payment")
		Example("Electricity Payment")
	})
	Attribute("ResponseCode", String, func() {
		Description("Posting leg response code")
		Example("0")
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})
	Required("AccountNumber", "Amount", "TransactionCurrency", "Narration", "ResponseCode", "ResponseDescription")
})
