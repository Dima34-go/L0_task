package service

import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/httpServer/repository"
)

type TodoOrderService struct {
	repo repository.TodoOrder
}

func NewTodoOrderService(repo repository.TodoOrder) *TodoOrderService {
	return &TodoOrderService{
		repo: repo,
	}
}
func (s TodoOrderService) GetOrderById(OrderId int) (todo.OrderItems, error) {
	return s.repo.GetOrderById(OrderId)
}
