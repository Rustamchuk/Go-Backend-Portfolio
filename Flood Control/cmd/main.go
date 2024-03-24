package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os/exec"
	"os/signal"
	"syscall"
	"task/internal/api"
	"task/internal/repository"
	"task/internal/repository/postgres"
	"task/internal/service"
	"task/pkg/generated/proto/flood_control"
)

func main() {
	run()
	var (
		appPort string
	)

	flag.StringVar(&appPort, "app_port", "8094", "application port")

	flag.Parse()

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	floodService := service.NewFloodControlService(10, 10, repo)

	floodServiceApi := api.NewFloodControlApi(floodService)

	lsn, err := net.Listen("tcp", fmt.Sprintf(":%s", appPort))

	if err != nil {
		log.Fatal(err)
	}

	var floodServiceServer = grpc.NewServer()

	reflection.Register(floodServiceServer)
	flood_control.RegisterFloodControlServiceServer(floodServiceServer, floodServiceApi)

	if err := floodServiceServer.Serve(lsn); err != nil {
		log.Fatalf("serving error %v", err)
	}
}

func run() {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGTERM)

	floodServiceStart := exec.Command("../bin/app", "localhost:8086", "-app_port", "8085")

	go func() {
		defer func() {
			err := floodServiceStart.Process.Kill()

			if err != nil {
				log.Println(err)
			}
		}()

		if err := floodServiceStart.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
}
