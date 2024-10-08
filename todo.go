package todoapi

import (
	"context"

	todo "github.com/too-common-name/todo-microservice/gen/todo"
	"goa.design/clue/log"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct{}

// NewTodo returns the todo service implementation.
func NewTodo() todo.Service {
	return &todosrvc{}
}

// Create implements create.
func (s *todosrvc) Create(ctx context.Context, p *todo.Todo) (res *todo.Todo, err error) {
	res = &todo.Todo{}
	log.Printf(ctx, "todo.create")
	return
}

// List implements list.
func (s *todosrvc) List(ctx context.Context) (res []*todo.Todo, err error) {
	log.Printf(ctx, "todo.list")
	return
}

// Get implements get.
func (s *todosrvc) Get(ctx context.Context, p *todo.GetPayload) (res *todo.Todo, err error) {
	res = &todo.Todo{}
	log.Printf(ctx, "todo.get")
	return
}

// Delete a ToDo by ID
func (s *todosrvc) Delete(ctx context.Context, p *todo.DeletePayload) (err error) {
	log.Printf(ctx, "todo.delete")
	return
}

// Partially update a ToDo
func (s *todosrvc) Update(ctx context.Context, p *todo.UpdatePayload) (res *todo.Todo, err error) {
	res = &todo.Todo{}
	log.Printf(ctx, "todo.update")
	return
}
