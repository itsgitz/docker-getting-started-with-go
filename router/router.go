package router

import (
	"go-docker/handler"

	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/getMusicPlaylist", handler.GetMusicPlaylist)
}
