package tasks

import (
	"log"
	"serotonin/business/messaging"
)

type Task struct {
	MessageService *messaging.MessagingService
}

func InitTask(msg *messaging.MessagingService) *Task {
	return &Task{
		MessageService: msg,
	}
}

// Init_signal method untuk menginisiasi transfer data dari elearning menggunakan message broker
func (task *Task) Init_signal() {
	err := task.MessageService.Start_signal()
	if err != nil {
		log.Printf("%s", err)
	}
}
