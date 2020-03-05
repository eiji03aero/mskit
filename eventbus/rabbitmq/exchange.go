package rabbitmq

import (
	"github.com/streadway/amqp"
)

type ExchangeOption struct {
	Name        string
	Type        string
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Arguments   amqp.Table
}

var DefaultExchangeOption = ExchangeOption{
	Name:        "",
	Type:        "fanout",
	Durable:     false,
	AutoDeleted: false,
	Internal:    false,
	NoWait:      false,
	Arguments:   nil,
}
