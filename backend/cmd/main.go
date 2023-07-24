package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/DanielTitkov/cryptgame/internal/api"
	"github.com/DanielTitkov/cryptgame/internal/app"
	"github.com/DanielTitkov/cryptgame/internal/configs"
	"github.com/DanielTitkov/cryptgame/internal/ent"
	"github.com/DanielTitkov/cryptgame/internal/ent/migrate"
	"github.com/DanielTitkov/cryptgame/internal/logger"
	"github.com/DanielTitkov/cryptgame/internal/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := repository.NewRepository(db, logger)

	app, err := app.NewApp(cfg, logger, repo)
	if err != nil {
		logger.Fatal("failed creating app", err)
	}

	server := NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}

func NewServer(cfg configs.Config, logger *logger.Logger, app *app.App) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	if cfg.Env != "dev" {
		e.Use(middleware.Recover())
	}

	api.NewHandler(e, cfg, logger, app)
	return e
}

func Migrate(ctx context.Context, client *ent.Client) error {
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}
	return nil
}
