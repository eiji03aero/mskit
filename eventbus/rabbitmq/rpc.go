package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	StatusCode_Success = 200
	StatusCode_Fail    = 400
)

func MakeSuccessResponse(p amqp.Publishing) amqp.Publishing {
	p.Headers["status_code"] = StatusCode_Success
	return p
}

func MakeFailResponse(p amqp.Publishing) amqp.Publishing {
	p.Headers["status_code"] = StatusCode_Fail
	return p
}

func IsSuccessResponse(p amqp.Publishing) bool {
	return p.Headers["status_code"] == StatusCode_Success
}
