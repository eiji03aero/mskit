package consumer

type Consumer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Consumer) ValidateOrder() error {
	return nil
}
