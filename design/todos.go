package design

import . "goa.design/goa/v3/dsl"

var _ = Service("todo", func() {
	Description("API for ToDos")
	Error("not_found", String, "ToDo not found")
	HTTP(func() {
		Path("/todos")
	})
	Method("create", func() {
		Payload(ToDo)
		Result(ToDo)
		HTTP(func() {
			POST("/")
		})
	})
	Method("list", func() {
		Result(ArrayOf(ToDo))
		HTTP(func() {
			GET("/")
		})
	})
	Method("get", func() {
		Payload(func() {
			Field(1, "id", String, "ToDo ID")
			Required("id")
		})
		Result(ToDo)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound, func() {
				Description("The ToDo with the given ID was not found.")
			})
		})
	})
	Method("delete", func() {
		Description("Delete a ToDo by ID")
		Payload(func() {
			Field(1, "id", String, "ToDo ID to delete")
			Required("id")
		})
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound, func() {
				Description("The ToDo with the given ID was not found.")
			})
		})
	})

	Method("update", func() {
		Description("Partially update a ToDo")
		Payload(func() {
			Attribute("id", String, "ID of the ToDo to update")
			Attribute("todo", ToDo, "Fields to update")
			Required("id", "todo")
		})
		Result(ToDo)
		HTTP(func() {
			PATCH("/{id}")
			Body("todo")
			Response(StatusOK)
			Response("not_found", StatusNotFound, func() {
				Description("The ToDo with the given ID was not found.")
			})
		})
	})
})
