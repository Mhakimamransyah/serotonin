package cron

import (
	"encoding/json"
	"fmt"
	"serotonin/business/pilgan"
	logging "serotonin/util/logging"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Handler func(interface{})

type Consumer struct {
	// regsiter all service handler here
	Queue   string
	Channel *amqp.Channel
	handler Handler
}

func InitConsumer(queue string, channel *amqp.Channel) *Consumer {
	return &Consumer{
		Queue:   queue,
		Channel: channel,
	}
}

func (consumer *Consumer) StartConsumer(service *pilgan.PilganService) {
	q, err := consumer.Channel.QueueDeclare(
		consumer.Queue, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		panic(err)
	}
	msgs, err := consumer.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}
	go func() {
		for d := range msgs {
			var data pilgan.PilganTone
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				// logging error format here
				logging.GetLogger("errorformat.log").WithFields(logrus.Fields{
					"data": string(d.Body),
				}).Error(err.Error())

			} else {
				service.SyncNewDataPilgan(&data)
			}
			d.Ack(true)
		}
	}()
	fmt.Println("Consumer Start", q.Name)
}
