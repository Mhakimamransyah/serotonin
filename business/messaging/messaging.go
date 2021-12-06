package messaging

type Publish struct {
	Queue     string
	Exchanger string
	Key       string
}

type Consume struct {
	queue string
}

type Messaging struct {
	Id      int    `json:"Id"`
	Message string `json:"Message"`
}
