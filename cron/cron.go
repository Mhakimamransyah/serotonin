package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

type Task func()

func StartCron(date_start string, task Task) {
	s := gocron.NewScheduler(time.Local)
	s.Every(1).Day().At(date_start).Do(task)
	s.StartAsync()
	fmt.Println("Cron Start : ", s.IsRunning())
}
