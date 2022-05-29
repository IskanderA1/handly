package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/IskanderA1/handly"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error env: %s", err.Error())
	}
	http.HandleFunc("/", initRoute)
	srv := new(handly.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), nil); err != nil {
			logrus.Fatalf("error http sever: %s", err.Error())
		}
	}()
	logrus.Print("Hadnly Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logrus.Print("Hadnly  Shutting Down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, i'm Handly!")
}
