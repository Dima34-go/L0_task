package repository

import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/storage"
)

type TodoOrder interface {
	GetOrderById(OrderId int) (todo.OrderItems, error)
}
type Repository struct {
	TodoOrder
}

func NewRepository(c *storage.Cache) *Repository {
	return &Repository{
		TodoOrder: NewTodoOrderPSql(c),
	}
}
