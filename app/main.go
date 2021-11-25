package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	Api "serotonin/api"
	ControllersUser "serotonin/api/v1/controllers/users"
	ServiceUser "serotonin/business/users"
	"serotonin/config"
	"serotonin/migrations"
	RepositoryUser "serotonin/repositories/users"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMongoDb() {

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

func main() {
	config := config.GetConfig()
	db := initDatabaseMysql(config)

	userRepository := RepositoryUser.InitRepository(db)
	userService := ServiceUser.InitUserService(userRepository)
	userController := ControllersUser.InitUserController(userService)

	e := echo.New()

	Api.RegisterPath(e, userController)

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
