package api

import (
	"net/http"

	"github.com/DanielTitkov/cryptgame/internal/app"
	"github.com/DanielTitkov/cryptgame/internal/configs"
	"github.com/DanielTitkov/cryptgame/internal/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	v1.GET("/game", h.GetGame)
	v1.POST("/key", h.PostKey)

}

func (h *Handler) GetGame(c echo.Context) error {
	game := &GetGameResponse{
		ID:        uuid.New(),
		KeyLength: 5,
		Challenge: `L5f2 5s wh1t h1pp2ns wh2n y2u'32 bus5 m1k5ng 2th32 pl1ns.`,
	}
	return c.JSON(http.StatusOK, game)
}

func (h *Handler) PostKey(c echo.Context) error {
	req := new(PostKeyRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &PostKeyResponse{
		Success: req.Key == "aerio",
		GameID:  req.GameID, // FIXME
		Mask:    "00000",    // FIXME
	})
}
