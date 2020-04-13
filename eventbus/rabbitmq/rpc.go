package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	StatusCode_Success  int32 = 200
	StatusCode_Fail     int32 = 400
	StatusCode_NotFound int32 = 404
)

func MakeSuccessResponse(p amqp.Publishing) amqp.Publishing {
	return setStatusCode(p, StatusCode_Success)
}

func MakeFailResponse(p amqp.Publishing, opts ...interface{}) (pr amqp.Publishing) {
	pr = setStatusCode(p, StatusCode_Fail)

	for _, opt := range opts {
		switch o := opt.(type) {
		case error:
			pr = setErrorMessage(pr, o.Error())
		}
	}

	return
}

func MakeNotFoundResponse(p amqp.Publishing, opts ...interface{}) (pr amqp.Publishing) {
	pr = setStatusCode(p, StatusCode_NotFound)

	for _, opt := range opts {
		switch o := opt.(type) {
		case error:
			pr = setErrorMessage(pr, o.Error())
		}
	}

	return
}

func IsSuccessResponse(d amqp.Delivery) bool {
	return d.Headers["status_code"] == StatusCode_Success
}

func IsNotFoundResponse(d amqp.Delivery) bool {
	return d.Headers["status_code"] == StatusCode_NotFound
}

func setStatusCode(p amqp.Publishing, code int32) amqp.Publishing {
	if p.Headers == nil {
		p.Headers = amqp.Table{}
	}

	p.Headers["status_code"] = code

	return p
}

func setErrorMessage(p amqp.Publishing, errMsg string) amqp.Publishing {
	if p.Headers == nil {
		p.Headers = amqp.Table{}
	}

	p.Headers["error_message"] = errMsg

	return p
}

func getErrorMessage(p amqp.Delivery) (errMsg string) {
	if p.Headers == nil || p.Headers["error_message"] == nil {
		return
	}

	errMsg = p.Headers["error_message"].(string)
	return
}
