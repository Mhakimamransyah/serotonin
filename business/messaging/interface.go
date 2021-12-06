package messaging

type Service interface {
	Start_signal() error
}

type Repository interface {
	Send(msg *Messaging, publish *Publish) error
}
