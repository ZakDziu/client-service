package main

//nolint:stylecheck
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-gateway/pkg/api"
	"api-gateway/pkg/api_builder"
	"api-gateway/pkg/authmiddleware/appauth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conf, err := config.New()
	if err != nil {
		logger.Fatalf("Can't read config file: %s", err)
	}

	apiBuilder := api_builder.New(conf.ServicesPath)

	middleware := appauth.NewAuthMiddleware(apiBuilder)

	apiServer := api.NewServer(&conf.Server, middleware, apiBuilder)
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
