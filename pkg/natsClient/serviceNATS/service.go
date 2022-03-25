package serviceNATS

import (
	"WB_GO_L0/pkg/natsClient/repositoryNATS"
	stan "github.com/nats-io/go-nats-streaming"
)

type TodoOrder interface {
	AddOrderItems(msg *stan.Msg)
}

type Service struct {
	TodoOrder
}

func NewService(repo repositoryNATS.TodoOrder) *Service {
	return &Service{
		NewTodoOrderService(repo),
	}
}

func (s *Service) InitHandler() stan.MsgHandler {
	return s.TodoOrder.AddOrderItems
}