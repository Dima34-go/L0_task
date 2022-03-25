package serviceNATS

import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/natsClient/repositoryNATS"
	"encoding/json"
	stan "github.com/nats-io/go-nats-streaming"
	"log"
)

type OrderService struct {
	repo repositoryNATS.TodoOrder
}

func NewTodoOrderService(repo repositoryNATS.TodoOrder) *OrderService {
	return &OrderService{
		repo: repo,
	}
}
func (s *OrderService) AddOrderItems(msg *stan.Msg) {
	Order:=new(todo.OrderItems)
	if err:=json.Unmarshal(msg.Data,&Order);err == nil{
		err=s.repo.AddOrderItems(Order)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}
