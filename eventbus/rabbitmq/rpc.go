package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	StatusCode_Success int32 = 200
	StatusCode_Fail    int32 = 400
)

func MakeSuccessResponse(p amqp.Publishing) amqp.Publishing {
	return setStatusCode(p, StatusCode_Success)
}

func MakeFailResponse(p amqp.Publishing) amqp.Publishing {
	return setStatusCode(p, StatusCode_Fail)
}

func IsSuccessResponse(d amqp.Delivery) bool {
	return d.Headers["status_code"] == StatusCode_Success
}

func setStatusCode(p amqp.Publishing, code int32) amqp.Publishing {
	if p.Headers == nil {
		p.Headers = amqp.Table{}
	}

	p.Headers["status_code"] = code

	return p
}
