package logging

import (
	"fmt"
	"os"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}
var logging *logrus.Logger

func GetLogger(filename string) *logrus.Logger {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.OpenFile("/log_serotonin/"+filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Println(file.Name())
	fmt.Println(file)
	fmt.Println(os.Environ())
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	if logging == nil {
		logging = logrus.New()
	}
	logging.SetOutput(file)
	logging.SetFormatter(&logrus.JSONFormatter{})
	return logging
}
