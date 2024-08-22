package main

//nolint:stylecheck
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"resources/pkg/api"
	"resources/pkg/config"
	"resources/pkg/logger"
)

func main() {
	conf, err := config.New()
	if err != nil {
		logger.Fatalf("Can't read config file: %s", err)
	}

	apiServer := api.NewServer(&conf.Server)
	runErr := make(chan error, 1)
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	logger.Infof("Start api-gateway: %s", time.Now())

	go func() {
		logger.Infof("HTTP API server start listen on: %s", apiServer.Addr)
		err = apiServer.ListenAndServe()
		if err != nil {
			runErr <- fmt.Errorf("can't start http server: %w", err)
		}
	}()

	select {
	case err = <-runErr:
		logger.Fatalf("Running error: %s", err)
	case s := <-quitCh:
		logger.Infof("Received signal: %v. Running graceful shutdown...", s)
		ctx := context.Background()

		err = apiServer.Shutdown(ctx)
		if err != nil {
			logger.Infof("Can't shutdown API server: %s", err)
		}
	}
}
