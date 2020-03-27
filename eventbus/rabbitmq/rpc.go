package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	StatusCode_Success = 200
	StatusCode_Fail    = 400
)

func MakeSuccessResponse(p amqp.Publishing) amqp.Publishing {
	return setStatusCode(p, StatusCode_Success)
}

func MakeFailResponse(p amqp.Publishing) amqp.Publishing {
	return setStatusCode(p, StatusCode_Fail)
}

func IsSuccessResponse(p amqp.Publishing) bool {
	return p.Headers["status_code"] == StatusCode_Success
}

func setStatusCode(p amqp.Publishing, code int) amqp.Publishing {
	if p.Headers == nil {
		p.Headers = amqp.Table{}
	}

	p.Headers["status_code"] = code

	return p
}
