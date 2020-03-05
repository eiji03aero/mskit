package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Client struct {
	conn *amqp.Connection
}

type Option struct {
	Host string
	Port string
}

func getUrl(opt Option) string {
	return fmt.Sprintf(
		"amqp://%s:%s/",
		opt.Host,
		opt.Port,
	)
}

func NewClient(opt Option) (cli *Client, err error) {
	conn, err := amqp.Dial(getUrl(opt))
	if err != nil {
		return
	}

	cli = &Client{
		conn: conn,
	}
	return
}

func (c *Client) NewPublisher() *Publisher {
	return NewPublisher(c.conn)
}

func (c *Client) NewConsumer() *Consumer {
	return NewConsumer(c.conn)
}

func (c *Client) NewRPCClient() *RPCClient {
	return NewRPCClient(c.conn)
}

func (c *Client) NewRPCServer() *RPCServer {
	return NewRPCServer(c.conn)
}
