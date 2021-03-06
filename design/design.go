package design

import "goa.design/goa/v3/dsl"
import _ "goa.design/plugins/v3/zerologger"

var Value = dsl.Type("value", dsl.Int, func() {
	dsl.Description("Value of the single num.")
	dsl.Minimum(3)
})

var GetResult = dsl.Type("get_result", func() {
	dsl.Attribute("value", Value)

	dsl.Required("value")
})

var SetResult = dsl.Type("set_result", func() {
	dsl.Attribute("success", dsl.Boolean)

	dsl.Required("success")
})

var _ = dsl.API("knowtfolio", func() {
	dsl.Title("Knowtfolio Backend")
	dsl.Description("HTTP service for multiplying numbers, a goa teaser")
	dsl.Server("knowtfolio", func() {
		dsl.Host("localhost", func() { dsl.URI("http://localhost:8080") })

		dsl.Services("SingleNumRegister")
	})
})

var _ = dsl.Service("SingleNumRegister", func() {
	dsl.Description("Call SingleNumRegister contract.")

	dsl.Method("GetNum", func() {
		dsl.Description("Get the value of the single num.")

		dsl.Result(GetResult)

		dsl.HTTP(func() {
			dsl.GET("/")

			dsl.Response(dsl.StatusOK)
		})
	})

	dsl.Method("SetNum", func() {
		dsl.Description("Set the value of the single num.")

		dsl.Payload(func() {
			dsl.Attribute("val", Value, "Value to set")

			dsl.Required("val")
		})
		dsl.Result(SetResult)

		dsl.HTTP(func() {
			dsl.GET("/set/{val}")

			dsl.Response(dsl.StatusOK)
		})
	})

	dsl.Method("Html", func() {
		dsl.Description("Get the value of the single num as HTML.")

		dsl.Result(dsl.Bytes)

		dsl.HTTP(func() {
			dsl.GET("/html")

			dsl.Response(dsl.StatusOK, func() {
				dsl.ContentType("text/html; charset=UTF-8")
			})
		})
	})
})
