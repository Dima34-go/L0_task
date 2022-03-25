package service

import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/httpServer/repository"
)

type TodoOrder interface {
	GetOrderById(OrderId int) (todo.OrderItems, error)
}

type Service struct {
	TodoOrder
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TodoOrder: NewTodoOrderService(repo),
	}
}

