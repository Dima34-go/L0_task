package todo

import (
	stan "github.com/nats-io/go-nats-streaming"
	"strconv"
)

type NatsConfig struct {
	StanClusterID string
	ClientID      string
	Subject 	  string
	QGroup        string
	DurableName   string
	SubsAmount    int
}

func Connect(cfg NatsConfig,cb stan.MsgHandler ) error {
	for i:=0;i < cfg.SubsAmount;i++{
		sc,err := stan.Connect(cfg.StanClusterID,cfg.ClientID+strconv.Itoa(i))
		if err != nil {
			return err
		}
		_,err = sc.QueueSubscribe(cfg.Subject,cfg.QGroup,cb,stan.DurableName(cfg.DurableName))
		if err != nil {
			return err
		}
	}
	return nil
}

