package app

import (
	"github.com/DanielTitkov/cryptgame/internal/configs"
	"github.com/DanielTitkov/cryptgame/internal/logger"
	"github.com/DanielTitkov/cryptgame/internal/repository"
)

type App struct {
	cfg    configs.Config
	logger *logger.Logger
	repo   *repository.Repository
}

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo *repository.Repository,
) (*App, error) {
	app := App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}

	return &app, nil
}
