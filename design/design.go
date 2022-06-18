package design

import "goa.design/goa/v3/dsl"

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
})
