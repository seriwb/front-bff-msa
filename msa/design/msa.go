package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("msa", func() {
	Title("Some function Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("msa-server", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("some-function", func() {
	Description("The calc service performs operations on numbers.")

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
