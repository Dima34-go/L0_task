package repositoryNATS

import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/storage"
	"github.com/jmoiron/sqlx"
)

type TodoOrder interface {
	AddOrderItems(Order *todo.OrderItems)  error
}
type Repository struct {
	TodoOrder
}

func NewRepository(c *storage.Cache,db *sqlx.DB) *Repository {
	return &Repository{
		TodoOrder: NewTodoOrderPSql(c,db),
	}
}

