package handler

import "github.com/gofiber/fiber/v2"

type MusicPlayList struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

func GetMusicPlaylist(c *fiber.Ctx) error {
	var data []MusicPlayList

	data = []MusicPlayList{
		{
			Title:  "21 Guns",
			Artist: "Green Day",
			Album:  "21st Century Breakdown",
		},
		{
			Title:  "Dear God",
			Artist: "Avenged Sevenfold",
			Album:  "Avenged Sevenfold",
		},
	}

	return c.JSON(data)
}
