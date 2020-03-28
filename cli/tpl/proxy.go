package tpl

func RootProxy() string {
	return `package {{ .PkgName }}
`
}

func ProxyTemplate() string {
	return `package {{ .LowerName }}

import (
	"{{ .PkgName }}"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type proxy struct {
	client *rabbitmq.Client
}

func New(
	client *rabbitmq.Client,
) {{ .PkgName }}.{{ .InterfaceName }} {
	return &proxy{
		client: client,
	}
}`
}

func ProxyImplTemplate() string {
	return `package {{ .LowerName }}

import (
	"log"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

func (p *proxy) Method() (err error) {
	log.Println("Method: ")

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "",
				Publishing: amqp.Publishing{},
			},
		).
		Exec()

	return
}`
}
