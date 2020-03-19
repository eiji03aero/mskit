package consumer

type CreateConsumer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ValidateOrder struct {
	Id    string `json:"id"`
	Total int    `json:"total"`
}
