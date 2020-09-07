package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/RodrigoDev/sword-task-manager/internal/config"
	"github.com/RodrigoDev/sword-task-manager/internal/logging"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/storage"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/user"
	"github.com/RodrigoDev/sword-task-manager/internal/transport"
)

const (
	version = "0.1.0"
	service = "sword-task-manager"
)

func Main(ctx context.Context) (err error) {
	logger := logging.Logger(ctx)

	defer func() {
		if err != nil {
			logger.Error("startup", zap.Error(err))
		}
		_ = logger.Sync()
	}()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal("error loading configuration", zap.Error(err))
	}

	mysqlStorage := storage.NewMySQLStorage(cfg.MySQLConfig)

	taskService := user.NewTaskService(mysqlStorage)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		h, err := transport.New(
			transport.Health(),
			transport.Task(taskService),
		)

		if err != nil {
			return err
		}
		return transport.ListenAndServe(ctx, ":80", h)
	})

	logger.Info(fmt.Sprintf("starting %s at port %d", service, 80))

	logger.Info("shutdown", zap.Error(g.Wait()))
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(c)

	go func() {
		select {
		case <-c:
		case <-ctx.Done():
		}
		cancel()
	}()

	if err := Main(ctx); err != nil {
		os.Exit(1)
	}
}
