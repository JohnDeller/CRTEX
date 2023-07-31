package app

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
	"transactionService/internal/broker/consumer"
	"transactionService/internal/broker/event"
	"transactionService/internal/broker/listener"
	"transactionService/internal/broker/rabbitmq"
	"transactionService/internal/config"
	"transactionService/internal/handler"
	"transactionService/internal/repository"
	"transactionService/internal/repository/postgres"
	"transactionService/internal/service"
	"transactionService/server"
)

const (
	configPath            = "configs/config.json"
	listenersIntervalTime = 1
	queueName             = "event_queue"
)

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg, err := config.InitConfig(configPath)

	if err != nil {
		logrus.Fatalf("failed to initialize config: %s", err.Error())
	}
	// todo вынести env в config.go , чтоб перед глазами не маячило
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		log.Fatalf("DB_PASSWORD environment variable is empty")
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Username: cfg.Database.UserName,
		Password: dbPassword,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})

	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	listeners := listener.NewListener(repos)
	brokerPassword := os.Getenv("BROKER_PASSWORD")

	if brokerPassword == "" {
		log.Fatalf("BROKER_PASSWORD environment variable is empty")
	}
	rabbitMqConnect, err := rabbitmq.NewRebbitMqConnection(rabbitmq.Config{
		User:     cfg.Broker.User,
		Password: brokerPassword,
		Host:     cfg.Broker.Host,
		Port:     cfg.Broker.Port,
	})
	if err != nil {
		logrus.Fatalf("failed to connect rabbit mq: %s", err.Error())
	}
	services := service.NewService(repos, rabbitMqConnect)
	handlers := handler.NewHandler(services)
	go func() {
		for {
			//logrus.Println("checking queue")
			consumer.Consume(rabbitMqConnect, event.Event{}, queueName, listeners)
			//time.Sleep(listenersIntervalTime * time.Second)
		}
	}()

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TransactionsService Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("TransactionsService Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
	if err := rabbitMqConnect.Close(); err != nil {
		logrus.Errorf("error occured on rabbbit mq connection close: %s", err.Error())
	}

}
