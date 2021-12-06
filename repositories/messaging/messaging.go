package messaging

import (
	"encoding/json"
	"serotonin/business/messaging"

	"github.com/streadway/amqp"
)

type MessagingRepository struct {
	channel *amqp.Channel
}

func InitMessagingRepository(channel *amqp.Channel) *MessagingRepository {
	return &MessagingRepository{
		channel: channel,
	}
}

func (repository *MessagingRepository) SendMessage(exchanger, queue_name, routing_key string, msg interface{}) error {
	_, err := repository.channel.QueueDeclare(
		queue_name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = repository.channel.Publish(
		exchanger,
		routing_key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (repository *MessagingRepository) Send(msg *messaging.Messaging, publish *messaging.Publish) error {

	_, err := repository.channel.QueueDeclare(
		publish.Queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = repository.channel.Publish(
		publish.Exchanger,
		publish.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
