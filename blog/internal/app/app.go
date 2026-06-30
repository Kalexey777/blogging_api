package app

import (
	"context"
	"easyapi/blog/internal/config"
	"easyapi/blog/internal/controller"
	"easyapi/blog/internal/repository"
	db "easyapi/blog/migrations"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Run(logger *zap.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.ConstructPostgresURL())

	if err != nil {
		logger.Error("can not create pgxpool", zap.Error(err))
		return
	}

	defer dbPool.Close()

	db.SetupPostgres(dbPool, logger)

	repo := repository.NewPostgresRepository(dbPool, logger)
	ctrl := controller.New(logger, repo)
	ctrl.Run("localhost:8080")
}
