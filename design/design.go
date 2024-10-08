package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("todo", func() {
	Title("Todo microservice")
	Description("ToDo Microservice's API")
	HTTP(func() {
		Path("/api/v1")
	})
	Server("todo", func() {
		Services("todo")
		Host("development", func() {
			URI("http://0.0.0.0:8000")
		})
	})
})

var ToDo = Type("Todo", func() {
	Attribute("id", String, func() {
	})

	Attribute("title", String, func() {
		MaxLength(100)
	})

	Attribute("description", String, func() {
		MaxLength(1000)
	})

	Attribute("category_id", String, func() {})

	Attribute("user_id", String, func() {})

	Attribute("due_date", String, func() {
		Format(FormatDate)
	})

	Attribute("priority", String, func() {
		Enum("LOW", "MEDIUM", "HIGH")
		Default("LOW")
	})

	Attribute("status", String, func() {
		Enum("PENDING", "COMPLETED")
		Default("PENDING")
	})

	Required("title", "description", "category_id", "user_id", "due_date")
})
