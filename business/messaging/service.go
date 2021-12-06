package messaging

import "math/rand"

type MessagingService struct {
	MessagingRepository Repository
}

func InitMessagingService(repos Repository) *MessagingService {
	return &MessagingService{
		MessagingRepository: repos,
	}
}

func (service *MessagingService) Start_signal() error {
	msg := Messaging{
		Id:      rand.Int(),
		Message: "Let's Start Transfer",
	}
	publish := Publish{
		Queue:     "init_ujian",
		Exchanger: "ujian_exchangers",
		Key:       "signal.init.ujian",
	}
	err := service.MessagingRepository.Send(&msg, &publish)
	if err != nil {
		return err
	}
	return nil
}
