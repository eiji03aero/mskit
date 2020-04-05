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
	"github.com/streadway/amqp"
)

type proxy struct {
	client *rabbitmq.Client
}

func New(c *rabbitmq.Client) {{ .PkgName }}.{{ .InterfaceName }} {
	return &proxy{
		client: c,
	}
}

func (p *proxy) Method() (err error) {
	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "target.rpc.",
				Publishing: amqp.Publishing{},
			},
		).
		Exec()

	return
}`
}
