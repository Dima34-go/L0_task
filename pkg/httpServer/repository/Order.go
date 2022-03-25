package repository

import "C"
import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/httpServer/appError"
	"WB_GO_L0/pkg/storage"
)

type OrderPSql struct {
	C *storage.Cache
}

func NewTodoOrderPSql(cache *storage.Cache) *OrderPSql {
	return &OrderPSql{C: cache}
}
func (r *OrderPSql) GetOrderById(OrderId int) (todo.OrderItems, error) {
	value, ok := r.C.Get(OrderId)
	if !ok{
		return value, appError.ErrOrderNotFound
	}
	return value, nil
}
