package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	delivery "github.com/IskanderA1/handly/iternal/delivery/http"
	"github.com/IskanderA1/handly/iternal/repository"
	"github.com/IskanderA1/handly/iternal/server"
	"github.com/IskanderA1/handly/iternal/service"
	"github.com/IskanderA1/handly/pkg/config"
	"github.com/IskanderA1/handly/pkg/token"
	"github.com/sirupsen/logrus"
)

func Run(configPath string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatalf("cannot load config", err.Error())
	}

	tokeManger, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		logrus.Fatalf("error occurred while connect to db: %s\n", err.Error())
	}

	repo := repository.NewRepositories(db)

	services := service.NewServices(repo, tokeManger, config)

	handlers := delivery.NewHandler(services)

	srv := server.NewServer(handlers.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Infof("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	logrus.Infof("Server stoped")
	if err := srv.Stop(ctx); err != nil {
		logrus.Fatalf("failed to stop server: %v", err)
	}
}
