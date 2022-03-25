package main

import (
	todo "WB_GO_L0"
	config "WB_GO_L0/configs"
	"WB_GO_L0/pkg/httpServer/handler"
	"WB_GO_L0/pkg/httpServer/repository"
	"WB_GO_L0/pkg/httpServer/service"
	"WB_GO_L0/pkg/natsClient/repositoryNATS"
	"WB_GO_L0/pkg/natsClient/serviceNATS"
	"WB_GO_L0/pkg/storage"
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Initializing config
	cfg,err:=config.InitConfig()
	if err!=nil{
		log.Fatalf("error initializing config: %s", err.Error())
	}
	//Initializing env variables
	if err:=godotenv.Load();err!=nil{
		log.Fatalf("error initializing env variables: %s", err.Error())
	}
	//Connecting to PostgreSQL Database
	db, err := storage.NewPgsqlDB(storage.ConfigDB{
		Host:     cfg.DB.Host,
		Port:    cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.DBName,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}
	//Creating Cache
	cache, err := storage.GetCache(db)
	if err != nil {
		log.Fatalf("error initializing cache: %s", err.Error())
	}

	//Creating an architecture httpServer
	repos := repository.NewRepository(cache)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	//Start Server
	srv := new(todo.Server)
	go func() {
		if err := srv.Run(cfg.ServerHTTP.Port, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error running http server: %v", err)
		}
	}()

	//Creating an architecture Nats Client
	natsRepo:=repositoryNATS.NewRepository(cache,db)
	natsService:=serviceNATS.NewService(natsRepo)
	//Create Client and connect to NatsServer
	natsCfg:= todo.NatsConfig{
		StanClusterID: cfg.NATS.ClusterID,
		ClientID: cfg.NATS.ClientID,
		Subject: cfg.NATS.Subject,
		QGroup:   cfg.NATS.QGroup,
		DurableName: cfg.NATS.DurableName,
		SubsAmount: 3,
	}
	if err=todo.Connect(natsCfg,natsService.InitHandler());err!=nil{
		log.Fatalf("error creating connect with Nats Server: %s", err.Error())
	}

	log.Print("Application Started")

	//Catch the shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("TodoApp Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}




