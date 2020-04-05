package consumer

import (
	consumerdmn "consumer/domain/consumer"
)

type Service interface {
	CreateConsumer(cmd consumerdmn.CreateConsumer) (id string, err error)
	GetConsumer(id string) (consumer *consumerdmn.Consumer, err error)
	ValidateOrder(cmd consumerdmn.ValidateOrder) error
}
