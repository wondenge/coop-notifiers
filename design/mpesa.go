package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// 11.SendToMPesa
var SendToMpesaCallbackRequest = Type("SendToMpesaCallbackRequest", func() {
	Description("Send to Mpesa callback request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("MessageDateTime", String, func() {
		Description("Acknowledgement message creation timestamp")
		Format(FormatDateTime)
		Example("2017-12-04T09:27:00")
	})
	Attribute("MessageCode", String, func() {
		Description("Message Response Code")
		Example("0")
	})
	Attribute("MessageDescription", String, func() {
		Description("Message Code description")
		Example("FULL SUCCESS")
	})
	Attribute("CallBackUrl", String, func() {
		Description("Your callback URL that will receive transaction processing results")
		Example("https://yourdomain.com/ft-callback")
	})
	Attribute("Source", SourceAccountCallbackRequest)
	Attribute("Destinations", DestinationAccountCallbackRequest)
	Required("MessageReference",
		"MessageDateTime",
		"MessageCode",
		"MessageDescription",
		"Source",
		"Destinations",
	)
})
