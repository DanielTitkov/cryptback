package repository

import (
	"github.com/DanielTitkov/cryptgame/internal/ent"
	"github.com/DanielTitkov/cryptgame/internal/logger"
)

type Repository struct {
	client *ent.Client
	logger *logger.Logger
}

func NewRepository(client *ent.Client, logger *logger.Logger) *Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}
