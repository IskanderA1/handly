package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	delivery "github.com/IskanderA1/handly/iternal/delivery/http"
	"github.com/IskanderA1/handly/iternal/server"
	"github.com/IskanderA1/handly/iternal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run(configPath string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	services := service.NewServices()

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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
