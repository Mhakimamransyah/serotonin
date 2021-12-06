package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	Api "serotonin/api"
	ControllerPilgan "serotonin/api/v1/controllers/pilgan"
	ServiceMessaging "serotonin/business/messaging"
	ServicePilgan "serotonin/business/pilgan"
	"serotonin/config"
	Worker "serotonin/cron"
	"serotonin/cron/tasks"
	"serotonin/migrations"
	RepositoryMessaging "serotonin/repositories/messaging"
	RepositoryPilgan "serotonin/repositories/pilgan"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMongoDb(appconfig *config.AppConfig) (*mongo.Database, context.Context) {
	var ctx = context.Background()
	configDB := map[string]string{
		"DB_Collection": appconfig.MongoCollection,
		"DB_Port":       strconv.Itoa(appconfig.MongoPort),
		"DB_Host":       appconfig.MongoHost,
		"DB_Username":   appconfig.MongoUsername,
		"DB_Password":   appconfig.MongoPassword,
	}
	fmt.Println(configDB)
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", configDB["DB_Username"], configDB["DB_Password"], configDB["DB_Host"], configDB["DB_Port"])
	clientOptions := options.Client()
	clientOptions.ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client.Database(configDB["DB_Collection"]), ctx
}

func initDatabaseMysql(appconfig *config.AppConfig) *gorm.DB {
	configDB := map[string]string{
		"DB_Username": appconfig.DbUsername,
		"DB_Password": appconfig.DbPassword,
		"DB_Port":     strconv.Itoa(appconfig.DbPort),
		"DB_Host":     appconfig.DbHost,
		"DB_Name":     appconfig.DbName,
	}
	fmt.Println(configDB)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	migrations.InitMigrate(db)
	fmt.Println("Connect to ", db.Migrator().CurrentDatabase())
	return db
}

func initMessaging(appconfig *config.AppConfig) *amqp.Channel {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/", appconfig.MsgBrokerUsername, appconfig.MsgBrokerPassword,
		appconfig.MsgBrokerHost, appconfig.MsgBrokerPort)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	fmt.Println("Messaging connect")
	return ch
}

func main() {
	config := config.GetConfig()
	// db := initDatabaseMysql(config)
	client, ctx := initMongoDb(config)
	channel := initMessaging(config)

	messagingRepository := RepositoryMessaging.InitMessagingRepository(channel)
	messagingService := ServiceMessaging.InitMessagingService(messagingRepository)

	pilganRepository := RepositoryPilgan.InitPilganRepository(client, &ctx, "siswa_pilgan")
	servicePilgan := ServicePilgan.InitPilganService(pilganRepository, *messagingRepository)
	pilganController := ControllerPilgan.InitPilganController(servicePilgan)

	task := tasks.InitTask(messagingService)

	Worker.StartCron(config.CronDate, task.Init_signal)
	Worker.InitConsumer("transfer_ujian", channel).StartConsumer(servicePilgan)

	e := echo.New()

	Api.RegisterPath(e, pilganController)

	go func() {
		address := fmt.Sprintf(":%d", config.AppPort)
		if err := e.Start(address); err != nil {
			log.Fatalf("error when starting echo %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}
